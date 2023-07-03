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

func TestPrintTypeAst(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "type.go", nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, file)
}

type StructField struct {
	Ident string
	Type  string
}

var S map[string][]StructField

func TestParseTypeAst(t *testing.T) {
	S = make(map[string][]StructField)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "type.go", nil, 0)
	if err != nil {
		panic(err)
	}
	for _, decl := range file.Decls {
		switch v := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range v.Specs {
				switch vv := spec.(type) {
				case *ast.TypeSpec:
					stuName := vv.Name.Name
					switch vvv := vv.Type.(type) {
					case *ast.StructType:
						for _, stuField := range vvv.Fields.List {
							for _, name := range stuField.Names {
								field := StructField{Ident: name.Name}
								switch pv := stuField.Type.(type) {
								case *ast.Ident:
									field.Type = pv.Name
								case *ast.StarExpr:
									switch pvv := pv.X.(type) {
									case *ast.Ident:
										field.Type = pvv.Name
									default:
										field.Type = "any"
									}
								default:
									field.Type = "any"
								}
								if S[stuName] == nil {
									S[stuName] = make([]StructField, 0)
								}
								S[stuName] = append(S[stuName], field)
							}
						}
					case *ast.Ident:
						S[stuName] = append(S[stuName], StructField{Type: vvv.Name})
					}
				}
			}
			// ast.Print(token.NewFileSet(), v)
		}
	}
	fmt.Println(S)
}
