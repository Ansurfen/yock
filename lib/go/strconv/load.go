// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strconv

import (
	"strconv"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadStrconv(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("strconv")
	lib.SetField(map[string]any{
		"Itoa": strconv.Itoa,
		"Atoi": strconv.Atoi,
	})
}
