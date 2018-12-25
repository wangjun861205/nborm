package nborm

type Union []table

func MakeUnion(tables ...table) Union {
	return Union(tables)
}

func (u Union) addrs() []uintptr {
	addrs := make([]uintptr, len(u))
	for i, tab := range u {
		addrs[i] = getTabAddr(tab)
	}
	return addrs
}

func (u Union) tabInfos() []*tableInfo {
	tabInfos := make([]*tableInfo, len(u))
	for i, tab := range u {
		tabInfos[i] = getTabInfo(tab)
	}
	return tabInfos
}
