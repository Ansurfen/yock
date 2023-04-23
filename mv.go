package yock

import (
	"flag"
	"os/exec"

	"github.com/ansurfen/cushion/utils"
)

type MoveCmd struct {
	Src string
	Dst string
}

func NewMoveCmd() Cmd {
	return &MoveCmd{}
}

func (mv *MoveCmd) Exec(args string) ([]byte, error) {
	initCmd(mv, args, func(cli *flag.FlagSet, cc *MoveCmd) {
	}, map[string]uint8{}, func(cc *MoveCmd, s string) error {
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
			return exec.Command("powershell", []string{"mv", mv.Src, mv.Dst}...).CombinedOutput()
		default:
			return exec.Command("cmd", []string{"/C", "move", mv.Src, mv.Dst}...).CombinedOutput()
		}
	case "linux":
		return exec.Command("/bin/bash", []string{"/C", "mv", mv.Src, mv.Dst}...).CombinedOutput()
	}
	return []byte(""), nil
}
