// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"fmt"
	"testing"

	"github.com/ansurfen/yock/util/test"
	lua "github.com/yuin/gopher-lua"
)

func TestCreateYockLib(t *testing.T) {
	testset := map[string][]any{
		"ver": {"10", lua.LTString},
		"echo": {func() {
			fmt.Println("Hello, I'm Golang")
		}, lua.LTFunction},
	}
	s := NewYState()
	lib := CreateYockLib(s, "os")
	for k, v := range testset {
		lib.SetField(map[string]any{
			k: v[0],
		})
	}
	lib.Meta().Value().ForEach(func(l1, l2 lua.LValue) {
		v := testset[l1.String()]
		test.Assert(l2.Type() == v[1].(lua.LValueType))
	})
	err := s.LState().DoString("print(os.ver); os.echo()")
	test.Assert(err == nil)
}

func TestOpenYockLib(t *testing.T) {
	s := NewYState()
	lib := OpenYockLib(s, "os")
	lib.SetField(map[string]any{
		"ver": "10",
		"echo": func() {
			fmt.Println("Hello, I'm Golang")
		},
	})
	test.Assert(lib.Name() == "os")
	lib.Meta().Value().ForEach(func(l1, l2 lua.LValue) {
		fmt.Println(l1, l2)
	})
}
