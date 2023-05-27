// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ypm

import (
	"fmt"
	"testing"
)

func TestModule(t *testing.T) {
	mod, err := CreateModule("module.json", "bash")
	if err != nil {
		panic(err)
	}
	fmt.Println(mod)
}
