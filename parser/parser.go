package parser

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type column struct {
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Null   bool        `json:"null"`
	Inc    bool        `json:"auto_increment"`
	DefVal interface{} `json:"default_value"`
	Fk     string      `json:"foreign_key"`
}

type table struct {
	Columns []column   `json:"columns"`
	Pk      string     `json:"primary_key"`
	Unis    [][]string `json:"unique_keys"`
}

type database map[string]table

type schema map[string]database

func parse(filepath string) (*schema, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	schema := &schema{}
	err = json.Unmarshal(b, schema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

type foreignKey struct {
	srcCol string
	dstCol string
}

var typeMap = map[string]string{
	"string":   "varchar(512)",
	"int":      "int",
	"float":    "decimal",
	"bool":     "bool",
	"bytes":    "blob",
	"date":     "date",
	"datetime": "datetime",
}

func create(schema *schema, username, password, address string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/information_schema", username, password, address))
	if err != nil {
		return err
	}
	defer db.Close()
	fks := make([]foreignKey, 0, 16)
	for dbName, tables := range *schema {
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return err
		}
		for tableName, table := range tables {
			colList := make([]string, len(table.Columns))
			for i, col := range table.Columns {
				l := make([]string, 0, 6)
				l = append(l, col.Name)
				l = append(l, typeMap[col.Type])
				if !col.Null {
					l = append(l, "not null")
				}
				if col.Inc {
					l = append(l, "auto_increment")
				}
				if col.DefVal != nil {
					switch d := col.DefVal.(type) {
					case string:
						if strings.HasPrefix(d, "@@") {
							l = append(l, fmt.Sprintf("default %s", strings.TrimLeft(d, "@")))
						} else {
							l = append(l, fmt.Sprintf("default %s", d))
						}
					default:
						l = append(l, fmt.Sprintf("default %#v", col.DefVal))
					}
				}
				if col.Fk != "" {
					fks = append(fks, foreignKey{fmt.Sprintf("%s.%s.%s", dbName, tableName, col.Name), col.Fk})
				}
				colList[i] = strings.Join(l, " ")
			}
			var unis []string
			if len(table.Unis) != 0 {
				for _, uks := range table.Unis {
					unis = append(unis, fmt.Sprintf("unique key (%s)", strings.Join(uks, ", ")))
				}
			}
			var uc string
			if len(unis) != 0 {
				uc = ", " + strings.Join(unis, ", ")
			}
			stmt := fmt.Sprintf("CREATE TABLE %s.%s (%s, primary key (%s) %s)", dbName, tableName, strings.Join(colList, ", "), table.Pk, uc)
			_, err := db.Exec(stmt)
			if err != nil {
				return err
			}
		}
	}
	for _, fk := range fks {
		srcL := strings.Split(fk.srcCol, ".")
		dstL := strings.Split(fk.dstCol, ".")
		stmt := fmt.Sprintf("ALTER TABLE %s.%s ADD FOREIGN KEY (%s) REFERENCES %s.%s(%s) ON DELETE CASCADE", srcL[0], srcL[1], srcL[2], dstL[0], dstL[1], dstL[2])
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
