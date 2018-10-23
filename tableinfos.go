package nborm

import (
	"encoding/json"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/wangjun861205/nbfmt"
)

var dbInfo = make(map[string]map[string]*tableInfo)

type OneToOneInfo struct {
	SrcDB  string     `json:"srcDB"`
	SrcTab tableInfo  `json:"srcTab"`
	SrcCol ColumnInfo `json:"srcCol"`
	DstDB  string     `json:"dstDB"`
	DstTab tableInfo  `json:"dstTab"`
	DstCol ColumnInfo `json:"dstCol"`
}

type ForeignKeyInfo struct {
	SrcDB  string     `json:"srcDB"`
	SrcTab tableInfo  `json:"srcTab"`
	SrcCol ColumnInfo `json:"srcCol"`
	DstDB  string     `json:"dstDB"`
	DstTab tableInfo  `json:"dstTab"`
	DstCol ColumnInfo `json:"dstCol"`
}

type ReverseForeignKeyInfo struct {
	SrcDB  string     `json:"srcDB"`
	SrcTab tableInfo  `json:"srcTab"`
	SrcCol ColumnInfo `json:"srcCol"`
	DstDB  string     `json:"dstDB"`
	DstTab tableInfo  `json:"dstTab"`
	DstCol ColumnInfo `json:"dstCol"`
}

type ManyToManyInfo struct {
	SrcDB       string     `json:"srcDB"`
	SrcTab      tableInfo  `json:"srcTab"`
	SrcCol      ColumnInfo `json:"srcCol"`
	MidDB       string     `json:"midDB"`
	MidTab      tableInfo  `json:"midTab"`
	MidLeftCol  ColumnInfo `json:"midLeftCol"`
	MidRightCol ColumnInfo `json:"midRightCol`
	DstDB       string     `json:"dstDB"`
	DstTab      tableInfo  `json:"dstTab"`
	DstCol      ColumnInfo `json:"dstCol"`
}

type ColumnInfo struct {
	Name           string `json:"name"`
	FieldName      string `json:"fieldName"`
	Type           string `json:"type"`
	TypeWithLength string `json:"length"`
	Key            string `json:"key"`
	Extra          string `json:"extra"`
	Position       int64  `json:"position"`
	Pk             bool   `json:"pk"`
	Inc            bool   `json:"inc"`
	Uni            bool   `json:"uni"`
}

type ColumnInfoList []*ColumnInfo

func (l *ColumnInfoList) Len() int {
	return len(*l)
}

func (l *ColumnInfoList) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *ColumnInfoList) Less(i, j int) bool {
	return (*l)[i].Position < (*l)[j].Position
}

type tableInfo struct {
	Name               string                   `json:"name"`
	ModelName          string                   `json:"modelName"`
	Columns            ColumnInfoList           `json:"columns"`
	ColumnMap          map[string]*ColumnInfo   `json:"columnMap"`
	OneToOnes          []*OneToOneInfo          `json:"oneToOne"`
	ForeignKeys        []*ForeignKeyInfo        `json:"foreignKeys"`
	ReverseForeignKeys []*ReverseForeignKeyInfo `json:"reverseForeignKeys"`
	ManyToManys        []*ManyToManyInfo        `json:"manyToManys"`
}

func _getTableNames(dbName string) (names []string) {
	tables := make(TableList, 0, 16)
	err := Query(&tables, TableExample.TABLE_SCHEMA.Eq(dbName), nil, nil)
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		name, _, _ := table.TABLE_NAME.Get()
		if dbInfo[dbName] == nil {
			dbInfo[dbName] = make(map[string]*tableInfo)
		}
		dbInfo[dbName][name] = &tableInfo{Name: name, ModelName: toGoName(name)}
		names = append(names, name)
	}
	return
}

func _getColumnInfos(dbName string, tableName string) {
	columns := make(ColumnList, 0, 16)
	columnInfos := make(ColumnInfoList, 0, 16)
	err := Query(&columns, ColumnExample.TABLE_SCHEMA.Eq(dbName).And(ColumnExample.TABLE_NAME.Eq(tableName)), nil, nil)
	if err != nil {
		panic(err)
	}
	for _, column := range columns {
		columnName, _, _ := column.COLUMN_NAME.Get()
		columnType, _, _ := column.DATA_TYPE.Get()
		columnTypeWithLength, _, _ := column.COLUMN_TYPE.Get()
		columnKey, _, _ := column.COLUMN_KEY.Get()
		columnExtra, _, _ := column.EXTRA.Get()
		columnPosition, _, _ := column.ORDINAL_POSITION.Get()
		var pk, inc, uni bool
		if columnKey == "PRI" {
			pk = true
		}
		if columnKey == "UNI" {
			uni = true
		}
		if columnExtra == "auto_increment" {
			inc = true
		}
		columnInfos = append(columnInfos,
			&ColumnInfo{
				columnName,
				toGoName(columnName),
				columnType,
				columnTypeWithLength,
				columnKey,
				columnExtra,
				columnPosition,
				pk,
				inc,
				uni})
	}
	sort.Sort(&columnInfos)
	colMap := make(map[string]*ColumnInfo)
	for _, col := range columnInfos {
		colMap[col.Name] = col
	}
	dbInfo[dbName][tableName].Columns = columnInfos
	dbInfo[dbName][tableName].ColumnMap = colMap
}

func GetDBInfo(dbNames ...string) {
	for _, dbName := range dbNames {
		tableNames := _getTableNames(dbName)
		for _, name := range tableNames {
			_getColumnInfos(dbName, name)
		}
	}
	for dbName, tabMap := range dbInfo {
		for tabName, _ := range tabMap {
			_getRelations(dbName, tabName)
		}
	}
}

func _getRelations(dbName, tableName string) {
	keys := make(KeyColumnUsageList, 0, 16)
	err := Query(&keys, KeyColumnUsageExample.TABLE_SCHEMA.Eq(dbName).And(
		KeyColumnUsageExample.TABLE_NAME.Eq(tableName)).And(
		KeyColumnUsageExample.REFERENCED_TABLE_SCHEMA.NotNull()), nil, nil)
	if err != nil {
		panic(err)
	}
	if strings.Contains(tableName, "__") {
		if keys.Len() != 2 {
			panic("nborm._getRelations() the middle table must contain 2 foreign key constained")
		}
		srcDB, _, _ := keys[0].REFERENCED_TABLE_SCHEMA.Get()
		srcTab, _, _ := keys[0].REFERENCED_TABLE_NAME.Get()
		srcCol, _, _ := keys[0].REFERENCED_COLUMN_NAME.Get()
		midDB := dbName
		midTab := tableName
		midLeftCol, _, _ := keys[0].COLUMN_NAME.Get()
		midRightCol, _, _ := keys[1].COLUMN_NAME.Get()
		dstDB, _, _ := keys[1].REFERENCED_TABLE_SCHEMA.Get()
		dstTab, _, _ := keys[1].REFERENCED_TABLE_NAME.Get()
		dstCol, _, _ := keys[1].REFERENCED_COLUMN_NAME.Get()
		dbInfo[srcDB][srcTab].ManyToManys = append(dbInfo[srcDB][srcTab].ManyToManys,
			&ManyToManyInfo{
				srcDB,
				*dbInfo[srcDB][srcTab],
				*dbInfo[srcDB][srcTab].ColumnMap[srcCol],
				midDB,
				*dbInfo[midDB][midTab],
				*dbInfo[midDB][midTab].ColumnMap[midLeftCol],
				*dbInfo[midDB][midTab].ColumnMap[midRightCol],
				dstDB,
				*dbInfo[dstDB][dstTab],
				*dbInfo[dstDB][dstTab].ColumnMap[dstCol]})
		dbInfo[dstDB][dstTab].ManyToManys = append(dbInfo[dstDB][dstTab].ManyToManys,
			&ManyToManyInfo{
				dstDB,
				*dbInfo[dstDB][dstTab],
				*dbInfo[dstDB][dstTab].ColumnMap[dstCol],
				midDB,
				*dbInfo[midDB][midTab],
				*dbInfo[midDB][midTab].ColumnMap[midRightCol],
				*dbInfo[midDB][midTab].ColumnMap[midLeftCol],
				srcDB,
				*dbInfo[srcDB][srcTab],
				*dbInfo[srcDB][srcTab].ColumnMap[srcCol]})
	} else {
		for _, key := range keys {
			dstDB, _, _ := key.REFERENCED_TABLE_SCHEMA.Get()
			dstTab, _, _ := key.REFERENCED_TABLE_NAME.Get()
			dstCol, _, _ := key.REFERENCED_COLUMN_NAME.Get()
			srcDB := dbName
			srcTab := tableName
			srcCol, _, _ := key.COLUMN_NAME.Get()
			if dbInfo[srcDB][srcTab].ColumnMap[srcCol].Key == "PRI" || dbInfo[srcDB][srcTab].ColumnMap[srcCol].Key == "UNI" {
				dbInfo[srcDB][srcTab].OneToOnes = append(dbInfo[srcDB][srcTab].OneToOnes,
					&OneToOneInfo{
						srcDB,
						*dbInfo[srcDB][srcTab],
						*dbInfo[srcDB][srcTab].ColumnMap[srcCol],
						dstDB,
						*dbInfo[dstDB][dstTab],
						*dbInfo[dstDB][dstTab].ColumnMap[dstCol]})
				dbInfo[dstDB][dstTab].OneToOnes = append(dbInfo[dstDB][dstTab].OneToOnes,
					&OneToOneInfo{
						dstDB,
						*dbInfo[dstDB][dstTab],
						*dbInfo[dstDB][dstTab].ColumnMap[dstCol],
						srcDB,
						*dbInfo[srcDB][srcTab],
						*dbInfo[srcDB][srcTab].ColumnMap[srcCol]})
			} else {
				dbInfo[srcDB][srcTab].ForeignKeys = append(dbInfo[srcDB][srcTab].ForeignKeys,
					&ForeignKeyInfo{
						srcDB,
						*dbInfo[srcDB][srcTab],
						*dbInfo[srcDB][srcTab].ColumnMap[srcCol],
						dstDB,
						*dbInfo[dstDB][dstTab],
						*dbInfo[dstDB][dstTab].ColumnMap[dstCol]})
				dbInfo[dstDB][dstTab].ReverseForeignKeys = append(dbInfo[dstDB][dstTab].ReverseForeignKeys,
					&ReverseForeignKeyInfo{
						dstDB,
						*dbInfo[dstDB][dstTab],
						*dbInfo[dstDB][dstTab].ColumnMap[dstCol],
						srcDB,
						*dbInfo[srcDB][srcTab],
						*dbInfo[srcDB][srcTab].ColumnMap[srcCol]})
			}
		}
	}
}

func MarshalDBInfo() {
	f, err := os.OpenFile("dbinfo.json", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := json.MarshalIndent(dbInfo, "\t", "\t")
	if err != nil {
		panic(err)
	}
	n, err := f.Write(b)
	if err != nil {
		panic(err)
	}
	err = f.Truncate(int64(n))
	if err != nil {
		panic(err)
	}
}

const template = `
package {{ pkgName }}

import (
	"github.com/wangjun861205/nborm"
	)

{{ for dbName, db in infos }}
	{{ for tabName, tab in db }}
		type {{ tab.ModelName }} struct {
			{{ for _, col in tab.Columns }}
				{{ switch col.Type }}
					{{ case "varchar", "char", "text" }}
						{{ col.FieldName }} *nborm.StringField` + " `json:\"{{ col.Name }}\"`" + `
					{{ case "int" }}
						{{ col.FieldName }} *nborm.IntField` + " `json:\"{{ col.Name }}\"`" + `
					{{ case "decimal", "float" }}
						{{ col.FieldName }} *nborm.FloatField` + " `json:\"{{ col.Name }}\"`" + `
					{{ case "tinyint" }}
						{{ if col.TypeWithLength == "tinyint(1)" }}
							{{ col.FieldName }} *nborm.BoolField` + " `json:\"{{ col.Name }}\"`" + `
						{{ else }}
							{{ col.FieldName }} *nborm.IntField` + " `json:\"{{ col.Name }}\"`" + `
						{{ endif }}
					{{ case "date" }}
						{{ col.FieldName }} *nborm.DateField` + " `json:\"{{ col.Name }}\"`" + `
					{{ case "datetime", "timestamp" }}
						{{ col.FieldName }} *nborm.DatetimeField` + " `json:\"{{ col.Name }}\"`" + `
				{{ endswitch }}
			{{ endfor }}
			{{ for _, oto in tab.OneToOnes }}
				{{ oto.DstTab.ModelName }} *nborm.OneToOne` + " `json:\"{{ oto.DstTab.Name }}\"`" + `
			{{ endfor }}
			{{ for _, fk in tab.ForeignKeys }}
				{{ fk.DstTab.ModelName }} *nborm.ForeignKey` + " `json:\"{{ fk.DstTab.Name }}\"`" + `
			{{ endfor }}
			{{ for _, rfk in tab.ReverseForeignKeys }}
				{{ rfk.DstTab.ModelName }} *nborm.ReverseForeignKey` + " `json:\"{{ rfk.DstTab.Name }}\"`" + `
			{{ endfor }}
			{{ for _, mtm in tab.ManyToManys }}
				{{ mtm.DstTab.ModelName }} *nborm.ManyToMany` + " `json:\"{{ mtm.DstTab.Name }}\"`" + `
			{{ endfor }}
		}

		func New{{ tab.ModelName }}() *{{ tab.ModelName }} {
			m := &{{ tab.ModelName }}{}
			{{ for _, col in tab.Columns }}
				{{ switch col.Type }}
					{{ case "varchar", "char", "text" }}
						m.{{ col.FieldName }} = nborm.NewStringField(m, "{{ col.Name }}", {{ col.Pk }}, {{ col.Inc }}, {{ col.Uni }})
					{{ case "int" }}
						m.{{ col.FieldName }} = nborm.NewIntField(m, "{{ col.Name }}", {{ col.Pk }}, {{ col.Inc }}, {{ col.Uni }})
					{{ case "tinyint" }}
						{{ if col.TypeWithLength == "tinyint(1)" }}
							m.{{ col.FieldName }} = nborm.NewBoolField(m, "{{ col.Name }}", {{ col.Pk }}, {{ col.Inc }}, {{ col.Uni }})
						{{ else }}
							m.{{ col.FieldName }} = nborm.NewIntField(m, "{{ col.Name }}", {{ col.Pk }}, {{ col.Inc }}, {{ col.Uni }})
						{{ endif }}
					{{ case "decimal", "float" }}
						m.{{ col.FieldName }} = nborm.NewFloatField(m, "{{ col.Name }}", {{ col.Pk }}, {{ col.Inc }}, {{ col.Uni }})
					{{ case "date" }}
						m.{{ col.FieldName }} = nborm.NewDateField(m, "{{ col.Name }}", {{ col.Pk }}, {{ col.Inc }}, {{ col.Uni }})
					{{ case "datetime", "timestamp" }}
						m.{{ col.FieldName }} = nborm.NewDatetimeField(m, "{{ col.Name }}", {{ col.Pk }}, {{ col.Inc }}, {{ col.Uni }})
				{{ endswitch }}
			{{ endfor }}
			{{ for i, oto in tab.OneToOnes }}
				otoDstMod{{ i }} := New{{ oto.DstTab.ModelName }}()
				m.{{ oto.DstTab.ModelName }} = nborm.NewOneToOne(m.{{ oto.SrcCol.FieldName }}, otoDstMod{{ i }}.{{ oto.DstCol.FieldName }}, otoDstMod{{ i }})
			{{ endfor }}
			{{ for i, fk in tab.ForeignKeys }}
				fkDstMod{{ i }} := New{{ fk.DstTab.ModelName }}()
				m.{{ fk.DstTab.ModelName }} = nborm.NewForeignKey(m.{{ fk.SrcCol.FieldName }}, fkDstMod{{ i }}.{{ fk.DstCol.FieldName }}, fkDstMod{{ i }})
			{{ endfor }}
			{{ for i, rfk in tab.ReverseForeignKeys }}
				rfkDstModList{{ i }} := New{{ rfk.DstTab.ModelName }}List()
				m.{{ rfk.DstTab.ModelName }} = nborm.NewReverseForeignKey(m.{{ rfk.SrcCol.FieldName }}, rfkDstModList{{ i }}.{{ rfk.DstCol.FieldName }}, rfkDstModList{{ i }})
			{{ endfor }}
			{{ for i, mtm in tab.ManyToManys }}
				mtmDstModList{{ i }} := New{{ mtm.DstTab.ModelName }}List()
				mtmMidMod{{ i }} := New{{ mtm.MidTab.ModelName }}()
				m.{{ mtm.DstTab.ModelName }} = nborm.NewManyToMany(m.{{ mtm.SrcCol.FieldName }}, mtmMidMod{{ i }}.{{ mtm.MidLeftCol.FieldName }}, 
					mtmMidMod{{ i }}.{{ mtm.MidRightCol.FieldName }}, mtmDstModList{{ i }}.{{ mtm.DstCol.FieldName }}, mtmDstModList{{ i }})
			{{ endfor }}	
			return m
		}

		func (m *{{ tab.ModelName }}) DB() string {
			return "{{ dbName }}"
		}

		func (m *{{ tab.ModelName }}) Tab() string {
			return "{{ tabName }}"
		}

		func (m *{{ tab.ModelName }}) Fields() []nborm.Field {
			return []nborm.Field{
				{{ for _, col in tab.Columns }}
					m.{{ col.FieldName }},
				{{ endfor }}
			}
		}

		type {{ tab.ModelName }}List struct {
			*{{ tab.ModelName }}
			List []*{{ tab.ModelName }}
		}

		func New{{ tab.ModelName }}List() *{{ tab.ModelName }}List {
			return &{{ tab.ModelName }}List{New{{ tab.ModelName }}(), make([]*{{ tab.ModelName }}, 0, 128)}
		}

		func (l *{{ tab.ModelName }}List) New() nborm.Model {
			m := New{{ tab.ModelName }}()
			l.List = append(l.List, m)
			return m
		}

		func (l *{{ tab.ModelName }}List) Len() int {
			return len(l.List)
		}

		func (l *{{ tab.ModelName }}List) Swap(i, j int) {
			l.List[i], l.List[j] = l.List[j], l.List[i]
		}

		func (l *{{ tab.ModelName }}List) Index(i int) nborm.Model {
			return l.List[i]
		}

		func (l *{{ tab.ModelName }}List) Delete(i int) {
			switch i {
			case 0:
				l.List = l.List[1:]
			case l.Len()-1:
				l.List = l.List[:l.Len()-1]
			default:
				l.List = append(l.List[:i], l.List[i+1:]...)
			}
		}
	{{ endfor }}
{{ endfor }}

func init() {
	{{ for dbName, _ in infos }}
		nborm.RegisterDB("{{ username }}", "{{ password }}", "{{ address }}", "{{ dbName }}")
	{{ endfor }}
}
`

func GenDef(pkg, username, password, address string) {
	f, err := os.OpenFile("definitions.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s, err := nbfmt.Fmt(template, map[string]interface{}{"pkgName": pkg, "username": username, "password": password, "address": address, "infos": dbInfo})
	if err != nil {
		panic(err)
	}
	n, err := f.WriteString(s)
	if err != nil {
		panic(err)
	}
	err = f.Truncate(int64(n))
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("go", "fmt")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func toGoName(s string) string {
	s = strings.Replace(s, "__", "$", -1)
	l := strings.Split(s, "_")
	for i, name := range l {
		l[i] = strings.Title(name)
	}
	return strings.Replace(strings.Join(l, ""), "$", "__", -1)
}
