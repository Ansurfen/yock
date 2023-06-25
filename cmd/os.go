// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import "github.com/ansurfen/yock/util"

// ExecOpt indicates configuration of exec
type ExecOpt struct {
	// Redirect set stdout, stderr, stdin stream
	Redirect bool
	// Debug prints output when it's true
	Debug bool
	// Caller is used to mark parent caller of HTTP function
	//
	// It'll printed on console when debug is true
	Caller string
	Quiet  bool
	// Strict will exit at once when error occur
	Strict bool
}

func Exec(opt ExecOpt, cmd string) (string, error) {
	var term *Terminal
	switch util.CurPlatform.OS {
	case "windows":
		term = WindowsTerm(cmd)
	default:
		term = PosixTerm()
	}
	if out, err := term.Exec(&opt); err != nil {
		return string(out), err
	} else {
		return string(out), nil
	}
}
