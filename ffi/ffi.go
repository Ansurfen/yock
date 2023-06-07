// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ffi

/*
 #cgo CFLAGS: -I.
 #cgo LDFLAGS: -L. -lffi
 #include "yockf.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type ffiABI C.ffi_abi

const (
	ABIFirst   ffiABI = C.FFI_FIRST_ABI
	ABIDefault ffiABI = C.FFI_DEFAULT_ABI
	ABILast    ffiABI = C.FFI_LAST_ABI
)

type ffiStatus uint32

func (s ffiStatus) String() string {
	switch s {
	case FFIStatusOk:
		return "C.FFI_OK"
	case FFIStatusBadTypeDef:
		return "C.FFI_BAD_TYPEDEF"
	case FFIStatusBadAbi:
		return "C.FFI_BAD_ABI"
	}
	panic("unknown error")
}

const (
	FFIStatusOk         ffiStatus = C.FFI_OK
	FFIStatusBadTypeDef ffiStatus = C.FFI_BAD_TYPEDEF
	FFIStatusBadAbi     ffiStatus = C.FFI_BAD_ABI
)

type Cif struct {
	ptr   *C.ffi_cif
	fPtr  unsafe.Pointer
	argn  int
	rType Type
}

func NewCif(fPtr unsafe.Pointer, rType Type, aTypes ...Type) (cif *Cif, err error) {
	cif = &Cif{
		ptr:   (*C.ffi_cif)(C.malloc(C.size_t(unsafe.Sizeof(C.ffi_cif{})))),
		fPtr:  fPtr,
		argn:  len(aTypes),
		rType: rType,
	}
	var argsPtr **C.ffi_type
	if cif.argn > 0 {
		typesArr := make([]*C.ffi_type, cif.argn)
		for index, aType := range aTypes {
			typesArr[index] = aType.ptr()
		}
		argsPtr = allocArrayOf(typesArr)
	}
	// free memeory by hand when object was destoryed
	runtime.SetFinalizer(cif, func(cif *Cif) {
		freePtr(unsafe.Pointer(cif.ptr))
		if argsPtr != nil {
			freePtr(unsafe.Pointer(argsPtr))
		}
	})
	status := C.ffi_prep_cif(
		cif.ptr,
		C.ffi_abi(ABIDefault),
		C.uint(cif.argn),
		rType.ptr(),
		argsPtr,
	)
	if status != C.FFI_OK {
		return nil, fmt.Errorf("error while preparing cif (%s)",
			ffiStatus(status))
	}
	return cif, nil
}

func (cif *Cif) Call(args ...any) *Result {
	if len(args) != cif.argn {
		panic("args count not match")
	}
	// copy into C memory space
	argp := allocParams(args)
	defer freeParams(argp)
	var resPtr unsafe.Pointer = nil
	resSize := cif.rType.size()
	var resArr []byte
	if resSize > 0 {
		// allocates space to hold temporary return values
		resPtr = alloc(resSize)
		defer freePtr(resPtr)
		// then allocates a go space to store the copied data,
		// which is managed by GC
		resArr = make([]byte, resSize)
	}
	C.ffi_call(cif.ptr, (FFI_FN)(cif.fPtr), resPtr, argp)
	if resSize > 0 {
		// returns copied address
		tmpArr := ptr2Arr[byte](resPtr, resSize)
		copy(resArr, tmpArr)
		return &Result{ptr: unsafe.Pointer(&resArr[0])}
	}
	return nil
}
