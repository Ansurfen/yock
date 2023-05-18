package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ansurfen/yock/lib/bash/parser"
)

func main() {
	if len(os.Args) < 2 {
		panic("args must more than or equal 2")
	}
	str := os.Args[1]
	if filepath.Ext(str) == ".sh" {
		parser.LoadBySh(str)
	} else {
		if out, err := parser.LoadByStr(strings.Join(os.Args[1:], " ")); err != nil {
			panic(err)
		} else {
			fmt.Println(string(out))
		}
	}
}
