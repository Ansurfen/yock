package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/ansurfen/yock"
)

func main() {
	out, err := yock.LoadByStr(strings.Join(os.Args[1:], " "))
	if err != nil {
		panic(err)
	}
	if len(out) > 0 {
		fmt.Println(string(out))
	}
}
