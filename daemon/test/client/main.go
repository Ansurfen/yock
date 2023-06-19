// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ansurfen/yock/daemon/client"
)

func main() {
	client.Gopt.Parse()
	client := client.New(client.Gopt)
	defer client.Close()
	if err := client.Ping(); err != nil {
		panic(err)
	}
	info, err := client.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}
