// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocke

import "testing"

func TestEnvVar(t *testing.T) {
	env := NewEnvVar()
	env.SetPath("sys")
	env.Export("")
}
