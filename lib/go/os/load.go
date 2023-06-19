// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package os

import (
	"os"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadOS(yocks yocki.YockScheduler) {
	lib := yocks.OpenLib("os")
	lib.SetField(map[string]any{
		"Stdin":  os.Stdin,
		"Stdout": os.Stdout,
		"Stderr": os.Stderr,

		"ReadDir": os.ReadDir,
		
	})
}
