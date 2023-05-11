package yock

import (
	"flag"
	"os/exec"

	"github.com/ansurfen/cushion/utils"
)

type RmCmd struct {
	Path string
	r    bool
}

func NewRmCmd() Cmd {
	return &RmCmd{}
}

func (rm *RmCmd) Exec(args string) ([]byte, error) {
	initCmd(rm, args, func(cli *flag.FlagSet, cc *RmCmd) {
		cli.BoolVar(&cc.r, "r", false, "")
	}, map[string]uint8{
		"-r": FlagBool,
	}, func(cc *RmCmd, s string) error {
		rm.Path = s
		return nil
	})
	switch utils.CurPlatform.OS {
	case "windows":
		switch utils.CurPlatform.Ver {
		case "10", "11":
			if rm.r {
				return exec.Command("powershell", []string{"rm", "-r", rm.Path}...).CombinedOutput()
			} else {
				return exec.Command("powershell", []string{"rm", rm.Path}...).CombinedOutput()
			}
		default:
			if rm.r {
				return exec.Command("cmd", []string{"rmdir", "/s", "/q", rm.Path}...).CombinedOutput()
			} else {
				return exec.Command("cmd", []string{"del", rm.Path}...).CombinedOutput()
			}
		}
	case "linux":
		if rm.r {
			return exec.Command("/bin/bash", []string{"/C", "rm", "-r", rm.Path}...).CombinedOutput()

		} else {
			return exec.Command("/bin/bash", []string{"/C", "rm", rm.Path}...).CombinedOutput()
		}
	}
	return []byte(""), nil
}

type RmdirCmd struct {
	Path string
}

func NewRmdirCmd() Cmd {
	return &RmdirCmd{}
}

func (rm *RmdirCmd) Exec(args string) ([]byte, error) {
	initCmd(rm, args, func(cli *flag.FlagSet, cc *RmdirCmd) {
	}, map[string]uint8{}, func(cc *RmdirCmd, s string) error {
		rm.Path = s
		return nil
	})
	switch utils.CurPlatform.OS {
	case "windows":
		return exec.Command("cmd", append([]string{"/C", "rmdir"}, rm.Path)...).CombinedOutput()
	case "linux":
		return exec.Command("/bin/bash", append([]string{"/C", "rmdir"}, rm.Path)...).CombinedOutput()
	}
	return []byte(""), nil
}
