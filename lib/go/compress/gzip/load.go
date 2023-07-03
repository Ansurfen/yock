// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gziplib

import (
	"compress/gzip"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadGzip(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("gzip")
	lib.SetField(map[string]any{
		// functions
		"NewReader":      gzip.NewReader,
		"NewWriter":      gzip.NewWriter,
		"NewWriterLevel": gzip.NewWriterLevel,
		// constants
		"NoCompression":      gzip.NoCompression,
		"BestSpeed":          gzip.BestSpeed,
		"BestCompression":    gzip.BestCompression,
		"DefaultCompression": gzip.DefaultCompression,
		"HuffmanOnly":        gzip.HuffmanOnly,
		// variable
		"ErrChecksum": gzip.ErrChecksum,
		"ErrHeader":   gzip.ErrHeader,
	})
}
