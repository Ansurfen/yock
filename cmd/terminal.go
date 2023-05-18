package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ansurfen/cushion/utils"
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
		util.YchoInfo(opt.Caller, fmt.Sprintf("exec: %s %s", name, strings.Join(args, " ")))
	}

	if opt.Redirect {
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return nil, cmd.Run()
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
	switch utils.CurPlatform.Ver {
	case "10", "11":
		return &Terminal{boot: []string{"powershell"}, cmd: cmds, this: TermPowershell}
	default:
		return &Terminal{boot: []string{"cmd", "/C"}, cmd: cmds, this: TermCmd}
	}
}

func PosixTerm(cmds ...string) *Terminal {
	return &Terminal{boot: []string{"/bin/bash", "/C"}, cmd: cmds, this: TermBash}
}
