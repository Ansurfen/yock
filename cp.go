package yock

import (
	"flag"
	"os/exec"

	"github.com/ansurfen/cushion/utils"
)

type CpCmd struct {
	Src string
	Dst string
}

func NewCpCmd() Cmd {
	return &CpCmd{}
}

func (cp *CpCmd) Exec(args string) ([]byte, error) {
	initCmd(cp, args, func(cli *flag.FlagSet, cc *CpCmd) {
	}, map[string]uint8{}, func(cc *CpCmd, s string) error {
		if len(cc.Src) == 0 {
			cc.Src = s
		} else {
			cc.Dst = s
		}
		return nil
	})
	switch utils.CurPlatform.OS {
	case "windows":
		switch utils.CurPlatform.Ver {
		case "10", "11":
			return exec.Command("powershell", []string{"cp", cp.Src, cp.Dst}...).CombinedOutput()
		default:
			return exec.Command("cmd", []string{"/C", "copy", cp.Src, cp.Dst}...).CombinedOutput()
		}
	case "linux", "darwin":
		return exec.Command("/bin/bash", []string{"/C", "cp", cp.Src, cp.Dst}...).CombinedOutput()
	}
	return []byte(""), nil
}
