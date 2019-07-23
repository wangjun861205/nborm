package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/wangjun861205/nbfmt"

	"github.com/wangjun861205/nbcolor"

	_ "github.com/go-sql-driver/mysql"
)

var modelInfos = make([]*ModelInfo, 0, 128)

var dbRe = regexp.MustCompile(`(?m)db:(\w+|\*)$`)
var tabRe = regexp.MustCompile(`(?m)tab:(\w+)$`)
var pkRe = regexp.MustCompile(`(?m)pk:([0-9a-zA-Z,]+)$`)
var unisRe = regexp.MustCompile(`(?m)uk:([0-9a-zA-Z,]+)$`)
var colRe = regexp.MustCompile(`col:"(\w+)"`)
var incRe = regexp.MustCompile(`auto_increment:"true"`)

var relWholeRe = regexp.MustCompile(`rel:"(.*?)"`)
var relModelRe = regexp.MustCompile(`@([\w@\$]+)\[(.+?)\](?:->|$)`)
var relWhereRe = regexp.MustCompile(`\[(.*?)\]`)
var relMidModelRe = regexp.MustCompile(`([a-zA-Z0-9]+)(_\d+)?`)
var relFieldRe = regexp.MustCompile(`@([\w@\$]+)\.(\w+)`)

func getRelTag(s string) string {
	return relWholeRe.FindStringSubmatch(s)[1]
}

type relModelStr struct {
	modelStr string
	onStr    string
}

type On struct {
	Expr       string
	Fields     []*RelField
	FieldsExpr string
}

type RelModel struct {
	ModelName string
	ModelType string
	Index     int
	On        *On
}

type RelField struct {
	Model     *RelModel
	FieldName string
}

func getRelModelStr(tag string) []relModelStr {
	gs := relModelRe.FindAllStringSubmatch(tag, -1)
	l := make([]relModelStr, 0, len(gs))
	for _, g := range gs {
		l = append(l, relModelStr{g[1], g[2]})
	}
	return l
}

type RelInfo struct {
	Field        string
	Type         string
	SrcModel     *RelModel
	SrcFields    []*RelField
	MidModels    []*RelModel
	DstModel     *RelModel
	DstFields    []*RelField
	IsList       bool
	MidModelsLen int
}

func parseRelation(dstModel, tag string) error {
	lastModel := modelInfos[len(modelInfos)-1]
	lastRel := lastModel.RelInfos[len(lastModel.RelInfos)-1]
	relWholeGroup := relWholeRe.FindStringSubmatch(tag)
	if len(relWholeGroup) == 0 {
		return fmt.Errorf("no rel tag: (%s)", tag)
	}
	relStr := relWholeGroup[1]
	relModelsGroup := relModelRe.FindAllStringSubmatch(relStr, -1)
	for _, mg := range relModelsGroup {
		switch {
		case strings.HasPrefix(mg[1], "@"):
			sm := RelModel{ModelName: "m", ModelType: lastModel.Name}
			lastRel.SrcModel = &sm
		case strings.HasPrefix(mg[1], "$"):
			dm := RelModel{ModelName: fmt.Sprintf("m.%s", lastRel.Field), ModelType: dstModel}
			lastRel.DstModel = &dm
		default:
			midModelGroup := relMidModelRe.FindStringSubmatch(mg[1])
			mm := RelModel{ModelName: midModelGroup[0], ModelType: midModelGroup[1]}
			lastRel.MidModels = append(lastRel.MidModels, &mm)
		}
	}
	if lastRel.SrcModel == nil {
		lastRel.SrcModel = &RelModel{ModelName: "m", ModelType: lastModel.Name}
	}
	fieldsRe := regexp.MustCompile(`@([@\$\w]+)\.(\w+)`)
	for _, mg := range relModelsGroup {
		onStr := mg[2]
		on := On{}
		fieldsGroup := fieldsRe.FindAllStringSubmatch(onStr, -1)
		for _, fieldGroup := range fieldsGroup {
			modelName := fieldGroup[1]
			fieldName := fieldGroup[2]
			switch {
			case strings.HasPrefix(modelName, "@"):
				field := RelField{lastRel.SrcModel, fieldName}
				on.Fields = append(on.Fields, &field)
				lastRel.SrcFields = append(lastRel.SrcFields, &field)
			case strings.HasPrefix(modelName, "$"):
				field := RelField{lastRel.DstModel, fieldName}
				on.Fields = append(on.Fields, &field)
				lastRel.DstFields = append(lastRel.DstFields, &field)
			default:
				var tarModel *RelModel
				for _, mm := range lastRel.MidModels {
					if modelName == mm.ModelName {
						tarModel = mm
						break
					}
				}
				field := RelField{tarModel, fieldName}
				on.Fields = append(on.Fields, &field)
			}
		}
		on.Expr = fieldsRe.ReplaceAllString(onStr, "@")
		switch {
		case strings.HasPrefix(mg[1], "@"):
			lastRel.SrcModel.On = &on
		case strings.HasPrefix(mg[1], "$"):
			lastRel.DstModel.On = &on
		default:
			for _, rm := range lastRel.MidModels {
				if rm.ModelName == mg[1] {
					rm.On = &on
					break
				}
			}
		}
	}
	lastRel.MidModelsLen = len(lastRel.MidModels)
	return nil
}

type FieldInfo struct {
	Type  string
	Col   string
	Field string
	IsInc bool
}

type MidWhere struct {
	Fields string
	Expr   string
}

type MidModel struct {
	Name  string
	Type  string
	Index int
}

type Field struct {
	MidModel  *MidModel
	ModelName string
	FieldName string
}

type ModelInfo struct {
	Name         string
	DB           string
	Tab          string
	FieldInfos   []*FieldInfo
	Pk           []string
	Unis         [][]string
	HasUk        bool
	Inc          string
	MidModels    []MidModel
	RelInfos     []*RelInfo
	HasRel       bool
	HasMidModels bool
}

func (m *ModelInfo) proc() {
	var incField string
	for _, field := range m.FieldInfos {
		if field.IsInc {
			incField = field.Field
			break
		}
	}
	if incField == "" {
		m.Inc = "nil"
	} else {
		m.Inc = incField
	}
	if len(m.Unis) > 0 {
		m.HasUk = true
	}
	if len(m.RelInfos) > 0 {
		m.HasRel = true
	}
	if len(m.MidModels) > 0 {
		m.HasMidModels = true
	}
	var midIdx int
	for _, relInfo := range m.RelInfos {
		for _, midMod := range relInfo.MidModels {
			midMod.Index = midIdx
			midMod.ModelName = fmt.Sprintf("mm%d", midMod.Index)
			midIdx++
		}
		for _, mm := range relInfo.MidModels {
			l := make([]string, 0, len(mm.On.Fields))
			for _, f := range mm.On.Fields {
				l = append(l, fmt.Sprintf("&%s.%s", f.Model.ModelName, f.FieldName))
			}
			mm.On.FieldsExpr = strings.Join(l, ", ")
		}
		l := make([]string, 0, len(relInfo.DstModel.On.Fields))
		for _, f := range relInfo.DstModel.On.Fields {
			l = append(l, fmt.Sprintf("&%s.%s", f.Model.ModelName, f.FieldName))
		}
		relInfo.DstModel.On.FieldsExpr = strings.Join(l, ", ")
	}
}

func (m *ModelInfo) newModelFunc() string {
	s, err := nbfmt.Fmt(`
	func New{{ model.Name }}() *{{ model.Name }} {
		m := &{{ model.Name }}{}
		nborm.InitModel(m)
		m.InitRel()
		return m
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) initRelFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) InitRel() {
		{{ for i, rel in model.RelInfos }}
			m.{{ rel.Field }} = &{{ rel.Type }}{}
			m.{{ rel.Field }}.SetParent(m)
			{{ if rel.IsList == true }}
				m.{{ rel.Field }}.dupMap = make(map[string]int)
			{{ endif }}
			nborm.InitModel(m.{{ rel.Field }})
			relInfo{{ i }} := nborm.RelationInfo{Name: "{{ rel.Field }}" }
			relInfo{{ i }}.SrcModel = m
			relInfo{{ i }}.DstModel = m.{{ rel.Field }}
			{{ for _, srcField in rel.SrcFields }}
				relInfo{{ i }}.SrcFields = append(relInfo{{ i }}.SrcFields, &m.{{ srcField.FieldName }})
			{{ endfor }}
			{{ for _, dstField in rel.DstFields }}
				relInfo{{ i }}.DstFields = append(relInfo{{ i }}.DstFields, &m.{{ rel.Field }}.{{ dstField.FieldName }})
			{{ endfor }}
			{{ for _, midModel in rel.MidModels }}
				mm{{ midModel.Index }} := &{{ midModel.ModelType }}{}
				mm{{ midModel.Index }}.SetParent(m)
				nborm.InitModel(mm{{ midModel.Index }})
				relInfo{{ i }}.MidModels = append(relInfo{{ i }}.MidModels, mm{{ midModel.Index }})
				mm{{ midModel.Index }}.SetOnCond(nborm.NewExpr("{{ midModel.On.Expr }}", {{ midModel.On.FieldsExpr }}))
			{{ endfor }}
			{{ rel.DstModel.ModelName }}.SetOnCond(nborm.NewExpr("{{ rel.DstModel.On.Expr }}", {{ rel.DstModel.On.FieldsExpr }}))
			m.Rels = append(m.Rels, relInfo{{ i }})
		{{ endfor }}
		m.AddRelInited()
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) dbFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) DB() string {
		return "{{ model.DB }}"
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) tabFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) Tab() string {
		return "{{ model.Tab }}"
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) fieldInfosFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) FieldInfos() nborm.FieldInfoList {
		return nborm.FieldInfoList {
			{{ for _, field in model.FieldInfos }}
				{"{{ field.Col }}", "{{ field.Field }}", &m.{{ field.Field }} },
			{{ endfor }}
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) autoIncFieldFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) AutoIncField() nborm.Field {
		{{ if model.Inc == "nil" }}
			return nil
		{{ else }}
			return &m.{{ model.Inc }}
		{{ endif }}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) primaryKeyFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) PrimaryKey() nborm.FieldList {
		return nborm.FieldList {
			{{ for _, pk in model.Pk }}
				&m.{{ pk }},
			{{ endfor }}
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) uniqueKeysFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) UniqueKeys() []nborm.FieldList {
			{{ if model.HasUk == false }}
				return nil
			{{ else }}
				return []nborm.FieldList{
					{{ for _, uk in model.Unis }}
					{
						{{ for _, f in uk }}
							&m.{{ f }},
						{{ endfor }}
					},
					{{ endfor }}
				}
			{{ endif }}
	}`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

// func (m *ModelInfo) relationsFunc() string {
// 	s, err := nbfmt.Fmt(`
// 	func (m *{{ model.Name }}) Relations() nborm.RelationInfoList {
// 		{{ if model.HasRel == false }}
// 			return nil
// 		{{ else }}
// 			if !m.IsRelInited() {
// 				m.InitRel()
// 			}
// 			{{ if model.HasMidModels == true }}
// 				{{ for i, mm in model.MidModels }}
// 					mm{{ i }} := m.GetMidTabs()[{{ i }}].(*{{ mm.Name }})
// 				{{ endfor }}
// 			{{ endif }}
// 			return nborm.RelationInfoList{
// 				{{ for _, info in model.RelInfos }}
// 					nborm.RelationInfo{
// 						nborm.FieldList{
// 							{{ for _, f in info.Fields }}
// 								&{{ f }},
// 							{{ endfor }}
// 						},
// 						m.{{ info.Field }},
// 						"{{ info.Field }}",
// 					},
// 				{{ endfor }}
// 			}
// 		{{ endif }}
// 	}
// 	`, map[string]interface{}{"model": m})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }

// func (m *ModelInfo) modelMarshalJSONFunc() string {
// 	s, err := nbfmt.Fmt(`
// 	func (m {{ model.Name }}) MarshalJSON() ([]byte, error) {
// 		if m.IsSynced() || m.IsContainValue() {
// 			return json.Marshal(struct{
// 				{{ for _, f in model.FieldInfos }}
// 					{{ f.Field }} interface{}
// 				{{ endfor }}
// 				{{ for _, r in model.RelInfos }}
// 					{{ r.Field }} *{{ r.Type }}
// 				{{ endfor }}
// 			}{
// 				{{ for _, f in model.FieldInfos }}
// 					{{ f.Field }}: m.{{ f.Field }}.JSONValue(),
// 				{{ endfor }}
// 				{{ for _, r in model.RelInfos }}
// 					{{ r.Field }}: m.{{ r.Field }},
// 				{{ endfor }}
// 			})
// 		}
// 		return []byte("null"), nil
// 	}
// 	`, map[string]interface{}{"model": m})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }

func (m *ModelInfo) modelMarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (m {{ model.Name }}) MarshalJSON() ([]byte, error) {
		return nborm.MarshalModel(&m), nil
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelUnmarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) UnmarshalJSON(data []byte) error {
		return nborm.UnmarshalModel(data, m)
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelCollapseFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) Collapse() {
		{{ for _, relInfo in model.RelInfos }}
			if m.{{ relInfo.Field }} != nil && m.{{ relInfo.Field }}.IsSynced() {
				m.{{ relInfo.Field }}.Collapse()
			}
		{{ endfor }}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelListType() string {
	s, err := nbfmt.Fmt(`
	type {{ model.Name }}List struct {
		{{ model.Name }}
		dupMap map[string]int
		List []*{{ model.Name }}
		Total int
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listCheckDupFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) checkDup() int {
		if l.Len() < 1 {
			return -1
		}
		var builder strings.Builder
		lastModel := l.List[l.Len()-1]
		{{ for _, f in model.FieldInfos }}
			if lastModel.{{ f.Field }}.IsValid() {
				builder.WriteString(fmt.Sprintf("%v", lastModel.{{ f.Field }}.Value()))
			}
		{{ endfor }}
		if idx, ok := l.dupMap[builder.String()]; ok {
			return idx
		}
		l.dupMap[builder.String()] = l.Len() - 1
		return -1
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) newListFunc() string {
	s, err := nbfmt.Fmt(`
	func New{{ model.Name }}List() *{{ model.Name }}List {
		l := &{{ model.Name }}List {
			{{ model.Name }}{},
			make(map[string]int),
			make([]*{{ model.Name }}, 0, 32),
			0,
		}
		nborm.InitModel(l)
		l.InitRel()
		return l
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listNewModelFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) NewModel() nborm.Model {
		m := &{{ model.Name }}{}
		m.SetParent(l.GetParent())
		m.SetConList(l)
		nborm.InitModel(m)
		m.InitRel()
		l.List = append(l.List, m)
		return m
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listSetTotalFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) SetTotal(total int) {
		l.Total = total
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listGetTotalFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) GetTotal() int {
		return l.Total
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listLenFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) Len() int {
		return len(l.List)
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) getInnerListFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) GetList() []nborm.Model {
		modelList := make([]nborm.Model, 0, l.Len())
		for _, m := range l.List {
			modelList = append(modelList, m)
		}
		return modelList
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listMarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (l {{ model.Name }}List) MarshalJSON() ([]byte, error) {
		return nborm.MarshalModel(&l), nil
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listUnmarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) UnmarshalJSON(b []byte) error {
		return nborm.UnmarshalModel(b, l)
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listCollapseFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) Collapse() {
		idx := l.checkDup()
		if idx >= 0 {
			{{ if model.HasRel == true }}
				{{ for _, relInfo in model.RelInfos }}
					{{ if relInfo.IsList == true }}
						l.List[idx].{{ relInfo.Field }}.checkDup()
						l.List[idx].{{ relInfo.Field }}.List = append(l.List[idx].{{ relInfo.Field }}.List, l.List[l.Len()-1].{{ relInfo.Field }}.List...)
					{{ else }}
						l.List[idx].{{ relInfo.Field }} = l.List[l.Len()-1].{{ relInfo.Field }}
					{{ endif }}
				{{ endfor }}
			{{ endif }}
			l.List = l.List[:len(l.List)-1]
			l.List[idx].Collapse()
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listFilterFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) Filter(f func(m *{{ model.Name }}) bool) []*{{ model.Name }} {
		ll := make([]*{{ model.Name }}, 0, l.Len())
		for _, m := range l.List {
			if f(m) {
				ll = append(ll, m)
			}
		}
		return ll
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listSliceFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) Slice(low, high int) {
		switch {
			case high <= l.Len():
				l.List = l.List[low: high]
			case low <= l.Len() && high > l.Len():
				l.List = l.List[low:]
			default:
				l.List = l.List[:0]
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func parseComment(com string) error {
	fmt.Println(nbcolor.Green(com))
	lastModelInfo := modelInfos[len(modelInfos)-1]
	dbGroup := dbRe.FindStringSubmatch(com)
	if len(dbGroup) != 2 {
		return fmt.Errorf("invalid db (%s)", com)
	}
	lastModelInfo.DB = dbGroup[1]
	tabGroup := tabRe.FindStringSubmatch(com)
	if len(tabGroup) != 2 {
		return fmt.Errorf("invalid tab (%s)", com)
	}
	lastModelInfo.Tab = tabGroup[1]
	pkGroup := pkRe.FindStringSubmatch(com)
	if len(pkGroup) != 2 {
		return fmt.Errorf("invalid primary key comment (%s)", com)
	}
	modelInfos[len(modelInfos)-1].Pk = strings.Split(pkGroup[1], ",")
	unisGroup := unisRe.FindAllStringSubmatch(com, -1)
	for _, unis := range unisGroup {
		modelInfos[len(modelInfos)-1].Unis = append(modelInfos[len(modelInfos)-1].Unis, strings.Split(unis[1], ","))
	}
	return nil
}

func parseFieldTag(field, tag string) error {
	fieldInfos := modelInfos[len(modelInfos)-1].FieldInfos
	lastField := fieldInfos[len(fieldInfos)-1]
	colGroup := colRe.FindStringSubmatch(tag)
	if len(colGroup) != 2 {
		lastField.Col = field
	} else {
		lastField.Col = colGroup[1]
	}
	lastField.IsInc = incRe.MatchString(tag)
	return nil
}

type arrayFlags []string

func (a *arrayFlags) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func (a *arrayFlags) String() string {
	return strings.Join(*a, ", ")
}

func newFile(filename string) *os.File {
	os.Remove(filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

type config struct {
	User         string
	Password     string
	Host         string
	Port         string
	Database     string
	Tables       []string
	AppendTables []string
	MatchAllDB   bool
}

func readConfig(filename string) config {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var conf config
	if err := json.Unmarshal(b, &conf); err != nil {
		panic(err)
	}
	return conf
}

func main() {
	var command, configFile string
	// tabs := make(arrayFlags, 0, 8)
	flag.StringVar(&command, "command", "", "specific command")
	flag.StringVar(&configFile, "config", "config", "specific config file")
	// flag.StringVar(&host, "host", "localhost", "specific mysql host")
	// flag.StringVar(&port, "port", "3306", "specific mysql port")
	// flag.StringVar(&user, "user", "", "specific mysql user")
	// flag.StringVar(&password, "password", "", "sepcific mysql password")
	// flag.StringVar(&db, "db", "information_schema", "specific database")
	// flag.Var(&tabs, "table", "specific tables")

	dir := flag.String("p", "./", "specific parse path")
	flag.Parse()
	if command == "model" {
		conf := readConfig(configFile)
		conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		if len(conf.AppendTables) > 0 {
			f, err := os.OpenFile("model.go", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			ts := getTabs(conn, conf.Database, conf.AppendTables...)
			getCols(conn, ts)
			getPK(conn, ts)
			getUK(conn, ts)
			procTables(ts)
			var s string
			if conf.MatchAllDB {
				s = createModelsMatchAllDB(ts)
			} else {
				s = createModels(ts)
			}
			if _, err := f.WriteString(s); err != nil {
				panic(err)
			}
			for _, tab := range conf.AppendTables {
				conf.Tables = append(conf.Tables, tab)
			}
			conf.AppendTables = conf.AppendTables[:0]
			if err := os.Remove(configFile); err != nil {
				panic(err)
			}
			b, err := json.MarshalIndent(conf, "", "\t")
			if err != nil {
				panic(err)
			}
			cf, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer cf.Close()
			if _, err := cf.Write(b); err != nil {
				panic(err)
			}
			cmd := exec.Command("gofmt", "-w", "./")
			if err := cmd.Run(); err != nil {
				log.Println(nbcolor.Red(err))
			}
			return
		}
		ts := getTabs(conn, conf.Database, conf.Tables...)
		getCols(conn, ts)
		getPK(conn, ts)
		getUK(conn, ts)
		procTables(ts)
		var s string
		if conf.MatchAllDB {
			s = createFileMatchAllDB(ts)
		} else {
			s = createFile(ts)
		}
		f := newFile("model.go")
		defer f.Close()
		if _, err := f.WriteString(s); err != nil {
			panic(err)
		}
		cmd := exec.Command("gofmt", "-w", "./")
		if err := cmd.Run(); err != nil {
			log.Println(nbcolor.Red(err))
		}
		return
	}
	os.Remove(path.Join(*dir, "methods.go"))
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, *dir, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(nbcolor.Red(err.Error()))
		return
	}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			if f.Name.Name == "methods.go" {
				continue
			}
			ast.Inspect(f, func(node ast.Node) bool {
				switch n := node.(type) {
				case *ast.GenDecl:
					if t, ok := n.Specs[0].(*ast.TypeSpec); ok {
						modelInfos = append(modelInfos, &ModelInfo{Name: t.Name.Name})
						parseComment(n.Doc.Text())
						return true
					}
					return false
				case *ast.TypeSpec:
					return true
				case *ast.StructType:
					for _, field := range n.Fields.List {
						switch ft := field.Type.(type) {
						case *ast.SelectorExpr:
							if ft.X.(*ast.Ident).String() == "nborm" {
								switch ft.Sel.Name {
								case "Meta":
									continue
								default:
									lastModel := modelInfos[len(modelInfos)-1]
									lastModel.FieldInfos = append(lastModel.FieldInfos, &FieldInfo{Field: field.Names[0].Name, Type: fmt.Sprintf("nborm.%s", ft.Sel.Name)})
									if field.Tag == nil {
										lastField := lastModel.FieldInfos[len(lastModel.FieldInfos)-1]
										lastField.Col = field.Names[0].Name
										continue
									}
									if err := parseFieldTag(field.Names[0].Name, field.Tag.Value); err != nil {
										panic(err)
									}
								}
							}
						case *ast.StarExpr:
							lastModel := modelInfos[len(modelInfos)-1]
							var isList bool
							if strings.HasSuffix(ft.X.(*ast.Ident).String(), "List") {
								isList = true
							}
							lastModel.RelInfos = append(lastModel.RelInfos, &RelInfo{Field: field.Names[0].Name, Type: ft.X.(*ast.Ident).String(), IsList: isList})
							if err := parseRelation(field.Names[0].Name, field.Tag.Value); err != nil {
								panic(err)
							}
						}
					}
					return false
				default:
					return true
				}
			})
		}
		nf, err := os.OpenFile(path.Join(*dir, "methods.go"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer nf.Close()
		nf.WriteString(fmt.Sprintf("package %s\n", pkg.Name))
		nf.WriteString(`
		import (
			"github.com/wangjun861205/nborm"
			"strings"
			"fmt"
		)
		`)
		for _, m := range modelInfos {
			m.proc()
			nf.WriteString(m.newModelFunc())
			nf.WriteString(m.initRelFunc())
			nf.WriteString(m.dbFunc())
			nf.WriteString(m.tabFunc())
			nf.WriteString(m.fieldInfosFunc())
			nf.WriteString(m.autoIncFieldFunc())
			nf.WriteString(m.primaryKeyFunc())
			nf.WriteString(m.uniqueKeysFunc())
			// nf.WriteString(m.relationsFunc())
			nf.WriteString(m.modelMarshalJSONFunc())
			nf.WriteString(m.modelUnmarshalJSONFunc())
			nf.WriteString(m.modelListType())
			nf.WriteString(m.modelCollapseFunc())
			nf.WriteString(m.newListFunc())
			nf.WriteString(m.listNewModelFunc())
			nf.WriteString(m.listSetTotalFunc())
			nf.WriteString(m.listGetTotalFunc())
			nf.WriteString(m.listLenFunc())
			nf.WriteString(m.getInnerListFunc())
			nf.WriteString(m.listMarshalJSONFunc())
			nf.WriteString(m.listUnmarshalJSONFunc())
			nf.WriteString(m.listCollapseFunc())
			nf.WriteString(m.listFilterFunc())
			nf.WriteString(m.listCheckDupFunc())
			nf.WriteString(m.listSliceFunc())
		}
		nf.Sync()
	}
	cmd := exec.Command("gofmt", "-w", *dir)
	if err := cmd.Run(); err != nil {
		log.Println(nbcolor.Red(err))
	}
}

type ConstraintType string

const (
	PrimaryKey ConstraintType = "PrimaryKey"
	UniqueKey  ConstraintType = "UniqeuKey"
)

type Constraint struct {
	Name    string
	Type    ConstraintType
	Columns []*Column
	Expr    string
}

type Column struct {
	Name      string
	Type      string
	IsInc     bool
	FieldName string
	FieldType string
}

type Table struct {
	DB        string
	Tab       string
	ModelName string
	Columns   []*Column
	Pk        *Constraint
	Uk        []*Constraint
}

func getTabs(db *sql.DB, dbname string, tableNames ...string) []*Table {
	if len(tableNames) > 0 {
		l := make([]*Table, 0, len(tableNames))
		for _, tabName := range tableNames {
			l = append(l, &Table{DB: dbname, Tab: tabName})
		}
		return l
	}
	stmt := fmt.Sprintf("SHOW TABLES FROM `%s`", dbname)
	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	l := make([]*Table, 0, 16)
	for rows.Next() {
		tab := Table{DB: dbname}
		if err := rows.Scan(&tab.Tab); err != nil {
			panic(err)
		}
		l = append(l, &tab)
	}
	return l
}

func getCols(db *sql.DB, tables []*Table) {
	for _, tab := range tables {
		stmt := "SELECT COLUMN_NAME, DATA_TYPE, EXTRA FROM information_schema.columns WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? ORDER BY ORDINAL_POSITION"
		rows, err := db.Query(stmt, tab.DB, tab.Tab)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var col, typ, extra string
			if err := rows.Scan(&col, &typ, &extra); err != nil {
				panic(err)
			}
			column := Column{
				Name: col,
				Type: typ,
			}
			if extra == "auto_increment" {
				column.IsInc = true
			}
			tab.Columns = append(tab.Columns, &column)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}
	}
}

func getPK(db *sql.DB, tables []*Table) {
	for _, tab := range tables {
		stmt := "SELECT COLUMN_NAME FROM information_schema.key_column_usage WHERE CONSTRAINT_NAME = 'PRIMARY' AND TABLE_SCHEMA = ? AND TABLE_NAME = ? ORDER BY ORDINAL_POSITION"
		rows, err := db.Query(stmt, tab.DB, tab.Tab)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		pk := &Constraint{
			Name:    "PRIMARY",
			Type:    PrimaryKey,
			Columns: make([]*Column, 0, 8),
		}
	OUTER:
		for rows.Next() {
			var colName string
			if err := rows.Scan(&colName); err != nil {
				panic(err)
			}
			for _, col := range tab.Columns {
				if col.Name == colName {
					pk.Columns = append(pk.Columns, col)
					continue OUTER
				}
			}
			panic(fmt.Errorf("cannot find column for primary key(%s)", colName))
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}
		tab.Pk = pk
	}
}

func getUK(db *sql.DB, tables []*Table) {
	for _, tab := range tables {
		stmt := "SELECT CONSTRAINT_NAME FROM information_schema.table_constraints WHERE TABLE_SCHEMA = ? and TABLE_NAME = ? AND CONSTRAINT_TYPE = 'UNIQUE'"
		rows, err := db.Query(stmt, tab.DB, tab.Tab)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var conName string
			if err := rows.Scan(&conName); err != nil {
				panic(err)
			}
			uk := &Constraint{
				Name:    conName,
				Type:    UniqueKey,
				Columns: make([]*Column, 0, 8),
			}
			tab.Uk = append(tab.Uk, uk)
			st := `SELECT COLUMN_NAME FROM information_schema.key_column_usage WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? AND CONSTRAINT_NAME = ? ORDER BY ORDINAL_POSITION`
			nr, err := db.Query(st, tab.DB, tab.Tab, conName)
			if err != nil {
				panic(err)
			}
			defer nr.Close()
		MID:
			for nr.Next() {
				var colName string
				if err := nr.Scan(&colName); err != nil {
					panic(err)
				}
				for _, col := range tab.Columns {
					if col.Name == colName {
						uk.Columns = append(uk.Columns, col)
						continue MID
					}
				}
				panic(fmt.Errorf("cannot find column for unique key (%s)", colName))
			}
			if err := nr.Err(); err != nil {
				panic(err)
			}
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}
	}
}

var colToField = map[string]string{
	"varchar":    "nborm.String",
	"bigint":     "nborm.Int",
	"longtext":   "nborm.String",
	"datetime":   "nborm.Datetime",
	"int":        "nborm.Int",
	"tinyint":    "nborm.Int",
	"decimal":    "nborm.Decimal",
	"double":     "nborm.Decimal",
	"text":       "nborm.String",
	"timestamp":  "nborm.Datetime",
	"char":       "nborm.String",
	"smallint":   "nborm.Int",
	"float":      "nborm.Decimal",
	"date":       "nborm.Date",
	"json":       "nborm.String",
	"time":       "nborm.Datetime",
	"enum":       "nborm.String",
	"tinytext":   "nborm.String",
	"mediumtext": "nborm.String",
}

func procTables(tables []*Table) {
	for _, tab := range tables {
		tab.ModelName = strings.Replace(strings.Title(strings.Replace(tab.Tab, "_", " ", -1)), " ", "", -1)
		for _, col := range tab.Columns {
			col.FieldName = strings.Replace(strings.Title(strings.Replace(col.Name, "_", " ", -1)), " ", "", -1)
			var ok bool
			col.FieldType, ok = colToField[col.Type]
			if !ok {
				panic(fmt.Errorf("unsupport column type (%s)", col.Type))
			}
		}
		pkFields := make([]string, 0, len(tab.Pk.Columns))
		for _, col := range tab.Pk.Columns {
			pkFields = append(pkFields, col.FieldName)
		}
		tab.Pk.Expr = strings.Join(pkFields, ",")
		for _, uk := range tab.Uk {
			ukFields := make([]string, 0, len(uk.Columns))
			for _, col := range uk.Columns {
				ukFields = append(ukFields, col.FieldName)
			}
			uk.Expr = strings.Join(ukFields, ",")
		}
	}
}

func createFile(tables []*Table) string {
	s, err := nbfmt.Fmt(`
	package model
	
	import (
		"github.com/wangjun861205/nborm"
	)
	
	{{ for _, tab in Tables }}
	//db:{{ tab.DB }}
	//tab:{{ tab.Tab }}
	//pk:{{ tab.Pk.Expr }}
	{{ for _, uk in tab.Uk }}
	//uk:{{ uk.Expr }}
	{{ endfor }}
		type {{ tab.ModelName }} struct {
			nborm.Meta
			{{ for _, col in tab.Columns }}
				{{ col.FieldName }} {{ col.FieldType }} `+"`col:\"{{ col.Name }}\"{{ if col.IsInc == true }} auto_increment:\"true\"{{ endif }}`"+`
			{{ endfor }}
		}
	{{ endfor }}
	`, map[string]interface{}{"Tables": tables})
	if err != nil {
		panic(err)
	}
	return s
}

func createFileMatchAllDB(tables []*Table) string {
	s, err := nbfmt.Fmt(`
	package model
	
	import (
		"github.com/wangjun861205/nborm"
	)
	
	{{ for _, tab in Tables }}
	//db:*
	//tab:{{ tab.Tab }}
	//pk:{{ tab.Pk.Expr }}
	{{ for _, uk in tab.Uk }}
	//uk:{{ uk.Expr }}
	{{ endfor }}
		type {{ tab.ModelName }} struct {
			nborm.Meta
			{{ for _, col in tab.Columns }}
				{{ col.FieldName }} {{ col.FieldType }} `+"`col:\"{{ col.Name }}\"{{ if col.IsInc == true }} auto_increment:\"true\"{{ endif }}`"+`
			{{ endfor }}
		}
	{{ endfor }}
	`, map[string]interface{}{"Tables": tables})
	if err != nil {
		panic(err)
	}
	return s
}

func createModels(tables []*Table) string {
	s, err := nbfmt.Fmt(`
	{{ for _, tab in Tables }}
	//db:{{ tab.DB }}
	//tab:{{ tab.Tab }}
	//pk:{{ tab.Pk.Expr }}
	{{ for _, uk in tab.Uk }}
	//uk:{{ uk.Expr }}
	{{ endfor }}
		type {{ tab.ModelName }} struct {
			nborm.Meta
			{{ for _, col in tab.Columns }}
				{{ col.FieldName }} {{ col.FieldType }} `+"`col:\"{{ col.Name }}\"{{ if col.IsInc == true }} auto_increment:\"true\"{{ endif }}`"+`
			{{ endfor }}
		}
	{{ endfor }}
	`, map[string]interface{}{"Tables": tables})
	if err != nil {
		panic(err)
	}
	return s
}

func createModelsMatchAllDB(tables []*Table) string {
	s, err := nbfmt.Fmt(`
	{{ for _, tab in Tables }}
	//db:*
	//tab:{{ tab.Tab }}
	//pk:{{ tab.Pk.Expr }}
	{{ for _, uk in tab.Uk }}
	//uk:{{ uk.Expr }}
	{{ endfor }}
		type {{ tab.ModelName }} struct {
			nborm.Meta
			{{ for _, col in tab.Columns }}
				{{ col.FieldName }} {{ col.FieldType }} `+"`col:\"{{ col.Name }}\"{{ if col.IsInc == true }} auto_increment:\"true\"{{ endif }}`"+`
			{{ endfor }}
		}
	{{ endfor }}
	`, map[string]interface{}{"Tables": tables})
	if err != nil {
		panic(err)
	}
	return s
}
