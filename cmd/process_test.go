// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"testing"

	"github.com/ansurfen/yock/util/test"
)

func TestPS(t *testing.T) {
	_, err := PS(PSOpt{User: true, CPU: true})
	test.Assert(err == nil)
}

func TestKill(t *testing.T) {
	err := KillByName("test")
	test.Assert(err == nil)
}

func TestPGrep(t *testing.T) {
	fmt.Println(PGrep("test"))
}

func TestNohup(t *testing.T) {
	err := Nohup("./test/test.exe")
	test.Assert(err == nil)
}
