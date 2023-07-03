// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import yocki "github.com/ansurfen/yock/interface"

var ycho yocki.Ycho

func init() {
	ycho = &vlog{}
}

func GetYcho() yocki.Ycho {
	return ycho
}

func SetYcho(y yocki.Ycho) {
	ycho = y
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
