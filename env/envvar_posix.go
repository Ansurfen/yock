//go:build !windows
// +build !windows

// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocke

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	yockc "github.com/ansurfen/yock/cmd"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
)

var _ yocki.EnvVar = &PosixEnvVar{}

type PosixEnvVar struct{}

func NewEnvVar() *PosixEnvVar {
	return &PosixEnvVar{}
}

// SetPath set operate target: posix: /etc/enviroment, this only is empty method.
func (env *PosixEnvVar) SetPath(path string) error { return nil }

// set global enviroment variable
func (env *PosixEnvVar) Set(k string, v any) error {
	if _, err := yockc.Exec(yockc.ExecOpt{}, fmt.Sprintf("export %s=%s", k, v)); err != nil {
		return err
	}
	return nil
}

// set global enviroment variable when key isn't exist
func (env *PosixEnvVar) SafeSet(k string, v any) error {
	vv, err := yockc.Exec(yockc.ExecOpt{}, fmt.Sprintf("echo $%s", k))
	if err != nil {
		return err
	}
	if string(vv) != "\n" {
		if _, err := yockc.Exec(yockc.ExecOpt{}, fmt.Sprintf("export %s=%s", k, v)); err != nil {
			return err
		}
	}
	return nil
}

// set local enviroment variable
func (env *PosixEnvVar) SetL(k, v string) error {
	return os.Setenv(k, v)
}

// set local enviroment variable when key isn't exist
func (env *PosixEnvVar) SafeSetL(k, v string) error {
	exist := false
	for _, e := range os.Environ() {
		if kk, _, ok := strings.Cut(e, "="); ok && kk == k {
			exist = true
			break
		}
	}
	if !exist {
		return os.Setenv(k, v)
	}
	return errors.New("var exist already")
}

// unset (delete) global enviroment variable
func (env *PosixEnvVar) Unset(k string) error {
	if _, err := yockc.Exec(yockc.ExecOpt{}, fmt.Sprintf("unset %s", k)); err != nil {
		return err
	}
	return nil
}

// export current enviroment string into specify file
func (env *PosixEnvVar) Export(file string) error {
	dict := make(map[string]string)
	for _, e := range os.Environ() {
		if k, v, ok := strings.Cut(e, "="); ok {
			dict[k] = v
		}
	}
	raw, err := json.Marshal(dict)
	if err != nil {
		return err
	}
	return util.WriteFile(file, raw)
}

// load exported env from disk
func (env *PosixEnvVar) Load(opt yocki.EnvVarLoadOpt) error {
	raw, err := util.ReadStraemFromFile(opt.File)
	if err != nil {
		return err
	}
	dict := make(map[string]string)
	if json.Unmarshal(raw, &dict) != nil {
		return err
	}
	for _, k := range opt.Keys {
		if v, ok := dict[k]; ok {
			if opt.Safe {
				env.SafeSet(k, v)
			} else {
				env.Set(k, v)
			}
		}
	}
	return nil
}

// Print enviroment variable
func (env *PosixEnvVar) Print() {
	for _, e := range os.Environ() {
		if k, v, ok := strings.Cut(e, "="); ok {
			fmt.Printf("%s: %s", k, v)
		}
	}
}
