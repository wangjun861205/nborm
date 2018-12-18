package nborm

import (
	"unsafe"
)

type sorter struct {
	slice table
	funcs []func(iaddr, jaddr uintptr) int
}

func (s sorter) Len() int {
	return len(**(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&s.slice)) + uintptr(8)))) - 1
}

func (s sorter) Swap(i, j int) {
	l := *(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&s.slice)) + uintptr(8)))
	(*l)[i+1], (*l)[j+1] = (*l)[j+1], (*l)[i+1]
}

func (s sorter) Less(i, j int) bool {
	l := **(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&s.slice)) + uintptr(8)))
	for _, f := range s.funcs {
		v := f(l[i+1], l[j+1])
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