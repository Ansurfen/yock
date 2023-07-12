// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/ansurfen/yock/ycho"
)

func init() {
	tlog, err := ycho.NewTLog(ycho.YchoOpt{})
	if err != nil {
		panic(err)
	}
	ycho.Set(tlog)
	// go ycho.Eventloop()
	ycho.Info("Init logger")
}

func main() {
	ycho.Infof("Hello %s!", "World")
	ycho.Warnf("Hello World!")
	ycho.Errorf("Hello World!")
	ycho.Fatalf("Hello World!")
}
