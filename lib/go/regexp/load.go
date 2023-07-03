// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package regexp

import (
	"regexp"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadRegexp(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("regexp")
	lib.SetYFunction(map[string]yocki.YGFunction{
		"Compile":     regexpCompile,
		"MustCompile": regexpMustCompile,
	})
}

// @param expr string
//
// @return userdata (*regexp.Regexp), err
func regexpCompile(l yocki.YockState) int {
	r, err := regexp.Compile(l.LState().CheckString(1))
	l.Pusha(r).PushError(err)
	return 0
}

// @param expr string
//
// @return userdata (*regexp.Regexp)
func regexpMustCompile(l yocki.YockState) int {
	r := regexp.MustCompile(l.LState().CheckString(1))
	l.Pusha(r)
	return 1
}
