// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ziplib

import (
	"archive/zip"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadZip(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("zip")
	lib.SetField(map[string]any{
		// functions
		"NewReader":            zip.NewReader,
		"RegisterDecompressor": zip.RegisterDecompressor,
		"RegisterCompressor":   zip.RegisterCompressor,
		"FileInfoHeader":       zip.FileInfoHeader,
		"NewWriter":            zip.NewWriter,
		"OpenReader":           zip.OpenReader,
		// constants
		"Store":   zip.Store,
		"Deflate": zip.Deflate,
		// variable
		"ErrFormat":       zip.ErrFormat,
		"ErrAlgorithm":    zip.ErrAlgorithm,
		"ErrChecksum":     zip.ErrChecksum,
		"ErrInsecurePath": zip.ErrInsecurePath,
	})
}
