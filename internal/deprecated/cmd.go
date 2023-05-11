package yock

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Cmd interface {
	Exec(string) ([]byte, error)
}

var globalCmds = map[string]func() Cmd{
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

type shCmd struct {
	cmd  string
	argv string
}

func parseCmd(cmd string) shCmd {
	cmd = strings.TrimSpace(cmd)
	res := strings.Split(cmd, " ")
	if len(res) > 1 {
		return shCmd{
			cmd:  res[0],
			argv: strings.Join(res[1:], " "),
		}
	} else if len(res) > 0 {
		return shCmd{
			cmd: res[0],
		}
	}
	return shCmd{}
}

func splitPipe(script string) []string {
	subscript := []string{}
	ban := false
	str := ""
	for _, ch := range script {
		if ch == '"' || ch == '\'' {
			ban = !ban
		}
		if ch != '|' || ban {
			str += string(ch)
		} else {
			subscript = append(subscript, str)
			str = ""
		}
	}
	if len(str) > 0 {
		subscript = append(subscript, str)
	}
	return subscript
}

const (
	FlagBool = iota
	FlagString
)

var NilByte = []byte(nil)

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
