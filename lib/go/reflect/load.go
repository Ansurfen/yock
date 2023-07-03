// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reflectlib

import (
	yocki "github.com/ansurfen/yock/interface"
	"reflect"
)

func LoadReflect(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("reflect")
	lib.SetField(map[string]any{
		// functions
		"Select":          reflect.Select,
		"Swapper":         reflect.Swapper,
		"FuncOf":          reflect.FuncOf,
		"MakeMap":         reflect.MakeMap,
		"MakeSlice":       reflect.MakeSlice,
		"MakeChan":        reflect.MakeChan,
		"ValueOf":         reflect.ValueOf,
		// "ArenaNew":        reflect.ArenaNew,
		"DeepEqual":       reflect.DeepEqual,
		"MakeFunc":        reflect.MakeFunc,
		"ChanOf":          reflect.ChanOf,
		"Zero":            reflect.Zero,
		"PointerTo":       reflect.PointerTo,
		"AppendSlice":     reflect.AppendSlice,
		"Copy":            reflect.Copy,
		"StructOf":        reflect.StructOf,
		"PtrTo":           reflect.PtrTo,
		"Indirect":        reflect.Indirect,
		"SliceOf":         reflect.SliceOf,
		"New":             reflect.New,
		"NewAt":           reflect.NewAt,
		"ArrayOf":         reflect.ArrayOf,
		"MakeMapWithSize": reflect.MakeMapWithSize,
		"Append":          reflect.Append,
		"VisibleFields":   reflect.VisibleFields,
		"MapOf":           reflect.MapOf,
		"TypeOf":          reflect.TypeOf,
		// constants
		"Invalid":       reflect.Invalid,
		"Bool":          reflect.Bool,
		"Int":           reflect.Int,
		"Int8":          reflect.Int8,
		"Int16":         reflect.Int16,
		"Int32":         reflect.Int32,
		"Int64":         reflect.Int64,
		"Uint":          reflect.Uint,
		"Uint8":         reflect.Uint8,
		"Uint16":        reflect.Uint16,
		"Uint32":        reflect.Uint32,
		"Uint64":        reflect.Uint64,
		"Uintptr":       reflect.Uintptr,
		"Float32":       reflect.Float32,
		"Float64":       reflect.Float64,
		"Complex64":     reflect.Complex64,
		"Complex128":    reflect.Complex128,
		"Array":         reflect.Array,
		"Chan":          reflect.Chan,
		"Func":          reflect.Func,
		"Interface":     reflect.Interface,
		"Map":           reflect.Map,
		"Pointer":       reflect.Pointer,
		"Slice":         reflect.Slice,
		"String":        reflect.String,
		"Struct":        reflect.Struct,
		"UnsafePointer": reflect.UnsafePointer,
		"Ptr":           reflect.Ptr,
		"RecvDir":       reflect.RecvDir,
		"SendDir":       reflect.SendDir,
		"BothDir":       reflect.BothDir,
		"SelectSend":    reflect.SelectSend,
		"SelectRecv":    reflect.SelectRecv,
		"SelectDefault": reflect.SelectDefault,
		// variable
	})
}
