// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import yockc "github.com/ansurfen/yock/cmd"

func main() {
	_, err := yockc.Curl(yockc.CurlOpt{
		Method: "GET",
		Save:   false,
		Dir:    ".",
		FilenameHandle: func(s string) string {
			return s
		},
	}, []string{"https://www.github.com"})
	if err != nil {
		panic(err)
	}
}
