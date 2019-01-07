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

var DefValRe = regexp.MustCompile(`'(.*)'`)
var tagFieldRe = regexp.MustCompile(`\w+:".*?"`)

func parseFieldTag(tag string) *ColumnInfo {
	colInfo := &ColumnInfo{}
	infoMap := make(map[string]string)
	fields := tagFieldRe.FindAllString(tag, -1)
	for _, f := range fields {
		l := strings.Split(f, ":")
		infoMap[l[0]] = l[1]
	}
	if SqlType, ok := infoMap["sql_type"]; ok {
		colInfo.SqlType = strings.Trim(SqlType, "\"")
	}
	if ColName, ok := infoMap["column"]; ok {
		colInfo.ColName = strings.Trim(ColName, "\"")
	}
	if Nullable, ok := infoMap["Nullable"]; ok {
		b, err := strconv.ParseBool(strings.Trim(Nullable, "\""))
		if err != nil {
			panic(err)
		}
		colInfo.Nullable = b
	} else {
		colInfo.Nullable = false
	}
	// if IsPk, ok := infoMap["primary_key"]; ok {
	// 	b, err := strconv.ParseBool(strings.Trim(IsPk, "\""))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	colInfo.IsPk = b
	// }
	if IsInc, ok := infoMap["auto_increment"]; ok {
		b, err := strconv.ParseBool(strings.Trim(IsInc, "\""))
		if err != nil {
			panic(err)
		}
		if b {
			colInfo.IsInc = b
			colInfo.IsUni = true
		}
	}
	// if IsUni, ok := infoMap["unique"]; ok {
	// 	b, err := strconv.ParseBool(strings.Trim(IsUni, "\""))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	colInfo.IsUni = b
	// }
	if DefVal, ok := infoMap["default_value"]; ok {
		expr := strings.Trim(DefVal, "\"")
		switch {
		case sqlConstRe.MatchString(expr):
			colInfo.DefVal = DefValRe.FindStringSubmatch(expr)[1]
		case sqlStringRe.MatchString(expr):
			colInfo.DefVal = fmt.Sprintf("%q", DefValRe.FindStringSubmatch(expr)[1])
		case sqlIntRe.MatchString(expr):
			d, err := strconv.ParseInt(DefValRe.FindStringSubmatch(expr)[1], 10, 64)
			if err != nil {
				panic(err)
			}
			colInfo.DefVal = d
		case sqlFloatRe.MatchString(expr):
			f, err := strconv.ParseFloat(DefValRe.FindStringSubmatch(expr)[1], 64)
			if err != nil {
				panic(err)
			}
			colInfo.DefVal = f
		case sqlBoolRe.MatchString(expr):
			b, err := strconv.ParseBool(DefValRe.FindStringSubmatch(expr)[1])
			if err != nil {
				panic(err)
			}
			colInfo.DefVal = b
		case sqlBinaryRe.MatchString(expr):
			colInfo.DefVal = expr
		default:
			panic(fmt.Errorf("nborm.parseFieldTag() error: invalid default value (%s)", expr))
		}
	}
	if Charset, ok := infoMap["Charset"]; ok {
		colInfo.Charset = Charset
	}
	if Collate, ok := infoMap["Collate"]; ok {
		colInfo.Collate = Collate
	}
	return colInfo
}

func parseField(field *ast.Field) *ColumnInfo {
	var colInfo *ColumnInfo
	if field.Tag != nil {
		colInfo = parseFieldTag(field.Tag.Value)
	} else {
		colInfo = &ColumnInfo{}
	}
	name := field.Names[0].String()
	colInfo.FieldName = name
	if colInfo.ColName == "" {
		colInfo.ColName = toSnakeCase(name)
	}
	if colInfo.SqlType == "" {
		typ := field.Type.(*ast.SelectorExpr).Sel.Name
		if SqlType, ok := goToSQLMap[typ]; !ok {
			panic(fmt.Errorf("nborm.parseField() error: unsupported field type (%s)", typ))
		} else {
			colInfo.SqlType = SqlType
		}
	}
	typ := field.Type.(*ast.SelectorExpr).Sel.Name
	ormType, ok := OrmTypeMap[typ]
	if !ok {
		panic(fmt.Errorf("nborm.parseField() error: unsupported field type (%s)", typ))
	}
	colInfo.OrmType = ormType
	return colInfo
}

type relInfo struct {
	SrcCol      string
	MidDB       string
	MidTab      string
	MidLeftCol  string
	MidRightCol string
	DstDB       string
	DstTab      string
	DstCol      string
}

func parseRelationInfo(decl *ast.GenDecl) (relInfo relInfo) {
	if typeSpec, ok := decl.Specs[0].(*ast.TypeSpec); ok {
		if stctType, ok := typeSpec.Type.(*ast.StructType); ok {
			ModelName := typeSpec.Name.Name
			dbName, tabName := dbTabMap[ModelName].db, dbTabMap[ModelName].tab
			tabInfo := SchemaCache.DatabaseMap[dbName].TableMap[tabName]
			for _, field := range stctType.Fields.List {
				if expr, ok := field.Type.(*ast.SelectorExpr); ok {
					switch expr.Sel.Name {
					case "OneToOne":
						otoInfo := parseOneToOneField(ModelName, field)
						tabInfo.OneToOnes = append(tabInfo.OneToOnes, otoInfo)
					case "ForeignKey":
						fkInfo := parseForeignKeyField(ModelName, field)
						tabInfo.ForeignKeys = append(tabInfo.ForeignKeys, fkInfo)
					case "ReverseForeignKey":
						rfkInfo := parseReverseForeignKeyField(ModelName, field)
						tabInfo.ReverseForeignKeys = append(tabInfo.ReverseForeignKeys, rfkInfo)
					case "ManyToMany":
						mtmInfo := parseManyToManyField(ModelName, field)
						tabInfo.ManyToManys = append(tabInfo.ManyToManys, mtmInfo)
					}
				}
			}
		}
	}
	return
}

func lookupTableInfoByModelName(ModelName string) *TableInfo {
	for _, dbInfo := range SchemaCache.Databases {
		for _, tabInfo := range dbInfo.Tables {
			if tabInfo.ModelName == ModelName {
				return tabInfo
			}
		}
	}
	return nil
}

func lookupTableInfoByTableName(tableName string) *TableInfo {
	for _, dbInfo := range SchemaCache.Databases {
		for _, tabInfo := range dbInfo.Tables {
			if tabInfo.Tab == tableName {
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

func parseBy(tag string) (MidTabName string, ok bool) {
	group := byRe.FindStringSubmatch(tag)
	if len(group) < 2 {
		return "", false
	}
	return group[1], true
}

func findTableAndColumnInfo(ModelName, FieldName string) (*TableInfo, *ColumnInfo) {
	tabInfo := lookupTableInfoByModelName(ModelName)
	if tabInfo == nil {
		panic(fmt.Errorf("nborm.findTableAndColumnInfo() error: cannot find table information (%s)", ModelName))
	}
	colInfo := tabInfo.lookupColInfoByFieldName(FieldName)
	if colInfo == nil {
		panic(fmt.Errorf("nborm.findTableAndColumnInfo() error: cannot find column information (%s.%s)", ModelName, FieldName))
	}
	return tabInfo, colInfo
}

func parseOneToOneField(srcModelName string, field *ast.Field) *OneToOneInfo {
	oto := &OneToOneInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseOneToOneField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	fieldName := field.Names[0].String()
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	_, SrcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	DstTabInfo, DstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	oto.SrcCol, oto.DstDB, oto.DstTab, oto.DstCol, oto.FieldName = SrcColInfo.ColName, DstTabInfo.DB, DstTabInfo.Tab, DstColInfo.ColName, fieldName
	return oto
}

func parseForeignKeyField(srcModelName string, field *ast.Field) *ForeignKeyInfo {
	fk := &ForeignKeyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseForeignKeyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	fieldName := field.Names[0].String()
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	_, SrcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	DstTabInfo, DstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	fk.SrcCol, fk.DstDB, fk.DstTab, fk.DstCol, fk.FieldName = SrcColInfo.ColName, DstTabInfo.DB, DstTabInfo.Tab, DstColInfo.ColName, fieldName
	return fk
}

func parseReverseForeignKeyField(srcModelName string, field *ast.Field) *ReverseForeignKeyInfo {
	rfk := &ReverseForeignKeyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseReverseForeignKeyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	fieldName := field.Names[0].String()
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	_, SrcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	DstTabInfo, DstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	rfk.SrcCol, rfk.DstDB, rfk.DstTab, rfk.DstCol, rfk.FieldName = SrcColInfo.ColName, DstTabInfo.DB, DstTabInfo.Tab, DstColInfo.ColName, fieldName
	return rfk
}

func parseManyToManyField(srcModelName string, field *ast.Field) *ManyToManyInfo {
	mtm := &ManyToManyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseManyToManyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	fieldName := field.Names[0].String()
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	srcTabInfo, SrcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	DstTabInfo, DstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	MidTabName, ok := parseBy(field.Tag.Value)
	var MidTabName1, MidTabName2 string
	var MidTabInfo *TableInfo
	if !ok {
		MidTabName1 = srcTabInfo.Tab + "__" + DstTabInfo.Tab
		MidTabName2 = DstTabInfo.Tab + "__" + srcTabInfo.Tab
		MidTabInfo1, MidTabInfo2 := lookupTableInfoByTableName(MidTabName1), lookupTableInfoByTableName(MidTabName2)
		switch {
		case MidTabInfo1 != nil:
			MidTabInfo = MidTabInfo1
		case MidTabInfo2 != nil:
			MidTabInfo = MidTabInfo2
		default:
			IDCol := &ColumnInfo{ColName: "id", IsInc: true, IsPk: true, SqlType: "INT", FieldName: "ID"}
			leftCol := &ColumnInfo{
				Charset:   SrcColInfo.Charset,
				Collate:   SrcColInfo.Collate,
				ColName:   srcTabInfo.Tab + "__" + SrcColInfo.ColName,
				SqlType:   SrcColInfo.SqlType,
				FieldName: srcTabInfo.ModelName + "__" + SrcColInfo.FieldName,
			}
			rightCol := &ColumnInfo{
				Charset:   DstColInfo.Charset,
				Collate:   DstColInfo.Collate,
				ColName:   DstTabInfo.Tab + "__" + DstColInfo.ColName,
				SqlType:   DstColInfo.SqlType,
				FieldName: DstTabInfo.ModelName + "__" + DstColInfo.FieldName,
			}
			MidTabInfo = &TableInfo{
				DB:        srcTabInfo.DB,
				Tab:       MidTabName1,
				ModelName: srcModelName + "__" + dstModelName,
				ColumnMap: map[string]*ColumnInfo{"id": IDCol, leftCol.ColName: leftCol, rightCol.ColName: rightCol},
				Columns:   []*ColumnInfo{IDCol, leftCol, rightCol},
				Pk:        PrimaryKey{IDCol},
				Unis:      UniqueKeys{[]*ColumnInfo{leftCol, rightCol}},
				ForeignKeys: []*ForeignKeyInfo{
					&ForeignKeyInfo{
						SrcCol: leftCol.ColName,
						DstDB:  srcTabInfo.DB,
						DstTab: srcTabInfo.Tab,
						DstCol: SrcColInfo.ColName,
					},
					&ForeignKeyInfo{
						SrcCol: rightCol.ColName,
						DstDB:  DstTabInfo.DB,
						DstTab: DstTabInfo.Tab,
						DstCol: DstColInfo.ColName,
					},
				},
				IsNewMiddleTable: true,
			}
			SchemaCache.DatabaseMap[MidTabInfo.DB].TableMap[MidTabInfo.Tab] = MidTabInfo
			SchemaCache.DatabaseMap[MidTabInfo.DB].Tables = append(SchemaCache.DatabaseMap[MidTabInfo.DB].Tables, MidTabInfo)
		}
	} else {
		MidTabInfo = lookupTableInfoByTableName(MidTabName)
		if MidTabInfo == nil {
			panic(fmt.Errorf("nborm.parseManyToMany() error: cannot find table (%s)", MidTabName))
		}
	}
	mtm.SrcCol = SrcColInfo.ColName
	mtm.MidDB = MidTabInfo.DB
	mtm.MidTab = MidTabInfo.Tab
	mtm.MidLeftCol = srcTabInfo.Tab + "__" + SrcColInfo.ColName
	mtm.MidRightCol = DstTabInfo.Tab + "__" + DstColInfo.ColName
	mtm.DstDB = DstTabInfo.DB
	mtm.DstTab = DstTabInfo.Tab
	mtm.DstCol = DstColInfo.ColName
	mtm.FieldName = fieldName
	return mtm
}

type dbAndTab struct {
	db  string
	tab string
}

var dbTabMap = map[string]*dbAndTab{}

func initPrimaryKey(tabInfo *TableInfo) {
	for _, name := range tabInfo.PkNames {
		var PkInfo *ColumnInfo
		for _, col := range tabInfo.Columns {
			if col.FieldName == name {
				PkInfo = col
				break
			}
		}
		if PkInfo == nil {
			panic(fmt.Errorf("nborm.initPrimaryKey() error: primary key field not exists (%s.%s)", tabInfo.ModelName, name))
		}
		tabInfo.Pk = append(tabInfo.Pk, PkInfo)
	}
}

func initUniqueKeys(tabInfo *TableInfo) {
	for _, key := range tabInfo.UniNames {
		l := make([]*ColumnInfo, 0, 4)
		for _, name := range key {
			var uni *ColumnInfo
			for _, col := range tabInfo.Columns {
				if col.FieldName == name {
					uni = col
					break
				}
			}
			if uni == nil {
				panic(fmt.Errorf("nborm.initUniqueKey() error: unique key field not exists (%s.%s)", tabInfo.ModelName, name))
			}
			l = append(l, uni)
		}
		tabInfo.Unis = append(tabInfo.Unis, l)
	}
}

func initKeys(tabInfo *TableInfo) {
	tabInfo.Keys = make(Keys, len(tabInfo.KeyNames))
	for i, key := range tabInfo.KeyNames {
		l := make([]*ColumnInfo, len(key))
		for j, FieldName := range key {
			colInfo := tabInfo.lookupColInfoByFieldName(FieldName)
			if colInfo == nil {
				panic(fmt.Errorf("nborm.initKey() error: key column not exists (%s.%s.%s)", tabInfo.DB, tabInfo.ModelName, FieldName))
			}
			l[j] = colInfo
		}
		tabInfo.Keys[i] = l
	}
}

func parseModel(decl *ast.GenDecl) {
	if typeSpec, ok := decl.Specs[0].(*ast.TypeSpec); ok {
		if stctType, ok := typeSpec.Type.(*ast.StructType); ok {
			ModelName := typeSpec.Name.Name
			dbName, tabName := dbTabMap[ModelName].db, dbTabMap[ModelName].tab
			tabInfo := SchemaCache.DatabaseMap[dbName].TableMap[tabName]
			tabInfo.DB, tabInfo.Tab, tabInfo.ModelName = dbName, tabName, ModelName
			tabInfo.KeyNames = parseKeys(commentMap[ModelName])
			tabInfo.PkNames = parsePrimaryKey(commentMap[ModelName])
			tabInfo.UniNames = parseUniqueKeys(commentMap[ModelName])
			for _, field := range stctType.Fields.List {
				if expr, ok := field.Type.(*ast.SelectorExpr); ok {
					switch expr.Sel.Name {
					case "StringField", "IntField", "FloatField", "BoolField", "BinaryField", "DateField", "DatetimeField":
						colInfo := parseField(field)
						tabInfo.Columns = append(tabInfo.Columns, colInfo)
						tabInfo.ColumnMap[colInfo.ColName] = colInfo
						if colInfo.IsInc {
							tabInfo.Inc = colInfo
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
		tabInfo := SchemaCache.getOrCreate(info.db).getOrCreate(info.tab)
		tabInfo.ModelName = m
		tabInfo.DB = info.db
		tabInfo.Tab = info.tab
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
	for dname, db := range SchemaCache.DatabaseMap {
		for tname, tab := range db.TableMap {
			for _, oto := range tab.OneToOnes {
				SchemaCache.DatabaseMap[oto.DstDB].TableMap[oto.DstTab].ColumnMap[oto.DstCol].IsUni = true
				tab.ColumnMap[oto.SrcCol].IsUni = true
			}
			for _, fk := range tab.ForeignKeys {
				SchemaCache.DatabaseMap[fk.DstDB].TableMap[fk.DstTab].ColumnMap[fk.DstCol].IsUni = true
			}
			for _, rfk := range tab.ReverseForeignKeys {
				tab.ColumnMap[rfk.SrcCol].IsUni = true
			}
			for _, mtm := range tab.ManyToManys {
				MidTab := SchemaCache.DatabaseMap[mtm.MidDB].TableMap[mtm.MidTab]
				if _, ok := MidTab.ColumnMap[mtm.MidLeftCol]; !ok {
					colInfo := *tab.ColumnMap[mtm.SrcCol]
					colInfo.ColName, colInfo.IsInc, colInfo.IsPk, colInfo.IsUni, colInfo.Nullable = mtm.MidLeftCol, false, false, false, false
					MidTab.ColumnMap[mtm.MidLeftCol] = &colInfo
					MidTab.Columns = append(MidTab.Columns, &colInfo)
				}
				if _, ok := MidTab.ColumnMap[mtm.MidRightCol]; !ok {
					colInfo := *SchemaCache.DatabaseMap[mtm.DstDB].TableMap[mtm.DstTab].ColumnMap[mtm.DstCol]
					colInfo.ColName, colInfo.IsInc, colInfo.IsPk, colInfo.IsUni, colInfo.Nullable = mtm.MidRightCol, false, false, false, false
					MidTab.ColumnMap[mtm.MidRightCol] = &colInfo
					MidTab.Columns = append(MidTab.Columns, &colInfo)
				}
				if len(MidTab.ForeignKeys) == 0 {
					leftForeignKey := &ForeignKeyInfo{SrcCol: mtm.MidLeftCol, DstDB: dname, DstTab: tname, DstCol: mtm.SrcCol}
					rightForeignKey := &ForeignKeyInfo{SrcCol: mtm.MidRightCol, DstDB: mtm.DstDB, DstTab: mtm.DstTab, DstCol: mtm.DstCol}
					MidTab.ForeignKeys = append(MidTab.ForeignKeys, leftForeignKey, rightForeignKey)
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
	for dname, db := range SchemaCache.DatabaseMap {
		if _, err := conn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", wrap(dname))); err != nil {
			return err
		}
		if _, err := conn.Exec(fmt.Sprintf("USE %s", wrap(dname))); err != nil {
			return err
		}
		for tname, tab := range db.TableMap {
			cols := make([]string, len(tab.Columns))
			// Pks := make([]string, len(tab.Pk))
			// Unis := make([]string, len(tab.Unis))
			for i, col := range tab.Columns {
				l := make([]string, 0, 8)
				l = append(l, wrap(col.ColName))
				l = append(l, col.SqlType)
				if col.Charset != "" {
					l = append(l, fmt.Sprintf("CHARACTER SET %s", col.Charset))
				}
				if col.Collate != "" {
					l = append(l, fmt.Sprintf("COLLATE %s", col.Collate))
				}
				if !col.Nullable {
					l = append(l, "NOT NULL")
				}
				if col.IsInc {
					l = append(l, "AUTO_INCREMENT")
				}
				if col.DefVal != nil {
					l = append(l, fmt.Sprintf("DEFAULT %v", col.DefVal))
				}
				cols[i] = strings.Join(l, " ")
			}
			// for i, Pk := range tab.Pks {
			// 	Pks[i] = wrap(Pk.ColName)
			// }
			// for i, uni := range tab.Unis {
			// 	l := make([]string, len(uni))
			// 	for i, col := range uni {
			// 		l[i] = wrap(col.ColName)
			// 	}
			// 	Unis[i] = fmt.Sprintf("UNIQUE KEY (%s)", strings.Join(l, ", "))
			// }
			// var uniClause string
			// if len(Unis) > 0 {
			// 	uniClause = ", " + strings.Join(Unis, ", ")
			// }
			// stmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s, PRIMARY KEY(%s) %s) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin`,
			// 	wrap(tname), strings.Join(cols, ", "), strings.Join(Pks, ", "), uniClause)
			stmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s%s%s%s) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin`,
				wrap(tname), strings.Join(cols, ", "), tab.Pk.genCreateClause(), tab.Unis.genCreateClause(), tab.Keys.genCreateClause())
			fmt.Println(stmt)
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			if _, err := conn.Exec(stmt); err != nil {
				return err
			}
		}
		for tname, tab := range db.TableMap {
			for _, fk := range tab.ForeignKeys {
				stmt := fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s.%s (%s) ON DELETE CASCADE", wrap(tname),
					// wrap(fmt.Sprintf("%s_%s_%s__%s_%s_%s", dname, tname, fk.SrcCol, fk.DstDB, fk.DstTab, fk.DstCol)), wrap(fk.SrcCol), wrap(fk.DstDB),
					wrap(fmt.Sprintf("%s_%s__%s_%s", tname, fk.SrcCol, fk.DstTab, fk.DstCol)), wrap(fk.SrcCol), wrap(fk.DstDB),
					wrap(fk.DstTab), wrap(fk.DstCol))
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
			for _, mtm := range tab.ManyToManys {
				stmt := fmt.Sprintf("ALTER TABLE %s.%s ADD CONSTRAINT UNIQUE KEY(%s, %s)", wrap(mtm.MidDB), wrap(mtm.MidTab), wrap(mtm.MidLeftCol),
					wrap(mtm.MidRightCol))
				fmt.Println(stmt)
				fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
				if _, err := conn.Exec(stmt); err != nil {
					return err
				}
			}
			for _, key := range tab.Keys {
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

var Pkg string
var commentMap = make(map[string]string)

//ParseComment parse package comments
func ParseComment(path string) error {
	pfset := token.NewFileSet()
	d, err := parser.ParseDir(pfset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	for _, f := range d {
		if Pkg == "" {
			Pkg = f.Name
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
	for ModelName, comment := range commentMap {
		db, tab := parseDBName(comment), parseTabName(comment)
		if db == "" || tab == "" {
			panic(fmt.Errorf("nborm.initDBTabMap() error: no database name or table name (%s:%s.%s, %s)", ModelName, db, tab, comment))
		}
		dbTabMap[ModelName] = &dbAndTab{db, tab}
	}
}

var dbRe = regexp.MustCompile(`DB:([\w_]+)`)
var tabRe = regexp.MustCompile(`Tab:([\w_]+)`)
var PrimaryKeyRe = regexp.MustCompile(`PrimaryKey:([\w_]+)`)
var KeysRe = regexp.MustCompile(`Index:([\w_]+)`)
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
	group := PrimaryKeyRe.FindStringSubmatch(comment)
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
	Keys := KeysRe.FindAllStringSubmatch(comment, -1)
	for _, key := range Keys {
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
	for dbName, db := range SchemaCache.DatabaseMap {
		infos[dbName] = make(map[string]MethodInfo)
		for tabName, tab := range db.TableMap {
			if !tab.IsNewMiddleTable {
				infos[dbName][tabName] = MethodInfo{tab.ModelName, tab.Pk.methodString(), tab.Unis.methodString(), tab.Keys.methodString()}
			}
		}
	}
	content, err := nbfmt.Fmt(methodTemplate, map[string]interface{}{"package": Pkg, "infos": infos})
	if err != nil {
		panic(fmt.Errorf("nborm.CreateMethodFile error(): %s", err.Error()))
	}
	f.WriteString(content)
	if err := exec.Command("go", "fmt", filepath.Join(path, "modelMethods.go")).Run(); err != nil {
		panic(fmt.Errorf("nborm.CreateMethodFile() error: %s", err.Error()))
	}
}

func CreateSchemaJSON(path string) {
	f, err := os.OpenFile(config.SchemaFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
	defer f.Close()
	b, err := json.MarshalIndent(SchemaCache, "", " ")
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
	_, err = f.Write(b)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
}
