package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/ansurfen/cushion/utils"
)

type ExecOpt struct {
	Redirect bool
	Debug    bool
	Quiet    bool
	Terminal struct{}
}

func Exec(opt ExecOpt, cmds []string) error {
	for _, raw := range cmds {
		var cmd *exec.Cmd
		switch utils.CurPlatform.OS {
		case "windows":
			switch utils.CurPlatform.Ver {
			case "10", "11":
				cmd = exec.Command("powershell", raw)
			default:
				cmd = exec.Command("cmd", []string{"/C", raw}...)
			}
		case "linux", "darwin":
			cmd = exec.Command("/bin/bash", []string{"/C", raw}...)
		default:
			return errors.New("not support platform")
		}
		if opt.Redirect {
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			out, _ := cmd.CombinedOutput()
			if !opt.Quiet {
				fmt.Print(string(out))
			}
		}
	}
	return nil
}
