// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestPrintf(t *testing.T) {
	Prinf(PrintfOpt{
		MaxLen: 20,
	}, []string{"Container", "Supplier", "Addr", "State", "Created"}, [][]string{
		{"go", "ark", "https://github.com/ansurfen/ark", "active", "2023-01-01"},
		{"nodejs", "ark", "https://github.com/ansurfen/ark", "unknown", "2023-01-08"},
	})
}

func TestExec(t *testing.T) {
	out, _ := ExecStr("go help build")
	fmt.Println(ConvertByte2String(out, GB18030))
	fmt.Println(PathIsExist("./abc"))
}

