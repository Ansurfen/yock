// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
)

const YockConfTmpl = `strict: false
lang: en_us
ycho:
  level: debug
  compress: false
  filename: yock.log
  fileMaxSize: 1024
  fileMaxBackups: 0
  stdout: true
  path: "@/log"
yockd:
  ip: localhost
  port: 1314
  self_boot: false
`

var conf *YockConf

type YockConf struct {
	Lang   string        `yaml:"lang"`
	Strict bool          `yaml:"strict"`
	Ycho   ycho.YchoOpt  `yaml:"ycho"`
	Yocks  yockScheduler `yaml:"yocks"`
	Yockd  yockDaemon    `yaml:"yockd"`
	Yockw  yockWatch     `yaml:"yockw"`
}

func (c *YockConf) Restore() error {
	return util.WriteFile(util.Pathf("@/conf.ymal"), []byte(YockConfTmpl))
}

func Instance() *YockConf {
	return conf
}
