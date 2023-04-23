package yock

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	var (
		out string
		err error
	)
	argv := strings.Split(echo.str, " ")
	var parsedArgs []string
	for i := 0; i < len(argv); i++ {
		arg := argv[i]
		if strings.HasPrefix(arg, "\"") {
			j := i + 1
			for ; j < len(args) && !strings.HasSuffix(argv[j], "\""); j++ {
				arg += "<&>" + argv[j]
			}
			if j < len(args) {
				arg += "<&>" + argv[j]
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
	isVar := regexp.MustCompile(`\$\w*`)
	for i := 0; i < len(parsedArgs); i++ {
		if s := isVar.FindString(parsedArgs[i]); len(s) > 0 && !strings.Contains(parsedArgs[i], "\n") {
			out += os.Getenv(s[1:])
		} else {
			out += parsedArgs[i]
		}
		out += " "
	}
	fmt.Println(out)
	return []byte(out), err
}
