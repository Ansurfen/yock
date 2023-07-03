// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Dependency analysis was Yock's early idea of package management.
// Scripters do not need to explicitly import dependencies,
// but yockpack completes dependency analysis by traversing the syntax tree,
// and automatically imports them at runtime. However, this design has a natural flaw,
// for the same name, the same parameter function, the compiler cannot distinguish the
// function that the user wants to use from the function registry.
// If you want to learn about the implementation to continue, you can read on.
// In this pattern, each global function will be treated as a driver.
// One driver specifically solves one type of problem, and every could be overrided according to optional table
// For added flexibility, each driver can mount multiple plugins,
// which you can think of as a callback function during driver execution.

package yockp

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/yuin/gopher-lua/ast"
)

// luaDependencyAnalyzer is used to analyze dependencies.
// It has now been deprecated in the latest version.
type luaDependencyAnalyzer struct {
	// Includes is responsible for recording the metadata of driver to be scanned
	Includes map[string][]LuaMethod `json:"includes"`
	yockp    YockPack[NilFrame]
}

func NewLuaDependencyAnalyzer() *luaDependencyAnalyzer {
	return &luaDependencyAnalyzer{
		Includes: make(map[string][]LuaMethod),
		yockp:    New(),
	}
}

// Completion returns undefined drivers based on the current environment comparison
func (analyzer *luaDependencyAnalyzer) Completion(file string) ([]string, map[string]bool) {
	ast := analyzer.yockp.ParseFile(file)
	dec, cal := parseFuncStmt(file, ast)
	deps := make(map[string]bool)
	libs := []string{}
	for name := range cal {
		if _, ok := dec[name]; ok {
			continue
		}
		if m, ok := analyzer.Includes[name]; !ok {
			libs = append(libs, name)
		} else {
			deps[m[0].Pkg] = true
		}
	}
	return libs, deps
}

// Lookup returns files of declare of methods according to specify name
func (analyzer *luaDependencyAnalyzer) Lookup(name string) []LuaMethod {
	return analyzer.Includes[name]
}

// Load to anaylse declare of method in specify lua script
func (analyzer *luaDependencyAnalyzer) Load(str string) {
	var ast []ast.Stmt
	scope := "g"
	if filepath.Ext(str) == ".lua" {
		ast = analyzer.yockp.ParseFile(str)
		scope = str
	} else {
		ast = analyzer.yockp.ParseStr(str)
	}
	dec, _ := parseFuncStmt(scope, ast)
	for name, method := range dec {
		analyzer.Includes[name] = append(analyzer.Includes[name], method)
	}
}

func (analyzer *luaDependencyAnalyzer) LoadG(str string) {
	var ast []ast.Stmt
	scope := "g"
	if filepath.Ext(str) == ".lua" {
		ast = analyzer.yockp.ParseFile(str)
	} else {
		ast = analyzer.yockp.ParseStr(str)
	}
	dec, _ := parseFuncStmt(scope, ast)
	for name, method := range dec {
		analyzer.Includes[name] = append(analyzer.Includes[name], method)
	}
}

func (analyzer *luaDependencyAnalyzer) Preload(name string, method LuaMethod) {
	analyzer.Includes[name] = append(analyzer.Includes[name], method)
}

// Export writes the analyzed metadata to a file as a cache,
// thereby skipping the process of repeated analysis.
func (analyzer *luaDependencyAnalyzer) Export(file string) {
	out, err := json.Marshal(analyzer)
	if err != nil {
		ycho.Fatal(err)
	}
	util.WriteFile(file, out)
}

// LuaMethod stores the metadata of the driver
type LuaMethod struct {
	// Argc indicates the number of parameters
	Argc int `json:"argc"`
	// Argv collects the name of each parameter
	Argv []string `json:"argv"`
	// Pkg indicates the scope for current driver.
	// Generally speaking, it is the file name,
	// and the way pkg is introduced for standard library functions and strings is g(global).
	Pkg string `json:"pkg"`
}

func parseFuncStmt(scope string, stmts []ast.Stmt) (declares, calls map[string]LuaMethod) {
	declares = make(map[string]LuaMethod)
	calls = make(map[string]LuaMethod)
	for _, stmt := range stmts {
		switch v := stmt.(type) {
		case *ast.FuncCallStmt:
			name := parseFuncExpr(v.Expr)
			calls[name] = LuaMethod{
				Pkg: scope,
			}
		case *ast.FuncDefStmt:
			name := ""
			if v.Name.Func != nil {
				name = fmt.Sprintf("%s()", parseFuncExpr(v.Name.Func))
			} else if v.Name.Receiver != nil {
				name = fmt.Sprintf("%s:%s()", parseFuncExpr(v.Name.Receiver), v.Name.Method)
			} else {
				continue
			}
			declares[name] = LuaMethod{
				Argv: v.Func.ParList.Names,
				Argc: len(v.Func.ParList.Names),
				Pkg:  scope,
			}
		}
	}
	return
}

func parseFuncExpr(expr ast.Expr) (ret string) {
	switch v := expr.(type) {
	case *ast.FuncCallExpr:
		if v.Func != nil {
			ret += fmt.Sprintf("%s()", parseFuncExpr(v.Func))
		} else if v.Receiver != nil {
			ret += fmt.Sprintf("%s:%s()", parseFuncExpr(v.Receiver), v.Method)
		}
	case *ast.AttrGetExpr:
		ret += fmt.Sprintf("%s.%s", parseFuncExpr(v.Object), parseFuncExpr(v.Key))
	case *ast.StringExpr:
		ret += v.Value
	case *ast.IdentExpr:
		ret += v.Value
	}
	return
}
