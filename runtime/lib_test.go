// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"fmt"
	"testing"

	lua "github.com/yuin/gopher-lua"
)

func TestCreateYockLib(t *testing.T) {
	s := NewYState()
	lib := CreateYockLib(s, "os")
	lib.SetField(map[string]any{
		"ver": "10",
		"echo": func() {
			fmt.Println("Hello, I'm Golang")
		},
	})
	lib.Meta().Value().ForEach(func(l1, l2 lua.LValue) {
		fmt.Println(l1, l2)
	})
	s.LState().DoString("print(os.ver); os.echo()")
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
	lib.Meta().Value().ForEach(func(l1, l2 lua.LValue) {
		fmt.Println(l1, l2)
	})
}
