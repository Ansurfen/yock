// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package archive

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

const (
	commentDoc    = "---{{.%s}}\n"
	commentParam  = "---@param %s %s\n"
	commentVargs  = "---@vararg %s\n"
	commentResult = "---@return %s\n"
	commentField  = "---@field %s %s\n"
	defineStruct  = "---@class %s\n%slocal %s = {}"
)

var nameFilter = map[string]string{
	"in":  "in_",
	"out": "out_",
}

func nameMap(s string) string {
	if v, ok := nameFilter[s]; ok {
		return v
	}
	return s
}

type Type struct {
	Name     string
	Ident    string
	Comments []string
}

type Function struct {
	Name     string
	Params   []Type
	Results  []Type
	Comments []string
}

func (fn *Function) GoString() string {
	return ""
}

func (fn *Function) LuaString() string {
	var (
		params        []string
		results       []string
		commentParams bytes.Buffer
		res           string
	)
	commentParams.WriteString(commentDoc)
	for _, param := range fn.Params {
		record := GetRecordWithReg(param.Ident)
		if record != nil && record.vargs {
			params = append(params, "...")
			commentParams.WriteString(fmt.Sprintf(commentVargs, record.luaType))
		} else {
			params = append(params, nameMap(param.Name))
			commentParams.WriteString(
				fmt.Sprintf(commentParam, nameMap(param.Name), GetWithReg(param.Ident)))
		}
	}
	for _, result := range fn.Results {
		results = append(results, Get(result.Ident))
	}
	if len(results) > 0 {
		res = fmt.Sprintf(commentResult, strings.Join(results, ", "))
	}
	return fmt.Sprintf("%sfunction %s%s(%s)\nend",
		commentParams.String()+res, "%s", fn.Name, strings.Join(params, ", "))
}

type Struct struct {
	Name     string
	Methods  map[string]*Function
	Fields   []Type
	Comments []string
}

func (stu *Struct) luaTypeString() string {
	commentFields := bytes.Buffer{}
	for _, field := range stu.Fields {
		if canExport(field.Name) {
			commentFields.WriteString(
				fmt.Sprintf(commentField, field.Name, Get(field.Ident)))
		}
	}
	return fmt.Sprintf(defineStruct,
		Get(stu.Name), commentFields.String(), Get(stu.Name))
}

func (stu *Struct) luaMethodString() string {
	methods := bytes.Buffer{}
	for name, method := range stu.Methods {
		if canExport(name) {
			methods.WriteString("\n" + fmt.Sprintf(method.LuaString(), Get(stu.Name)+name, Get(stu.Name)+":") + "\n")
		}
	}
	return methods.String()
}

func (stu *Struct) LuaString() string {
	return ""
}

type File struct {
	file      string
	constants []string
	variable  []string
	functions map[string]*Function
	structs   map[string]*Struct
	types     map[string]Type
}

func canExport(str string) bool {
	return len(str) > 0 && unicode.IsUpper(rune(str[0]))
}
