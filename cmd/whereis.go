// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"os/exec"
)

func Whereis(path string) (string, error) {
	return exec.LookPath(path)
}
