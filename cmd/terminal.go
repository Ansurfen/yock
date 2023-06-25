// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ansurfen/yock/util"
)

const (
	TermUndefined = iota
	TermPowershell
	TermCmd
	TermBash
)

// Terminal is a struct to abstract different termnial
type Terminal struct {
	cmd  []string
	boot []string
	this uint8
}

func (term *Terminal) Exec(opt *ExecOpt) ([]byte, error) {
	name := term.boot[0]
	args := []string{}
	if len(term.boot) > 1 {
		args = append(args, term.boot[1:]...)
	}
	args = append(args, term.cmd...)
	cmd := exec.Command(name, args...)

	if opt.Debug {
		if len(opt.Caller) > 0 {
			util.Ycho.Info(fmt.Sprintf("%s\t%s", opt.Caller, fmt.Sprintf("%s %s", name, strings.Join(args, " "))))
		} else {
			util.Ycho.Info(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
		}
	}

	if opt.Redirect {
		var out bytes.Buffer
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = io.MultiWriter(os.Stdout, &out)
		err := cmd.Run()
		if err != nil {
			return nil, err
		}
		return out.Bytes(), err
	}

	out, err := cmd.CombinedOutput()
	if !opt.Quiet {
		fmt.Print(string(out))
	}
	return out, err
}

func (term *Terminal) SetCmds(cmds ...string) {
	term.cmd = cmds
}

func (term *Terminal) AppendCmds(cmds ...string) {
	term.cmd = append(term.cmd, cmds...)
}

func (term *Terminal) Clear() {
	term.cmd = []string{}
}

func (term *Terminal) Type() uint8 {
	return term.this
}

func WindowsTerm(cmds ...string) *Terminal {
	switch util.CurPlatform.Ver {
	case "10", "11":
		return &Terminal{boot: []string{"powershell"}, cmd: cmds, this: TermPowershell}
	default:
		return &Terminal{boot: []string{"cmd", "/C"}, cmd: cmds, this: TermCmd}
	}
}

func PosixTerm(cmds ...string) *Terminal {
	return &Terminal{boot: []string{"/bin/sh", "-c"}, cmd: cmds, this: TermBash}
}

func UpgradeBackend(term *Terminal) {

}
