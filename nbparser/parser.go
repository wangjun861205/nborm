package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/wangjun861205/nbfmt"

	"github.com/wangjun861205/nbcolor"
)

var modelInfos = make([]*ModelInfo, 0, 128)

var dbRe = regexp.MustCompile(`(?m)db:(\w+|\*)$`)
var tabRe = regexp.MustCompile(`(?m)tab:(\w+)$`)
var pkRe = regexp.MustCompile(`(?m)pk:([0-9a-zA-Z,]+)$`)
var unisRe = regexp.MustCompile(`(?m)uk:([0-9a-zA-Z,]+)$`)
var colRe = regexp.MustCompile(`col:"(\w+)"`)
var incRe = regexp.MustCompile(`auto_increment:"true"`)

// var relRe = regexp.MustCompile(`rel:"(.*?)"`)
// var dstFieldRelRe = regexp.MustCompile(`(\w+)(\[.+\])?`)
// var masterRelFieldRe = regexp.MustCompile(`(\w+)(\[.+?\])?\.(\w+)`)
// var midWhereFieldRe = regexp.MustCompile(`@(\w+)`)

var relWholeRe = regexp.MustCompile(`rel:"(.*?)"`)
var relModelRe = regexp.MustCompile(`(\w+)(?:->|\[.*?\])`)
var relWhereRe = regexp.MustCompile(`\[(.*?)\]`)
var relMidModelRe = regexp.MustCompile(`([a-zA-Z0-9]+)(_\d+)?`)
var relFieldRe = regexp.MustCompile(`@([\w@\$]+)\.(\w+)`)

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

// type MidModel struct {
// 	Name        string
// 	MidWhere    MidWhere
// 	HasMidWhere bool
// }

// type RelInfo struct {
// 	Field  string
// 	Type   string
// 	Models []string
// 	Fields []string
// 	//=======================================
// 	HasAddedCond bool
// 	AddedCond    MidWhere
// 	IsList       bool
// }

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

type RelInfo struct {
	Field        string
	Type         string
	SrcModel     string
	SrcFields    []string
	MidModels    []*MidModel
	DstModel     string
	DstFields    []string
	Fields       []Field
	FieldsStr    string
	JoinWheres   string
	IsList       bool
	MidModelsLen int
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
			midIdx++
		}
	}
	for _, relInfo := range m.RelInfos {
		for _, field := range relInfo.Fields {
			switch field.ModelName {
			case "@":
				relInfo.FieldsStr += fmt.Sprintf("&m.%s, ", field.FieldName)
			case "$":
				relInfo.FieldsStr += fmt.Sprintf("&m.%s.%s, ", relInfo.Field, field.FieldName)
			default:
				relInfo.FieldsStr += fmt.Sprintf("&mm%d.%s, ", field.MidModel.Index, field.FieldName)
			}
		}
		relInfo.FieldsStr = strings.TrimSuffix(relInfo.FieldsStr, ", ")
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
				relInfo{{ i }}.SrcFields = append(relInfo{{ i }}.SrcFields, &m.{{ srcField }})
			{{ endfor }}
			{{ for _, dstField in rel.DstFields }}
				relInfo{{ i }}.DstFields = append(relInfo{{ i }}.DstFields, &m.{{ rel.Field }}.{{ dstField }})
			{{ endfor }}
			{{ for _, midModel in rel.MidModels }}
				mm{{ midModel.Index }} := &{{ midModel.Type }}{}
				mm{{ midModel.Index }}.SetParent(m)
				nborm.InitModel(mm{{ midModel.Index }})
				relInfo{{ i }}.MidModels = append(relInfo{{ i }}.MidModels, mm{{ midModel.Index }})
			{{ endfor }}
			relInfo{{ i }}.JoinWheres = append(relInfo{{ i }}.JoinWheres, nborm.NewExpr("{{ rel.JoinWheres }}", {{ rel.FieldsStr }}))
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
			if m.{{ relInfo.Field }}.IsSynced() {
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

// func parseRelation(dstModel, tag string) error {
// 	relGroup := relRe.FindStringSubmatch(tag)
// 	if len(relGroup) != 2 {
// 		return fmt.Errorf("no rel tag: (%s)", tag)
// 	}
// 	fields := strings.Split(relGroup[1], "->")
// 	if len(fields) < 2 || len(fields)%2 != 0 {
// 		return fmt.Errorf("invalid rel tag: (%s)", tag)
// 	}
// 	lastModel := modelInfos[len(modelInfos)-1]
// 	lastRel := lastModel.RelInfos[len(lastModel.RelInfos)-1]
// 	for i, field := range fields {
// 		switch {
// 		case i == 0:
// 			lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("m.%s", field))
// 		case i == len(fields)-1:
// 			fieldGroup := dstFieldRelRe.FindStringSubmatch(field)
// 			fieldName, addedConds := fieldGroup[1], fieldGroup[2]
// 			condGroup := midWhereFieldRe.FindAllStringSubmatch(addedConds, -1)
// 			if len(condGroup) == 0 {
// 				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("m.%s.%s", dstModel, fieldName))
// 			} else {
// 				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("m.%s.%s", dstModel, fieldName))
// 				lastRel.HasAddedCond = true
// 				fields := make([]string, 0, len(condGroup))
// 				for _, c := range condGroup {
// 					fields = append(fields, fmt.Sprintf("&m.%s.%s", lastRel.Field, c[1]))
// 				}
// 				fieldsStr := strings.Join(fields, ", ")
// 				expr := midWhereFieldRe.ReplaceAllString(strings.Trim(addedConds, "[]"), "@")
// 				lastRel.AddedCond = MidWhere{fieldsStr, expr}
// 			}
// 		default:
// 			if i%2 == 1 {
// 				fieldGroup := masterRelFieldRe.FindStringSubmatch(field)
// 				modelName, midWhereStr, fieldName := fieldGroup[1], fieldGroup[2], fieldGroup[3]
// 				midWhereGroup := midWhereFieldRe.FindAllStringSubmatch(midWhereStr, -1)
// 				var midModel MidModel
// 				if len(midWhereGroup) == 0 {
// 					midModel = MidModel{
// 						Name: modelName,
// 					}
// 				} else {
// 					fields := make([]string, 0, len(midWhereGroup))
// 					for _, w := range midWhereGroup {
// 						fields = append(fields, fmt.Sprintf("&%s.%s", fmt.Sprintf("mm%d", i/2), w[1]))
// 					}
// 					fieldsStr := strings.Join(fields, ", ")
// 					whereExpr := midWhereFieldRe.ReplaceAllString(strings.Trim(midWhereStr, "[]"), "@")
// 					midModel = MidModel{
// 						Name:        modelName,
// 						MidWhere:    MidWhere{fieldsStr, whereExpr},
// 						HasMidWhere: true,
// 					}
// 				}
// 				lastModel.MidModels = append(lastModel.MidModels, midModel)
// 				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("mm%d.%s", len(lastModel.MidModels)-1, fieldName))
// 			} else {
// 				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("mm%d.%s", len(lastModel.MidModels)-1, strings.Split(field, ".")[1]))
// 			}
// 		}
// 	}
// 	return nil
// }

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
		midModelGroup := relMidModelRe.FindStringSubmatch(mg[1])
		midModel := MidModel{
			Name: midModelGroup[0],
			Type: midModelGroup[1],
		}
		lastRel.MidModels = append(lastRel.MidModels, &midModel)
	}
	joinWhereGroup := relWhereRe.FindStringSubmatch(relStr)
	lastRel.MidModelsLen = len(lastRel.MidModels)
	lastRel.DstModel = dstModel
	fieldsGroup := relFieldRe.FindAllStringSubmatch(joinWhereGroup[1], -1)
	for _, fg := range fieldsGroup {
		switch fg[1] {
		case "@":
			lastRel.SrcFields = append(lastRel.SrcFields, fg[2])
			lastRel.Fields = append(lastRel.Fields, Field{nil, "@", fg[2]})
		case "$":
			lastRel.DstFields = append(lastRel.DstFields, fg[2])
			lastRel.Fields = append(lastRel.Fields, Field{nil, "$", fg[2]})
		default:
			midModName, midFieldName := fg[1], fg[2]
			var midMod *MidModel
			for i, midModel := range lastRel.MidModels {
				if midModel.Name == midModName {
					midMod = lastRel.MidModels[i]
				}
			}
			if midMod == nil {
				panic(fmt.Errorf("invalid middle model name for field(%s)", midModName))
			}
			lastRel.Fields = append(lastRel.Fields, Field{midMod, "", midFieldName})
		}
	}
	lastRel.JoinWheres = relFieldRe.ReplaceAllString(joinWhereGroup[1], "@")
	return nil
}

func main() {
	dir := flag.String("p", "./", "specific parse path")
	flag.Parse()
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
