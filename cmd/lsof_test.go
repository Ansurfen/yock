// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"testing"

	"github.com/ansurfen/yock/util/test"
)

func TestLsof(t *testing.T) {
	fmt.Println(Lsof())
}

func TestLsof4Linux(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/lsof/linux",
	}).Range(func(data string) error {
		fmt.Println(lsofLinux(data))
		return nil
	})
}
