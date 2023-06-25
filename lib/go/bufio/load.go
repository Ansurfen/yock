// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bufio

import (
	"bufio"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadBufio(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("bufio")
	lib.SetField(map[string]any{
		"NewScanner":    bufio.NewScanner,
		"ScanLines":     bufio.ScanLines,
		"NewReaderSize": bufio.NewReaderSize,
		"NewReader":     bufio.NewReader,
		"NewWriterSize": bufio.NewWriterSize,
		"NewWriter":     bufio.NewWriter,
		"NewReadWriter": bufio.NewReadWriter,
	})
}
