// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package test

func Fn1() {}

func Fn2() error {
	return nil
}

func Fn3(a, b string, c int) error {
	return nil
}

func Fn4(a string, b StrArr) (*Stu1, error) {
	return nil, nil
}

func Fn5(a StrArr, b *Stu2, c ...[]map[*Stu1]string) (s Stu3, e error) {
	return
}

func Fn6(a ...string) (s Stu3, e error) {
	return
}
