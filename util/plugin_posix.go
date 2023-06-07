//go:build !windows
// +build !windows

// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"plugin"
)

var (
	_ Plugin     = (*PosixPlugin)(nil)
	_ PluginFunc = (*PosixPluginFunc)(nil)
)

type PosixPlugin struct {
	*plugin.Plugin
}

func NewPlugin(path string) (Plugin, error) {
	plugin, err := plugin.Open(path)
	return &PosixPlugin{
		Plugin: plugin,
	}, err
}

// Func return PluginFunc which is an abstract function to be exported dynamic library
// according to funcName. You can use PluginFunc to call function from dynamic library.
func (pp *PosixPlugin) Func(name string) (PluginFunc, error) {
	sym, err := pp.Lookup(name)
	if err != nil {
		return nil, err
	}
	return &PosixPluginFunc{
		Symbol: sym,
	}, nil
}

type PosixPluginFunc struct {
	plugin.Symbol
}

// Call return excuted result from dynamic library
func (ppf *PosixPluginFunc) Call(params ...uintptr) (uintptr, error) {
	switch f := ppf.Symbol.(type) {
	case func(...uintptr) uintptr:
		return f(params...), nil
	default:
		return uintptr(0), errors.New("fail to assert func")
	}
}

// Call return excuted result from dynamic library
func (ppf *PosixPluginFunc) Addr() uintptr {
	return ppf.Symbol.(uintptr)
}
