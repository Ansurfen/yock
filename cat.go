package yock

import (
	"flag"
	"os/exec"

	"github.com/ansurfen/cushion/utils"
)

type CatCmd struct {
	path string
}

func NewCatCmd() Cmd {
	return &CatCmd{}
}

func (cat *CatCmd) Exec(args string) ([]byte, error) {
	initCmd(cat, args, func(cli *flag.FlagSet, cc *CatCmd) {
	}, map[string]uint8{}, func(cc *CatCmd, s string) error {
		cc.path = s
		return nil
	})
	switch utils.CurPlatform.OS {
	case "windows":
		switch utils.CurPlatform.Ver {
		case "10", "11":
			return exec.Command("powershell", []string{"cat", cat.path}...).CombinedOutput()
		default:
			return exec.Command("cmd", []string{"/C", "type", cat.path}...).CombinedOutput()
		}
	case "linux", "darwin":
		return exec.Command("/bin/bash", []string{"cat", cat.path}...).CombinedOutput()
	}
	return []byte(""), nil
}
