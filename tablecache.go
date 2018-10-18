package nborm

type tableCache struct {
	fieldMap map[string]int
	pk       int
	inc      int
	unis     []int
}

var tableCacheMap = map[string]map[string]*tableCache{}

func getTableCache(m Model) *tableCache {
	if tableCacheMap[m.DB()] == nil {
		tableCacheMap[m.DB()] = make(map[string]*tableCache)
	}
	if tableCacheMap[m.DB()][m.Tab()] == nil {
		fields := m.Fields()
		fieldMap := make(map[string]int)
		pk, inc := -1, -1
		unis := make([]int, 0, len(fields))
		for i, field := range fields {
			fieldMap[field.Column()] = i
			if field.IsPk() {
				pk = i
			}
			if field.IsInc() {
				inc = i
			}
			if field.IsUni() {
				unis = append(unis, i)
			}
		}
		tableCacheMap[m.DB()][m.Tab()] = &tableCache{fieldMap, pk, inc, unis}
	}
	return tableCacheMap[m.DB()][m.Tab()]
}

func initCacheByInfo() {
	for dbName, tabMap := range dbInfo {
		tableCacheMap[dbName] = make(map[string]*tableCache)
		for tabName, tab := range tabMap {
			fieldMap := make(map[string]int)
			pk, inc := -1, -1
			unis := make([]int, 0, 8)
			for i, col := range tab.Columns {
				fieldMap[col.Name] = i
				if col.Key == "PRI" {
					pk = i
				}
				if col.Key == "UNI" {
					unis = append(unis, i)
				}
				if col.Extra == "auto_increment" {
					inc = i
				}
			}
			cache := &tableCache{fieldMap, pk, inc, unis}
			tableCacheMap[dbName][tabName] = cache
		}
	}
}
