// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ansurfen/yock/util"
)

type EchoOpt struct {
	Fd   []string
	Mode string
}

// Echo prints str to the screen and returns str.
// Similar to GNU's echo, you can also print environment variables by $GOPATH
func Echo(opt EchoOpt, str string) (string, error) {
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
				return "", util.ErrGeneral
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
	mode := 0
	for _, m := range strings.Split(opt.Mode, "|") {
		switch m {
		case "c":
			mode |= os.O_CREATE
		case "t":
			mode |= os.O_TRUNC
		case "r":
			mode |= os.O_RDONLY
		case "w":
			mode |= os.O_WRONLY
		case "rw":
			mode |= os.O_RDWR
		case "a":
			mode |= os.O_APPEND
		case "e":
			mode |= os.O_EXCL
		case "s":
			mode |= os.O_SYNC
		}
	}
	if mode == 0 {
		mode = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	for _, fd := range opt.Fd {
		if fd == "stdout" {
			fmt.Fprint(os.Stdout, out)
		} else if fd == "stderr" {
			fmt.Fprint(os.Stderr, out)
		} else {
			fp, err := os.OpenFile(fd, mode, 0666)
			if err != nil {
				return out, err
			}
			_, err = fp.Write([]byte(out))
			if err != nil {
				return out, err
			}
		}
	}
	return out, nil
}
