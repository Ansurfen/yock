// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import "github.com/ansurfen/yock/ycho"

func init() {
	zlog, err := ycho.NewZLog(ycho.YchoOpt{
		Stdout: true,
	})
	if err != nil {
		panic(err)
	}
	ycho.Set(zlog)
}

func main() {
	ycho.Info("Hello World!")
	ycho.Fatalf("1 == 2 -> %v", 1 == 2)
}
