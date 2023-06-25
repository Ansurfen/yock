// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"strconv"
)

// JsonValue is an interface to abstract object of json type.
type JsonValue interface {
	Value() string
}

var (
	JsonTrue  = NewJsonBool(true)
	JsonFalse = NewJsonBool(false)
	JsonNil   = NewJsonNull()
)

type JsonString struct {
	v string
}

func NewJsonString(v string) JsonString {
	return JsonString{
		v: v,
	}
}

func (obj JsonString) Value() string {
	return fmt.Sprintf(`"%s"`, obj.v)
}

type JsonNumber struct {
	v string
}

func NewJsonNumber(v int64) JsonNumber {
	return JsonNumber{
		v: strconv.Itoa(int(v)),
	}
}

func (obj JsonNumber) Value() string {
	return obj.v
}

type JsonNull struct{}

func NewJsonNull() JsonNull {
	return JsonNull{}
}

func (JsonNull) Value() string {
	return "null"
}

type JsonBool struct {
	b bool
}

func NewJsonBool(b bool) JsonBool {
	return JsonBool{b: b}
}

func (obj JsonBool) Value() string {
	if obj.b {
		return "true"
	}
	return "false"
}

type JsonArray struct {
	v []JsonValue
}

func NewJsonArray(v []JsonValue) *JsonArray {
	return &JsonArray{
		v: v,
	}
}

func (obj *JsonArray) Value() string {
	res := ""
	for i := 0; i < len(obj.v); i++ {
		res += obj.v[i].Value()
		if len(obj.v)-1 != i {
			res += ", "
		}
	}
	return fmt.Sprintf("[%s]", res)
}

type JsonObject struct {
	v map[string]JsonValue
}

func NewJsonObject(v map[string]JsonValue) *JsonObject {
	return &JsonObject{
		v: v,
	}
}

func (obj *JsonObject) Value() string {
	res := ""
	idx := 0
	objLen := len(obj.v)
	for key, value := range obj.v {
		res += fmt.Sprintf(`"%s": %s`, key, value.Value())
		if objLen-1 != idx {
			res += ", "
		}
		idx++
	}
	return fmt.Sprintf(`{%s}`, res)
}

// JsonStr return json string according to JsonValue.
func JsonStr(v JsonValue) string {
	return v.Value()
}
