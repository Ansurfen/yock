// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	yocki "github.com/ansurfen/yock/interface"
)

const (
	Y_STRICT = 1
	Y_DEBUG  = 2
)

type YockModeManager int32

func (ym *YockModeManager) Mode() int32 {
	return int32(*ym)
}

func (ym *YockModeManager) SetMode(m int32) {
	*ym = YockModeManager(m | ym.Mode())
}

func (ym *YockModeManager) Strict() bool {
	return (ym.Mode()>>0)&1 == 1
}

func (ym *YockModeManager) Debug() bool {
	return (ym.Mode()>>1)&1 == 1
}

func init() {
	yocki.Y_MODE = new(YockModeManager)
}
