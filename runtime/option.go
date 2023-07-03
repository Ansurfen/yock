// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	yocki "github.com/ansurfen/yock/interface"
	lua "github.com/yuin/gopher-lua"
)

func OptionLState(opt lua.Options) YockrOption {
	return func(yockr yocki.YockRuntime) error {
		state := lua.NewState(opt)
		yockr.SetState(UpgradeLState(state))
		return nil
	}
}

func OptionEnableInterpPool() YockrOption {
	return func(yockr yocki.YockRuntime) error {
		yockr = UpgradeInterpPool(yockr)
		return nil
	}
}
