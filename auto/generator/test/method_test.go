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

func TestPrintMethodAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "method.go", nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, file)
}

type Struct struct {
	Methods map[string]*StructMethod
}

type StructMethod struct {
	Params  []MethodParameter
	Results []MethodParameter
}

type MethodParameter struct {
	Ident string
	Type  string
}

var stus map[string]*Struct

func TestParseMethodAst(t *testing.T) {
	stus = make(map[string]*Struct)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "method.go", nil, 0)
	if err != nil {
		panic(err)
	}
	for _, decl := range file.Decls {
		switch v := decl.(type) {
		case *ast.GenDecl:
		case *ast.FuncDecl:
			stuName := ""
			if v.Recv != nil {
				switch vv := v.Recv.List[0].Type.(type) {
				case *ast.StarExpr:
					stuName = vv.X.(*ast.Ident).Name
				case *ast.Ident:
					stuName = vv.Name
				}
				if len(stuName) > 0 {
					if _, ok := stus[stuName]; !ok {
						stus[stuName] = &Struct{}
					}
				} else {
					continue
				}
			}
			fnName := v.Name.Name
			if stus[stuName].Methods == nil {
				stus[stuName].Methods = make(map[string]*StructMethod)
			}
			stus[stuName].Methods[fnName] = &StructMethod{}
			for _, p := range v.Type.Params.List {
				t := ""
				switch pv := p.Type.(type) {
				case *ast.Ident:
					t = pv.Name
				case *ast.StarExpr:
					t = pv.X.(*ast.Ident).Name
				case *ast.Ellipsis:
				case *ast.SelectorExpr:
					t = pv.X.(*ast.Ident).Name + "." + pv.Sel.Name
				case *ast.IndexExpr:
					t = pv.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + pv.X.(*ast.SelectorExpr).Sel.Name
				}
				for _, po := range p.Names {
					param := MethodParameter{}
					param.Ident = po.Name
					param.Type = t
					stus[stuName].Methods[fnName].Params = append(stus[stuName].Methods[fnName].Params, param)
				}
			}
		}
	}
	fmt.Println(stus["Stu4"].Methods["Fn6"])
}
