package parser

import (
	"regexp"
	"strings"

	. "github.com/ansurfen/yock/lib/bash/cmd"
	"github.com/ansurfen/yock/util"
)

func LoadBySh(file string) {
	isVar := regexp.MustCompile(`^\w+ *=`)
	cmds := [][]string{}
	vars := make(map[string]string)
	isRef := regexp.MustCompile(`\$\w*`)
	_, err := util.ReadLineFromFile(file, func(s string) string {
		if isVar.Match([]byte(s)) {
			if name, value, ok := strings.Cut(s, "="); ok {
				vars["$"+strings.TrimSpace(name)] = value
			}
		} else {
			if strings.Contains(s, "$") {
				for _, v := range isRef.FindAllString(s, -1) {
					if vv, ok := vars[v]; ok {
						s = strings.ReplaceAll(s, v, vv)
					}
				}
			}
			cmds = append(cmds, splitPipe(s))
		}
		return ""
	})
	if err != nil {
		panic(err)
	}
	for _, cmd := range cmds {
		last := ""
		for _, c := range cmd {
			sc := parseCmd(c)
			if _cmd, ok := GlobalCmds[sc.cmd]; ok {
				out, _ := _cmd().Exec(last + " " + sc.argv)
				// if err != nil {
				// 	panic(err)
				// }
				last = string(out)
			}
		}
	}
}

func LoadByStr(cmd string) ([]byte, error) {
	sc := parseCmd(cmd)
	if c, ok := GlobalCmds[sc.cmd]; ok {
		return c().Exec(sc.argv)
	}
	return nil, nil
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
