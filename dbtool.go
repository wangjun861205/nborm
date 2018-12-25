package nborm

import (
	"database/sql"
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
	"StringField":   "VARCHAR(128)",
	"IntField":      "INT",
	"FloatField":    "DECIMAL(65, 2)",
	"BoolField":     "BOOL",
	"BinaryField":   "BLOB",
	"DateField":     "DATE",
	"DatetimeField": "DATETIME",
}

var defValRe = regexp.MustCompile(`'(.*)'`)
var tagFieldRe = regexp.MustCompile(`\w+:".*?"`)

func parseFieldTag(tag string) *columnInfo {
	colInfo := &columnInfo{}
	infoMap := make(map[string]string)
	// fields := strings.Fields(strings.Trim(tag, "`"))
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
	if isPk, ok := infoMap["primary_key"]; ok {
		b, err := strconv.ParseBool(strings.Trim(isPk, "\""))
		if err != nil {
			panic(err)
		}
		colInfo.isPk = b
	}
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
	if isUni, ok := infoMap["unique"]; ok {
		b, err := strconv.ParseBool(strings.Trim(isUni, "\""))
		if err != nil {
			panic(err)
		}
		colInfo.isUni = b
	}
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
	name := field.Names[0].Name
	colInfo := &columnInfo{}
	if field.Tag != nil {
		colInfo = parseFieldTag(field.Tag.Value)
	}
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

// func parseRelationInfo(tag string) (relInfo relInfo) {
// 	fields := strings.Fields(strings.Trim(tag, "`"))
// 	for _, f := range fields {
// 		l := strings.Split(f, ":")
// 		k, v := l[0], strings.Trim(l[1], "\"")
// 		switch k {
// 		case "source_column":
// 			relInfo.srcCol = v
// 		case "middle_database":
// 			relInfo.midDB = v
// 		case "middle_table":
// 			relInfo.midTab = v
// 		case "middle_left_column":
// 			relInfo.midLeftCol = v
// 		case "middle_right_column":
// 			relInfo.midRightCol = v
// 		case "destination_column":
// 			dtc := strings.Split(v, ".")
// 			relInfo.dstDB = dtc[0]
// 			relInfo.dstTab = dtc[1]
// 			relInfo.dstCol = dtc[2]
// 		}
// 	}
// 	return
// }

func parseRelationInfo(tag string) (relInfo relInfo) {
	fields := strings.Fields(strings.Trim(tag, "`"))
	for _, f := range fields {
		l := strings.Split(f, ":")
		k, v := l[0], strings.Trim(l[1], "\"")
		switch k {
		case "src_col":
			relInfo.srcCol = v
		case "mid_db":
			relInfo.midDB = v
		case "mid_tab":
			relInfo.midTab = v
		case "mid_left_col":
			relInfo.midLeftCol = v
		case "mid_right_col":
			relInfo.midRightCol = v
		case "dst_db":
			relInfo.dstDB = v
		case "dst_tab":
			relInfo.dstTab = v
		case "dst_col":
			relInfo.dstCol = v
		}
	}
	return
}

func parseOneToOneField(field *ast.Field) *oneToOneInfo {
	oto := &oneToOneInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseOneToOneField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	relInfo := parseRelationInfo(field.Tag.Value)
	if relInfo.srcCol == "" || relInfo.dstDB == "" {
		panic(fmt.Errorf("nborm.parseOneToOneField() error: invalid relation information tag (%s)", field.Tag.Value))
	}
	oto.srcCol, oto.dstDB, oto.dstTab, oto.dstCol = relInfo.srcCol, relInfo.dstDB, relInfo.dstTab, relInfo.dstCol
	return oto
}

func parseForeignKeyField(field *ast.Field) *foreignKeyInfo {
	fk := &foreignKeyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseForeignKeyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	relInfo := parseRelationInfo(field.Tag.Value)
	if relInfo.srcCol == "" || relInfo.dstDB == "" {
		panic(fmt.Errorf("nborm.parseForeignKeyField() error: invalid relation information tag (%s)", field.Tag.Value))
	}
	fk.srcCol, fk.dstDB, fk.dstTab, fk.dstCol = relInfo.srcCol, relInfo.dstDB, relInfo.dstTab, relInfo.dstCol
	return fk
}

func parseReverseForeignKeyField(field *ast.Field) *reverseForeignKeyInfo {
	rfk := &reverseForeignKeyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseReverseForeignKeyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	relInfo := parseRelationInfo(field.Tag.Value)
	if relInfo.srcCol == "" || relInfo.dstDB == "" {
		panic(fmt.Errorf("nborm.parseReverseForeignKeyField() error: invalid relation information tag (%s)", field.Tag.Value))
	}
	rfk.srcCol, rfk.dstDB, rfk.dstTab, rfk.dstCol = relInfo.srcCol, relInfo.dstDB, relInfo.dstTab, relInfo.dstCol
	return rfk
}

func parseManyToManyField(field *ast.Field) *manyToManyInfo {
	mtm := &manyToManyInfo{}
	if field.Tag == nil {
		panic(fmt.Errorf("nborm.parseManyToManyField() error: no relation information tag (%s)", field.Names[0].Name))
	}
	relInfo := parseRelationInfo(field.Tag.Value)
	if relInfo.srcCol == "" || relInfo.midDB == "" || relInfo.midTab == "" || relInfo.midLeftCol == "" || relInfo.midRightCol == "" ||
		relInfo.dstDB == "" {
		panic(fmt.Errorf("nborm.parseManyToManyField() error: invalid relation information tag (%s)", field.Tag.Value))
	}
	mtm.srcCol, mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.midRightCol, mtm.dstDB, mtm.dstTab, mtm.dstCol = relInfo.srcCol,
		relInfo.midDB, relInfo.midTab, relInfo.midLeftCol, relInfo.midRightCol, relInfo.dstDB, relInfo.dstTab, relInfo.dstCol
	return mtm
}

type dbAndTab struct {
	db  string
	tab string
}

var dbTabMap = map[string]*dbAndTab{}

// func parseDBAndTab(fn *ast.FuncDecl) {
// 	if fn.Recv != nil {
// 		if recvType, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
// 			recv := recvType.X.(*ast.Ident).Name
// 			var db, tab string
// 			switch fn.Name.Name {
// 			case "DB":
// 				if retStmt, ok := fn.Body.List[len(fn.Body.List)-1].(*ast.ReturnStmt); ok {
// 					db = strings.Trim(retStmt.Results[0].(*ast.BasicLit).Value, "\"")
// 				}
// 			case "Tab":
// 				if retStmt, ok := fn.Body.List[len(fn.Body.List)-1].(*ast.ReturnStmt); ok {
// 					tab = strings.Trim(retStmt.Results[0].(*ast.BasicLit).Value, "\"")
// 				}
// 			}
// 			if _, ok := dbTabMap[recv]; !ok {
// 				dbTabMap[recv] = &dbAndTab{}
// 			}
// 			if db != "" {
// 				dbTabMap[recv].db = db
// 			}
// 			if tab != "" {
// 				dbTabMap[recv].tab = tab
// 			}
// 		}
// 	}
// }

func parseModel(decl *ast.GenDecl) {
	if typeSpec, ok := decl.Specs[0].(*ast.TypeSpec); ok {
		if stctType, ok := typeSpec.Type.(*ast.StructType); ok {
			modelName := typeSpec.Name.Name
			dbName, tabName := parseDBName(modelName), parseTabName(modelName)
			// if dt, ok := dbTabMap[modelName]; !ok {
			// 	panic(fmt.Errorf("nborm.parseModel() error: no database name and table name for (%s)", modelName))
			// } else {
			if _, ok := schemaCache.databaseMap[dbName]; !ok {
				schemaCache.databaseMap[dbName] = &databaseInfo{tableMap: map[string]*tableInfo{tabName: &tableInfo{columnMap: make(map[string]*columnInfo)}}}
			}
			if _, ok := schemaCache.databaseMap[dbName].tableMap[tabName]; !ok {
				schemaCache.databaseMap[dbName].tableMap[tabName] = &tableInfo{columnMap: make(map[string]*columnInfo)}
			}
			tabInfo := schemaCache.databaseMap[dbName].tableMap[tabName]
			tabInfo.db, tabInfo.tab, tabInfo.modelName = dbName, tabName, modelName
			tabInfo.keys = parseKeys(modelName)
			for _, field := range stctType.Fields.List {
				if expr, ok := field.Type.(*ast.SelectorExpr); ok {
					switch expr.Sel.Name {
					case "StringField", "IntField", "FloatField", "BoolField", "BinaryField", "DateField", "DatetimeField":
						colInfo := parseField(field)
						tabInfo.columns = append(tabInfo.columns, colInfo)
						tabInfo.columnMap[colInfo.colName] = colInfo
						if colInfo.isPk {
							tabInfo.pks = append(tabInfo.pks, colInfo)
						}
						if colInfo.isUni {
							tabInfo.unis = append(tabInfo.unis, colInfo)
						}
						if colInfo.isInc {
							tabInfo.inc = colInfo
						}
					case "OneToOne":
						otoInfo := parseOneToOneField(field)
						tabInfo.oneToOnes = append(tabInfo.oneToOnes, otoInfo)
					case "ForeignKey":
						fkInfo := parseForeignKeyField(field)
						tabInfo.foreignKeys = append(tabInfo.foreignKeys, fkInfo)
					case "ReverseForeignKey":
						rfkInfo := parseReverseForeignKeyField(field)
						tabInfo.reverseForeignKeys = append(tabInfo.reverseForeignKeys, rfkInfo)
					case "ManyToMany":
						mtmInfo := parseManyToManyField(field)
						tabInfo.manyToManys = append(tabInfo.manyToManys, mtmInfo)
					}
				}
			}
			for _, mtm := range tabInfo.manyToManys {
				if _, ok := schemaCache.databaseMap[mtm.midDB]; !ok {
					schemaCache.databaseMap[mtm.midDB] = &databaseInfo{tableMap: make(map[string]*tableInfo)}
				}
				if _, ok := schemaCache.databaseMap[mtm.midDB].tableMap[mtm.midTab]; !ok {
					schemaCache.databaseMap[mtm.midDB].tableMap[mtm.midTab] = &tableInfo{columnMap: make(map[string]*columnInfo)}
				}
				midTab := schemaCache.databaseMap[mtm.midDB].tableMap[mtm.midTab]
				if len(midTab.columns) == 0 {
					idCol := &columnInfo{colName: "id", sqlType: "INT", isInc: true, isPk: true}
					midTab.columnMap["id"] = idCol
					midTab.columns = append(midTab.columns, idCol)
					midTab.inc = idCol
					midTab.pks = append(midTab.pks, idCol)
				}
			}
			// }
		}
	}
}

func parseDB(filename string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	// ast.Inspect(f, func(node ast.Node) bool {
	// 	if funcDecl, ok := node.(*ast.FuncDecl); ok {
	// 		parseDBAndTab(funcDecl)
	// 	}
	// 	return true
	// })
	ast.Inspect(f, func(node ast.Node) bool {
		if genDecl, ok := node.(*ast.GenDecl); ok {
			parseModel(genDecl)
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
			pks := make([]string, len(tab.pks))
			unis := make([]string, len(tab.unis))
			for i, col := range tab.columns {
				l := make([]string, 0, 8)
				l = append(l, wrap(col.colName), col.sqlType)
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
			for i, pk := range tab.pks {
				pks[i] = wrap(pk.colName)
			}
			for i, uni := range tab.unis {
				unis[i] = fmt.Sprintf("UNIQUE KEY (%s)", wrap(uni.colName))
			}
			var uniClause string
			if len(unis) > 0 {
				uniClause = ", " + strings.Join(unis, ", ")
			}
			stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s, PRIMARY KEY (%s) %s)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
				wrap(tname), strings.Join(cols, ", "), strings.Join(pks, ", "), uniClause)
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
	return nil
}

var dbRe = regexp.MustCompile(`DB:(.*?)\n`)
var tabRe = regexp.MustCompile(`Tab:(.*?)\n`)
var keysRe = regexp.MustCompile(`key:(.*?)\n`)

func parseDBName(name string) string {
	group := dbRe.FindStringSubmatch(commentMap[name])
	if len(group) < 2 {
		panic(fmt.Errorf("nborm.parseDBName() error: the database name of %s not exists", name))
	}
	return group[1]
}

func parseTabName(name string) string {
	group := tabRe.FindStringSubmatch(commentMap[name])
	if len(group) < 2 {
		panic(fmt.Errorf("nborm.parseTabName() error: the table name of %s not exists", name))
	}
	return group[1]
}

func parseKeys(name string) []string {
	l := make([]string, 0, 8)
	if comment, ok := commentMap[name]; ok {
		keys := keysRe.FindAllStringSubmatch(comment, -1)
		for _, key := range keys {
			l = append(l, key[1])
		}
	} else {
		panic(fmt.Errorf("nborm.parseKeys() error: model name not exists"))
	}
	return l
}

func CreateMethodFile(path string) {
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		os.Remove(filepath.Join(path, "modelMethods.go"))
	// 		panic(err)
	// 	}
	// }()
	f, err := os.OpenFile(filepath.Join(path, "modelMethods.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Errorf("nborm.createMethodFile() error: %s", err.Error()))
	}
	defer f.Close()
	infos := make(map[string]map[string]string)
	for dbName, db := range schemaCache.databaseMap {
		infos[dbName] = make(map[string]string)
		for tabName, tab := range db.tableMap {
			infos[dbName][tabName] = tab.modelName
		}
	}
	content, err := nbfmt.Fmt(methodTemplate, map[string]interface{}{"package": pkg, "infos": infos})
	if err != nil {
		panic(fmt.Errorf("nborm.createMethodFile error(): %s", err.Error()))
	}
	f.WriteString(content)
	if err := exec.Command("go", "fmt", filepath.Join(path, "modelMethods.go")).Run(); err != nil {
		panic(fmt.Errorf("nborm.CreateMethodFile() error: %s", err.Error()))
	}
}
