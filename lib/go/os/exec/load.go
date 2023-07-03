// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package execlib

import (
	"os/exec"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadExec(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("exec")
	lib.SetField(map[string]any{
		// functions
		"CommandContext": exec.CommandContext,
		"Command":        exec.Command,
		"LookPath":       exec.LookPath,
		// constants
		// variable
		"ErrWaitDelay": exec.ErrWaitDelay,
		"ErrDot":       exec.ErrDot,
		"ErrNotFound":  exec.ErrNotFound,
	})
}
