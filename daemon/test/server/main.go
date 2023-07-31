// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/ansurfen/yock/daemon/server"
)

func main() {
	s := api.New()
	defer s.Close()
	s.Run()
}
