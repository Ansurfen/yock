// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import (
	"fmt"
	"time"

	yocki "github.com/ansurfen/yock/interface"
)

var _ yocki.Ycho = (*Ylog)(nil)

type Ylog struct {
	Vlog
}

func NewYLog() *Ylog {
	return new(Ylog)
}

func (y *Ylog) Info(msg string) {
	y.logger("INFO", msg)
}

func (y *Ylog) Infof(msg string, a ...any) {
	y.logger("INFO", msg, a...)
}

func (y *Ylog) Debug(msg string) {
	y.logger("DEBUG", msg)
}

func (y *Ylog) Debugf(msg string, a ...any) {
	y.logger("DEBUG", msg, a...)
}

func (y *Ylog) Warn(msg string) {
	y.logger("WARN", msg)
}

func (y *Ylog) Warnf(msg string, a ...any) {
	y.logger("WARN", msg, a...)
}

func (y *Ylog) Error(msg string) {
	y.logger("ERROR", msg)
}

func (y *Ylog) Errorf(msg string, a ...any) {
	y.logger("ERROR", msg, a...)
}

func (y *Ylog) Fatal(msg string) {
	y.logger("FATAL", msg)
}

func (y *Ylog) Fatalf(msg string, a ...any) {
	y.logger("FATAL", msg, a...)
}

func (y *Ylog) logger(level, msg string, a ...any) {
	fr := getTopCaller(3)
	fmt.Printf("%s %s %s:%d %s\n",
		time.Now().Format(defaultTimeFormat),
		level, fr.name, fr.line, fmt.Sprintf(msg, a...))
}
