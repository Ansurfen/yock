package scheduler

import lua "github.com/yuin/gopher-lua"

func tableDeepCopy(L *lua.LState, tbl *lua.LTable) *lua.LTable {
	table := tbl
	newTable := &lua.LTable{}
	copyTable(L, table, newTable)
	return newTable
}

func copyTable(L *lua.LState, srcTable *lua.LTable, dstTable *lua.LTable) {
	srcTable.ForEach(func(key lua.LValue, value lua.LValue) {
		if tbl, ok := value.(*lua.LTable); ok {
			newTbl := L.NewTable()
			copyTable(L, tbl, newTbl)
			dstTable.RawSet(key, newTbl)
		} else {
			dstTable.RawSet(key, value)
		}
	})
}
