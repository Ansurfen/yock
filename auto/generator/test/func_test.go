// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestPrintFuncAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "func.go", nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, file)
}

func TestParseFuncAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "func.go", nil, 0)
	if err != nil {
		panic(err)
	}
	for _, decl := range file.Decls {
		switch v := decl.(type) {
		case *ast.FuncDecl:
			for _, param := range v.Type.Params.List {
				switch tv := param.Type.(type) {
				case *ast.Ellipsis:
					switch ev := tv.Elt.(type) {
					case *ast.Ident:
						fmt.Println(ev)
					default: // any
					}
				}
			}
		}
	}
}
