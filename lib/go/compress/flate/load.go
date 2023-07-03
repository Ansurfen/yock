// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package flatelib

import (
	"compress/flate"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadFlate(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("flate")
	lib.SetField(map[string]any{
		// functions
		"NewReaderDict": flate.NewReaderDict,
		"NewWriter":     flate.NewWriter,
		"NewWriterDict": flate.NewWriterDict,
		"NewReader":     flate.NewReader,
		// constants
		"NoCompression":      flate.NoCompression,
		"BestSpeed":          flate.BestSpeed,
		"BestCompression":    flate.BestCompression,
		"DefaultCompression": flate.DefaultCompression,
		"HuffmanOnly":        flate.HuffmanOnly,
		// variable
	})
}
