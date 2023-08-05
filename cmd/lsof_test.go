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
	_, err := Lsof()
	test.Assert(err == nil)
}

func ExampleLsofInfo() {
	test.Batch(test.BatchOpt{
		Path: "./testdata/lsof/linux",
	}).Range(func(data string) error {
		fmt.Println(lsofLinux(data))
		return nil
	})
	// Output:
	// [{tcp 0.0.0.0:22 0.0.0.0:* LISTEN 0} {tcp 127.0.0.1:631 0.0.0.0:* LISTEN 0} {tcp 127.0.0.53:53 0.0.0.0:* LISTEN 0} {tcp6 :::8080 :::* LISTEN 2717}]
}
