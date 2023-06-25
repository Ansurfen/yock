// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"os/user"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

var env *BaseEnv

const (
	BlankConf = `workdir: ""`
	ConfFile  = "conf.yaml"
)

// GetEnv returns global BaseEnv pointer
func GetEnv() *BaseEnv {
	return env
}

type BaseEnv struct {
	workdir string `yaml:"workdir"`
	user    *user.User
	conf    *viper.Viper
	file    string
}

type EnvOpt[T any] struct {
	Payload   T
	Workdir   string
	Subdirs   []string
	BlankConf string
}

// NewEnv to init BaseEnv and return Payload pointer which will be initialized from specify configure file in opt.
func NewEnv[T any](opt EnvOpt[T]) *T {
	env.workdir = filepath.ToSlash(path.Join(env.workdir, opt.Workdir))
	if ok, err := PathIsExist(env.workdir); err != nil {
		panic(err)
	} else if !ok {
		if err := env.initWorkspace(opt.Subdirs); err != nil {
			panic(err)
		}
	}
	env.file = path.Join(env.workdir, ConfFile)
	if ok, err := PathIsExist(env.file); err != nil {
		panic(err)
	} else if ok {
		env.ReadWithBind(env.file, &opt.Payload)
	} else {
		bc := opt.BlankConf
		if len(bc) == 0 {
			bc = BlankConf
		}
		err := WriteFile(env.file, []byte(bc))
		if err != nil {
			panic(err)
		}
		env.ReadWithBind(env.file, &opt.Payload)
	}

	return &opt.Payload
}

func NewBaseEnv() *BaseEnv {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return &BaseEnv{
		user:    curUser,
		workdir: curUser.HomeDir,
	}
}

func (env *BaseEnv) Workdir() string {
	return env.workdir
}

// initWorkspace batch to create directory when BaseEnv initializes.
func (env *BaseEnv) initWorkspace(dirs []string) error {
	workdir := env.workdir
	if err := SafeMkdirs(workdir); err != nil {
		return err
	}
	for _, dir := range dirs {
		dir = path.Join(workdir, dir)
		if err := SafeMkdirs(dir); err != nil {
			return err
		}
	}
	return nil
}

// Dump reprensent env on console
func (env *BaseEnv) Dump() {
	fmt.Println("workdir: ", env.workdir)
}

func (env *BaseEnv) Read(path string) {
	var err error
	env.conf, err = OpenConf(path)
	if err != nil {
		panic(err)
	}
	if wd := env.conf.GetString("workdir"); len(wd) > 0 {
		env.workdir = filepath.ToSlash(wd)
	}
	if err := env.conf.Unmarshal(env); err != nil {
		panic(err)
	}
}

func (env *BaseEnv) ReadWithBind(path string, payload any) {
	var err error
	env.conf, err = OpenConf(path)
	if err != nil {
		panic(err)
	}
	if wd := env.conf.GetString("workdir"); len(wd) > 0 {
		env.workdir = filepath.ToSlash(wd)
	}
	if err := env.conf.Unmarshal(payload); err != nil {
		panic(err)
	}
}

// Commit push key to configure file. It'll be write in disk throught Write.
func (env *BaseEnv) Commit(key string, value any) *BaseEnv {
	env.conf.Set(key, value)
	return env
}

// Write to persist configure file in disk
func (env *BaseEnv) Write() {
	if err := env.conf.WriteConfig(); err != nil {
		panic(err)
	}
}
