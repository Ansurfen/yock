// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"

	"github.com/ansurfen/cushion/utils"
)

var conf *yockConf

type yockConf struct {
	Ycho YchoOpt `yaml:"ycho"`
	Lang string  `yaml:"lang"`
}

// Restore configuration file to initial state
func (c *yockConf) Restore() error {
	return utils.WriteFile(Pathf("@/conf.ymal"), []byte(fmt.Sprintf(yockConfTmpl, WorkSpace)))
}

func Conf() *yockConf {
	return conf
}
