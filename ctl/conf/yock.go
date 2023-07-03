// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
)

const YockConfTmpl = `lang: en_us
ycho:
  level: debug
  compress: false
  filename: yock.log
  fileMaxSize: 1024
  fileMaxBackups: 0
  stdout: true
  path: %s/log
`

var conf *YockConf

type YockConf struct {
	Lang  string        `yaml:"lang"`
	Ycho  ycho.YchoOpt  `yaml:"ycho"`
	Yocks yockScheduler `yaml:"yocks"`
	Yockd yockDaemon    `yaml:"yockd"`
}

func (c *YockConf) String() string {

	return ""
}

func (c *YockConf) Restore() error {
	return util.WriteFile(util.Pathf("@/conf.ymal"), []byte(fmt.Sprintf(YockConfTmpl, util.WorkSpace)))
}

func Instance() *YockConf {
	return conf
}
