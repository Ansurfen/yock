// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package test

type Stu1 struct {
	S string
	I int
	N *Stu1
	M map[string]*Stu1
}

type Stu2 struct {
	S1  Stu1
	Ss1 *Stu1
	Ms  map[*Stu1][]string
}

type StrArr []string

type ID int64

type (
	Name   string
	Gender bool
	Stu3   struct {
		Str string
		Stu struct {
			S1 Stu1
			S2 *Stu2
		}
	}
)
