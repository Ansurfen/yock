// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	yockr "github.com/ansurfen/yock/runtime"
)

type user struct {
	Name string
	Pwd  string
}

func main() {
	u := user{}
	r := yockr.New()
	if err := r.Eval(`return {name = "root", pwd = "123456"}`); err != nil {
		panic(err)
	}
	if err := r.State().CheckTable(1).Bind(&u); err != nil {
		panic(err)
	}
	fmt.Println(u)
}
