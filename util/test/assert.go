// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package test

import (
	"fmt"
	"runtime"
	"strings"
)

func getTopCaller(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	file, line := runtime.FuncForPC(pc).FileLine(pc)
	str := strings.Split(file, "/")
	name := str[len(str)-2] + "/" + str[len(str)-1]
	return fmt.Sprintf("%s:%d", name, line)
}

func Assert(ok bool, msg ...string) {
	if !ok {
		v := "fail asserted!"
		if len(msg) > 0 {
			v = msg[0]
		}
		panic(getTopCaller(2) + "\n" + v)
	}
}
