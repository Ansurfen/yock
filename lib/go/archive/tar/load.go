// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tarlib

import (
	"archive/tar"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadTar(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("tar")
	lib.SetField(map[string]any{
		// functions
		"FileInfoHeader": tar.FileInfoHeader,
		"NewReader":      tar.NewReader,
		"NewWriter":      tar.NewWriter,
		// constants
		"TypeReg":           tar.TypeReg,
		// "TypeRegA":          tar.TypeRegA,
		"TypeLink":          tar.TypeLink,
		"TypeSymlink":       tar.TypeSymlink,
		"TypeChar":          tar.TypeChar,
		"TypeBlock":         tar.TypeBlock,
		"TypeDir":           tar.TypeDir,
		"TypeFifo":          tar.TypeFifo,
		"TypeCont":          tar.TypeCont,
		"TypeXHeader":       tar.TypeXHeader,
		"TypeXGlobalHeader": tar.TypeXGlobalHeader,
		"TypeGNUSparse":     tar.TypeGNUSparse,
		"TypeGNULongName":   tar.TypeGNULongName,
		"TypeGNULongLink":   tar.TypeGNULongLink,
		"FormatUnknown":     tar.FormatUnknown,
		"FormatUSTAR":       tar.FormatUSTAR,
		"FormatPAX":         tar.FormatPAX,
		"FormatGNU":         tar.FormatGNU,
		// variable
		"ErrHeader":          tar.ErrHeader,
		"ErrWriteTooLong":    tar.ErrWriteTooLong,
		"ErrFieldTooLong":    tar.ErrFieldTooLong,
		"ErrWriteAfterClose": tar.ErrWriteAfterClose,
		"ErrInsecurePath":    tar.ErrInsecurePath,
	})
}
