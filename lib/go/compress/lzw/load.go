// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lzwlib

import (
	"compress/lzw"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadLzw(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("lzw")
	lib.SetField(map[string]any{
		// functions
		"NewReader": lzw.NewReader,
		"NewWriter": lzw.NewWriter,
		// constants
		"LSB": lzw.LSB,
		"MSB": lzw.MSB,
		// variable
	})
}
