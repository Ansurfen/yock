// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package test

import (
	"io/fs"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/ansurfen/yock/util"
)

type BatchOpt struct {
	Path    string
	Filters []string
}

func Batch(opt BatchOpt) *TestSet {
	set := &TestSet{}
	filepath.Walk(opt.Path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		for _, filter := range opt.Filters {
			re := regexp.MustCompile(filter)
			tmp := util.Filename(filepath.Base(path))
			if re.MatchString(tmp) {
				return nil
			}
		}
		str, err := util.ReadStraemFromFile(path)
		if err != nil {
			return err
		}
		set.files = append(set.files, file{
			name: path,
			data: string(str),
		})
		return nil
	})
	return set
}

type TestSet struct {
	files []file
}

type file struct {
	name string
	data string
}

func (set *TestSet) Range(fn func(data string) error) {
	for _, file := range set.files {
		if err := fn(file.data); err != nil {
			panic(err)
		}
	}
}

func (set *TestSet) TRange(fn func(s *TestSet, data string) error) {
	for _, file := range set.files {
		if err := fn(set, file.data); err != nil {
			panic(err)
		}
	}
}

func (set *TestSet) FRange(fn func(s *TestSet, filename, data string) error) {
	for _, file := range set.files {
		if err := fn(set, file.name, file.data); err != nil {
			panic(err)
		}
	}
}

func (set *TestSet) Assert(err error) {
	if err != nil {
		panic(err)
	}
}

func (set *TestSet) AssertNil(v any) {
	switch vv := reflect.ValueOf(v); vv.Kind() {
	case reflect.Slice, reflect.Map:
		if !vv.IsNil() {
			panic("slice want nil")
		}
	default:

	}
}

func (set *TestSet) AssertEqual(want, got int) {
	if want != got {
		panic("length no match")
	}
}
