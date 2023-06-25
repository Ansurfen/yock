// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ansurfen/yock/util"
	"github.com/yuin/gopher-lua/parse"
)

func TestParseAST(t *testing.T) {
	yockp := New()
	ast := yockp.ParseStr(`
	A():B().C:D().E.F(a,b)
	A:B().C:D().E.F()
	A().B().C:D().E.F()
	A.B.C.D()
	A.B.C:D()
	function A.B.C.D() end
	function A.B.C:D(a,b) end
	`)
	fmt.Println(parse.Dump(ast))
	fmt.Println(parseFuncStmt("g", ast))
}

func TestDependencyAnalyer(t *testing.T) {
	anlyzer := NewLuaDependencyAnalyzer()
	anlyzer.Load("../sdk/yock.lua")
	anlyzer.Preload("print()", LuaMethod{Pkg: "g"})
	anlyzer.Export(util.RandString(8) + ".json")
	fmt.Println(anlyzer.Completion("../ctl/test/check_env_test.lua"))
}

func TestExportStdlib(t *testing.T) {
	anlyzer := NewLuaDependencyAnalyzer()
	root := ``
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			anlyzer.LoadG(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	anlyzer.Export("stdlib.json")
}

func TestReload(t *testing.T) {
	out, err := util.ReadStraemFromFile("stdlib.json")
	if err != nil {
		panic(err)
	}
	anlyzer := NewLuaDependencyAnalyzer()
	if err = json.Unmarshal(out, anlyzer); err != nil {
		panic(err)
	}
	fmt.Println(anlyzer)
}
