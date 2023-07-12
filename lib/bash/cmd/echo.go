package cmd

import (
	"flag"

	yockc "github.com/ansurfen/yock/cmd"
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
	out, err := yockc.Echo(yockc.EchoOpt{}, echo.str)
	return []byte(out), err
}
