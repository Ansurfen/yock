package scheduler

import (
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
)

func loadPath(yocks *YockScheduler) lua.LValue {
	return yocks.registerLib(pathLib)
}

var pathLib = luaFuncs{
	"exist":    pathExist,
	"filename": pathFilename,
	"join":     pathJoin,
	"dir":      pathDir,
	"base":     pathBase,
	"clean":    pathClean,
	"ext":      pathExt,
	"abs":      pathAbs,
}

func pathExist(l *lua.LState) int {
	ok := utils.IsExist(l.CheckString(1))
	if ok {
		l.Push(lua.LTrue)
	} else {
		l.Push(lua.LFalse)
	}
	return 1
}

func pathFilename(l *lua.LState) int {
	l.Push(lua.LString(utils.Filename(l.CheckString(1))))
	return 1
}

func pathJoin(l *lua.LState) int {
	elem := []string{}
	for i := 1; i <= l.GetTop(); i++ {
		elem = append(elem, l.CheckString(i))
	}
	l.Push(lua.LString(filepath.Join(elem...)))
	return 1
}

func pathDir(l *lua.LState) int {
	l.Push(lua.LString(filepath.Dir(l.CheckString(1))))
	return 1
}

func pathBase(l *lua.LState) int {
	l.Push(lua.LString(filepath.Base(l.CheckString(1))))
	return 1
}

func pathClean(l *lua.LState) int {
	l.Push(lua.LString(filepath.Clean(l.CheckString(1))))
	return 1
}

func pathExt(l *lua.LState) int {
	l.Push(lua.LString(filepath.Ext(l.CheckString(1))))
	return 1
}

func pathAbs(l *lua.LState) int {
	abs, err := filepath.Abs(l.CheckString(1))
	l.Push(lua.LString(abs))
	handleErr(l, err)
	return 2
}
