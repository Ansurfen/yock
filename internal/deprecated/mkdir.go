package yock

import (
	"flag"
	"os/exec"

	"github.com/ansurfen/cushion/utils"
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
	switch utils.CurPlatform.OS {
	case "windows":
		switch utils.CurPlatform.Ver {
		case "10", "11":
			if mkdir.p {
				return exec.Command("powershell", []string{"mkdir", "-p", mkdir.path}...).CombinedOutput()
			} else {
				return exec.Command("powershell", []string{"mkdir", mkdir.path}...).CombinedOutput()
			}
		default:
			return exec.Command("cmd", []string{"mkdir", mkdir.path, "/p"}...).CombinedOutput()
		}
	case "linux":
		if mkdir.p {
			return exec.Command("/bin/bash", []string{"/C", "mkdir", "-p", mkdir.path}...).CombinedOutput()
		} else {
			return exec.Command("powershell", []string{"mkdir", mkdir.path}...).CombinedOutput()
		}
	}
	return NilByte, nil
}
