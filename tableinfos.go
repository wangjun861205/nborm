package nborm

var dbInfo = make(map[string]map[string]*TableInfo)

// //OneToOneInfo OneToOne relation information in database
// type OneToOneInfo struct {
// 	SrcDB  string     `json:"srcDB"`
// 	SrcTab tableInfo  `json:"srcTab"`
// 	SrcCol ColumnInfo `json:"srcCol"`
// 	DstDB  string     `json:"dstDB"`
// 	DstTab tableInfo  `json:"dstTab"`
// 	DstCol ColumnInfo `json:"dstCol"`
// }

// //ForeignKeyInfo ForeignKey relation information in database
// type ForeignKeyInfo struct {
// 	SrcDB  string     `json:"srcDB"`
// 	SrcTab tableInfo  `json:"srcTab"`
// 	SrcCol ColumnInfo `json:"srcCol"`
// 	DstDB  string     `json:"dstDB"`
// 	DstTab tableInfo  `json:"dstTab"`
// 	DstCol ColumnInfo `json:"dstCol"`
// }

// //ReverseForeignKeyInfo ReverseForeignKey relation information in database
// type ReverseForeignKeyInfo struct {
// 	SrcDB  string     `json:"srcDB"`
// 	SrcTab tableInfo  `json:"srcTab"`
// 	SrcCol ColumnInfo `json:"srcCol"`
// 	DstDB  string     `json:"dstDB"`
// 	DstTab tableInfo  `json:"dstTab"`
// 	DstCol ColumnInfo `json:"dstCol"`
// }

// //ManyToManyInfo ManyToMany relation information in database
// type ManyToManyInfo struct {
// 	SrcDB       string     `json:"srcDB"`
// 	SrcTab      tableInfo  `json:"srcTab"`
// 	SrcCol      ColumnInfo `json:"srcCol"`
// 	MidDB       string     `json:"midDB"`
// 	MidTab      tableInfo  `json:"midTab"`
// 	MidLeftCol  ColumnInfo `json:"midLeftCol"`
// 	MidRightCol ColumnInfo `json:"midRightCol"`
// 	DstDB       string     `json:"dstDB"`
// 	DstTab      tableInfo  `json:"dstTab"`
// 	DstCol      ColumnInfo `json:"dstCol"`
// }

// //ColumnInfo column information in database
// type ColumnInfo struct {
// 	Name           string `json:"name"`
// 	FieldName      string `json:"fieldName"`
// 	Type           string `json:"type"`
// 	TypeWithLength string `json:"length"`
// 	Key            string `json:"key"`
// 	Extra          string `json:"extra"`
// 	Position       int64  `json:"position"`
// 	Pk             bool   `json:"pk"`
// 	Inc            bool   `json:"inc"`
// 	Uni            bool   `json:"uni"`
// }

// //ColumnInfoList column information list
// type ColumnInfoList []*ColumnInfo

// //Len implement sort.sorter interface
// func (l *ColumnInfoList) Len() int {
// 	return len(*l)
// }

// //Swap implement sort.sorter interface
// func (l *ColumnInfoList) Swap(i, j int) {
// 	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
// }

// //Less implement sort.sorter interface
// func (l *ColumnInfoList) Less(i, j int) bool {
// 	return (*l)[i].Position < (*l)[j].Position
// }

// type tableInfo struct {
// 	Name               string                   `json:"name"`
// 	ModelName          string                   `json:"modelName"`
// 	Columns            ColumnInfoList           `json:"columns"`
// 	ColumnMap          map[string]*ColumnInfo   `json:"columnMap"`
// 	OneToOnes          []*OneToOneInfo          `json:"oneToOne"`
// 	ForeignKeys        []*ForeignKeyInfo        `json:"foreignKeys"`
// 	ReverseForeignKeys []*ReverseForeignKeyInfo `json:"reverseForeignKeys"`
// 	ManyToManys        []*ManyToManyInfo        `json:"manyToManys"`
// }

// func _getTableNames(dbName string) (names []string) {
// 	tables := NewTableList()
// 	err := Query(tables, tables.TABLE_SCHEMA.Eq(dbName), nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, table := range tables.List {
// 		name, _, _ := table.TABLE_NAME.Get()
// 		if dbInfo[dbName] == nil {
// 			dbInfo[dbName] = make(map[string]*tableInfo)
// 		}
// 		dbInfo[dbName][name] = &tableInfo{Name: name, ModelName: toGoName(name)}
// 		names = append(names, name)
// 	}
// 	return
// }

// func _getColumnInfos(dbName string, tableName string) {
// 	columns := NewColumnList()
// 	columnInfos := make(ColumnInfoList, 0, 16)
// 	err := Query(columns, columns.TABLE_SCHEMA.Eq(dbName).And(columns.TABLE_NAME.Eq(tableName)), nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, column := range columns.List {
// 		columnName, _, _ := column.COLUMN_NAME.Get()
// 		columnType, _, _ := column.DATA_TYPE.Get()
// 		columnTypeWithLength, _, _ := column.COLUMN_TYPE.Get()
// 		columnKey, _, _ := column.COLUMN_KEY.Get()
// 		columnExtra, _, _ := column.EXTRA.Get()
// 		columnPosition, _, _ := column.ORDINAL_POSITION.Get()
// 		var pk, inc, uni bool
// 		if columnKey == "PRI" {
// 			pk = true
// 		}
// 		if columnKey == "UNI" {
// 			uni = true
// 		}
// 		if columnExtra == "auto_increment" {
// 			inc = true
// 		}
// 		columnInfos = append(columnInfos,
// 			&ColumnInfo{
// 				columnName,
// 				toGoName(columnName),
// 				columnType,
// 				columnTypeWithLength,
// 				columnKey,
// 				columnExtra,
// 				columnPosition,
// 				pk,
// 				inc,
// 				uni})
// 	}
// 	sort.Sort(&columnInfos)
// 	colMap := make(map[string]*ColumnInfo)
// 	for _, col := range columnInfos {
// 		colMap[col.Name] = col
// 	}
// 	dbInfo[dbName][tableName].Columns = columnInfos
// 	dbInfo[dbName][tableName].ColumnMap = colMap
// }

// //GetDBInfo get database information
// func GetDBInfo(dbNames ...string) {
// 	for _, dbName := range dbNames {
// 		tableNames := _getTableNames(dbName)
// 		for _, name := range tableNames {
// 			_getColumnInfos(dbName, name)
// 		}
// 	}
// 	for dbName, tabMap := range dbInfo {
// 		for tabName := range tabMap {
// 			_getRelations(dbName, tabName)
// 		}
// 	}
// }

// func _getRelations(dbName, tableName string) {
// 	keys := NewKeyColumnUsageList()
// 	err := Query(keys, keys.TABLE_SCHEMA.Eq(dbName).And(
// 		keys.TABLE_NAME.Eq(tableName)).And(
// 		keys.REFERENCED_TABLE_SCHEMA.NotNull()), nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if strings.Contains(tableName, "__") {
// 		if keys.Len() != 2 {
// 			panic("nborm._getRelations() the middle table must contain 2 foreign key constained")
// 		}
// 		srcDB, _, _ := keys.List[0].REFERENCED_TABLE_SCHEMA.Get()
// 		srcTab, _, _ := keys.List[0].REFERENCED_TABLE_NAME.Get()
// 		srcCol, _, _ := keys.List[0].REFERENCED_COLUMN_NAME.Get()
// 		midDB := dbName
// 		midTab := tableName
// 		midLeftCol, _, _ := keys.List[0].COLUMN_NAME.Get()
// 		midRightCol, _, _ := keys.List[1].COLUMN_NAME.Get()
// 		dstDB, _, _ := keys.List[1].REFERENCED_TABLE_SCHEMA.Get()
// 		dstTab, _, _ := keys.List[1].REFERENCED_TABLE_NAME.Get()
// 		dstCol, _, _ := keys.List[1].REFERENCED_COLUMN_NAME.Get()
// 		dbInfo[srcDB][srcTab].ManyToManys = append(dbInfo[srcDB][srcTab].ManyToManys,
// 			&ManyToManyInfo{
// 				srcDB,
// 				*dbInfo[srcDB][srcTab],
// 				*dbInfo[srcDB][srcTab].ColumnMap[srcCol],
// 				midDB,
// 				*dbInfo[midDB][midTab],
// 				*dbInfo[midDB][midTab].ColumnMap[midLeftCol],
// 				*dbInfo[midDB][midTab].ColumnMap[midRightCol],
// 				dstDB,
// 				*dbInfo[dstDB][dstTab],
// 				*dbInfo[dstDB][dstTab].ColumnMap[dstCol]})
// 		dbInfo[dstDB][dstTab].ManyToManys = append(dbInfo[dstDB][dstTab].ManyToManys,
// 			&ManyToManyInfo{
// 				dstDB,
// 				*dbInfo[dstDB][dstTab],
// 				*dbInfo[dstDB][dstTab].ColumnMap[dstCol],
// 				midDB,
// 				*dbInfo[midDB][midTab],
// 				*dbInfo[midDB][midTab].ColumnMap[midRightCol],
// 				*dbInfo[midDB][midTab].ColumnMap[midLeftCol],
// 				srcDB,
// 				*dbInfo[srcDB][srcTab],
// 				*dbInfo[srcDB][srcTab].ColumnMap[srcCol]})
// 	} else {
// 		for _, key := range keys.List {
// 			dstDB, _, _ := key.REFERENCED_TABLE_SCHEMA.Get()
// 			dstTab, _, _ := key.REFERENCED_TABLE_NAME.Get()
// 			dstCol, _, _ := key.REFERENCED_COLUMN_NAME.Get()
// 			srcDB := dbName
// 			srcTab := tableName
// 			srcCol, _, _ := key.COLUMN_NAME.Get()
// 			if dbInfo[srcDB][srcTab].ColumnMap[srcCol].Key == "PRI" || dbInfo[srcDB][srcTab].ColumnMap[srcCol].Key == "UNI" {
// 				dbInfo[srcDB][srcTab].OneToOnes = append(dbInfo[srcDB][srcTab].OneToOnes,
// 					&OneToOneInfo{
// 						srcDB,
// 						*dbInfo[srcDB][srcTab],
// 						*dbInfo[srcDB][srcTab].ColumnMap[srcCol],
// 						dstDB,
// 						*dbInfo[dstDB][dstTab],
// 						*dbInfo[dstDB][dstTab].ColumnMap[dstCol]})
// 				dbInfo[dstDB][dstTab].OneToOnes = append(dbInfo[dstDB][dstTab].OneToOnes,
// 					&OneToOneInfo{
// 						dstDB,
// 						*dbInfo[dstDB][dstTab],
// 						*dbInfo[dstDB][dstTab].ColumnMap[dstCol],
// 						srcDB,
// 						*dbInfo[srcDB][srcTab],
// 						*dbInfo[srcDB][srcTab].ColumnMap[srcCol]})
// 			} else {
// 				dbInfo[srcDB][srcTab].ForeignKeys = append(dbInfo[srcDB][srcTab].ForeignKeys,
// 					&ForeignKeyInfo{
// 						srcDB,
// 						*dbInfo[srcDB][srcTab],
// 						*dbInfo[srcDB][srcTab].ColumnMap[srcCol],
// 						dstDB,
// 						*dbInfo[dstDB][dstTab],
// 						*dbInfo[dstDB][dstTab].ColumnMap[dstCol]})
// 				dbInfo[dstDB][dstTab].ReverseForeignKeys = append(dbInfo[dstDB][dstTab].ReverseForeignKeys,
// 					&ReverseForeignKeyInfo{
// 						dstDB,
// 						*dbInfo[dstDB][dstTab],
// 						*dbInfo[dstDB][dstTab].ColumnMap[dstCol],
// 						srcDB,
// 						*dbInfo[srcDB][srcTab],
// 						*dbInfo[srcDB][srcTab].ColumnMap[srcCol]})
// 			}
// 		}
// 	}
// }

// //MarshalDBInfo write database information to json file
// func MarshalDBInfo() {
// 	f, err := os.OpenFile("dbinfo.json", os.O_CREATE|os.O_WRONLY, 0664)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	b, err := json.MarshalIndent(dbInfo, "\t", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
// 	n, err := f.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = f.Truncate(int64(n))
// 	if err != nil {
// 		panic(err)
// 	}
// }

// //GenDef generate definitions.go
// func GenDef(pkg, username, password, address string) {
// 	f, err := os.OpenFile("definitions.go", os.O_CREATE|os.O_WRONLY, 0755)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	s, err := nbfmt.Fmt(template, map[string]interface{}{"pkgName": pkg, "username": username, "password": password, "address": address, "infos": dbInfo})
// 	if err != nil {
// 		panic(err)
// 	}
// 	n, err := f.WriteString(s)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = f.Truncate(int64(n))
// 	if err != nil {
// 		panic(err)
// 	}
// 	cmd := exec.Command("go", "fmt")
// 	if err := cmd.Run(); err != nil {
// 		panic(err)
// 	}
// }

// func toGoName(s string) string {
// 	s = strings.Replace(s, "__", "$", -1)
// 	l := strings.Split(s, "_")
// 	for i, name := range l {
// 		l[i] = strings.Title(name)
// 	}
// 	return strings.Replace(strings.Join(l, ""), "$", "__", -1)
// }
