// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import (
	"io"

	yocki "github.com/ansurfen/yock/interface"
)

var ycho yocki.Ycho

const defaultTimeFormat = "2006-01-02 15:04:05.000 -0700"

func init() {
	ycho = &Vlog{}
}

func Get() yocki.Ycho {
	return ycho
}

func Set(y yocki.Ycho) {
	ycho = y
}

func Progress(total int64, r io.Reader) io.Reader {
	return io.TeeReader(r, ycho.Progress(total, r))
}

func Eventloop() {
	ycho.Eventloop()
}

func Info(msg string) {
	ycho.Info(msg)
}

func Infof(msg string, v ...any) {
	ycho.Infof(msg, v...)
}

func Fatal(msg error) {
	ycho.Fatal(msg.Error())
}

func Fatalf(msg string, v ...any) {
	ycho.Fatalf(msg, v...)
}

func Debug(msg string) {
	ycho.Debug(msg)
}

func Debugf(msg string, v ...any) {
	ycho.Debugf(msg, v...)
}

func Warn(msg error) {
	ycho.Warn(msg.Error())
}

func Warnf(msg string, v ...any) {
	ycho.Warnf(msg, v...)
}

func Error(msg error) {
	ycho.Error(msg.Error())
}

func Errorf(msg string, v ...any) {
	ycho.Errorf(msg, v...)
}
