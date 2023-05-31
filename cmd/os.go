package cmd

import (
	"github.com/ansurfen/cushion/utils"
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
}

func Exec(opt ExecOpt, cmd string) (string, error) {
	var term *Terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = WindowsTerm(cmd)
	default:
		term = PosixTerm()
	}
	if out, err := term.Exec(&opt); err != nil {
		return string(out), err
	} else {
		return string(out), nil
	}
}
