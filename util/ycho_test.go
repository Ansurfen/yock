// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"testing"
)

func TestLog(t *testing.T) {
	Ycho.Info("Hello World")
	Ycho.Warn("This is warn")
	Ycho.Error("This is error")
	Ycho.Fatal("This is fatal")
}
