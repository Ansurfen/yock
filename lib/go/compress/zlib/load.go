// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zliblib

import (
	"compress/zlib"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadZlib(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("zlib")
	lib.SetField(map[string]any{
		// functions
		"NewReader":          zlib.NewReader,
		"NewReaderDict":      zlib.NewReaderDict,
		"NewWriter":          zlib.NewWriter,
		"NewWriterLevel":     zlib.NewWriterLevel,
		"NewWriterLevelDict": zlib.NewWriterLevelDict,
		// constants
		"NoCompression":      zlib.NoCompression,
		"BestSpeed":          zlib.BestSpeed,
		"BestCompression":    zlib.BestCompression,
		"DefaultCompression": zlib.DefaultCompression,
		"HuffmanOnly":        zlib.HuffmanOnly,
		// variable
		"ErrChecksum":   zlib.ErrChecksum,
		"ErrDictionary": zlib.ErrDictionary,
		"ErrHeader":     zlib.ErrHeader,
	})
}
