// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ansurfen/yock/util"
)

type FindOpt struct {
	Search  bool
	Pattern string
	Case    bool
	Dir     bool
	File    bool
}

func Find(opt FindOpt, path string) ([]string, error) {
	if !util.IsExist(path) {
		return nil, util.ErrPathNotFound
	}
	if !opt.Search {
		return nil, nil
	}
	if len(opt.Pattern) == 0 {
		return nil, nil
	}
	re := regexp.MustCompile(opt.Pattern)
	res := []string{}
	if opt.Case {
		path = strings.ToLower(path)
	}
	return res, filepath.Walk(path, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if !opt.Dir {
				return nil
			}
		} else {
			if !opt.File {
				return nil
			}
		}
		if opt.Case {
			p = strings.ToLower(p)
		}
		if re.MatchString(p) {
			res = append(res, p)
		}
		return nil
	})
}
