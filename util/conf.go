// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
)

var conf *yockConf

type yockConf struct {
	Lang  string        `yaml:"lang"`
	Ycho  YchoOpt       `yaml:"ycho"`
	Yocks yockScheduler `yaml:"yocks"`
	Yockd yockDaemon    `yaml:"yockd"`
	Yocki yockInterface `yaml:"yocki"`
}

type yockScheduler struct {
	Goroutine  yockGoroutine `yaml:"goroutine"`
	InterpPool bool          `yaml:"interPool"`
	MaxInterp  int           `yaml:"maxInterp"`
}

type yockGoroutine struct {
	MaxGoroutine int64 `yaml:"maxGoroutine"`
	MaxWaitRound int64 `yaml:"maxWaitRound"`
	RoundStep    int   `yaml:"roundStep"`
}

type yockDaemon struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	// MTL is abbreviation to max transfer length for file
	MTL  int    `yaml:"MTL"`
	Name string `yaml:"name"`
}

type yockInterface struct{}

// Restore configuration file to initial state
func (c *yockConf) Restore() error {
	return WriteFile(Pathf("@/conf.ymal"), []byte(fmt.Sprintf(yockConfTmpl, WorkSpace)))
}

func Conf() *yockConf {
	return conf
}
