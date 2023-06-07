// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ffi

import (
	"path/filepath"
	"reflect"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/runtime"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func LoadFFI(l *lua.LState) lua.LValue {
	yockf = New()
	ffi := &lua.LTable{}
	ffi.RawSetString("lib", yockf.tbl)
	ffi.RawSetString("library", l.NewFunction(ffiLibrary))
	return ffi
}

type yockffi struct {
	libs map[string]*yockfLib
	tbl  *lua.LTable
}

type yockfLib struct {
	instance Lib
	funcs    *lua.LTable
}

func newYockfLib(name string) *yockfLib {
	lib, err := NewLibray(name)
	if err != nil {
		panic(err)
	}
	return &yockfLib{
		instance: lib,
		funcs:    &lua.LTable{},
	}
}

func (lib *yockfLib) registerFn(name string, fn *lua.LFunction) {
	lib.funcs.RawSetString(name, fn)
}

func (lib *yockfLib) close() {
	lib.funcs = &lua.LTable{}
}

func (yockf *yockffi) Open(path string) *yockfLib {
	path = pathf(path)
	name := utils.Filename(path)
	if lib, ok := yockf.libs[name]; ok {
		return lib
	}
	lib := newYockfLib(path)
	yockf.libs[name] = lib
	yockf.tbl.RawSetString(name, lib.funcs)
	return lib
}

func (yockf *yockffi) Close(path string) {
	path = utils.Filename(path)
	if lib, ok := yockf.libs[path]; ok {
		lib.close()
		delete(yockf.libs, path)
	}
}

func (yockf *yockffi) Free() {
	for _, lib := range yockf.libs {
		lib.close()
	}
}

func pathf(path string) string {
	if filepath.Ext(path) == "" {
		switch utils.CurPlatform.OS {
		case "windows":
			path += ".dll"
		case "darwin":
			path += ".dylib"
		default:
			path += ".so"
		}
	}
	return path
}

func ffiLibrary(l *lua.LState) int {
	name := l.CheckString(1)
	lib := yockf.Open(name)

	l.CheckTable(2).ForEach(func(fn, declare lua.LValue) {
		def := runtime.UpgradeTable(declare.(*lua.LTable))
		fname := fn.String()
		var (
			go_args       []reflect.Type
			go_rtype      []reflect.Type
			ffi_args      []Type
			args          []string
			rtype_literal = def.ToString(1)
			rtype         = Archive.MustFindRecord(rtype_literal)
		)

		def.ToTable(2).ForEach(func(_, arg lua.LValue) {
			args = append(args, arg.String())
		})

		v := lib.instance.RawFunc(
			fname, rtype_literal,
			args, &go_rtype, &go_args, &ffi_args)

		n := 0
		if len(go_rtype) > 0 && go_rtype[0].Kind() != typeVoid.Kind() {
			n = 1
		}

		lib.registerFn(fname, l.NewClosure(func(l *lua.LState) int {
			lua_args := []reflect.Value{}
			for i := 1; i <= l.GetTop(); i++ {
				switch l.CheckAny(i).Type() {
				case lua.LTBool:
					lua_args = append(lua_args, reflect.ValueOf(l.CheckBool(i)))
				case lua.LTString:
					lua_args = append(lua_args, reflect.ValueOf(l.CheckString(i)))
				case lua.LTNumber:
					num := l.CheckNumber(i)
					switch go_args[i-1] {
					case typeInt:
						lua_args = append(lua_args, reflect.ValueOf(int(num)))
					case typeUint:
						lua_args = append(lua_args, reflect.ValueOf(uint(num)))
					case typeInt32:
						lua_args = append(lua_args, reflect.ValueOf(int32(num)))
					case typeInt64:
						lua_args = append(lua_args, reflect.ValueOf(int64(num)))
					case typeFloat32:
						lua_args = append(lua_args, reflect.ValueOf(float32(num)))
					case typeFloat64:
						lua_args = append(lua_args, reflect.ValueOf(float64(num)))
					}
				}
			}
			ret := v.Call(lua_args)
			if n > 0 {
				retv := ret[0]
				switch rtype.value() {
				case typeBool:
					if retv.Bool() {
						l.Push(lua.LTrue)
					} else {
						l.Push(lua.LFalse)
					}
				case typeInt:
					l.Push(lua.LNumber(retv.Int()))
				case typeStr:
					l.Push(lua.LString(retv.String()))
				case typeFloat32:
					l.Push(lua.LNumber(retv.Float()))
				case typeFloat64:
					l.Push(lua.LNumber(retv.Float()))
				case typeUnsafePointer:
					l.Push(luar.New(l, retv.Interface()))
				}
			}
			return n
		}))
	})
	return 0
}

var yockf *yockffi

func New() *yockffi {
	return &yockffi{
		tbl:  &lua.LTable{},
		libs: make(map[string]*yockfLib),
	}
}
