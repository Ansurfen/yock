// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// comments the file to disable ffi compile implement cross build
package yocks

import "github.com/ansurfen/yock/ffi"

func init() {
	libyock = append(libyock, ffi.LoadFFI)
}
