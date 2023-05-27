// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ypm

import (
	"encoding/json"
	"os"
	"path/filepath"

	yockc "github.com/ansurfen/yock/cmd"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/cushion/utils/build"
	"github.com/ansurfen/yock/util"
)

type YpmNewOpt struct {
	Module    string
	Lang      string
	Version   string
	CreateDir bool
}

func New(opt YpmNewOpt) error {
	out, err := utils.ReadStraemFromFile(util.Pathf("~/ypm/index.tpl"))
	if err != nil {
		return err
	}
	tmpl := build.NewTemplate()
	str, err := tmpl.OnceParse(string(out), map[string]string{
		"version": opt.Version,
	})
	if err != nil {
		return err
	}
	if opt.CreateDir {
		if err := utils.Mkdirs(opt.Module); err != nil {
			return err
		}
		if err = yockc.Cd(opt.Module); err != nil {
			return err
		}
	}
	if err = utils.WriteFile("index.lua", []byte(str)); err != nil {
		return err
	}
	include_path := util.Pathf("~/lib/include")
	files, err := os.ReadDir(include_path)
	if err != nil {
		return err
	}
	out, err = utils.ReadStraemFromFile(filepath.Join(include_path, "lang", "zh_cn.json"))
	if err != nil {
		return err
	}
	var doc map[string]any
	if err = json.Unmarshal(out, &doc); err != nil {
		return err
	}
	if err := utils.SafeMkdirs("include"); err != nil {
		return err
	}
	for _, file := range files {
		if fn := file.Name(); filepath.Ext(fn) == ".lua" {
			out, err = utils.ReadStraemFromFile(filepath.Join(include_path, fn))
			if err != nil {
				return err
			}
			str, err = tmpl.OnceParse(string(out), doc)
			if err != nil {
				return err
			}
			if err = utils.WriteFile("include/"+fn, []byte(str)); err != nil {
				return err
			}
		}
	}
	return nil
}
