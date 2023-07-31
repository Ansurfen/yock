// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"net"
	"testing"
)

func TestID(t *testing.T) {
	fmt.Println(ID)
	interfaces, err := net.Interfaces()
	if err != nil || len(interfaces) == 0 {
		panic("fail to init machine id")
	}
}
