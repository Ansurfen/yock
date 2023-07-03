// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package archive

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/ansurfen/yock/util"
)

func parse(path string) *File {
	fp := &File{
		structs:   make(map[string]*Struct),
		functions: make(map[string]*Function),
		types:     make(map[string]Type),
		file:      util.Filename(path),
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, decl := range file.Decls {
		switch v := decl.(type) {
		case *ast.GenDecl:
			specs := []ast.Spec{}
			isConst := true
			switch v.Tok.String() {
			case "const":
				specs = v.Specs
			case "var":
				specs = v.Specs
				isConst = false
			case "type":
				comments := []string{}
				if v.Doc != nil {
					for _, comment := range v.Doc.List {
						comments = append(comments, comment.Text)
					}
				}
				for _, spec := range v.Specs {
					switch vv := spec.(type) {
					case *ast.TypeSpec:
						typeName := vv.Name.Name
						if fp.structs[typeName] == nil {
							fp.structs[typeName] = &Struct{
								Methods:  make(map[string]*Function),
								Name:     typeName,
								Comments: comments,
							}
						}
						switch vvv := vv.Type.(type) {
						case *ast.StructType:
							for _, stuField := range vvv.Fields.List {
								for _, name := range stuField.Names {
									field := Type{Name: name.Name}
									switch pv := stuField.Type.(type) {
									case *ast.Ident:
										field.Ident = pv.Name
									case *ast.StarExpr:
										switch pvv := pv.X.(type) {
										case *ast.Ident:
											field.Ident = pvv.Name
										default:
											field.Ident = "any"
										}
									default:
										field.Ident = "any"
									}
									fp.structs[typeName].Fields = append(fp.structs[typeName].Fields, field)
								}
							}
						case *ast.Ident:
							fp.types[typeName] = Type{Name: typeName, Ident: vvv.Name, Comments: comments}
						}
					}
				}
			default:
			}
			for _, spec := range specs {
				if isConst {
					fp.constants = append(fp.constants, spec.(*ast.ValueSpec).Names[0].Name)
				} else {
					fp.variable = append(fp.variable, spec.(*ast.ValueSpec).Names[0].Name)
				}
			}
		case *ast.FuncDecl:
			fnName := v.Name.Name
			stuName := ""
			comments := []string{}
			if v.Doc != nil {
				for _, comment := range v.Doc.List {
					comments = append(comments, comment.Text)
				}
			}
			if v.Recv != nil {
				switch vv := v.Recv.List[0].Type.(type) {
				case *ast.StarExpr:
					switch pvv := vv.X.(type) {
					case *ast.Ident:
						stuName = pvv.Name
					case *ast.SelectorExpr:
						stuName = pvv.X.(*ast.Ident).Name + "." + pvv.Sel.Name
					}
				case *ast.Ident:
					stuName = vv.Name
				}
				if len(stuName) > 0 {
					if _, ok := fp.structs[stuName]; !ok {
						fp.structs[stuName] = &Struct{
							Methods: make(map[string]*Function),
						}
					}
					fp.structs[stuName].Methods[fnName] = &Function{Name: fnName, Comments: comments}
					for _, param := range v.Type.Params.List {
						t := ""
						switch pv := param.Type.(type) {
						case *ast.Ident:
							t = pv.Name
						case *ast.StarExpr:
							switch pvv := pv.X.(type) {
							case *ast.Ident:
								t = pvv.Name
							case *ast.SelectorExpr:
								t = pvv.X.(*ast.Ident).Name + "." + pvv.Sel.Name
							}
						case *ast.Ellipsis:
							switch ev := pv.Elt.(type) {
							case *ast.Ident:
								t = "..." + ev.Name
							default: // any
								t = "...any"
							}
						case *ast.SelectorExpr:
							t = pv.X.(*ast.Ident).Name + "." + pv.Sel.Name
						case *ast.IndexExpr:
							t = pv.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + pv.X.(*ast.SelectorExpr).Sel.Name
						case *ast.ArrayType:
							switch ev := pv.Elt.(type) {
							case *ast.Ident:
								t = "[]" + ev.Name
							default: // any
								t = "[]any"
							}
						}
						for _, pn := range param.Names {
							fp.structs[stuName].Methods[fnName].Params =
								append(fp.structs[stuName].Methods[fnName].Params, Type{Name: pn.Name, Ident: t})
						}
					}
					if v.Type.Results != nil {
						for _, param := range v.Type.Results.List {
							t := ""
							switch pv := param.Type.(type) {
							case *ast.Ident:
								t = pv.Name
							case *ast.StarExpr:
								switch pvv := pv.X.(type) {
								case *ast.Ident:
									t = pvv.Name
								case *ast.SelectorExpr:
									t = pvv.X.(*ast.Ident).Name + "." + pvv.Sel.Name
								}
							case *ast.Ellipsis:
								switch ev := pv.Elt.(type) {
								case *ast.Ident:
									t = "..." + ev.Name
								default: // any
									t = "...any"
								}
							case *ast.SelectorExpr:
								t = pv.X.(*ast.Ident).Name + "." + pv.Sel.Name
							case *ast.IndexExpr:
								t = pv.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + pv.X.(*ast.SelectorExpr).Sel.Name
							case *ast.ArrayType:
								switch ev := pv.Elt.(type) {
								case *ast.Ident:
									t = "[]" + ev.Name
								default: // any
									t = "[]any"
								}
							}
							fp.structs[stuName].Methods[fnName].Results =
								append(fp.structs[stuName].Methods[fnName].Results, Type{Ident: t})
						}
					}
				}
			} else {
				fp.functions[fnName] = &Function{Name: fnName, Comments: comments}
				for _, param := range v.Type.Params.List {
					t := ""
					switch pv := param.Type.(type) {
					case *ast.Ident:
						t = pv.Name
					case *ast.StarExpr:
						switch pvv := pv.X.(type) {
						case *ast.Ident:
							t = pvv.Name
						case *ast.SelectorExpr:
							t = pvv.X.(*ast.Ident).Name + "." + pvv.Sel.Name
						}
					case *ast.Ellipsis:
						switch ev := pv.Elt.(type) {
						case *ast.Ident:
							t = "..." + ev.Name
						default: // any
							t = "...any"
						}
					case *ast.SelectorExpr:
						t = pv.X.(*ast.Ident).Name + "." + pv.Sel.Name
					case *ast.IndexExpr:
						t = pv.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + pv.X.(*ast.SelectorExpr).Sel.Name
					case *ast.ArrayType:
						switch ev := pv.Elt.(type) {
						case *ast.Ident:
							t = "[]" + ev.Name
						default: // any
							t = "[]any"
						}
					case *ast.FuncType:
						t = "func"
					}
					for _, pn := range param.Names {
						fp.functions[fnName].Params =
							append(fp.functions[fnName].Params, Type{Name: pn.Name, Ident: t})
					}
				}
				if v.Type.Results != nil {
					for _, param := range v.Type.Results.List {
						t := ""
						switch pv := param.Type.(type) {
						case *ast.Ident:
							t = pv.Name
						case *ast.StarExpr:
							switch pvv := pv.X.(type) {
							case *ast.Ident:
								t = pvv.Name
							case *ast.SelectorExpr:
								t = pvv.X.(*ast.Ident).Name + "." + pvv.Sel.Name
							}
						case *ast.Ellipsis:
							switch ev := pv.Elt.(type) {
							case *ast.Ident:
								t = "..." + ev.Name
							default: // any
								t = "...any"
							}
						case *ast.SelectorExpr:
							t = pv.X.(*ast.Ident).Name + "." + pv.Sel.Name
						case *ast.IndexExpr:
							t = pv.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + pv.X.(*ast.SelectorExpr).Sel.Name
						case *ast.ArrayType:
							switch ev := pv.Elt.(type) {
							case *ast.Ident:
								t = "[]" + ev.Name
							default: // any
								t = "[]any"
							}
						}
						fp.functions[fnName].Results =
							append(fp.functions[fnName].Results, Type{Ident: t})
					}
				}
			}
		}
	}
	return fp
}
