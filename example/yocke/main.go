// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	yocke "github.com/ansurfen/yock/env"
)

type YockeTestConf struct {
	Author  string `yaml:"author"`
	Version string `yaml:"version"`
}

func main() {
	// defer yocke.FreeEnv[YockeTestConf]()
	yocke.InitEnv(&yocke.EnvOpt[YockeTestConf]{
		Workdir: ".yocke",
		Subdirs: []string{"tmp", "log", "mnt", "unmnt"},
		Conf:    YockeTestConf{},
	})
	env := yocke.GetEnv[YockeTestConf]()
	defer env.Save()
	env.SetValue("author", "a")
	fmt.Println(env.Viper())
}
