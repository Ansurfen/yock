package scheduler

import (
	"github.com/ansurfen/cushion/runtime"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type Boolean struct {
	v *bool
}

func (b *Boolean) Ptr() *bool {
	return b.v
}

func (b *Boolean) Var() bool {
	return *b.v
}

type String struct {
	v *string
}

func (s *String) Ptr() *string {
	return s.v
}

func (s *String) Var() string {
	return *s.v
}

type StringArray struct {
	v *[]string
}

func (arr *StringArray) Ptr() *[]string {
	return arr.v
}

func (arr *StringArray) Var() *lua.LTable {
	res := &lua.LTable{}
	for i, v := range *arr.v {
		res.Insert(i+1, lua.LString(v))
	}
	return res
}

func loadCtl(vm *YockScheduler) runtime.Handles {
	return runtime.LuaFuncs{
		"new_command": func(l *lua.LState) int {
			l.Push(luar.New(l, &cobra.Command{}))
			return 1
		},
		"Boolean": func(l *lua.LState) int {
			b := l.CheckBool(1)
			l.Push(luar.New(l, &Boolean{v: &b}))
			return 1
		},
		"String": func(l *lua.LState) int {
			s := l.CheckString(1)
			l.Push(luar.New(l, &String{v: &s}))
			return 1
		},
		"StringArray": func(l *lua.LState) int {
			s := []string{}
			for i := 1; i <= l.GetTop(); i++ {
				s = append(s, l.CheckString(i))
			}
			l.Push(luar.New(l, &StringArray{v: &s}))
			return 1
		},
	}
}
