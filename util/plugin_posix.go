//go:build !windows
// +build !windows

// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

// #include <dlfcn.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var (
	_ Plugin     = (*PosixPlugin)(nil)
	_ PluginFunc = (*PosixPluginFunc)(nil)
)

type PosixPlugin struct {
	handle unsafe.Pointer
}

func NewPlugin(path string) (Plugin, error) {
	c_str := C.CString(path)
	defer C.free(unsafe.Pointer(c_str))

	h := C.dlopen(c_str, C.int(C.RTLD_NOW))
	if h == nil {
		return nil, fmt.Errorf("%s", C.GoString(C.dlerror()))
	}

	return &PosixPlugin{
		handle: h,
	}, nil
}

// Func return PluginFunc which is an abstract function to be exported dynamic library
// according to funcName. You can use PluginFunc to call function from dynamic library.
func (pp *PosixPlugin) Func(name string) (PluginFunc, error) {
	c_sym := C.CString(name)
	defer C.free(unsafe.Pointer(c_sym))
	c_addr := C.dlsym(pp.handle, c_sym)
	if c_addr == nil {
		return nil, fmt.Errorf("%s not found", name)
	}
	return &PosixPluginFunc{
		sym: uintptr(c_addr),
	}, nil
}

type PosixPluginFunc struct {
	sym uintptr
}

// Call return excuted result from dynamic library
func (ppf *PosixPluginFunc) Call(params ...uintptr) (uintptr, error) {
	return uintptr(0), errors.New("fail to assert func")
}

// Call return excuted result from dynamic library
func (ppf *PosixPluginFunc) Addr() uintptr {
	return ppf.sym
}
