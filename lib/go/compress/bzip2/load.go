// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bzip2lib

import (
	"compress/bzip2"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadBzip2(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("bzip2")
	lib.SetField(map[string]any{
		// functions
		"NewReader": bzip2.NewReader,
		// constants
		// variable
	})
}
