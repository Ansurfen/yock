// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ffi

import (
	"errors"
	"reflect"
)

type Builder struct {
	fields []reflect.StructField
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) addField(field string, typ reflect.Type) *Builder {
	b.fields = append(b.fields, reflect.StructField{Name: field, Type: typ})
	return b
}

func (b *Builder) Build() *Struct {
	stu := reflect.StructOf(b.fields)
	index := make(map[string]int)
	for i := 0; i < stu.NumField(); i++ {
		index[stu.Field(i).Name] = i
	}
	return &Struct{stu, index}
}

func (b *Builder) AddString(name string) *Builder {
	return b.addField(name, reflect.TypeOf(""))
}

func (b *Builder) AddBool(name string) *Builder {
	return b.addField(name, reflect.TypeOf(true))
}

func (b *Builder) AddInt64(name string) *Builder {
	return b.addField(name, reflect.TypeOf(int64(0)))
}

func (b *Builder) AddInt(name string) *Builder {
	return b.addField(name, reflect.TypeOf(int(0)))
}

func (b *Builder) AddFloat64(name string) *Builder {
	return b.addField(name, reflect.TypeOf(float64(1.2)))
}

type Struct struct {
	typ   reflect.Type
	index map[string]int
}

func (s Struct) New() *Instance {
	return &Instance{reflect.New(s.typ).Elem(), s.index}
}

type Instance struct {
	instance reflect.Value
	index    map[string]int
}

var (
	ErrFieldNoExist error = errors.New("field no exist")
)

func (in Instance) Field(name string) (reflect.Value, error) {
	if i, ok := in.index[name]; ok {
		return in.instance.Field(i), nil
	} else {
		return reflect.Value{}, ErrFieldNoExist
	}
}
func (in *Instance) SetString(name, value string) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetString(value)
	}
}

func (in *Instance) SetBool(name string, value bool) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetBool(value)
	}
}

func (in *Instance) SetInt64(name string, value int64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetInt(value)
	}
}

func (in *Instance) SetInt(name string, value int) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetInt(int64(value))
	}
}

func (in *Instance) SetFloat64(name string, value float64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetFloat(value)
	}
}
func (i *Instance) Interface() interface{} {
	return i.instance.Interface()
}

func (i *Instance) Addr() interface{} {
	return i.instance.Addr().Interface()
}
