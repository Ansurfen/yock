// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import "github.com/ansurfen/yock/util"

// Clear clears the output on the screen
func Clear() error {
	var term *Terminal
	switch util.CurPlatform.OS {
	case "windows":
		term = WindowsTerm("cls")
	default:
		term = PosixTerm("clear")
	}
	if _, err := term.Exec(&ExecOpt{Quiet: true}); err != nil {
		return err
	}
	return nil
}
