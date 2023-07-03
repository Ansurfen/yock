// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package archive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ansurfen/yock/util"
)

func init() {
	archive = &TypeArchive{
		records: map[string]*TypeRecord{
			"[]byte":     {luaType: "integer[]"},
			"bool":       {luaType: "boolean"},
			"error":      {luaType: "err"},
			"byte":       {luaType: "integer"},
			"int64":      {luaType: "number"},
			"int":        {luaType: "number"},
			"float32":    {luaType: "number"},
			"float64":    {luaType: "number"},
			"string":     {luaType: "string"},
			"[]string":   {luaType: "string[]"},
			"uint64":     {luaType: "number"},
			"complex128": {luaType: "any"},
			"func":       {luaType: "function"},
			"any":        {luaType: "any"},
		},
	}
}

type TypeRecord struct {
	luaType string
	vargs   bool
}

type TypeArchive struct {
	records map[string]*TypeRecord
}

var archive *TypeArchive

func Put(lib, name, ident string) {
	if _, ok := archive.records[name]; !ok {
		_type := ""
		if strings.Contains(ident, ".") {
			_type = strings.ReplaceAll(ident, ".", "")
		} else if strings.HasPrefix(ident, "...") {
			_type := Get(strings.ReplaceAll(ident, ".", ""))
			archive.records[ident] = &TypeRecord{luaType: _type, vargs: true}
		} else {
			_type = lib + ident
		}
		archive.records[name] = &TypeRecord{luaType: _type}
	}
}

func Get(ident string) string {
	if v, ok := archive.records[ident]; ok {
		return v.luaType
	}
	return "any"
}

func GetRecord(ident string) *TypeRecord {
	if v, ok := archive.records[ident]; ok {
		return v
	}
	return nil
}

func GetRecordWithReg(ident string) *TypeRecord {
	if v, ok := archive.records[ident]; ok {
		return v
	}
	if strings.Contains(ident, "...") {
		_type := Get(strings.ReplaceAll(ident, ".", ""))
		archive.records[ident] = &TypeRecord{luaType: _type, vargs: true}
	} else if strings.Contains(ident, ".") {
		_type := strings.ReplaceAll(ident, ".", "")
		archive.records[ident] = &TypeRecord{luaType: _type}
	}
	return archive.records[ident]
}

func GetWithReg(ident string) string {
	if v, ok := archive.records[ident]; ok {
		return v.luaType
	}
	if strings.Contains(ident, "...") {
		_type := Get(strings.ReplaceAll(ident, ".", ""))
		archive.records[ident] = &TypeRecord{luaType: _type, vargs: true}
		return _type
	} else if strings.Contains(ident, ".") {
		_type := strings.ReplaceAll(ident, ".", "")
		archive.records[ident] = &TypeRecord{luaType: _type}
		return _type
	}
	return "any"
}

func EnableYockComment() {
	archive.records["byte"] = &TypeRecord{luaType: "byte"}
	archive.records["[]byte"] = &TypeRecord{luaType: "byte[]"}
}

func LoadFile(libn, path string) {
	lib := getLib(libn)
	file := parse(path)
	for name, stu := range file.structs {
		if canExport(name) {
			Put(libn, name, name)
			lib.structs[name] = stu
		}
	}
	for name := range file.types {
		if canExport(name) {
			Put(libn, name, name)
		}
	}
	for name, cb := range file.functions {
		if canExport(name) {
			lib.functions[name] = cb
		}
	}
	for _, cns := range file.constants {
		if canExport(cns) {
			lib.constants = append(lib.constants, cns)
		}
	}
	for _, cns := range file.variable {
		if canExport(cns) {
			lib.variable = append(lib.variable, cns)
		}
	}
}

func LoadDir(libn, dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			if !strings.Contains(filename, "_test.go") && filepath.Ext(filename) == ".go" {
				LoadFile(libn, filepath.Join(dir, filename))
			}
		}
	}
}

func PrintLua(name string) {
	if lib, ok := libs[name]; ok {
		fmt.Println(lib.LuaString())
	}
}

func PrintGo(name string) {
	if lib, ok := libs[name]; ok {
		fmt.Println(lib.GoString())
	}
}

func Export(path string) {
	util.Mkdirs(path)
	for name, lib := range libs {
		mod := filepath.Join(path, name)
		util.Mkdirs(mod)
		util.WriteFile(filepath.Join(mod, name+".lua"), []byte(lib.LuaString()))
		util.WriteFile(filepath.Join(mod, name+".go"), []byte(lib.GoString()))
		doc := make(map[string]string)
		for name, fn := range lib.functions {
			doc[lib.name+name] = commentf(fn.Comments)
		}
		for stuName, stu := range lib.structs {
			for name, fn := range stu.Methods {
				doc[Get(stuName)+name] = commentf(fn.Comments)
			}
		}
		raw, err := json.Marshal(doc)
		if err != nil {
			panic(err)
		}
		var out bytes.Buffer
		err = json.Indent(&out, raw, "", "\t")
		if err != nil {
			panic(err)
		}
		util.WriteFile(filepath.Join(mod, name+".json"), out.Bytes())
	}
}

func SetInfo(a, l string) {
	author = a
	license = l
}

func commentf(comments []string) string {
	buf := ""
	for idx, comment := range comments {
		text := strings.Replace(comment, "//", "---", 1)
		if idx != len(comments)-1 {
			text += "\n"
		}
		buf += text
	}
	return buf
}
