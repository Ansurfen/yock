package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/ansurfen/yock/auto/generator/archive"
	"github.com/ansurfen/yock/auto/generator/walk"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	methods := make([]*archive.Method, 0)
	gowalk := walk.GoWalk{}
	gowalk.VisitDir("D:/D/langs/go/src/strings", walk.VisitDirHandle{
		walk.DeclFunc: func(pkg string, decl walk.GoDecl) bool {
			fn := decl.(walk.FuncDecl)
			funcName := fn.Name.Name
			if unicode.IsLower(rune(funcName[0])) ||
				strings.HasPrefix(funcName, "Bench") ||
				strings.HasPrefix(funcName, "Test") ||
				strings.HasPrefix(funcName, "Example") {
				return false
			}
			if fn.Recv == nil {
				method := &archive.Method{
					Package: pkg,
					Name:    funcName,
				}
				methods = append(methods, method)
				if fn.Doc != nil {
					for _, comment := range fn.Doc.List {
						method.Comment = append(method.Comment,
							strings.Replace(comment.Text, "//", "--", 1))
					}
					if fn.Type.Params != nil {
						for idx, field := range fn.Type.Params.List {
							for _, name := range field.Names {
								var argument archive.MethodArgument
								gowalk.VisitExpr(idx, field.Type, walk.VisitExprHandle{
									walk.SymbolIdent: func(idx int, expr walk.GoExpr) {
										argument = archive.MethodArgument{
											Name:      name.String(),
											TypeIdent: expr.(walk.IdentSymbol).Name,
										}
									},
									walk.SymbolEllipsis: func(idx int, expr walk.GoExpr) {
										// argument = archive.MethodArgument{
										// 	Name:      name.String(),
										// 	TypeIdent: expr.(walk.Ellipsis).Elt.(walk.IdentSymbol).String(),
										// }
									},
									walk.HandleDefault: func(idx int, expr walk.GoExpr) {
										fmt.Println(reflect.TypeOf(expr))
									},
								})
								method.Params = append(method.Params, argument)
							}
						}
					}
					if fn.Type.Results != nil {
						for idx, field := range fn.Type.Results.List {
							var argument archive.MethodArgument
							gowalk.VisitExpr(idx, field.Type, walk.VisitExprHandle{
								walk.TypeArray: func(idx int, expr walk.GoExpr) {
									argument = archive.MethodArgument{
										TypeIdent: "any",
									}
								},
								walk.TypeMap: func(idx int, expr walk.GoExpr) {
									argument = archive.MethodArgument{
										TypeIdent: "any",
									}
								},
								walk.ExprStar: func(idx int, expr walk.GoExpr) {
									argument = archive.MethodArgument{
										TypeIdent: "any",
									}
								},
								walk.SymbolIdent: func(idx int, expr walk.GoExpr) {
									argument = archive.MethodArgument{
										TypeIdent: expr.(walk.IdentSymbol).String(),
									}
								},
							})
							method.Results = append(method.Results, argument)
						}
					}
				}
			}
			return true
		},
		walk.PackageHandle: func(pkg string, decl walk.GoDecl) bool {
			return !strings.HasSuffix(pkg, "_test")
		},
	})
	archives := archive.GetArchive()
	for _, method := range methods {
		fmt.Println(method)
		varName := 'a'
		buf := bytes.Buffer{}
		for i := range method.Results {
			method.Results[i].Name = string(varName)
			buf.WriteString(string(varName))
			if i != len(method.Results)-1 {
				buf.WriteString(", ")
			}
			varName += 1
		}
		if varName != 'a' {
			buf.WriteString(" := ")
		}
		buf.WriteString(method.Package + "." + method.Name + "(")
		for i, arg := range method.Params {
			buf.WriteString(archives.Lookup(arg.TypeIdent).Check(i + 1))
			if i != len(method.Params)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString(")\n")
		for _, arg := range method.Results {
			record := archives.Lookup(arg.TypeIdent)
			if record.CheckType() == lua.LTBool {
				buf.WriteString(record.Type(arg.Name) + "\n")
			} else {
				buf.WriteString(fmt.Sprintf("l.Push(%s)\n", record.Type(arg.Name)))
			}
		}
		buf.WriteString(fmt.Sprintf("return %d", len(method.Results)))
		// fmt.Println(buf.String())
		// fmt.Println()
	}
}