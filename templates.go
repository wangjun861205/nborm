package nborm

const template = `
package {{ pkgName }}

import (
	"github.com/wangjun861205/nborm"
	"encoding/json"
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
			{{ for _, oto in tab.OneToOnes }}
				m.{{ oto.DstTab.ModelName }} = nborm.NewOneToOne(m.{{ oto.SrcCol.FieldName }}, "{{ oto.DstCol.Name }}", New{{ oto.DstTab.ModelName }})
			{{ endfor }}
			{{ for _, fk in tab.ForeignKeys }}
				m.{{ fk.DstTab.ModelName }} = nborm.NewForeignKey(m.{{ fk.SrcCol.FieldName }}, "{{ fk.DstCol.Name }}", New{{ fk.DstTab.ModelName }})
			{{ endfor }}
			{{ for _, rfk in tab.ReverseForeignKeys }}
				m.{{ rfk.DstTab.ModelName }} = nborm.NewReverseForeignKey(m.{{ rfk.SrcCol.FieldName }}, "{{ rfk.DstCol.Name }}", New{{ rfk.DstTab.ModelName }}List)
			{{ endfor }}
			{{ for i, mtm in tab.ManyToManys }}
				mtmMidMod{{ i }} := New{{ mtm.MidTab.ModelName }}()
				m.{{ mtm.DstTab.ModelName }} = nborm.NewManyToMany(m.{{ mtm.SrcCol.FieldName }}, mtmMidMod{{ i }}.{{ mtm.MidLeftCol.FieldName }}, mtmMidMod{{ i }}.{{ mtm.MidRightCol.FieldName }}, "{{ mtm.DstCol.Name }}", New{{ mtm.DstTab.ModelName }}List)
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

		func (m *{{ tab.ModelName }}) NewList() nborm.ModelList {
			return New{{ tab.ModelName }}List()
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

		func (l *{{ tab.ModelName }}List) MarshalJSON() ([]byte, error) {
			return json.MarshalIndent(l.List, "\t", "\t")
		}

	{{ endfor }}
{{ endfor }}

func init() {
	{{ for dbName, _ in infos }}
		nborm.RegisterDB("{{ username }}", "{{ password }}", "{{ address }}", "{{ dbName }}")
	{{ endfor }}
}
`
