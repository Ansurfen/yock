// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

const (
	PermRead  = iota // download
	PermWrite        // upload

	PermCount
)

type Perm int

func (p Perm) String() string {
	switch p {
	case PermRead:
		return "read"
	case PermWrite:
		return "write"
	default:
		panic("unreachable")
	}
}

var findPerm = map[string]int{
	"read":  PermRead,
	"write": PermWrite,
	"all":   PermCount,
}
