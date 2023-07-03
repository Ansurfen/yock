// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package archive

import (
	"bytes"
	"fmt"

	"github.com/ansurfen/yock/util"
)

var (
	libs    map[string]*Lib
	author  string
	license string
)

const (
	luaStringHead = `-- Copyright 2023 %s. All rights reserved.
-- Use of this source code is governed by a %s-style
-- license that can be found in the LICENSE file.

---@meta _

%s%s = {}
`
	goStringHead = `// Copyright 2023 %s. All rights reserved.
// Use of this source code is governed by a %s-style
// license that can be found in the LICENSE file.

package %slib

import yocki "github.com/ansurfen/yock/interface"
`
)

func init() {
	libs = make(map[string]*Lib)
}

type Lib struct {
	constants []string
	variable  []string
	functions map[string]*Function
	structs   map[string]*Struct
	types     map[string]Type
	name      string
}

func (lib *Lib) GoString() string {
	head := fmt.Sprintf(goStringHead, author, license, lib.name)
	buf := bytes.Buffer{}
	buf.WriteString(
		fmt.Sprintf("func Load%s(yocks yocki.YockScheduler) {\n\tlib := yocks.CreateLib(\"%s\")\n", util.Title(lib.name), lib.name))
	buf.WriteString("\tlib.SetField(map[string]any{\n")
	buf.WriteString("\t\t// functions\n")
	for name, _ := range lib.functions {
		buf.WriteString(fmt.Sprintf("\t\t\"%s\": %s.%s,\n", name, lib.name, name))
	}
	buf.WriteString("\t\t// constants\n")
	for _, name := range lib.constants {
		buf.WriteString(fmt.Sprintf("\t\t\"%s\": %s.%s,\n", name, lib.name, name))
	}
	buf.WriteString("\t\t// variable\n")
	for _, name := range lib.variable {
		buf.WriteString(fmt.Sprintf("\t\t\"%s\": %s.%s,\n", name, lib.name, name))
	}
	buf.WriteString("\t})\n")
	buf.WriteString("}")
	return head + buf.String()
}

func (lib *Lib) LuaString() string {
	libType := bytes.Buffer{}
	if len(lib.constants) != 0 || len(lib.variable) != 0 {
		libType.WriteString(fmt.Sprintf("---@class %s\n", lib.name))
	}
	for _, v := range lib.constants {
		libType.WriteString(fmt.Sprintf(commentField, v, "any"))
	}
	for _, v := range lib.variable {
		libType.WriteString(fmt.Sprintf(commentField, v, "any"))
	}
	head := fmt.Sprintf(luaStringHead, author, license, libType.String(), lib.name)
	buf := bytes.Buffer{}
	for name, fn := range lib.functions {
		buf.WriteString("\n" + fmt.Sprintf(fn.LuaString(), lib.name+name, lib.name+".") + "\n")
	}
	for _, stu := range lib.structs {
		buf.WriteString("\n" + stu.luaTypeString() + "\n")
		buf.WriteString(stu.luaMethodString())
	}
	return head + buf.String()
}

func getLib(name string) *Lib {
	if v, ok := libs[name]; ok {
		return v
	}
	lib := &Lib{
		functions: make(map[string]*Function),
		structs:   make(map[string]*Struct),
		types:     make(map[string]Type),
		name:      name,
	}
	libs[name] = lib
	return lib
}
