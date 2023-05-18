package cmd

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Echo(str string) (string, error) {
	out := ""
	argv := strings.Split(str, " ")
	var parsedArgs []string
	for i := 0; i < len(argv); i++ {
		arg := argv[i]
		if strings.HasPrefix(arg, "\"") {
			j := i + 1
			for ; j < len(str) && !strings.HasSuffix(argv[j], "\""); j++ {
				arg += "<&>" + argv[j]
			}
			if j < len(str) {
				arg += "<&>" + argv[j]
				i = j
			}
			unquoted, err := strconv.Unquote(arg)
			if err == nil {
				parsedArgs = append(parsedArgs, strings.ReplaceAll(unquoted, "<&>", " "))
			} else {
				return "", ErrGeneral
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
	return out, nil
}
