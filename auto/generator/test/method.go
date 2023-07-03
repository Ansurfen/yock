// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package test

import (
	"archive/zip"

	yocke "github.com/ansurfen/yock/env"
)

type Stu4 struct {
}

func (Stu1) Fn1(a, b string, c int) error {
	return nil
}

func (s *Stu4) Fn2() error {
	return nil
}

func (s *Stu4) Fn3(a, b string, c int) error {
	return nil
}

func (s *Stu4) Fn4(a string, b StrArr) (*Stu1, error) {
	return nil, nil
}

func (s *Stu4) Fn5(a StrArr, b *Stu2, c ...[]map[*Stu1]string) (s3 Stu3, e error) {
	return
}

func (s *Stu4) Fn6(z zip.Compressor, a yocke.Env[int]) {

}
