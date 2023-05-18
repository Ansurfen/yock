package cmd

import (
	"flag"

	"github.com/ansurfen/yock/cmd"
)

type EchoCmd struct {
	str string
}

func NewEchoCmd() Cmd {
	return &EchoCmd{}
}

func (echo *EchoCmd) Exec(args string) ([]byte, error) {
	initCmd(echo, args, func(cli *flag.FlagSet, cc *EchoCmd) {}, map[string]uint8{},
		func(cc *EchoCmd, s string) error {
			cc.str += s
			return nil
		})
	out, err := cmd.Echo(echo.str)
	return []byte(out), err
}
