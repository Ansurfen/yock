package cmd

import (
	"flag"

	"github.com/ansurfen/cushion/utils"
)

type CatCmd struct {
	file string
}

func NewCatCmd() Cmd {
	return &CatCmd{}
}

func (cat *CatCmd) Exec(args string) ([]byte, error) {
	initCmd(cat, args, func(cli *flag.FlagSet, cc *CatCmd) {}, map[string]uint8{},
		func(cc *CatCmd, s string) error {
			cc.file = s
			return nil
		})
	out, err := utils.ReadStraemFromFile(cat.file)
	return out, err
}
