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

func TestPrintVarAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "var.go", nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, file)
}

func TestParseVarAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "var.go", nil, 0)
	if err != nil {
		panic(err)
	}
	for _, decl := range file.Decls {
		switch v := decl.(type) {
		case *ast.GenDecl:
			fmt.Println(v.Tok)
			for _, spec := range v.Specs {
				fmt.Println(spec.(*ast.ValueSpec).Names[0].Name)
			}
		}
	}
}
