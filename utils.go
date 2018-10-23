package nborm

import (
	"database/sql"
	"fmt"
	"strings"
)

type Pager [2]int

func NewPager(numPerPage, pageNum int) *Pager {
	p := Pager([2]int{numPerPage, pageNum})
	return &p
}

func (p *Pager) toSQL() string {
	if p == nil {
		return ""
	}
	offset := (*p)[0] * ((*p)[1] - 1)
	return fmt.Sprintf("LIMIT %d, %d", offset, p[0])
}

func (p *Pager) NextPage() {
	(*p)[1]++
}

func (p *Pager) PrevPage() (ok bool) {
	if (*p)[1] != 0 {
		(*p)[1]--
		return true
	}
	return false
}

func (p *Pager) Index(i int) {
	(*p)[1] = i
}

type Sorter string

func NewSorter(infos ...string) *Sorter {
	s := Sorter(fmt.Sprintf("ORDER BY %s", strings.Join(infos, ", ")))
	return &s
}

func (s *Sorter) toSQL() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

func scanRows(l ModelList, rows *sql.Rows) error {
	defer rows.Close()
	for rows.Next() {
		m := l.New()
		fields := m.Fields()
		scanList := make([]interface{}, len(fields))
		for i, f := range fields {
			scanList[i] = f
		}
		err := rows.Scan(scanList...)
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func scanRow(m Model, row *sql.Row) error {
	fields := m.Fields()
	scanList := make([]interface{}, len(fields))
	for i, f := range fields {
		scanList[i] = f
	}
	err := row.Scan(scanList...)
	if err != nil {
		return err
	}
	return nil
}

func getPk(m Model) (pk Field, others []Field) {
	info := getTableCache(m)
	fields := m.Fields()
	pk = fields[info.pk]
	others = append(fields[:info.pk], fields[info.pk+1:]...)
	return
}

func getInc(m Model) (inc Field, others []Field) {
	info := getTableCache(m)
	fields := m.Fields()
	inc = fields[info.inc]
	others = append(fields[:info.inc], fields[info.inc+1:]...)
	return
}

func getUni(m Model) (unis []Field, others []Field) {
	fields := m.Fields()
	for _, f := range fields {
		if f.IsPk() || f.IsInc() || f.IsUni() {
			unis = append(unis, f)
		} else {
			others = append(others, f)
		}
	}
	return
}

func filterValid(fields []Field) (validFields []Field) {
	for _, f := range fields {
		if f.IsValid() {
			validFields = append(validFields, f)
		}
	}
	return
}

func getByName(m Model, name string) Field {
	info := getTableCache(m)
	fields := m.Fields()
	return fields[info.fieldMap[name]]
}

type sortObj struct {
	ModelList
	funcs []func(Model, Model) int
}

func (o *sortObj) Less(i, j int) bool {
	iMod, jMod := o.ModelList.Index(i), o.ModelList.Index(j)
	for _, f := range o.funcs {
		v := f(iMod, jMod)
		switch {
		case v < 0:
			return true
		case v > 0:
			return false
		default:
			continue
		}
	}
	return false
}

func filterList(l ModelList, f func(Model) bool) {
	var i int
	for i < l.Len() {
		if f(l.Index(i)) {
			l.Delete(i)
			continue
		}
		i++
	}
}

func updateModel(m Model, values ...*UpdateValue) {
	for _, val := range values {
		getByName(m, val.column).SetByUpdateValue(val)
	}
}

func invalidateModel(m Model) {
	for _, field := range m.Fields() {
		field.Invalidate()
	}
}

func genWhere(m Model) *Where {
	fields := m.Fields()
	info := getTableCache(m)
	switch {
	case info.inc != -1 && fields[info.inc].IsValid():
		return fields[info.inc].Where()
	case info.pk != -1 && fields[info.pk].IsValid():
		return fields[info.pk].Where()
	case len(info.unis) > 0:
		for _, index := range info.unis {
			if fields[index].IsValid() {
				return fields[index].Where()
			}
		}
		fallthrough
	default:
		var w *Where
		for _, f := range fields {
			if f.IsValid() {
				if w == nil {
					w = f.Where()
				} else {
					w = w.And(f.Where())
				}
			}
		}
		if w == nil {
			panic("nborm.genWhere() error: empty where")
		}
		return w
	}
}
