// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"fmt"
	"testing"
)

func TestYocki(t *testing.T) {
	c := New("localhost", 9090)
	defer c.Close()
	fmt.Println(c.Call("SayHello", "Hello, I'm Golang"))
}
