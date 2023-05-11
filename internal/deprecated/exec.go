package yock

import (
	"regexp"
	"strings"

	"github.com/ansurfen/cushion/utils"
)

func LoadBySh(file string) {
	isVar := regexp.MustCompile(`^\w+ *=`)
	cmds := [][]string{}
	vars := make(map[string]string)
	isRef := regexp.MustCompile(`\$\w*`)
	_, err := utils.ReadLineFromFile(file, func(s string) string {
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
			if _cmd, ok := globalCmds[sc.cmd]; ok {
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
	if c, ok := globalCmds[sc.cmd]; ok {
		return c().Exec(sc.argv)
	}
	return NilByte, nil
}
