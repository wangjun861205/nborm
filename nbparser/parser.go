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
var relRe = regexp.MustCompile(`rel:"(.*?)"`)
var dstFieldRelRe = regexp.MustCompile(`(\w+)(\(.+\))?`)
var masterRelFieldRe = regexp.MustCompile(`(\w+)(\[.+?\])?\.(\w+)`)
var midWhereFieldRe = regexp.MustCompile(`@(\w+)`)

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
	Name        string
	MidWhere    MidWhere
	HasMidWhere bool
}

type RelInfo struct {
	Field  string
	Type   string
	Models []string
	Fields []string
	//=======================================
	HasAddedCond bool
	AddedCond    MidWhere
	IsList       bool
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
		{{ for _, rel in model.RelInfos }}
			m.{{ rel.Field }} = &{{ rel.Type }}{}
			m.{{ rel.Field }}.SetParent(m)
			nborm.InitModel(m.{{ rel.Field }})
			{{ if rel.HasAddedCond == true }}
				{{ for _, aw in rel.AddedConds }}
					m.{{ rel.Field }}.{{ aw.Field }}.AndWhere({{ aw.Op }}, {{ aw.Value }})
				{{ endfor }}
				m.{{ rel.Field }}.AndExprWhere(nborm.NewExpr("{{ rel.AddedCond.Expr }}", &{{ rel.Field }}.{{ rel.AddedCond.Fields }}))
			{{ endif }}
		{{ endfor }}
		{{ if model.HasMidModels == true }}
			{{ for i, mm in model.MidModels }}
				var mm{{ i }} *{{ mm.Name }}
			{{ endfor }}
			{{ for i, mm in model.MidModels }}
				mm{{ i }} = &{{ mm.Name }}{}
				mm{{ i }}.SetParent(m)
				nborm.InitModel(mm{{ i }})
				{{ if mm.HasMidWhere == true }}
					mm{{ i }}.AndExprWhere(nborm.NewExpr("{{ mm.MidWhere.Expr }}", &mm{{ i }}.{{ mm.MidWhere.Fields }}))
				{{ endif }}
				m.AppendMidTab(mm{{ i }})
			{{ endfor }}
		{{ endif }}
		m.AddRelInited()
	}`, map[string]interface{}{"model": m})
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

func (m *ModelInfo) relationsFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) Relations() nborm.RelationInfoList {
		{{ if model.HasRel == false }}
			return nil
		{{ else }}
			if !m.IsRelInited() {
				m.InitRel()
			}
			{{ if model.HasMidModels == true }}
				{{ for i, mm in model.MidModels }}
					mm{{ i }} := m.GetMidTabs()[{{ i }}].(*{{ mm.Name }})
				{{ endfor }}
			{{ endif }}
			return nborm.RelationInfoList{
				{{ for _, info in model.RelInfos }}
					nborm.RelationInfo{
						nborm.FieldList{
							{{ for _, f in info.Fields }}
								&{{ f }},
							{{ endfor }}
						},
						m.{{ info.Field }},
					},
				{{ endfor }}
			}
		{{ endif }}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelMarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (m {{ model.Name }}) MarshalJSON() ([]byte, error) {
		if m.IsSynced() || m.IsContainValue() {
			return json.Marshal(struct{
				{{ for _, f in model.FieldInfos }}
					{{ f.Field }} interface{}
				{{ endfor }}
				{{ for _, r in model.RelInfos }}
					{{ r.Field }} *{{ r.Type }}
				{{ endfor }}
			}{
				{{ for _, f in model.FieldInfos }}
					{{ f.Field }}: m.{{ f.Field }}.JSONValue(),
				{{ endfor }}
				{{ for _, r in model.RelInfos }}
					{{ r.Field }}: m.{{ r.Field }},
				{{ endfor }}
			})
		}
		return []byte("null"), nil
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
		List []*{{ model.Name }}
		Total int
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

func (m *ModelInfo) listMarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (l {{ model.Name }}List) MarshalJSON() ([]byte, error) {
		s := struct{
			List []*{{ model.Name }}
			Total int
		} {
			make([]*{{ model.Name }}, 0, 1),
			l.Total,
		}
		if l.Len() > 0 {
			s.List = l.List
		}
		return json.Marshal(s)
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
		ll := struct {
			List []*{{ model.Name }}
			Total int
		} {
			l.List,
			0,
		}
		if err := json.Unmarshal(b, &ll); err != nil {
			return err
		}
		l.List = ll.List
		l.Total = ll.Total
		return nil
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
		{{ if model.HasRel == false }}
			return
		{{ else }}
			if len(l.List) < 2 {
				return
			}
			lm := l.List[len(l.List)-2]
			rm := l.List[len(l.List)-1]
			if nborm.IsPrimaryKeyEqual(lm, rm) {
				{{ for _, relInfo in model.RelInfos }}
					{{ if relInfo.IsList == true }}
						lm.{{ relInfo.Field }}.List = append(lm.{{ relInfo.Field }}.List, rm.{{ relInfo.Field }}.List...)
					{{ else }}
						lm.{{ relInfo.Field }} = rm.{{ relInfo.Field }}
					{{ endif }}
				{{ endfor }}
				l.List = l.List[:len(l.List)-1]
			}
		{{ endif }}
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

func parseRelation(dstModel, tag string) error {
	relGroup := relRe.FindStringSubmatch(tag)
	if len(relGroup) != 2 {
		return fmt.Errorf("no rel tag: (%s)", tag)
	}
	fields := strings.Split(relGroup[1], "->")
	if len(fields) < 2 || len(fields)%2 != 0 {
		return fmt.Errorf("invalid rel tag: (%s)", tag)
	}
	lastModel := modelInfos[len(modelInfos)-1]
	lastRel := lastModel.RelInfos[len(lastModel.RelInfos)-1]
	for i, field := range fields {
		switch {
		case i == 0:
			lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("m.%s", field))
		case i == len(fields)-1:
			fieldGroup := dstFieldRelRe.FindStringSubmatch(field)
			fieldName, addedConds := fieldGroup[1], fieldGroup[2]
			condGroup := midWhereFieldRe.FindAllStringSubmatch(addedConds, -1)
			if len(condGroup) == 0 {
				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("m.%s.%s", dstModel, fieldName))
			} else {
				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("m.%s.%s", dstModel, fieldName))
				lastRel.HasAddedCond = true
				fields := make([]string, 0, len(condGroup))
				for _, c := range condGroup {
					fields = append(fields, c[1])
				}
				fieldsStr := strings.Join(fields, ", ")
				expr := midWhereFieldRe.ReplaceAllString(strings.Trim(addedConds, "[]"), "@")
				lastRel.AddedCond = MidWhere{fieldsStr, expr}
			}
		default:
			if i%2 == 1 {
				fieldGroup := masterRelFieldRe.FindStringSubmatch(field)
				modelName, midWhereStr, fieldName := fieldGroup[1], fieldGroup[2], fieldGroup[3]
				midWhereGroup := midWhereFieldRe.FindAllStringSubmatch(midWhereStr, -1)
				var midModel MidModel
				if len(midWhereGroup) == 0 {
					midModel = MidModel{
						Name: modelName,
					}
				} else {
					fields := make([]string, 0, len(midWhereGroup))
					for _, w := range midWhereGroup {
						fields = append(fields, w[1])
					}
					fieldsStr := strings.Join(fields, ", ")
					whereExpr := midWhereFieldRe.ReplaceAllString(strings.Trim(midWhereStr, "[]"), "@")
					midModel = MidModel{
						Name:        modelName,
						MidWhere:    MidWhere{fieldsStr, whereExpr},
						HasMidWhere: true,
					}
				}
				lastModel.MidModels = append(lastModel.MidModels, midModel)
				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("mm%d.%s", len(lastModel.MidModels)-1, fieldName))
			} else {
				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("mm%d.%s", len(lastModel.MidModels)-1, strings.Split(field, ".")[1]))
			}
		}
	}
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
			"encoding/json"
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
			nf.WriteString(m.relationsFunc())
			nf.WriteString(m.modelMarshalJSONFunc())
			nf.WriteString(m.modelListType())
			nf.WriteString(m.newListFunc())
			nf.WriteString(m.listNewModelFunc())
			nf.WriteString(m.listSetTotalFunc())
			nf.WriteString(m.listLenFunc())
			nf.WriteString(m.listMarshalJSONFunc())
			nf.WriteString(m.listUnmarshalJSONFunc())
			nf.WriteString(m.listCollapseFunc())
			nf.WriteString(m.listFilterFunc())
		}
		nf.Sync()
	}
	cmd := exec.Command("gofmt", "-w", *dir)
	if err := cmd.Run(); err != nil {
		log.Println(nbcolor.Red(err))
	}
}
