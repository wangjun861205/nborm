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

var dbRe = regexp.MustCompile(`(?m)db:(\w+)$`)
var tabRe = regexp.MustCompile(`(?m)tab:(\w+)$`)
var pkRe = regexp.MustCompile(`(?m)pk:([0-9a-zA-Z,]+)$`)
var unisRe = regexp.MustCompile(`(?m)uk:([0-9a-zA-Z,]+)$`)
var colRe = regexp.MustCompile(`col:"(\w+)"`)
var incRe = regexp.MustCompile(`auto_increment:"true"`)
var relRe = regexp.MustCompile(`rel:"(.*?)"`)

type FieldInfo struct {
	Col   string
	Field string
	IsInc bool
}

type RelInfo struct {
	Field  string
	Type   string
	Models []string
	Fields []string
}

type ModelInfo struct {
	Name       string
	DB         string
	Tab        string
	FieldInfos []*FieldInfo
	Pk         []string
	Unis       [][]string
	HasUk      bool
	Inc        string
	MidModels  []string
	RelInfos   []*RelInfo
	HasRel     bool
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
}

func (m *ModelInfo) newModelFunc() string {
	s, err := nbfmt.Fmt(`
	func New{{ model.Name }}() *{{ model.Name }} {
		m := &{{ model.Name }}{
			{{ for _, rel in model.RelInfos }}
				{{ rel.Field }}: &{{ rel.Type }}{},
			{{ endfor }}
		}
		model.InitModel(m)
		return m
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
	func (m *{{ model.Name }}) FieldInfos() model.FieldInfoList {
		return model.FieldInfoList {
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
	func (m *{{ model.Name }}) AutoIncField() model.Field {
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
	func (m *{{ model.Name }}) PrimaryKey() model.FieldList {
		return model.FieldList {
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
	func (m *{{ model.Name }}) UniqueKeys() []model.FieldList {
			{{ if model.HasUk == false }}
				return nil
			{{ else }}
				return []model.FieldList{
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
	func (m *{{ model.Name }}) Relations() model.RelationInfoList {
		{{ if model.HasRel == false }}
			return nil
		{{ else }}
			{{ for i, mm in model.MidModels }}
				mm{{ i }} := {{ mm }}{}
			{{ endfor }}
			return model.RelationInfoList{
				{{ for _, info in model.RelInfos }}
					model.RelationInfo{
						model.FieldList{
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
		if m.IsSynced() {
			return json.Marshal(m)
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
		*{{ model.Name }}
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
	func New{{ model.Name }}List(cap int) *{{ model.Name }}List {
		return &{{ model.Name }}List {
			New{{ model.Name }}(),
			make([]*{{ model.Name }}, 0, cap),
			0,
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listNewModelFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) NewModel() model.Model {
		m := New{{ model.Name }}()
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
		if l.IsSynced() {
			s := struct{
				List []*{{ model.Name }}
				Total int
			} {
				l.List,
				l.Total,
			}
			return json.Marshal(s)
		}
		return []byte("null"), nil
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
			lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("m.%s.%s", dstModel, field))
		default:
			if i%2 == 1 {
				lastModel.MidModels = append(lastModel.MidModels, strings.Split(field, ".")[0])
				lastRel.Fields = append(lastRel.Fields, fmt.Sprintf("mm%d.%s", len(lastModel.MidModels)-1, strings.Split(field, ".")[1]))
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
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		os.Remove(path.Join(*dir, "methods.go"))
	// 		panic(err)
	// 	}
	// }()
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
									lastModel.FieldInfos = append(lastModel.FieldInfos, &FieldInfo{Field: field.Names[0].Name})
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
							lastModel.RelInfos = append(lastModel.RelInfos, &RelInfo{Field: field.Names[0].Name, Type: ft.X.(*ast.Ident).String()})
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
		)
		`)
		for _, m := range modelInfos {
			m.proc()
			nf.WriteString(m.newModelFunc())
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
		}
		nf.Sync()
	}
	cmd := exec.Command("gofmt", "-w", *dir)
	if err := cmd.Run(); err != nil {
		log.Println(nbcolor.Red(err))
	}
}
