// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/ycho"
)

type mylog struct{}

func (m *mylog) Info(msg string) {
	log.Println(msg)
}

func (m *mylog) Infof(msg string, a ...any) {
	log.Printf(msg, a...)
}

func (m *mylog) Debug(msg string) {
	log.Println(msg)
}

func (m *mylog) Debugf(msg string, a ...any) {
	log.Printf(msg, a...)
}

func (m *mylog) Fatal(msg string) {
	log.Fatal(msg)
}

func (m *mylog) Fatalf(msg string, a ...any) {
	log.Fatalf(msg, a...)
}

func (m *mylog) Warn(msg string) {
	log.Println(msg)
}

func (m *mylog) Warnf(msg string, a ...any) {
	log.Printf(msg, a...)
}

func (m *mylog) Error(msg string) {
	log.Println(msg)
}

func (m *mylog) Errorf(msg string, a ...any) {
	log.Printf(msg, a...)
}

var _ yocki.Ycho = (*mylog)(nil)

func main() {
	ycho.SetYcho(&mylog{})
	ycho.Info("Hello World!")
	ycho.Fatalf("1 == 2 -> %v", 1 == 2)
}
