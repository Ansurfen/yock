// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocki

type Ycho interface {
	Info(msg string)
	Debug(msg string)
	Warn(msg string)
	Fatal(msg string)
}
