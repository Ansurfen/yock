package cmd

import (
	"fmt"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
)

// ExecOpt indicates configuration of exec
type ExecOpt struct {
	// Redirect set stdout, stderr, stdin stream
	Redirect bool
	// Debug prints output when it's true
	Debug bool
	// Caller is used to mark parent caller of HTTP function
	//
	// It'll printed on console when debug is true
	Caller string
	Quiet  bool
	// Strict will exit at once when error occur
	Strict bool

	err error
}

func Exec(opt ExecOpt, cmds []string) error {
	for _, cmd := range cmds {
		var term *terminal
		switch utils.CurPlatform.OS {
		case "windows":
			term = windowsTerm(cmd)
		default:
			term = posixTerm()
		}
		if err := term.exec(&opt); err != nil {
			if opt.Debug {
				util.YchoWarn(opt.Caller, fmt.Sprintf("%s err: %s", cmd, err.Error()))
			}
			if opt.Strict {
				return err
			} else {
				opt.err = ErrGeneral
			}
		}
	}
	return opt.err
}