// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocke

import (
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/ansurfen/yock/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const defaultConfName = "config.yaml"

type Env[T any] struct {
	u    *user.User
	conf *viper.Viper
	opt  *EnvOpt[T]
}

func (e *Env[T]) String() string {
	return ""
}

type EnvOpt[T any] struct {
	Conf     T
	Workdir  string
	Subdirs  []string
	ConfTmpl string
	Filename string
}

func (opt *EnvOpt[T]) String() string {
	return ""
}

func InitEnv[T any](opt *EnvOpt[T], opts ...viper.Option) *Env[T] {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	e := &Env[T]{
		u:   u,
		opt: opt,
	}

	opt.Workdir = filepath.ToSlash(path.Join(u.HomeDir, opt.Workdir))
	if exist, err := util.PathIsExist(opt.Workdir); err != nil {
		panic(err)
	} else if !exist {
		util.Mkdirs(opt.Workdir)
	}
	for _, sub := range opt.Subdirs {
		if err := util.SafeMkdirs(filepath.Join(opt.Workdir, sub)); err != nil {
			panic(err)
		}
	}

	filename := opt.Filename
	if len(filename) == 0 {
		filename = defaultConfName
		opt.Filename = defaultConfName
	}
	filename = filepath.Join(opt.Workdir, filename)

	if exist, err := util.PathIsExist(filename); err != nil {
		panic(err)
	} else if !exist {
		var raw []byte
		if len(opt.ConfTmpl) != 0 {
			raw = []byte(opt.ConfTmpl)
		} else {
			raw, err = yaml.Marshal(&opt.Conf)
			if err != nil {
				panic(err)
			}
		}
		err = util.WriteFile(filename, raw)
		if err != nil {
			panic(err)
		}
	}
	e.bind(&opt.Conf, opts...)
	env = e
	return e
}

func (e *Env[T]) bind(d any, opts ...viper.Option) {
	conf, err := util.OpenConf(filepath.Join(e.opt.Workdir, e.opt.Filename), opts...)
	if err != nil {
		panic(err)
	}
	e.conf = conf
	if err := e.conf.Unmarshal(d); err != nil {
		panic(err)
	}
}

func (e *Env[T]) User() *user.User {
	return e.u
}

func (e *Env[T]) Conf() T {
	return e.opt.Conf
}

func (e *Env[T]) Viper() *viper.Viper {
	return e.conf
}

func (e *Env[T]) SetValue(k string, v any) {
	e.conf.Set(k, v)
}

func (e *Env[T]) Save(files ...string) error {
	if len(files) == 0 {
		return e.conf.WriteConfig()
	}
	for _, file := range files {
		err := e.conf.WriteConfigAs(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetEnv[T any]() *Env[T] {
	return env.(*Env[T])
}

func FreeEnv[T any]() {
	e := env.(*Env[T])
	if err := os.RemoveAll(e.opt.Workdir); err != nil {
		panic(err)
	}
}

var env any
