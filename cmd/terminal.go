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

// terminal is a struct to abstract different termnial
type terminal struct {
	cmd  []string
	boot []string
	this uint8
}

func (term *terminal) exec(opt *ExecOpt) error {
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
		return cmd.Run()
	}

	out, err := cmd.CombinedOutput()
	if !opt.Quiet {
		fmt.Print(string(out))
	}
	return err
}

func (term *terminal) setCmds(cmds ...string) {
	term.cmd = cmds
}

func windowsTerm(cmds ...string) *terminal {
	switch utils.CurPlatform.Ver {
	case "10", "11":
		return &terminal{boot: []string{"powershell"}, cmd: cmds, this: TermPowershell}
	default:
		return &terminal{boot: []string{"cmd", "/C"}, cmd: cmds, this: TermCmd}
	}
}

func posixTerm(cmds ...string) *terminal {
	return &terminal{boot: []string{"/bin/bash", "/C"}, cmd: cmds, this: TermBash}
}
