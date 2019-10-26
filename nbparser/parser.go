package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
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

type RelInfoElem struct {
	DstModel  *RelModel
	On        string
	Fields    []*RelField
	IsMid     bool
	Index     int
	FieldsStr string
}

type RelInfo struct {
	FieldName string
	FieldType string
	IsList    bool
	Elems     []*RelInfoElem
}

const (
	neutral    = "neutral"
	on         = "on"
	field      = "field"
	value      = "value"
	quote      = "quote"
	modelName  = "modelName"
	modelIndex = "modelIndex"
)

func parseRelation(tag string) error {
	lastModel := modelInfos[len(modelInfos)-1]
	info := lastModel.RelInfos[len(lastModel.RelInfos)-1]
	statStack := make([]string, 0, 8)
	statStack = append(statStack, neutral)
	// popStack := func() string {
	// 	var stat string
	// 	statStack, stat = statStack[:len(statStack)-1], statStack[len(statStack)-1]
	// 	return stat
	// }
	// pushStat := func(stat string) {
	// 	statStack = append(statStack, stat)
	// }
	indexStat := func(index int) string {
		if index >= 0 {
			return statStack[index]
		}
		return statStack[len(statStack)+index]
	}
	var builder strings.Builder
	reader := strings.NewReader(tag)

	lastElem := func() *RelInfoElem {
		return info.Elems[len(info.Elems)-1]
	}

	lastField := func() *RelField {
		e := lastElem()
		return e.Fields[len(e.Fields)-1]
	}

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch r {
		case '[':
			switch statStack[len(statStack)-1] {
			case neutral:
				statStack = append(statStack, on)
				info.Elems = append(info.Elems, &RelInfoElem{
					DstModel: new(RelModel),
				})
			case modelName:
				statStack[len(statStack)-1] = on
				info.Elems = append(info.Elems, &RelInfoElem{
					DstModel: new(RelModel),
				})
			case modelIndex:
				statStack = statStack[:len(statStack)-1]
				statStack[len(statStack)-1] = on
				info.Elems = append(info.Elems, &RelInfoElem{
					DstModel: new(RelModel),
				})
			case quote:
				builder.WriteRune(r)
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case ']':
			switch statStack[len(statStack)-1] {
			case value, field:
				statStack = []string{neutral}
				lastElem := info.Elems[len(info.Elems)-1]
				lastElem.On = builder.String()
				builder.Reset()
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case '\'':
			switch statStack[len(statStack)-1] {
			case value:
				statStack = append(statStack, quote)
				builder.WriteRune(r)
			case quote:
				statStack = statStack[:len(statStack)-1]
				builder.WriteRune(r)
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case '@':
			switch statStack[len(statStack)-1] {
			case neutral:
				statStack = append(statStack, modelName)
				lastElem := info.Elems[len(info.Elems)-1]
				lastElem.DstModel = new(RelModel)
			case modelName:
				builder.WriteRune(r)
				switch statStack[len(statStack)-2] {
				case neutral:
					lastElem := info.Elems[len(info.Elems)-1]
					lastElem.DstModel.ModelName = "@"
					lastElem.DstModel.ModelType = "@"
					statStack = []string{neutral}
				case on:
					lastElem := info.Elems[len(info.Elems)-1]
					lastField := lastElem.Fields[len(lastElem.Fields)-1]
					lastField.Model.ModelName = "@"
				default:
					panic(fmt.Errorf("invalid rel tag(%s)", tag))
				}
				if statStack[len(statStack)-2] == neutral {
				}
			case quote:
				builder.WriteRune(r)
			case on:
				statStack = append(statStack, modelName)
				builder.WriteRune(r)
				lastElem := info.Elems[len(info.Elems)-1]
				lastElem.Fields = append(lastElem.Fields, &RelField{
					Model: new(RelModel),
				})
			case value:
				switch statStack[len(statStack)-3] {
				case modelName:
					statStack = statStack[:len(statStack)-2]
					builder.WriteRune(r)
				case modelIndex:
					statStack = statStack[:len(statStack)-3]
					builder.WriteRune(r)
				default:
					panic(fmt.Errorf("invalid rel tag(%s)", tag))
				}
				lastElem := info.Elems[len(info.Elems)-1]
				lastElem.Fields = append(lastElem.Fields, &RelField{
					Model: new(RelModel),
				})
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case '$':
			switch statStack[len(statStack)-1] {
			case modelName:
				builder.WriteRune(r)
				switch statStack[len(statStack)-2] {
				case neutral:
					lastElem := info.Elems[len(info.Elems)-1]
					lastElem.DstModel.ModelName = "$"
					lastElem.DstModel.ModelType = "$"
					statStack = []string{neutral}
				case on:
					lastElem := info.Elems[len(info.Elems)-1]
					lastField := lastElem.Fields[len(lastElem.Fields)-1]
					lastField.Model.ModelName = "$"
				default:
					panic(fmt.Errorf("invalid rel tag(%s)", tag))
				}
			case quote:
				builder.WriteRune(r)
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case '<', '>', '^':
			switch indexStat(-1) {
			case value, quote:
				builder.WriteRune(r)
			case modelName:
				if indexStat(-2) == on {
					builder.WriteRune(r)
					f := lastField()
					f.Model.ModelName += string(r)
				} else {
					panic(fmt.Errorf("invalid rel tag(%s)", tag))
				}
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case '.':
			switch statStack[len(statStack)-1] {
			case quote:
				builder.WriteRune(r)
			case modelName:
				if statStack[len(statStack)-2] == on {
					statStack = append(statStack, field)
					builder.WriteRune(r)
				} else {
					panic(fmt.Errorf("invalid rel tag(%s)", tag))
				}
			case modelIndex:
				if statStack[len(statStack)-3] == on {
					statStack = append(statStack, field)
					lastElem := info.Elems[len(info.Elems)-1]
					lastElem.Fields = append(lastElem.Fields, new(RelField))
					builder.WriteRune(r)
				} else {
					panic(fmt.Errorf("invalid rel tag(%s)", tag))
				}
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case '_':
			switch statStack[len(statStack)-1] {
			case value:
				builder.WriteRune(r)
			case modelName:
				statStack = append(statStack, modelIndex)
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			switch statStack[len(statStack)-1] {
			case value, quote:
				builder.WriteRune(r)
			case modelIndex:
				switch statStack[len(statStack)-2] {
				case neutral:
					lastElem := info.Elems[len(info.Elems)-1]
					lastElem.DstModel.ModelName += string(r)
				case on:
					lastElem := info.Elems[len(info.Elems)-1]
					lastField := lastElem.Fields[len(lastElem.Fields)-1]
					lastField.Model.ModelName += string(r)
				}
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))

			}
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			switch statStack[len(statStack)-1] {
			case value, quote:
				builder.WriteRune(r)
			case modelName:
				switch statStack[len(statStack)-2] {
				case neutral:
					lastElem := info.Elems[len(info.Elems)-1]
					lastElem.DstModel.ModelName += string(r)
					lastElem.DstModel.ModelType += string(r)
				case on:
					builder.WriteRune(r)
					lastElem := info.Elems[len(info.Elems)-1]
					lastField := lastElem.Fields[len(lastElem.Fields)-1]
					lastField.Model.ModelName += string(r)
				default:
					panic(fmt.Errorf("invalid rel tag(%s)", tag))
				}
			case field:
				builder.WriteRune(r)
				lastElem := info.Elems[len(info.Elems)-1]
				lastField := lastElem.Fields[len(lastElem.Fields)-1]
				lastField.FieldName += string(r)
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		case ' ':
			switch statStack[len(statStack)-1] {
			case value, quote:
				builder.WriteRune(r)
			case modelName:
				if statStack[len(statStack)-2] == neutral {
					statStack = []string{neutral}
				}
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			case modelIndex:
				if statStack[len(statStack)-3] == neutral {
					statStack = []string{neutral}
				}
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			case field:
				builder.WriteRune(r)
				statStack = append(statStack, value)
			case neutral:
				continue
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		default:
			switch statStack[len(statStack)-1] {
			case value, quote:
				builder.WriteRune(r)
			case field:
				builder.WriteRune(r)
				statStack = append(statStack, value)
			default:
				panic(fmt.Errorf("invalid rel tag(%s)", tag))
			}
		}
	}
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
	var midIndex int
	for _, relInfo := range m.RelInfos {
		for i, elem := range relInfo.Elems {
		INNER:
			for _, field := range elem.Fields {
				switch field.Model.ModelName {
				case "@":
					field.Model.ModelName = "m"
					field.Model.ModelType = m.Name
				case "$":
					field.Model.ModelName = fmt.Sprintf("m.%s", relInfo.FieldName)
					field.Model.ModelType = relInfo.FieldType
				case "<":
					field.Model = relInfo.Elems[i-1].DstModel
				case ">":
					field.Model = relInfo.Elems[i+1].DstModel
				case "^":
					field.Model = elem.DstModel
				default:
					for _, e := range relInfo.Elems {
						if field.Model.ModelName == e.DstModel.ModelName {
							field.Model = e.DstModel
							continue INNER
						}
					}
					panic(fmt.Errorf("cannot find model (%s)", field.Model.ModelName))
				}
			}
		}
		for _, elem := range relInfo.Elems {
			switch elem.DstModel.ModelName {
			case "@":
				elem.DstModel.ModelName = "m"
				elem.DstModel.ModelType = m.Name
			case "$", "":
				elem.DstModel.ModelName = fmt.Sprintf("m.%s", relInfo.FieldName)
				elem.DstModel.ModelType = relInfo.FieldType
			default:
				elem.IsMid = true
				elem.Index = midIndex
				elem.DstModel.ModelName = fmt.Sprintf("mm%d", elem.Index)
				midIndex++
			}
			elem.On = regexp.MustCompile(`@[@\$<>\^\w]+\.[\w]+`).ReplaceAllString(elem.On, "@")
		}
		for _, elem := range relInfo.Elems {
			var builder strings.Builder
			for _, field := range elem.Fields {
				builder.WriteString(fmt.Sprintf("&%s.%s, ", field.Model.ModelName, field.FieldName))
			}
			elem.FieldsStr = strings.Trim(builder.String(), ", ")
		}
	}
}

func (m *ModelInfo) newModelFunc() string {
	s, err := nbfmt.Fmt(`
	func New{{ model.Name }}() *{{ model.Name }} {
		m := &{{ model.Name }}{}
		m.Init(m, nil, nil)
		{{ for i, field in model.FieldInfos }}
			m.{{ field.Field }}.Init(m, "{{ field.Col }}", "{{ field.Field }}", {{ i }})
		{{ endfor }}
		m.InitRel()
		return m
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) newSubModelFunc() string {
	s, err := nbfmt.Fmt(`
	func newSub{{ model.Name }}(parent nborm.Model) *{{ model.Name }} {
		m := &{{ model.Name }}{}
		m.Init(m, parent, nil)
		{{ for i, field in model.FieldInfos }}
			m.{{ field.Field }}.Init(m, "{{ field.Col }}", "{{ field.Field }}", {{ i }})
		{{ endfor }}
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
			m.{{ rel.FieldName }} = newSub{{ rel.FieldType }}(m)
			{{ if rel.IsList == true }}
				m.{{ rel.FieldName }}.dupMap = make(map[string]int)
			{{ endif }}
			var relInfo{{ i }} *nborm.RelationInfo
			{{ for _, elem in rel.Elems }}
				{{ if elem.IsMid == true }}
					mm{{ elem.Index }} := newSub{{ elem.DstModel.ModelType }}(m)
					relInfo{{ i }} = relInfo{{ i }}.Append("{{ rel.FieldName }}", mm{{ elem.Index }}, nborm.NewExpr("{{ elem.On }}", {{ elem.FieldsStr }}))
				{{ else }}
					relInfo{{ i }} = relInfo{{ i }}.Append("{{ rel.FieldName }}", {{ elem.DstModel.ModelName }}, nborm.NewExpr("{{ elem.On }}", {{ elem.FieldsStr }}))
				{{ endif }}
			{{ endfor }}
			m.AppendRelation(relInfo{{ i }})
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
			{{ for i, field in model.FieldInfos }}
				{"{{ field.Col }}", "{{ field.Field }}", &m.{{ field.Field }}, {{ i }} },
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

func (m *ModelInfo) modelMarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (m {{ model.Name }}) MarshalJSON() ([]byte, error) {
		if !m.IsSynced() {
			return []byte("null"), nil
		}
		buffer := bytes.NewBuffer(make([]byte, 0, 1024))
		buffer.WriteString("{\n\"Aggs\": ")
		metaB, err := json.MarshalIndent(m.Meta, "", "\t")
		if err != nil {
			return nil, err
		}
		buffer.Write(metaB)
		{{ for _, field in model.FieldInfos }}
			if m.{{ field.Field }}.IsValid() {
				buffer.WriteString(",\n\"{{ field.Field }}\": ")
				{{ field.Field }}B, err := json.MarshalIndent(m.{{ field.Field }}, "", "\t")
				if err != nil {
					return nil, err
				}
				buffer.Write({{ field.Field }}B)
			}
		{{ endfor }}
		{{ for _, relInfo in model.RelInfos }}
			{{ if relInfo.IsList == true }}
				if m.{{ relInfo.FieldName}} != nil && m.{{ relInfo.FieldName }}.Len() > 0 {
					buffer.WriteString(",\n\"{{ relInfo.FieldName }}\": ")
					{{ relInfo.FieldName }}B, err := json.MarshalIndent(m.{{ relInfo.FieldName }}, "", "\t")
					if err != nil {
						return nil, err
					}
					buffer.Write({{ relInfo.FieldName }}B)
				}
			{{ else }}
				if m.{{ relInfo.FieldName }} != nil && m.{{ relInfo.FieldName }}.IsSynced() {
					buffer.WriteString(",\n\"{{ relInfo.FieldName }}\": ")
					{{ relInfo.FieldName }}B, err := json.MarshalIndent(m.{{ relInfo.FieldName }}, "", "\t")
					if err != nil {
						return nil, err
					}
					buffer.Write({{ relInfo.FieldName }}B)
				}
			{{ endif }}
		{{ endfor }}
		buffer.WriteString("\n}")
		return buffer.Bytes(), nil
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelStringMethod() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) String() string {
		b, _ := json.Marshal(m)
		return string(b)
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

// func (m *ModelInfo) modelUnmarshalJSONFunc() string {
// 	s, err := nbfmt.Fmt(`
// 	func (m *{{ model.Name }}) UnmarshalJSON(data []byte) error {
// 		return nborm.UnmarshalModel(data, m)
// 	}
// 	`, map[string]interface{}{"model": m})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }

func (m *ModelInfo) modelCollapseFunc() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) Collapse() {
		{{ for _, relInfo in model.RelInfos }}
			if m.{{ relInfo.FieldName }} != nil && m.{{ relInfo.FieldName }}.IsSynced() {
				m.{{ relInfo.FieldName }}.Collapse()
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
		{{ model.Name }} `+"`json:\"-\"`"+`
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
				builder.WriteString(fmt.Sprintf("%v", lastModel.{{ f.Field }}.AnyValue()))
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
		l.Init(l, nil, nil)
		{{ for i, field in model.FieldInfos }}
			l.{{ field.Field }}.Init(l, "{{ field.Col }}", "{{ field.Field }}", {{ i }})
		{{ endfor }}
		l.InitRel()
		return l
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) newSubListFunc() string {
	s, err := nbfmt.Fmt(`
	func newSub{{ model.Name }}List(parent nborm.Model) *{{ model.Name }}List {
		l := &{{ model.Name }}List {
			{{ model.Name }}{},
			make(map[string]int),
			make([]*{{ model.Name }}, 0, 32),
			0,
		}
		l.Init(l, parent, nil)
		{{ for i, field in model.FieldInfos }}
			l.{{ field.Field }}.Init(l, "{{ field.Col }}", "{{ field.Field }}", {{ i }})
		{{ endfor }}
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
		m.Init(m, nil, l)
		l.CopyAggs(m)
		{{ for i, field in model.FieldInfos }}
			m.{{ field.Field }}.Init(m, "{{ field.Col }}", "{{ field.Field }}", {{ i }})
			l.{{ field.Field }}.CopyStatus(&m.{{ field.Field }})
		{{ endfor }}
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
		bs := make([]byte, 0, 1024)
		bs = append(bs, []byte("{")...)
		ListB, err := json.MarshalIndent(l.List, "", "\t")
		if err != nil {
			return nil, err
		}
		ListB = append([]byte("\"List\": "), ListB...)
		bs = append(bs, ListB...)
		bs = append(bs, []byte(", ")...)
		TotalB, err := json.MarshalIndent(l.Total, "", "\t")
		if err != nil {
			return nil, err
		}
		TotalB = append([]byte("\"Total\": "), TotalB...)
		bs = append(bs, TotalB...)
		bs = append(bs, []byte("}")...)
		return bs, nil
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) listStringMethod() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) String() string {
		b, _ := json.Marshal(l)
		return string(b)
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

// func (m *ModelInfo) listUnmarshalJSONFunc() string {
// 	s, err := nbfmt.Fmt(`
// 	func (l *{{ model.Name }}List) UnmarshalJSON(b []byte) error {
// 		return nborm.UnmarshalModel(b, l)
// 	}
// 	`, map[string]interface{}{"model": m})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }

func (m *ModelInfo) listUnmarshalJSONFunc() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) UnmarshalJSON(b []byte) error {
		if string(b) == "[]" {
			return nil
		}
		jl := struct {
			List *[]*{{ model.Name }}
			Total *int
		} {
			&l.List,
			&l.Total,
		}
		return json.Unmarshal(b, &jl)
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
						l.List[idx].{{ relInfo.FieldName }}.checkDup()
						l.List[idx].{{ relInfo.FieldName }}.List = append(l.List[idx].{{ relInfo.FieldName }}.List, l.List[l.Len()-1].{{ relInfo.FieldName }}.List...)
					{{ else }}
						l.List[idx].{{ relInfo.FieldName }} = l.List[l.Len()-1].{{ relInfo.FieldName }}
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

func (m *ModelInfo) modelCacheElemType() string {
	s, err := nbfmt.Fmt(`
	type {{ model.Name }}CacheElem struct {
		hashValue string
		model *{{ model.Name }}
		modifyTime time.Time
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelListCacheElemType() string {
	s, err := nbfmt.Fmt(`
	type {{ model.Name }}ListCacheElem struct {
		hashValue string
		list *{{ model.Name }}List
		modifyTime time.Time
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelCacheManagerType() string {
	s, err := nbfmt.Fmt(`
	type {{ model.Name }}CacheManager struct {
		container map[string]*{{ model.Name }}CacheElem
		query chan string
		in chan *{{ model.Name }}CacheElem
		out chan *{{ model.Name }}CacheElem
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) newModelCacheManagerFunc() string {
	s, err := nbfmt.Fmt(`
	func new{{ model.Name }}CacheManager() *{{ model.Name }}CacheManager {
		return &{{ model.Name }}CacheManager{
			make(map[string]*{{ model.Name }}CacheElem),
			make(chan string),
			make(chan *{{ model.Name }}CacheElem),
			make(chan *{{ model.Name }}CacheElem),
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelCacheManagerRunMethod() string {
	s, err := nbfmt.Fmt(`
	func (mgr *{{ model.Name }}CacheManager) run() {
		for {
			select {
				case h := <-mgr.query:
					mgr.out <- mgr.container[h]
				case elem := <-mgr.in:
					mgr.container[elem.hashValue] = elem
			}
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelListCacheMangagerType() string {
	s, err := nbfmt.Fmt(`
	type {{ model.Name }}ListCacheManager struct {
		container map[string]*{{ model.Name }}ListCacheElem
		query chan string
		in chan *{{ model.Name }}ListCacheElem
		out chan *{{ model.Name }}ListCacheElem
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) newModelListCacheManagerFunc() string {
	s, err := nbfmt.Fmt(`
	func new{{ model.Name }}ListCacheManager() *{{ model.Name }}ListCacheManager {
		return &{{ model.Name }}ListCacheManager {
			make(map[string]*{{ model.Name }}ListCacheElem),
			make(chan string),
			make(chan *{{ model.Name }}ListCacheElem),
			make(chan *{{ model.Name }}ListCacheElem),
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelListCacheManagerRunMethod() string {
	s, err := nbfmt.Fmt(`
	func (mgr *{{ model.Name }}ListCacheManager) run() {
		for {
			select {
			case h := <-mgr.query:
				mgr.out <- mgr.container[h]
			case elem := <-mgr.in:
				mgr.container[elem.hashValue] = elem
			}
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) globalModelCacheVar() string {
	s, err := nbfmt.Fmt(`
	var {{ model.Name }}Cache = new{{ model.Name }}CacheManager()
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) globalModelListCacheVar() string {
	s, err := nbfmt.Fmt(`
	var {{ model.Name }}ListCache = new{{ model.Name }}ListCacheManager()
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func initFunc(l []*ModelInfo) string {
	s, err := nbfmt.Fmt(`
	func init() {
		{{ for _, m in modelList }}
			go {{ m.Name }}Cache.run()
			go {{ m.Name }}ListCache.run()
		{{ endfor }}
	}
	`, map[string]interface{}{"modelList": l})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelGetCacheMethod() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) GetCache(hashVal string, timeout time.Duration) bool {
		{{ model.Name }}Cache.query <- hashVal
		elem := <-{{ model.Name }}Cache.out
		if elem == nil || time.Since(elem.modifyTime) > timeout {
			return false
		}
		*m = *elem.model
		return true
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelSetCacheMethod() string {
	s, err := nbfmt.Fmt(`
	func (m *{{ model.Name }}) SetCache(hashValue string) {
		{{ model.Name }}Cache.in <- &{{ model.Name }}CacheElem {
			hashValue,
			m,
			time.Now(),
		}
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelListGetCacheMethod() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) GetListCache(hashValue string, timeout time.Duration) bool {
		{{ model.Name }}ListCache.query <- hashValue
		elem := <-{{ model.Name }}ListCache.out
		if elem == nil || time.Since(elem.modifyTime) > timeout {
			return false
		}
		*l = *elem.list
		return true
	}
	`, map[string]interface{}{"model": m})
	if err != nil {
		panic(err)
	}
	return s
}

func (m *ModelInfo) modelListSetCacheMethod() string {
	s, err := nbfmt.Fmt(`
	func (l *{{ model.Name }}List) SetListCache(hashValue string) {
		{{ model.Name }}ListCache.in <- &{{ model.Name }}ListCacheElem {
			hashValue,
			l,
			time.Now(),
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

const configTemplete = `
{
	"User": "",
	"Password": "",
	"Host": "",
	"Port": "",
	"Database": "",
	"Tables": [],
	"AppendTables": [],
	"MatchAllDB": true
}
`

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
	flag.StringVar(&command, "command", "", "specific command")
	flag.StringVar(&configFile, "config", "config.json", "specific config file")

	dir := flag.String("p", "./", "specific parse path")
	flag.Parse()
	if command == "init" {
		configFile, err := os.OpenFile("config.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
		if _, err := configFile.WriteString(configTemplete); err != nil {
			panic(err)
		}
		return
	}
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
									if field.Tag == nil {
										field.Tag = &ast.BasicLit{}
										field.Tag.Kind = token.STRING
										field.Tag.Value = "`json:\"Aggs\"`"
										field.Tag.ValuePos = field.End() + 4
									}
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
							lastModel.RelInfos = append(lastModel.RelInfos, &RelInfo{FieldName: field.Names[0].Name, FieldType: ft.X.(*ast.Ident).String(), IsList: isList})
							tagStr := regexp.MustCompile(`rel:"(.*?)"`).FindStringSubmatch(field.Tag.Value)[1]
							if err := parseRelation(tagStr); err != nil {
								panic(err)
							}
						}
					}
					return false
				default:
					return true
				}
			})
			os.Remove(path.Join(*dir, fmt.Sprintf("%s.go", f.Name.Name)))
			modelFile, err := os.OpenFile(path.Join(*dir, fmt.Sprintf("%s.go", f.Name.Name)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				panic(err)
			}
			if err := printer.Fprint(modelFile, fs, f); err != nil {
				panic(err)
			}
			modelFile.Close()
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
			"time"
			"encoding/json"
			"bytes"
		)
		`)
		for _, m := range modelInfos {
			m.proc()
			nf.WriteString(m.newModelFunc())
			nf.WriteString(m.newSubModelFunc())
			nf.WriteString(m.initRelFunc())
			nf.WriteString(m.dbFunc())
			nf.WriteString(m.tabFunc())
			nf.WriteString(m.fieldInfosFunc())
			nf.WriteString(m.autoIncFieldFunc())
			nf.WriteString(m.primaryKeyFunc())
			nf.WriteString(m.uniqueKeysFunc())
			// nf.WriteString(m.relationsFunc())
			nf.WriteString(m.modelMarshalJSONFunc())
			// nf.WriteString(m.modelUnmarshalJSONFunc())
			nf.WriteString(m.modelListType())
			nf.WriteString(m.modelCollapseFunc())
			nf.WriteString(m.newListFunc())
			nf.WriteString(m.newSubListFunc())
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
			nf.WriteString(m.modelStringMethod())
			nf.WriteString(m.listStringMethod())

			nf.WriteString(m.modelCacheElemType())
			nf.WriteString(m.modelListCacheElemType())
			nf.WriteString(m.modelCacheManagerType())
			nf.WriteString(m.modelListCacheMangagerType())
			nf.WriteString(m.newModelCacheManagerFunc())
			nf.WriteString(m.newModelListCacheManagerFunc())
			nf.WriteString(m.modelCacheManagerRunMethod())
			nf.WriteString(m.modelListCacheManagerRunMethod())

			nf.WriteString(m.globalModelCacheVar())
			nf.WriteString(m.globalModelListCacheVar())
			nf.WriteString(m.modelGetCacheMethod())
			nf.WriteString(m.modelSetCacheMethod())
			nf.WriteString(m.modelListGetCacheMethod())
			nf.WriteString(m.modelListSetCacheMethod())

		}
		nf.WriteString(initFunc(modelInfos))
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
	"time":       "nborm.Time",
	"enum":       "nborm.String",
	"tinytext":   "nborm.String",
	"mediumtext": "nborm.String",
	"binary":     "nborm.Bytes",
	"varbinary":  "nborm.Bytes",
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
