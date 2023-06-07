// Places not otherwise noted in code are MIT licenses.
// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ffi

// #include "yockf.h"
import "C"
import (
	"reflect"
	"unsafe"
)

/* Mozilla Public License {{{ */

/*
* This is a modified version of the original code (https://github.com/gogogoghost/libffigo)
* under the terms of the MPL license.
* The modifications are licensed under the MPL license.
 */

type anyStruct struct {
	typePtr uintptr
	dataPtr uintptr
}

const NilPtr uintptr = 0

var PtrSize = unsafe.Sizeof(NilPtr)

// []T of go -> *T of C
//
// only support potiner convert
func allocArrayOf[T any](src []T) *T {
	length := len(src)
	ptr := allocArray(length)
	// convert to an array of the corresponding type
	arr := ptr2Arr[T](ptr, length)
	copy(arr[:], src)
	return &arr[0]
}

// allocArray allocates a size-sized C pointer array space and copy nothing
func allocArray(size int) unsafe.Pointer {
	// allocate an array
	ptr := C.malloc(C.size_t(int(PtrSize) * (size + 1)))
	// the last byte of the array is filled with 0
	arr := ptr2Arr[uintptr](ptr, size)
	arr[size] = uintptr(0)
	return ptr
}

// getPtrFromAny fetches the contents from the pointer memory space
// and generates a new pointer
func getPtrFromAny(ptr *any) unsafe.Pointer {
	anyPtr := (*anyStruct)(unsafe.Pointer(ptr))
	return unsafe.Pointer(anyPtr.dataPtr)
}

// allocValOf converts any type into unsafe.Pointer and copy into C memoery space
func allocValOf(src any) unsafe.Pointer {
	// get real the pointer to data
	dataPtr := getPtrFromAny(&src)

	val := reflect.ValueOf(src)
	if val.Kind() == reflect.UnsafePointer || val.Kind() == reflect.Pointer {
		// if the type is a pointer, dataPtr is actually the pointer content
		ptrValue := uintptr(dataPtr)
		dataPtr = unsafe.Pointer(&ptrValue)
	}

	realSize := val.Type().Size()

	destPtr := C.malloc(C.size_t(realSize))
	destArr := ptr2Arr[byte](destPtr, int(realSize))

	srcArr := ptr2Arr[byte](dataPtr, int(realSize))
	// it's easy to copy by array, and the raw data is'nt array.
	// thus, the last 0 isn't exist in memory layout,
	// so it's essential for specifing length to be copied,
	// otherwise it's easy to copy out of bounds.
	copy(destArr[0:realSize], srcArr[0:realSize])
	return destPtr
}

func alloc(size int) unsafe.Pointer {
	ptr := C.malloc(C.size_t(size))
	return ptr
}

// freePtr frees the memory of pointer
func freePtr(ptr unsafe.Pointer) {
	C.free(ptr)
}

// allocParams converts any[] into void*
func allocParams(args []any) *unsafe.Pointer {
	count := len(args)
	var argp *unsafe.Pointer

	arrPtr := allocArray(count)

	arr := ptr2Arr[unsafe.Pointer](arrPtr, count)
	// write array into C memory space
	for index, arg := range args {
		// allocates memory space for each variable
		ptr := allocValOf(arg)
		arr[index] = ptr
	}
	argp = &(arr[0])
	return argp
}

// freeParams frees void** and frees all memory within the array
func freeParams(ptr *unsafe.Pointer) {
	arrPtr := unsafe.Pointer(ptr)
	ptrAddr := uintptr(arrPtr)
	for {
		// take out the pointer to the data
		dataPtr := *(*unsafe.Pointer)(unsafe.Pointer(ptrAddr))
		// 0 indicates that the array is over
		if uintptr(dataPtr) == 0 {
			break
		}
		C.free(dataPtr)
		ptrAddr += PtrSize
	}
	C.free(arrPtr)
}

// ptr2Arr converts ptr into []T
func ptr2Arr[T any](ptr unsafe.Pointer, length int) []T {
	sliceHeader := struct {
		p   unsafe.Pointer
		len int
		cap int
	}{ptr, length + 1, length + 1}
	return *(*[]T)(unsafe.Pointer(&sliceHeader))
}

/* Mozilla Public License }}} */
