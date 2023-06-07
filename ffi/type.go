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
)

type Record = Type

type typeArchive struct {
	records map[string]Record
}

type FFI_FN = *[0]byte

var Archive *typeArchive

type c_void struct{}

var (
	typeError         = reflect.TypeOf((*error)(nil)).Elem()
	typeUint          = reflect.TypeOf(uint(0))
	typeUint8         = reflect.TypeOf(uint8(0))
	typeUint16        = reflect.TypeOf(uint16(0))
	typeUint32        = reflect.TypeOf(uint32(0))
	typeUint64        = reflect.TypeOf(uint64(0))
	typeInt           = reflect.TypeOf(int(0))
	typeInt8          = reflect.TypeOf(int8(0))
	typeInt16         = reflect.TypeOf(int16(0))
	typeInt32         = reflect.TypeOf(int32(0))
	typeInt64         = reflect.TypeOf(int64(0))
	typeFloat32       = reflect.TypeOf(float32(0))
	typeFloat64       = reflect.TypeOf(float64(0))
	typeBool          = reflect.TypeOf(true)
	typeUintptr       = reflect.TypeOf(uintptr(0))
	typeUnsafePointer = reflect.TypeOf(unsafe.Pointer(uintptr(NilPtr)))
	typeVoid          = reflect.TypeOf(&c_void{})
	typeStr           = reflect.TypeOf("")
)

func init() {
	Archive = &typeArchive{
		records: map[string]Record{
			"void":   &basicType{typePtr: &C.ffi_type_void, typeSize: 0, goType: typeVoid},
			"uint8":  &basicType{typePtr: &C.ffi_type_uint8, typeSize: 1, goType: typeUint},
			"int8":   &basicType{typePtr: &C.ffi_type_sint8, typeSize: 1, goType: typeInt8},
			"uint16": &basicType{typePtr: &C.ffi_type_uint16, typeSize: 2, goType: typeUint16},
			"uint32": &basicType{typePtr: &C.ffi_type_uint32, typeSize: 4, goType: typeUint32},
			"int32":  &basicType{typePtr: &C.ffi_type_sint32, typeSize: 4, goType: typeInt32},
			"uint64": &basicType{typePtr: &C.ffi_type_uint64, typeSize: 8, goType: typeUint64},
			"int64":  &basicType{typePtr: &C.ffi_type_sint64, typeSize: 8, goType: typeInt64},
			"float":  &basicType{typePtr: &C.ffi_type_float, typeSize: 4, goType: typeFloat32},
			"double": &basicType{typePtr: &C.ffi_type_double, typeSize: 8, goType: typeFloat64},
			"ptr":    &basicType{typePtr: &C.ffi_type_pointer, typeSize: int(PtrSize), goType: typeUnsafePointer},
			//
			"bool": &basicType{typePtr: &C.ffi_type_uint8, typeSize: 1, goType: typeBool},
			"str":  &basicType{typePtr: &C.ffi_type_pointer, typeSize: int(PtrSize), goType: reflect.TypeOf("")},
		},
	}
	Archive.AddRecord("int", Archive.FindRecord("int32"))
	Archive.AddRecord("long", Archive.FindRecord("int64"))
}

func (ar *typeArchive) FindRecord(name string) Type {
	if record, ok := ar.records[name]; ok {
		return record
	}
	return nil
}

func (ar *typeArchive) MustFindRecord(name string) Type {
	if record, ok := ar.records[name]; ok {
		return record
	}
	panic(fmt.Errorf("%s's record not found", name))
}

func (ar *typeArchive) AddRecord(name string, record Record) {
	ar.records[name] = record
}

func (ar *typeArchive) MustAddRecord(name string, record Record) {
	if _, ok := ar.records[name]; ok {
		panic(fmt.Errorf("%s's record is exist already", name))
	}
	ar.records[name] = record
}

type Type interface {
	ptr() *C.ffi_type

	size() int

	value() reflect.Type
}

var (
	_ Type = (*basicType)(nil)
	_ Type = (*structType)(nil)
)

type basicType struct {
	typePtr  *C.ffi_type
	typeSize int
	goType   reflect.Type
}

func (t *basicType) ptr() *C.ffi_type {
	return t.typePtr
}

func (t *basicType) size() int {
	return t.typeSize
}

func (t *basicType) value() reflect.Type {
	return t.goType
}

type structField struct {
	t Type
}

type structType struct {
	basicType
	fields []structField
}

type Field struct {
	Name string
	Type string
}

func NewStructType(name string, fields []Field) Type {
	st := &structType{}
	var element **C.ffi_type = nil
	tmp_fields := make([]*C.ffi_type, len(fields))
	for i, field := range fields {
		record := Archive.MustFindRecord(field.Name)
		st.basicType.typeSize += record.size()
		st.fields = append(st.fields, structField{t: record})
		tmp_fields[i] = record.ptr()
	}
	if len(fields) > 0 {
		element = &tmp_fields[0]
	}
	st.basicType.typePtr.elements = element
	st.basicType.typePtr._type = 13 // FFI_TYPE_STRUCT
	Archive.AddRecord(name, st)
	return st
}
