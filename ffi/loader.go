// Places not otherwise noted in code are MIT licenses.
// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ffi

// #include "yockf.h"
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/ansurfen/yock/util"
)

type Lib struct {
	plugin util.Plugin
}

func NewLibray(name string) (lib Lib, err error) {
	lib.plugin, err = util.NewPlugin(name)
	return
}

type callableFunction func(args ...any) []reflect.Value

func (lib *Lib) rawCif(name string, rType Type, aTypes ...Type) (*Cif, error) {
	str := C.CString(name)
	defer C.free(unsafe.Pointer(str))
	fn, err := lib.plugin.Func(name)
	if err != nil {
		return nil, err
	}
	cif, err := NewCif(unsafe.Pointer(fn.Addr()), rType, aTypes...)
	if err != nil {
		return nil, err
	}
	return cif, nil
}

func (lib *Lib) rawCifMust(name string, rType Type, aTypes ...Type) *Cif {
	cif, err := lib.rawCif(name, rType, aTypes...)
	if err != nil {
		panic(err)
	}
	return cif
}

func (lib *Lib) RawFunc(
	name string, rType string, argTypes []string,
	go_rtype *[]reflect.Type, go_args *[]reflect.Type,
	ffi_args *[]Type) reflect.Value {
	rtype := Archive.MustFindRecord(rType)
	if rtype.value() != typeVoid {
		*go_rtype = append(*go_rtype, rtype.value())
	}
	for _, arg := range argTypes {
		record := Archive.MustFindRecord(arg)
		*ffi_args = append(*ffi_args, record)
		*go_args = append(*go_args, record.value())
	}
	fnTmpl := reflect.FuncOf(*go_args, *go_rtype, false)
	cif := lib.rawCifMust(name, rtype, *ffi_args...)
	var out reflect.Type

	if fnTmpl.NumOut() > 1 {
		panic(fmt.Errorf("c functions can return 0 or 1 values, not %d", fnTmpl.NumOut()))
	} else if fnTmpl.NumOut() == 1 {
		out = fnTmpl.Out(0)
	}

	return reflect.MakeFunc(fnTmpl, NewFunction(cif, out).Call)
}

func (lib *Lib) Func(name string, rType string, argTypes []string) callableFunction {
	var (
		go_args  []reflect.Type
		go_rtype []reflect.Type
		ffi_args []Type
	)
	v := lib.RawFunc(name, rType, argTypes, &go_rtype, &go_args, &ffi_args)
	call := func(args ...any) []reflect.Value {
		if len(args) != len(go_args) {
			panic("args count not match")
		}
		fargs := []reflect.Value{}
		for i, arg := range args {
			if go_args[i].Kind() != reflect.TypeOf(arg).Kind() {
				panic(fmt.Errorf("type not match, want: %v, got: %v",
					go_args[i].Kind(), reflect.TypeOf(arg).Kind()))
			}
			fargs = append(fargs, reflect.ValueOf(arg))
		}
		return v.Call(fargs)
	}
	return call
}

/* Mozilla Public License {{{ */

/*
* This is a modified version of the original code (https://github.com/gogogoghost/libffigo)
* under the terms of the MPL license.
* The modifications are licensed under the MPL license.
 */

type Result struct {
	ptr unsafe.Pointer
}

func (res *Result) Pointer() unsafe.Pointer {
	return *(*unsafe.Pointer)(res.ptr)
}

func (res *Result) Uint8() uint8 {
	return *(*uint8)(res.ptr)
}

func (res *Result) Int8() int8 {
	return *(*int8)(res.ptr)
}

func (res *Result) Uint16() uint16 {
	return *(*uint16)(res.ptr)
}

func (res *Result) Int16() int16 {
	return *(*int16)(res.ptr)
}

func (res *Result) Uint32() uint32 {
	return *(*uint32)(res.ptr)
}

func (res *Result) Int32() int32 {
	return *(*int32)(res.ptr)
}

func (res *Result) Int() int {
	return *(*int)(res.ptr)
}

func (res *Result) Uint() uint {
	return *(*uint)(res.ptr)
}

func (res *Result) Uint64() uint64 {
	return *(*uint64)(res.ptr)
}

func (res *Result) Int64() int64 {
	return *(*int64)(res.ptr)
}

func (res *Result) Float() float32 {
	return *(*float32)(res.ptr)
}

func (res *Result) Double() float64 {
	return *(*float64)(res.ptr)
}

func (res *Result) String() string {
	ptr := (*C.char)(res.Pointer())
	return C.GoString(ptr)
}

type Function struct {
	cif     *Cif
	outType reflect.Type
}

func NewFunction(cif *Cif, outType reflect.Type) *Function {
	return &Function{
		cif:     cif,
		outType: outType,
	}
}

func (obj *Function) Call(rawArgs []reflect.Value) []reflect.Value {
	args := []any{}
	for _, arg := range rawArgs {
		var r any
		switch arg.Kind() {
		case reflect.Int8:
			r = int8(arg.Int())
		case reflect.Int16:
			r = int16(arg.Int())
		case reflect.Int32:
			r = int32(arg.Int())
		case reflect.Int:
			r = int32(arg.Int())
		case reflect.Int64:
			r = int64(arg.Int())
		case reflect.Uint8:
			r = uint8(arg.Uint())
		case reflect.Uint16:
			r = uint16(arg.Uint())
		case reflect.Uint32:
			r = uint32(arg.Uint())
		case reflect.Uint:
			r = uint32(arg.Uint())
		case reflect.Uint64:
			r = uint64(arg.Uint())
		case reflect.Float32:
			r = float32(arg.Float())
		case reflect.Float64:
			r = float64(arg.Float())
		case reflect.String:
			strPtr := C.CString(arg.String())
			defer freePtr(unsafe.Pointer(strPtr))
			r = strPtr
		case reflect.Pointer, reflect.UnsafePointer:
			r = arg.Pointer()
		default:
		}
		args = append(args, r)
	}
	res := obj.cif.Call(args...)
	if obj.outType == nil {
		return []reflect.Value{}
	}
	var r any
	switch obj.outType.Kind() {
	case reflect.Int8:
		r = res.Int8()
	case reflect.Int16:
		r = res.Int16()
	case reflect.Int32:
		r = res.Int32()
	case reflect.Int:
		r = int(res.Int32())
	case reflect.Int64:
		r = res.Int64()
	case reflect.Uint8:
		r = res.Uint8()
	case reflect.Uint16:
		r = res.Uint16()
	case reflect.Uint32:
		r = res.Uint32()
	case reflect.Uint:
		r = uint(res.Uint32())
	case reflect.Uint64:
		r = res.Uint64()
	case reflect.Float32:
		r = res.Float()
	case reflect.Float64:
		r = res.Double()
	case reflect.String:
		r = res.String()
	case reflect.Pointer, reflect.UnsafePointer:
		r = res.Pointer()
	default:
		r = res
	}
	return []reflect.Value{
		reflect.ValueOf(r),
	}
}

/* Mozilla Public License }}} */
