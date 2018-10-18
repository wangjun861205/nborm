package nborm

type Where struct {
	str string
}

func (w *Where) String() string {
	return w.str
}

func (w *Where) And(other *Where) *Where {
	if w == nil && other != nil {
		return other
	} else if w != nil && other == nil {
		return w
	} else if w == nil && other == nil {
		return nil
	}
	return &Where{w.str + " AND " + other.str}
}

func (w *Where) Or(other *Where) *Where {
	if w == nil && other != nil {
		return other
	} else if w != nil && other == nil {
		return w
	} else if w == nil && other == nil {
		return nil
	}
	return &Where{w.str + " OR " + other.str}
}

func Group(w *Where) *Where {
	if w == nil {
		return nil
	}
	return &Where{"(" + w.str + ")"}
}
