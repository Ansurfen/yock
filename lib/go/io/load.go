// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package io

import (
	"io/ioutil"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadIO(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("ioutil")
	lib.SetField(map[string]any{
		"ReadAll": ioutil.ReadAll,
		"ToString": func(str []byte) string {
			return string(str)
		},
	})
}
