package cmd

import (
	"flag"

	"github.com/ansurfen/yock/util"
)

type TouchCmd struct {
	file string
	c    bool
}

func NewTouchCmd() Cmd {
	return &TouchCmd{}
}

func (touch *TouchCmd) Exec(args string) ([]byte, error) {
	initCmd(touch, args, func(cli *flag.FlagSet, cc *TouchCmd) {
		cli.BoolVar(&cc.c, "c", false, "")
	}, map[string]uint8{
		"-c": FlagBool,
	}, func(cc *TouchCmd, s string) error {
		cc.file = s
		return nil
	})
	var err error
	if touch.c {
		err = util.SafeWriteFile(touch.file, NilByte)
	} else {
		err = util.WriteFile(touch.file, NilByte)
	}
	return NilByte, err
}
