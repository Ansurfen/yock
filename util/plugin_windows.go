//go:build windows
// +build windows

// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "syscall"

var (
	_ Plugin     = (*WindowsPlugin)(nil)
	_ PluginFunc = (*WindowsPluginFunc)(nil)
)

type WindowsPlugin struct {
	*syscall.LazyDLL
}

func NewPlugin(path string) (Plugin, error) {
	return &WindowsPlugin{
		LazyDLL: syscall.NewLazyDLL(path),
	}, nil
}

// Func return PluginFunc which is an abstract function to be exported dynamic library
// according to funcName. You can use PluginFunc to call function from dynamic library.
func (wp *WindowsPlugin) Func(plugin string) (PluginFunc, error) {
	return &WindowsPluginFunc{
		LazyProc: wp.LazyDLL.NewProc(plugin),
	}, nil
}

type WindowsPluginFunc struct {
	*syscall.LazyProc
}

// Addr returns the address of function pointer
func (wpf *WindowsPluginFunc) Addr() uintptr {
	return wpf.LazyProc.Addr()
}

// Call return excuted result from dynamic library
func (wpf *WindowsPluginFunc) Call(params ...uintptr) (uintptr, error) {
	ret, _, err := wpf.LazyProc.Call(params...)
	return ret, err
}
