// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unicode

import (
	"unicode"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadUnicode(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("unicode")
	lib.SetField(map[string]any{
		"IsGraphic": unicode.IsGraphic,
		"IsPrint":   unicode.IsPrint,
		"IsControl": unicode.IsControl,
		"IsLetter":  unicode.IsLetter,
		"IsMark":    unicode.IsMark,
		"IsNumber":  unicode.IsNumber,
		"IsPunct":   unicode.IsPunct,
		"IsSpace":   unicode.IsSpace,
		"IsSymbol":  unicode.IsSymbol,
		"IsDigit":   unicode.IsDigit,
	})
}
