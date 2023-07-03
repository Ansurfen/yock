// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
)

var _ yocki.YockLoader = (*yockLoader)(nil)

type yockLoader struct {
	yocki.YockState
}

func NewYockLoader(s yocki.YockState) *yockLoader {
	return &yockLoader{s}
}

func (loader *yockLoader) CreateLib(name string) yocki.YockLib {
	return yockr.CreateYockLib(loader.YockState, name)
}

func (loader *yockLoader) OpenLib(name string) yocki.YockLib {
	return yockr.OpenYockLib(loader.YockState, name)
}

func (loader *yockLoader) RegLuaFn(v yocki.LuaFuncs) {
	for name, fn := range v {
		loader.LState().SetGlobal(name, loader.LState().NewFunction(fn))
	}
}

func (loader *yockLoader) RegYockFn(v yocki.YockFuns) {
	for name, fn := range v {
		loader.LState().SetGlobal(name, loader.NewYFunction(fn))
	}
}

func (loader *yockLoader) State() yocki.YockState {
	return loader.YockState
}
