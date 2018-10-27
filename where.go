package nborm

//Where is used for query
type Where struct {
	str string
}

//String string represention of Where
func (w *Where) String() string {
	return w.str
}

//And and operation
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

//Or or operation
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

//Group group operation
func Group(w *Where) *Where {
	if w == nil {
		return nil
	}
	return &Where{"(" + w.str + ")"}
}
