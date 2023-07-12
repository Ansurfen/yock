// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"testing"
)

func TestPS(t *testing.T) {
	fmt.Println(PS(PSOpt{User: true, CPU: true}))
}

func TestKill(t *testing.T) {
	KillByName("test")
}

func TestPGrep(t *testing.T) {
	fmt.Println(PGrep("test"))
}

func TestNohup(t *testing.T) {
	Nohup("./test/test.exe")
}
