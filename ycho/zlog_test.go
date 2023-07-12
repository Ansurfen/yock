// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import (
	"testing"
)

func TestZLog(t *testing.T) {
	ycho, err := NewZLog(YchoOpt{Stdout: true})
	if err != nil {
		panic(err)
	}
	ycho.Info("Hello World")
	ycho.Warn("This is warn")
	ycho.Errorf("This is error")
	ycho.Fatalf("This is fatal")
}
