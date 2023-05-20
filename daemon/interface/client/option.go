package client

import (
	"flag"
	"os"

	"github.com/ansurfen/yock/util"
)

const HelpOpt = `Usage:
  -p specify daemon port to conn`

type DaemonOption struct {
	IP   *string
	Port *int
	// MTL is abbreviation to max transfer length for file
	MTL  *int
	Name *string
}

func (opt *DaemonOption) Parse() {
	if len(os.Args) == 1 {
		panic(util.ErrArgsTooLittle)
	}
	flag.Parse()
	if *opt.Port == 0 {
		panic(util.ErrInvalidPort)
	}
}

var Gopt = &DaemonOption{
	Port: flag.Int("p", 0, ""),
	MTL:  flag.Int("mtl", 1024, ""),
	IP:   flag.String("ip", "localhost", ""),
	Name: flag.String("name", "", ""),
}
