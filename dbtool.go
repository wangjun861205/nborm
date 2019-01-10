package nborm

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
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

func parseField(dbName, tabName string, field *ast.Field) *ColumnInfo {
	var colInfo *ColumnInfo
	if field.Tag != nil {
		colInfo = parseFieldTag(field.Tag.Value)
	} else {
		colInfo = &ColumnInfo{}
	}
	name := field.Names[0].String()
	colInfo.FieldName = name
	colInfo.DBName = dbName
	colInfo.TabName = tabName
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
			if tabInfo.TabName == tableName {
				return tabInfo
			}
		}
	}
	return nil
}

var onRe = regexp.MustCompile(`on:"(\w+)\s*=\s*(\w+)\.(\w+)"`)
var byRe = regexp.MustCompile(`by:"(\w+)\[(\d+),(\d+)\]"`)

func parseOn(tag string) (srcFieldName, dstModelName, dstFieldName string) {
	group := onRe.FindStringSubmatch(tag)
	if len(group) < 4 {
		panic(fmt.Errorf("nborm.parseOn() error: no valid ON condition (%s)", tag))
	}
	return group[1], group[2], group[3]
}

func parseBy(tag string) (midModelName string, leftColIndex, rightColIndex int, ok bool) {
	group := byRe.FindStringSubmatch(tag)
	if len(group) < 4 {
		return
	}
	midModelName = group[1]
	leftIdx, _ := strconv.ParseInt(group[2], 10, 64)
	rightIdx, _ := strconv.ParseInt(group[3], 10, 64)
	leftColIndex, rightColIndex = int(leftIdx), int(rightIdx)
	ok = true
	return
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
	srcFieldName := field.Names[0].String()
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	srcTabInfo, srcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	dstTabInfo, dstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	oto.SrcCol = srcColInfo
	oto.DstCol = dstColInfo
	oto.FieldName = srcFieldName
	midModelName, leftIndex, rightIndex, ok := parseBy(field.Tag.Value)
	if !ok {
		midTabInfo, isCreated := getOrCreateMiddleTableInfo(srcTabInfo, dstTabInfo, srcColInfo, dstColInfo, oneToOneMiddleTable)
		if isCreated {
			SchemaCache.DatabaseMap[midTabInfo.DBName].TableMap[midTabInfo.TabName] = midTabInfo
			SchemaCache.DatabaseMap[midTabInfo.DBName].Tables = append(SchemaCache.DatabaseMap[midTabInfo.DBName].Tables, midTabInfo)
			oto.MidLeftCol = midTabInfo.Columns[1]
			oto.MidRightCol = midTabInfo.Columns[2]
		} else {
			oto.MidLeftCol = midTabInfo.Columns[2]
			oto.MidRightCol = midTabInfo.Columns[1]
		}
	} else {
		midTabInfo := lookupTableInfoByModelName(midModelName)
		if midTabInfo == nil {
			panic(fmt.Errorf("nborm.parseOneToOneField() error: middle table model not exists (%s)", midModelName))
		}
		oto.MidLeftCol = midTabInfo.Columns[leftIndex]
		oto.MidRightCol = midTabInfo.Columns[rightIndex]
	}
	var srcColUnique, dstColUnique bool
	if len(srcTabInfo.Pk) == 1 && srcTabInfo.Pk[0].ColName == oto.SrcCol.ColName {
		srcColUnique = true
	}
	if !srcColUnique {
		for _, uk := range srcTabInfo.Unis {
			if len(uk) == 1 && uk[0].ColName == oto.SrcCol.ColName {
				srcColUnique = true
			}
		}
	}
	if !srcColUnique {
		srcTabInfo.Unis = append(srcTabInfo.Unis, []*ColumnInfo{srcColInfo})
	}
	if len(dstTabInfo.Pk) == 1 && dstTabInfo.Pk[0].ColName == oto.DstCol.ColName {
		dstColUnique = true
	}
	if !dstColUnique {
		for _, uk := range dstTabInfo.Unis {
			if len(uk) == 1 && uk[0].ColName == oto.DstCol.ColName {
				dstColUnique = true
			}
		}
	}
	if !dstColUnique {
		dstTabInfo.Unis = append(dstTabInfo.Unis, []*ColumnInfo{dstColInfo})
	}
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
	_, DstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	fk.SrcCol, fk.DstCol, fk.FieldName = SrcColInfo, DstColInfo, fieldName
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
	_, DstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	rfk.SrcCol = SrcColInfo
	rfk.DstCol = DstColInfo
	rfk.FieldName = fieldName
	return rfk
}

func parseManyToManyField(srcModelName string, field *ast.Field) *ManyToManyInfo {
	mtm := &ManyToManyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseManyToManyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	fieldName := field.Names[0].String()
	srcFieldName, dstModelName, dstFieldName := parseOn(field.Tag.Value)
	srcTabInfo, srcColInfo := findTableAndColumnInfo(srcModelName, srcFieldName)
	dstTabInfo, dstColInfo := findTableAndColumnInfo(dstModelName, dstFieldName)
	mtm.SrcCol = srcColInfo
	mtm.DstCol = dstColInfo
	mtm.FieldName = fieldName
	midModelName, leftColIndex, rightColIndex, ok := parseBy(field.Tag.Value)
	if !ok {
		midTabInfo, isCreated := getOrCreateMiddleTableInfo(srcTabInfo, dstTabInfo, srcColInfo, dstColInfo, manyToManyMiddleTable)
		if isCreated {
			SchemaCache.DatabaseMap[midTabInfo.DBName].TableMap[midTabInfo.TabName] = midTabInfo
			SchemaCache.DatabaseMap[midTabInfo.DBName].Tables = append(SchemaCache.DatabaseMap[midTabInfo.DBName].Tables, midTabInfo)
			mtm.MidLeftCol = midTabInfo.Columns[1]
			mtm.MidRightCol = midTabInfo.Columns[2]

		} else {
			mtm.MidLeftCol = midTabInfo.Columns[2]
			mtm.MidRightCol = midTabInfo.Columns[1]
		}
	} else {
		midTabInfo := lookupTableInfoByModelName(midModelName)
		if midTabInfo == nil {
			panic(fmt.Errorf("nborm.parseManyToMany() error: cannot find table (%s)", midModelName))
		}
		mtm.MidLeftCol = midTabInfo.Columns[leftColIndex]
		mtm.MidRightCol = midTabInfo.Columns[rightColIndex]
	}
	var srcColUnique, dstColUnique bool
	if len(srcTabInfo.Pk) == 1 && srcTabInfo.Pk[0].ColName == mtm.SrcCol.ColName {
		srcColUnique = true
	}
	if !srcColUnique {
		for _, uk := range srcTabInfo.Unis {
			if len(uk) == 1 && uk[0].ColName == mtm.SrcCol.ColName {
				srcColUnique = true
			}
		}
	}
	if !srcColUnique {
		srcTabInfo.Unis = append(srcTabInfo.Unis, []*ColumnInfo{srcColInfo})
	}
	if len(dstTabInfo.Pk) == 1 && dstTabInfo.Pk[0].ColName == mtm.DstCol.ColName {
		dstColUnique = true
	}
	if !dstColUnique {
		for _, uk := range dstTabInfo.Unis {
			if len(uk) == 1 && uk[0].ColName == mtm.DstCol.ColName {
				dstColUnique = true
			}
		}
	}
	if !dstColUnique {
		dstTabInfo.Unis = append(dstTabInfo.Unis, []*ColumnInfo{dstColInfo})
	}
	return mtm
}

type middleTableType int

const (
	manyToManyMiddleTable middleTableType = iota
	oneToOneMiddleTable
)

func getOrCreateMiddleTableInfo(srcTabInfo, dstTabInfo *TableInfo, srcColInfo, dstColInfo *ColumnInfo, tableType middleTableType) (midTabInfo *TableInfo, isCreated bool) {
	midTabName := genMiddleTableName(srcTabInfo.ModelName, srcColInfo.FieldName, dstTabInfo.ModelName, dstColInfo.FieldName)
	switch tableType {
	case manyToManyMiddleTable:
		midTabName = "mtm" + "_" + midTabName
	case oneToOneMiddleTable:
		midTabName = "oto" + "_" + midTabName
	}
	midTabInfo = lookupTableInfoByTableName(midTabName)
	if midTabInfo == nil {
		midTabInfo = &TableInfo{}
		isCreated = true
		midTabInfo.DBName = srcTabInfo.DBName
		midTabInfo.TabName = midTabName
		idCol := &ColumnInfo{}
		idCol.DBName = midTabInfo.DBName
		idCol.TabName = midTabInfo.TabName
		idCol.ColName = "id"
		idCol.SqlType = "int"
		idCol.IsInc = true
		leftCol := &ColumnInfo{}
		leftCol.DBName = midTabInfo.DBName
		leftCol.TabName = midTabInfo.TabName
		leftCol.Charset = srcColInfo.Charset
		leftCol.ColName = srcTabInfo.TabName + "__" + srcColInfo.ColName
		leftCol.Collate = srcColInfo.Collate
		leftCol.SqlType = srcColInfo.SqlType
		rightCol := &ColumnInfo{}
		rightCol.DBName = midTabInfo.DBName
		leftCol.TabName = midTabInfo.TabName
		rightCol.Charset = dstColInfo.Charset
		rightCol.ColName = dstTabInfo.TabName + "__" + dstColInfo.ColName
		rightCol.Collate = dstColInfo.Collate
		rightCol.SqlType = dstColInfo.SqlType
		midTabInfo.ColumnMap = map[string]*ColumnInfo{
			"id":             idCol,
			leftCol.ColName:  leftCol,
			rightCol.ColName: rightCol,
		}
		midTabInfo.Columns = []*ColumnInfo{idCol, leftCol, rightCol}
		midTabInfo.Pk = PrimaryKey{idCol}
		midTabInfo.ForeignKeys = []*ForeignKeyInfo{
			&ForeignKeyInfo{
				SrcCol: leftCol,
				DstCol: srcColInfo,
			},
			&ForeignKeyInfo{
				SrcCol: rightCol,
				DstCol: dstColInfo,
			},
		}
		midTabInfo.IsNewMiddleTable = true
		switch tableType {
		case manyToManyMiddleTable:
			midTabInfo.Unis = UniqueKeys{[]*ColumnInfo{leftCol, rightCol}}
		case oneToOneMiddleTable:
			midTabInfo.Unis = UniqueKeys{[]*ColumnInfo{leftCol}, []*ColumnInfo{rightCol}}
		}
		return
	}
	return
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
				panic(fmt.Errorf("nborm.initKey() error: key column not exists (%s.%s.%s)", tabInfo.DBName, tabInfo.ModelName, FieldName))
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
			tabInfo.DBName, tabInfo.TabName, tabInfo.ModelName = dbName, tabName, ModelName
			tabInfo.KeyNames = parseKeys(commentMap[ModelName])
			tabInfo.PkNames = parsePrimaryKey(commentMap[ModelName])
			tabInfo.UniNames = parseUniqueKeys(commentMap[ModelName])
			for _, field := range stctType.Fields.List {
				if expr, ok := field.Type.(*ast.SelectorExpr); ok {
					switch expr.Sel.Name {
					case "StringField", "IntField", "FloatField", "BoolField", "BinaryField", "DateField", "DatetimeField":
						colInfo := parseField(dbName, tabName, field)
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
		tabInfo.DBName = info.db
		tabInfo.TabName = info.tab
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
			for i, col := range tab.Columns {
				l := make([]string, 0, 8)
				l = append(l, col.colName())
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
				fkName := wrap(fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s_%s__%s_%s", tname, fk.SrcCol.ColName, fk.DstCol.DBName,
					fk.DstCol.ColName)))))
				stmt := fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s.%s (%s) ON DELETE CASCADE", wrap(tname),
					fkName, wrap(fk.SrcCol.ColName), wrap(fk.DstCol.DBName),
					wrap(fk.DstCol.TabName), wrap(fk.DstCol.ColName))
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
			log.Printf("nborm.initDBTabMap() warnning: no database name or table name (%s:%s.%s, %s)\n", ModelName, db, tab, comment)
			continue
		}
		dbTabMap[ModelName] = &dbAndTab{db, tab}
	}
}

var dbRe = regexp.MustCompile(`DB:(.*?)\n`)
var tabRe = regexp.MustCompile(`Tab:(.*?)\n`)
var PrimaryKeyRe = regexp.MustCompile(`PrimaryKey:(.*?)\n`)
var KeysRe = regexp.MustCompile(`Index:(.*?)\n`)
var uniqueKeyRe = regexp.MustCompile(`UniqueKey:(.*?)\n`)

func parseDBName(comment string) string {
	group := dbRe.FindStringSubmatch(comment)
	if len(group) < 2 {
		// panic(fmt.Errorf("nborm.parseDBName() error: database name not exists(%s)", comment))
		log.Printf("nborm.parseDBName() warnning: database name not exists(%s)", comment)
		return ""
	}
	return group[1]
}

func parseTabName(comment string) string {
	group := tabRe.FindStringSubmatch(comment)
	if len(group) < 2 {
		// panic(fmt.Errorf("nborm.parseTabName() error: the table name not exists(%s)", comment))
		log.Printf("nborm.parseTabName() error: the table name not exists(%s)", comment)
		return ""
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
	err = f.Truncate(0)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
	b, err := json.MarshalIndent(SchemaCache, "", " ")
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
	_, err = f.Write(b)
	if err != nil {
		panic(fmt.Errorf("nborm.CreateSchemaJSON() error: %v", err))
	}
}
