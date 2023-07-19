// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocki

import "io"

type Ycho interface {
	Progress(total int64, r io.Reader) io.Writer
	Eventloop()
	Info(msg string)
	Infof(msg string, v ...any)
	Debug(msg string)
	Debugf(msg string, v ...any)
	Warn(msg string)
	Warnf(msg string, v ...any)
	Fatal(msg string)
	Fatalf(msg string, v ...any)
	Error(msg string)
	Errorf(msg string, v ...any)
	Print(msg string)
}
