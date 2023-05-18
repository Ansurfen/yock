package cmd

import (
	"flag"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/cmd"
)

type CurlCmd struct {
	urls   []string
	body   string
	method string
	o      string
	O      string
}

func NewCurlCmd() Cmd {
	return &CurlCmd{}
}

func (curl *CurlCmd) Exec(arg string) ([]byte, error) {
	initCmd(curl, arg, func(cli *flag.FlagSet, cc *CurlCmd) {
		cli.StringVar(&cc.body, "d", "", "")
		cli.StringVar(&cc.method, "x", "GET", "")
		cli.StringVar(&cc.O, "O", "", "")
		cli.StringVar(&cc.o, "o", ".", "")
	}, map[string]uint8{
		"-d": FlagString,
		"-x": FlagString,
		"-O": FlagString,
		"-o": FlagString,
	}, func(cc *CurlCmd, s string) error {
		if utils.IsURL(s) {
			cc.urls = append(cc.urls, s)
		}
		return nil
	})
	save := false
	if len(curl.O) > 0 {
		save = true
	}
	return NilByte, cmd.HTTP(cmd.HttpOpt{
		Method: curl.method,
		Data:   curl.body,
		Save:   save,
		Filename: func(s string) string {
			if len(curl.O) > 0 {
				return curl.o
			}
			return s
		},
	}, curl.urls)
}
