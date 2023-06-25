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

		"ReadDir":   os.ReadDir,
		"Open":      os.Open,
		"Create":    os.Create,
		"OpenFile":  os.OpenFile,
		"ReadFile":  os.ReadFile,
		"WriteFile": os.WriteFile,

		"O_RDONLY": os.O_RDONLY,
		"O_WRONLY": os.O_WRONLY,
		"O_RDWR":   os.O_RDWR,

		"O_APPEND": os.O_APPEND,
		"O_CREATE": os.O_CREATE,
		"O_EXCL":   os.O_EXCL,
		"O_SYNC":   os.O_SYNC,
		"O_TRUNC":  os.O_TRUNC,

		"Rename":        os.Rename,
		"TempDir":       os.TempDir,
		"UserCacheDir":  os.UserCacheDir,
		"UserConfigDir": os.UserConfigDir,
		"UserHomeDir":   os.UserHomeDir,
		"Chmod":         os.Chmod,
		"Chdir":         os.Chdir,
		"DirFS":         os.DirFS,
		"Mkdir":         os.Mkdir,
	})
}
