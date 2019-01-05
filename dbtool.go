package nborm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/wangjun861205/nbfmt"

	"github.com/go-sql-driver/mysql"
)

var goToSQLMap = map[string]string{
	"StringField":   "varchar(128)",
	"IntField":      "int",
	"FloatField":    "decimal(65, 2)",
	"BoolField":     "bool",
	"BinaryField":   "blob",
	"DateField":     "date",
	"DatetimeField": "datetime",
}

var defValRe = regexp.MustCompile(`'(.*)'`)
var tagFieldRe = regexp.MustCompile(`\w+:".*?"`)

func parseFieldTag(tag string) *columnInfo {
	colInfo := &columnInfo{}
	infoMap := make(map[string]string)
	fields := tagFieldRe.FindAllString(tag, -1)
	for _, f := range fields {
		l := strings.Split(f, ":")
		infoMap[l[0]] = l[1]
	}
	if sqlType, ok := infoMap["sql_type"]; ok {
		colInfo.sqlType = strings.Trim(sqlType, "\"")
	}
	if colName, ok := infoMap["column"]; ok {
		colInfo.colName = strings.Trim(colName, "\"")
	}
	if nullable, ok := infoMap["nullable"]; ok {
		b, err := strconv.ParseBool(strings.Trim(nullable, "\""))
		if err != nil {
			panic(err)
		}
		colInfo.nullable = b
	} else {
		colInfo.nullable = false
	}
	// if isPk, ok := infoMap["primary_key"]; ok {
	// 	b, err := strconv.ParseBool(strings.Trim(isPk, "\""))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	colInfo.isPk = b
	// }
	if isInc, ok := infoMap["auto_increment"]; ok {
		b, err := strconv.ParseBool(strings.Trim(isInc, "\""))
		if err != nil {
			panic(err)
		}
		if b {
			colInfo.isInc = b
			colInfo.isUni = true
		}
	}
	// if isUni, ok := infoMap["unique"]; ok {
	// 	b, err := strconv.ParseBool(strings.Trim(isUni, "\""))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	colInfo.isUni = b
	// }
	if defVal, ok := infoMap["default_value"]; ok {
		expr := strings.Trim(defVal, "\"")
		switch {
		case sqlConstRe.MatchString(expr):
			colInfo.defVal = defValRe.FindStringSubmatch(expr)[1]
		case sqlStringRe.MatchString(expr):
			colInfo.defVal = fmt.Sprintf("%q", defValRe.FindStringSubmatch(expr)[1])
		case sqlIntRe.MatchString(expr):
			d, err := strconv.ParseInt(defValRe.FindStringSubmatch(expr)[1], 10, 64)
			if err != nil {
				panic(err)
			}
			colInfo.defVal = d
		case sqlFloatRe.MatchString(expr):
			f, err := strconv.ParseFloat(defValRe.FindStringSubmatch(expr)[1], 64)
			if err != nil {
				panic(err)
			}
			colInfo.defVal = f
		case sqlBoolRe.MatchString(expr):
			b, err := strconv.ParseBool(defValRe.FindStringSubmatch(expr)[1])
			if err != nil {
				panic(err)
			}
			colInfo.defVal = b
		case sqlBinaryRe.MatchString(expr):
			colInfo.defVal = expr
		default:
			panic(fmt.Errorf("nborm.parseFieldTag() error: invalid default value (%s)", expr))
		}
	}
	if charset, ok := infoMap["charset"]; ok {
		colInfo.charset = charset
	}
	if collate, ok := infoMap["collate"]; ok {
		colInfo.collate = collate
	}
	return colInfo
}

func parseField(field *ast.Field) *columnInfo {
	var colInfo *columnInfo
	if field.Tag != nil {
		colInfo = parseFieldTag(field.Tag.Value)
	} else {
		colInfo = &columnInfo{}
	}
	name := field.Names[0].String()
	colInfo.fieldName = name
	if colInfo.colName == "" {
		colInfo.colName = toSnakeCase(name)
	}
	if colInfo.sqlType == "" {
		typ := field.Type.(*ast.SelectorExpr).Sel.Name
		if sqlType, ok := goToSQLMap[typ]; !ok {
			panic(fmt.Errorf("nborm.parseField() error: unsupported field type (%s)", typ))
		} else {
			colInfo.sqlType = sqlType
		}
	}
	return colInfo
}

type relInfo struct {
	srcCol      string
	midDB       string
	midTab      string
	midLeftCol  string
	midRightCol string
	dstDB       string
	dstTab      string
	dstCol      string
}

func parseRelationInfo(decl *ast.GenDecl) (relInfo relInfo) {
	if typeSpec, ok := decl.Specs[0].(*ast.TypeSpec); ok {
		if stctType, ok := typeSpec.Type.(*ast.StructType); ok {
			modelName := typeSpec.Name.Name
			dbName, tabName := dbTabMap[modelName].db, dbTabMap[modelName].tab
			tabInfo := schemaCache.databaseMap[dbName].tableMap[tabName]
			for _, field := range stctType.Fields.List {
				if expr, ok := field.Type.(*ast.SelectorExpr); ok {
					switch expr.Sel.Name {
					case "OneToOne":
						otoInfo := parseOneToOneField(modelName, field)
						tabInfo.oneToOnes = append(tabInfo.oneToOnes, otoInfo)
					case "ForeignKey":
						fkInfo := parseForeignKeyField(modelName, field)
						tabInfo.foreignKeys = append(tabInfo.foreignKeys, fkInfo)
					case "ReverseForeignKey":
						rfkInfo := parseReverseForeignKeyField(modelName, field)
						tabInfo.reverseForeignKeys = append(tabInfo.reverseForeignKeys, rfkInfo)
					case "ManyToMany":
						mtmInfo := parseManyToManyField(modelName, field)
						tabInfo.manyToManys = append(tabInfo.manyToManys, mtmInfo)
					}
				}
			}
		}
	}
	return
}

func lookupTableInfoByModelName(modelName string) *tableInfo {
	for _, dbInfo := range schemaCache.databases {
		for _, tabInfo := range dbInfo.tables {
			if tabInfo.modelName == modelName {
				return tabInfo
			}
		}
	}
	return nil
}

func lookupTableInfoByTableName(tableName string) *tableInfo {
	for _, dbInfo := range schemaCache.databases {
		for _, tabInfo := range dbInfo.tables {
			if tabInfo.tab == tableName {
				return tabInfo
			}
		}
	}
	return nil
}

var onRe = regexp.MustCompile(`on:"(\w+)\s*=\s*(\w+)\.(\w+)"`)
var byRe = regexp.MustCompile(`by:"(\w+)"`)

func parseOn(tag string) (srcFieldName, dstModelName, dstFieldName string) {
	group := onRe.FindStringSubmatch(tag)
	if len(group) < 4 {
		panic(fmt.Errorf("nborm.parseOn() error: no valid ON condition (%s)", tag))
	}
	return group[1], group[2], group[3]
}

func parseBy(tag string) (midTabName string, ok bool) {
	group := byRe.FindStringSubmatch(tag)
	if len(group) < 2 {
		return "", false
	}
	return group[1], true
}

func findTableAndColumnInfo(modelName, fieldName string) (*tableInfo, *columnInfo) {
	tabInfo := lookupTableInfoByModelName(modelName)
	if tabInfo == nil {
		panic(fmt.Errorf("nborm.findTableAndColumnInfo() error: cannot find table information (%s)", modelName))
	}
	colInfo := tabInfo.lookupColInfoByFieldName(fieldName)
	if colInfo == nil {
		panic(fmt.Errorf("nborm.findTableAndColumnInfo() error: cannot find column information (%s.%s)", modelName, fieldName))
	}
	return tabInfo, colInfo
}

func parseOneToOneField(srcModelName string, field *ast.Field) *oneToOneInfo {
	oto := &oneToOneInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseOneToOneField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	_, srcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	dstTabInfo, dstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	oto.srcCol, oto.dstDB, oto.dstTab, oto.dstCol = srcColInfo.colName, dstTabInfo.db, dstTabInfo.tab, dstColInfo.colName
	return oto
}

func parseForeignKeyField(srcModelName string, field *ast.Field) *foreignKeyInfo {
	fk := &foreignKeyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseForeignKeyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	_, srcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	dstTabInfo, dstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	fk.srcCol, fk.dstDB, fk.dstTab, fk.dstCol = srcColInfo.colName, dstTabInfo.db, dstTabInfo.tab, dstColInfo.colName
	return fk
}

func parseReverseForeignKeyField(srcModelName string, field *ast.Field) *reverseForeignKeyInfo {
	rfk := &reverseForeignKeyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseReverseForeignKeyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	_, srcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	dstTabInfo, dstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	rfk.srcCol, rfk.dstDB, rfk.dstTab, rfk.dstCol = srcColInfo.colName, dstTabInfo.db, dstTabInfo.tab, dstColInfo.colName
	return rfk
}

func parseManyToManyField(srcModelName string, field *ast.Field) *manyToManyInfo {
	mtm := &manyToManyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseManyToManyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	srcTabInfo, srcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	dstTabInfo, dstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	midTabName, ok := parseBy(field.Tag.Value)
	var midTabName1, midTabName2 string
	var midTabInfo *tableInfo
	if !ok {
		midTabName1 = srcTabInfo.tab + "__" + dstTabInfo.tab
		midTabName2 = dstTabInfo.tab + "__" + srcTabInfo.tab
		midTabInfo1, midTabInfo2 := lookupTableInfoByTableName(midTabName1), lookupTableInfoByTableName(midTabName2)
		switch {
		case midTabInfo1 != nil:
			midTabInfo = midTabInfo1
		case midTabInfo2 != nil:
			midTabInfo = midTabInfo2
		default:
			IDCol := &columnInfo{colName: "id", isInc: true, isPk: true, sqlType: "INT", fieldName: "ID"}
			leftCol := &columnInfo{
				charset:   srcColInfo.charset,
				collate:   srcColInfo.collate,
				colName:   srcTabInfo.tab + "__" + srcColInfo.colName,
				sqlType:   srcColInfo.sqlType,
				fieldName: srcTabInfo.modelName + "__" + srcColInfo.fieldName,
			}
			rightCol := &columnInfo{
				charset:   dstColInfo.charset,
				collate:   dstColInfo.collate,
				colName:   dstTabInfo.tab + "__" + dstColInfo.colName,
				sqlType:   dstColInfo.sqlType,
				fieldName: dstTabInfo.modelName + "__" + dstColInfo.fieldName,
			}
			midTabInfo = &tableInfo{
				db:        srcTabInfo.db,
				tab:       midTabName1,
				modelName: srcModelName + "__" + dstModelName,
				columnMap: map[string]*columnInfo{"id": IDCol, leftCol.colName: leftCol, rightCol.colName: rightCol},
				columns:   []*columnInfo{IDCol, leftCol, rightCol},
				pk:        primaryKey{IDCol},
				unis:      uniqueKeys{[]*columnInfo{leftCol, rightCol}},
				foreignKeys: []*foreignKeyInfo{
					&foreignKeyInfo{
						srcCol: leftCol.colName,
						dstDB:  srcTabInfo.db,
						dstTab: srcTabInfo.tab,
						dstCol: srcColInfo.colName,
					},
					&foreignKeyInfo{
						srcCol: rightCol.colName,
						dstDB:  dstTabInfo.db,
						dstTab: dstTabInfo.tab,
						dstCol: dstColInfo.colName,
					},
				},
			}
			schemaCache.databaseMap[midTabInfo.db].tableMap[midTabInfo.tab] = midTabInfo
			schemaCache.databaseMap[midTabInfo.db].tables = append(schemaCache.databaseMap[midTabInfo.db].tables, midTabInfo)
		}
	} else {
		midTabInfo = lookupTableInfoByTableName(midTabName)
		if midTabInfo == nil {
			panic(fmt.Errorf("nborm.parseManyToMany() error: cannot find table (%s)", midTabName))
		}
	}
	mtm.srcCol = srcColInfo.colName
	mtm.midDB = midTabInfo.db
	mtm.midTab = midTabInfo.tab
	mtm.midLeftCol = srcTabInfo.tab + "__" + srcColInfo.colName
	mtm.midRightCol = dstTabInfo.tab + "__" + dstColInfo.colName
	mtm.dstDB = dstTabInfo.db
	mtm.dstTab = dstTabInfo.tab
	mtm.dstCol = dstColInfo.colName
	return mtm
}

type dbAndTab struct {
	db  string
	tab string
}

var dbTabMap = map[string]*dbAndTab{}

func initPrimaryKey(tabInfo *tableInfo) {
	for _, name := range tabInfo.pkNames {
		var pkInfo *columnInfo
		for _, col := range tabInfo.columns {
			if col.fieldName == name {
				pkInfo = col
				break
			}
		}
		if pkInfo == nil {
			panic(fmt.Errorf("nborm.initPrimaryKey() error: primary key field not exists (%s.%s)", tabInfo.modelName, name))
		}
		tabInfo.pk = append(tabInfo.pk, pkInfo)
	}
}

func initUniqueKeys(tabInfo *tableInfo) {
	for _, key := range tabInfo.uniNames {
		l := make([]*columnInfo, 0, 4)
		for _, name := range key {
			var uni *columnInfo
			for _, col := range tabInfo.columns {
				if col.fieldName == name {
					uni = col
					break
				}
			}
			if uni == nil {
				panic(fmt.Errorf("nborm.initUniqueKey() error: unique key field not exists (%s.%s)", tabInfo.modelName, name))
			}
			l = append(l, uni)
		}
		tabInfo.unis = append(tabInfo.unis, l)
	}
}

func initKeys(tabInfo *tableInfo) {
	tabInfo.keys = make(keys, len(tabInfo.keyNames))
	for i, key := range tabInfo.keyNames {
		l := make([]*columnInfo, len(key))
		for j, fieldName := range key {
			colInfo := tabInfo.lookupColInfoByFieldName(fieldName)
			if colInfo == nil {
				panic(fmt.Errorf("nborm.initKey() error: key column not exists (%s.%s.%s)", tabInfo.db, tabInfo.modelName, fieldName))
			}
			l[j] = colInfo
		}
		tabInfo.keys[i] = l
	}
}

func parseModel(decl *ast.GenDecl) {
	if typeSpec, ok := decl.Specs[0].(*ast.TypeSpec); ok {
		if stctType, ok := typeSpec.Type.(*ast.StructType); ok {
			modelName := typeSpec.Name.Name
			dbName, tabName := dbTabMap[modelName].db, dbTabMap[modelName].tab
			tabInfo := schemaCache.databaseMap[dbName].tableMap[tabName]
			tabInfo.db, tabInfo.tab, tabInfo.modelName = dbName, tabName, modelName
			tabInfo.keyNames = parseKeys(commentMap[modelName])
			tabInfo.pkNames = parsePrimaryKey(commentMap[modelName])
			tabInfo.uniNames = parseUniqueKeys(commentMap[modelName])
			for _, field := range stctType.Fields.List {
				if expr, ok := field.Type.(*ast.SelectorExpr); ok {
					switch expr.Sel.Name {
					case "StringField", "IntField", "FloatField", "BoolField", "BinaryField", "DateField", "DatetimeField":
						colInfo := parseField(field)
						tabInfo.columns = append(tabInfo.columns, colInfo)
						tabInfo.columnMap[colInfo.colName] = colInfo
						if colInfo.isInc {
							tabInfo.inc = colInfo
						}
					}
				}
			}
			initPrimaryKey(tabInfo)
			initUniqueKeys(tabInfo)
			initKeys(tabInfo)
		}
	}
}

func initSchema() {
	for m, info := range dbTabMap {
		if info.db == "" || info.tab == "" {
			panic(fmt.Errorf("nborm.initSchema() error: no database name or no table name (model %s)", m))
		}
		tabInfo := schemaCache.getOrCreate(info.db).getOrCreate(info.tab)
		tabInfo.modelName = m
		tabInfo.db = info.db
		tabInfo.tab = info.tab
	}
}

func parseDB(filename string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	initSchema()
	ast.Inspect(f, func(node ast.Node) bool {
		if genDecl, ok := node.(*ast.GenDecl); ok {
			parseModel(genDecl)
		}
		return true
	})
	ast.Inspect(f, func(node ast.Node) bool {
		if genDecl, ok := node.(*ast.GenDecl); ok {
			parseRelationInfo(genDecl)
		}
		return true
	})
	for dname, db := range schemaCache.databaseMap {
		for tname, tab := range db.tableMap {
			for _, oto := range tab.oneToOnes {
				schemaCache.databaseMap[oto.dstDB].tableMap[oto.dstTab].columnMap[oto.dstCol].isUni = true
				tab.columnMap[oto.srcCol].isUni = true
			}
			for _, fk := range tab.foreignKeys {
				schemaCache.databaseMap[fk.dstDB].tableMap[fk.dstTab].columnMap[fk.dstCol].isUni = true
			}
			for _, rfk := range tab.reverseForeignKeys {
				tab.columnMap[rfk.srcCol].isUni = true
			}
			for _, mtm := range tab.manyToManys {
				midTab := schemaCache.databaseMap[mtm.midDB].tableMap[mtm.midTab]
				if _, ok := midTab.columnMap[mtm.midLeftCol]; !ok {
					colInfo := *tab.columnMap[mtm.srcCol]
					colInfo.colName, colInfo.isInc, colInfo.isPk, colInfo.isUni, colInfo.nullable = mtm.midLeftCol, false, false, false, false
					midTab.columnMap[mtm.midLeftCol] = &colInfo
					midTab.columns = append(midTab.columns, &colInfo)
				}
				if _, ok := midTab.columnMap[mtm.midRightCol]; !ok {
					colInfo := *schemaCache.databaseMap[mtm.dstDB].tableMap[mtm.dstTab].columnMap[mtm.dstCol]
					colInfo.colName, colInfo.isInc, colInfo.isPk, colInfo.isUni, colInfo.nullable = mtm.midRightCol, false, false, false, false
					midTab.columnMap[mtm.midRightCol] = &colInfo
					midTab.columns = append(midTab.columns, &colInfo)
				}
				if len(midTab.foreignKeys) == 0 {
					leftForeignKey := &foreignKeyInfo{srcCol: mtm.midLeftCol, dstDB: dname, dstTab: tname, dstCol: mtm.srcCol}
					rightForeignKey := &foreignKeyInfo{srcCol: mtm.midRightCol, dstDB: mtm.dstDB, dstTab: mtm.dstTab, dstCol: mtm.dstCol}
					midTab.foreignKeys = append(midTab.foreignKeys, leftForeignKey, rightForeignKey)
				}
			}
		}
	}
	return nil
}

func create() error {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/information_schema", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		return err
	}
	for dname, db := range schemaCache.databaseMap {
		if _, err := conn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", wrap(dname))); err != nil {
			return err
		}
		if _, err := conn.Exec(fmt.Sprintf("USE %s", wrap(dname))); err != nil {
			return err
		}
		for tname, tab := range db.tableMap {
			cols := make([]string, len(tab.columns))
			// pks := make([]string, len(tab.pk))
			// unis := make([]string, len(tab.unis))
			for i, col := range tab.columns {
				l := make([]string, 0, 8)
				l = append(l, wrap(col.colName))
				l = append(l, col.sqlType)
				if col.charset != "" {
					l = append(l, fmt.Sprintf("CHARACTER SET %s", col.charset))
				}
				if col.collate != "" {
					l = append(l, fmt.Sprintf("COLLATE %s", col.collate))
				}
				if !col.nullable {
					l = append(l, "NOT NULL")
				}
				if col.isInc {
					l = append(l, "AUTO_INCREMENT")
				}
				if col.defVal != nil {
					l = append(l, fmt.Sprintf("DEFAULT %v", col.defVal))
				}
				cols[i] = strings.Join(l, " ")
			}
			// for i, pk := range tab.pks {
			// 	pks[i] = wrap(pk.colName)
			// }
			// for i, uni := range tab.unis {
			// 	l := make([]string, len(uni))
			// 	for i, col := range uni {
			// 		l[i] = wrap(col.colName)
			// 	}
			// 	unis[i] = fmt.Sprintf("UNIQUE KEY (%s)", strings.Join(l, ", "))
			// }
			// var uniClause string
			// if len(unis) > 0 {
			// 	uniClause = ", " + strings.Join(unis, ", ")
			// }
			// stmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s, PRIMARY KEY(%s) %s) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin`,
			// 	wrap(tname), strings.Join(cols, ", "), strings.Join(pks, ", "), uniClause)
			stmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s%s%s%s) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin`,
				wrap(tname), strings.Join(cols, ", "), tab.pk.genCreateClause(), tab.unis.genCreateClause(), tab.keys.genCreateClause())
			fmt.Println(stmt)
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			if _, err := conn.Exec(stmt); err != nil {
				return err
			}
		}
		for tname, tab := range db.tableMap {
			for _, fk := range tab.foreignKeys {
				stmt := fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s.%s (%s) ON DELETE CASCADE", wrap(tname),
					// wrap(fmt.Sprintf("%s_%s_%s__%s_%s_%s", dname, tname, fk.srcCol, fk.dstDB, fk.dstTab, fk.dstCol)), wrap(fk.srcCol), wrap(fk.dstDB),
					wrap(fmt.Sprintf("%s_%s__%s_%s", tname, fk.srcCol, fk.dstTab, fk.dstCol)), wrap(fk.srcCol), wrap(fk.dstDB),
					wrap(fk.dstTab), wrap(fk.dstCol))
				fmt.Println(stmt)
				fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
				if _, err := conn.Exec(stmt); err != nil {
					if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1826 {
						fmt.Printf("warning: %v\n", e)
						fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
						continue
					}
					return err
				}
			}
			for _, mtm := range tab.manyToManys {
				stmt := fmt.Sprintf("ALTER TABLE %s.%s ADD CONSTRAINT UNIQUE KEY(%s, %s)", wrap(mtm.midDB), wrap(mtm.midTab), wrap(mtm.midLeftCol),
					wrap(mtm.midRightCol))
				fmt.Println(stmt)
				fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
				if _, err := conn.Exec(stmt); err != nil {
					return err
				}
			}
			for _, key := range tab.keys {
				stmt := fmt.Sprintf("ALTER TABLE %s.%s ADD KEY %s", wrap(dname), wrap(tname), key)
				fmt.Println(stmt)
				fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
				if _, err := conn.Exec(stmt); err != nil {
					if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1061 {
						fmt.Printf("warning: %v\n", e)
						fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
						continue
					}
					return err
				}
			}
		}
	}
	return nil
}

func ParseAndCreate(filename string) error {
	if err := parseDB(filename); err != nil {
		return err
	}
	if err := create(); err != nil {
		return err
	}
	return nil
}

var pkg string
var commentMap = make(map[string]string)

//ParseComment parse package comments
func ParseComment(path string) error {
	pfset := token.NewFileSet()
	d, err := parser.ParseDir(pfset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	for _, f := range d {
		if pkg == "" {
			pkg = f.Name
		}
		p := doc.New(f, path, 0)
		for _, t := range p.Types {
			commentMap[t.Name] = t.Doc
		}
	}
	initDBTabMap()
	return nil
}

func initDBTabMap() {
	for modelName, comment := range commentMap {
		db, tab := parseDBName(comment), parseTabName(comment)
		if db == "" || tab == "" {
			panic(fmt.Errorf("nborm.initDBTabMap() error: no database name or table name (%s:%s.%s, %s)", modelName, db, tab, comment))
		}
		dbTabMap[modelName] = &dbAndTab{db, tab}
	}
}

var dbRe = regexp.MustCompile(`DB:([\w_]+)`)
var tabRe = regexp.MustCompile(`Tab:([\w_]+)`)
var primaryKeyRe = regexp.MustCompile(`PrimaryKey:([\w_]+)`)
var keysRe = regexp.MustCompile(`Index:([\w_]+)`)
var uniqueKeyRe = regexp.MustCompile(`UniqueKey:([\w_]+)`)

func parseDBName(comment string) string {
	group := dbRe.FindStringSubmatch(comment)
	if len(group) < 2 {
		panic(fmt.Errorf("nborm.parseDBName() error: database name not exists(%s)", comment))
	}
	return group[1]
}

func parseTabName(comment string) string {
	group := tabRe.FindStringSubmatch(comment)
	if len(group) < 2 {
		panic(fmt.Errorf("nborm.parseTabName() error: the table name not exists(%s)", comment))
	}
	return group[1]
}

func parsePrimaryKey(comment string) []string {
	group := primaryKeyRe.FindStringSubmatch(comment)
	if len(group) < 2 {
		panic(fmt.Errorf("nborm.parsePrimaryKey() error: primary key not exist (%s)", comment))
	}
	l := strings.Split(group[1], ",")
	for i, name := range l {
		l[i] = strings.Trim(name, " ")
	}
	return l
}

func parseKeys(comment string) [][]string {
	l := make([][]string, 0, 8)
	keys := keysRe.FindAllStringSubmatch(comment, -1)
	for _, key := range keys {
		keyStr := strings.Replace(key[1], " ", "", -1)
		cols := strings.Split(keyStr, ",")
		l = append(l, cols)
	}
	return l
}

func parseUniqueKeys(comment string) [][]string {
	l := make([][]string, 0, 8)
	group := uniqueKeyRe.FindAllStringSubmatch(comment, -1)
	for _, uni := range group {
		tl := strings.Split(uni[1], ",")
		for i, uni := range tl {
			tl[i] = strings.Trim(uni, " ")
		}
		l = append(l, tl)
	}
	return l
}

type MethodInfo struct {
	ModelName string
	Pk        string
	Unis      string
	Keys      string
}

func CreateMethodFile(path string) {
	f, err := os.OpenFile(filepath.Join(path, "modelMethods.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateMethodFile() error: %s", err.Error()))
	}
	defer f.Close()
	infos := make(map[string]map[string]MethodInfo)
	for dbName, db := range schemaCache.databaseMap {
		infos[dbName] = make(map[string]MethodInfo)
		for tabName, tab := range db.tableMap {
			infos[dbName][tabName] = MethodInfo{tab.modelName, tab.pk.methodString(), tab.unis.methodString(), tab.keys.methodString()}
		}
	}
	content, err := nbfmt.Fmt(methodTemplate, map[string]interface{}{"package": pkg, "infos": infos})
	if err != nil {
		panic(fmt.Errorf("nborm.CreateMethodFile error(): %s", err.Error()))
	}
	f.WriteString(content)
	if err := exec.Command("go", "fmt", filepath.Join(path, "modelMethods.go")).Run(); err != nil {
		panic(fmt.Errorf("nborm.CreateMethodFile() error: %s", err.Error()))
	}
}

func CreateSchemaJSON(path string) {
	f, err := os.OpenFile(filepath.Join(path, "schema.json"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
	defer f.Close()
	b, err := json.Marshal(schemaCache)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
	_, err = f.Write(b)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
}
