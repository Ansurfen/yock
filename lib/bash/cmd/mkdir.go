package cmd

import (
	"flag"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
)

type MkdirCmd struct {
	path string
	p    bool
}

func NewMkdirCmd() Cmd {
	return &MkdirCmd{}
}

func (mkdir *MkdirCmd) Exec(args string) ([]byte, error) {
	initCmd(mkdir, args, func(cli *flag.FlagSet, mkdir *MkdirCmd) {
		cli.BoolVar(&mkdir.p, "p", false, "")
	}, map[string]uint8{
		"-p": FlagBool,
	}, func(cc *MkdirCmd, s string) error {
		cc.path = s
		return nil
	})
	var term *cmd.Terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = cmd.WindowsTerm("mkdir")
		if mkdir.p {
			if term.Type() == cmd.TermPowershell {
				term.AppendCmds("-p", mkdir.path)
			} else {
				term.AppendCmds(mkdir.path, "/p")
			}
		} else {
			term.AppendCmds(mkdir.path)
		}
	default:
		term = cmd.PosixTerm()
		if mkdir.p {
			term.AppendCmds("-p", mkdir.path)
		} else {
			term.AppendCmds(mkdir.path)
		}
	}
	return term.Exec(&cmd.ExecOpt{})
}
