// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import yocki "github.com/ansurfen/yock/interface"

func LoadBit(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("bit")
	lib.SetField(map[string]any{
		"And": func(a, b uint32) uint32 {
			return a & b
		},
		"Or": func(a, b uint32) uint32 {
			return a | b
		},
	})
}
