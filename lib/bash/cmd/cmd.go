package cmd

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var GlobalCmds = map[string]func() Cmd{
	"mv":     NewMoveCmd,
	"cp":     NewCpCmd,
	"echo":   NewEchoCmd,
	"rm":     NewRmCmd,
	"rmdir":  NewRmdirCmd,
	"curl":   NewCurlCmd,
	"mkdir":  NewMkdirCmd,
	"cat":    NewCatCmd,
	"whoami": NewWhoamiCmd,
	"pwd":    NewPwdCmd,
	"touch":  NewTouchCmd,
}

type Cmd interface {
	Exec(string) ([]byte, error)
}

var NilByte = []byte("")

const (
	FlagBool = iota
	FlagString
)

func initCmd[c Cmd](
	cc c, input string,
	register func(cli *flag.FlagSet, cc c),
	options map[string]uint8,
	raw func(cc c, s string) error,
) c {
	args := strings.Split(input, " ")
	var parsedArgs []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "\"") {
			j := i + 1
			for ; j < len(args) && !strings.HasSuffix(args[j], "\""); j++ {
				arg += "<&>" + args[j]
			}
			if j < len(args) {
				arg += "<&>" + args[j]
				i = j
			}
			unquoted, err := strconv.Unquote(arg)
			if err == nil {
				parsedArgs = append(parsedArgs, strings.ReplaceAll(unquoted, "<&>", " "))
			} else {
				fmt.Printf("Error unquoting argument: %s\n", arg)
			}
		} else {
			parsedArgs = append(parsedArgs, arg)
		}
	}
	cli := flag.NewFlagSet("", flag.ContinueOnError)
	register(cli, cc)
	for i := 0; i < len(parsedArgs); i++ {
		if ft, ok := options[parsedArgs[i]]; ok {
			switch ft {
			case FlagString:
				i++
			case FlagBool:
			}
		} else {
			if err := raw(cc, parsedArgs[i]); err != nil {
				panic(err)
			}
			parsedArgs = append(parsedArgs[:i], parsedArgs[i+1:]...)
			i--
		}
	}
	cli.Parse(parsedArgs)
	return cc
}
