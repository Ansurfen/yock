// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ansurfen/yock/ffi"
)

func main() {
	mylib, err := ffi.NewLibray("libmylib.dll")
	if err != nil {
		panic(err)
	}
	hello := mylib.Func("hello", "void", []string{})
	hello()
	hello2 := mylib.Func("hello2", "str", []string{"str", "int"})
	fmt.Println(hello2("ansurfen", int32(60)))
}
