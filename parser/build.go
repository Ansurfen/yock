package parser

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/cushion/utils/build"
	"github.com/yuin/gopher-lua/ast"
)

type exprFrame struct {
	fn bool
}

func exprBuilder(buf *bytes.Buffer, exp ast.Expr, frame exprFrame) {
	switch v := exp.(type) {
	case *ast.FunctionExpr:
		buf.WriteString("function (")
		for idx, arg := range v.ParList.Names {
			buf.WriteString(arg)
			if idx != len(v.ParList.Names)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString(")\n")
		stmtBuilder(buf, v.Stmts, nil)
		buf.WriteString("end\n")
	case *ast.FuncCallExpr:
		exprBuilder(buf, v.Func, exprFrame{})
		buf.WriteString("(")
		for idx, arg := range v.Args {
			exprBuilder(buf, arg, exprFrame{fn: true})
			if idx != len(v.Args)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString(")\n")
	case *ast.StringExpr:
		if frame.fn {
			buf.WriteString(fmt.Sprintf(`"%s"`, v.Value))
		} else {
			buf.WriteString(v.Value)
		}
	case *ast.NumberExpr:
		buf.WriteString(v.Value)
	case *ast.IdentExpr:
		buf.WriteString(v.Value)
	case *ast.AttrGetExpr:
		exprBuilder(buf, v.Object, exprFrame{})
		buf.Write([]byte{'.'})
		exprBuilder(buf, v.Key, exprFrame{})
	case *ast.TrueExpr:
		buf.WriteString("true")
	case *ast.FalseExpr:
		buf.WriteString("false")
	case *ast.TableExpr:
		buf.WriteString("{")
		for idx, field := range v.Fields {
			if field.Key != nil {
				exprBuilder(buf, field.Key, exprFrame{})
				buf.WriteString("=")
			}
			exprBuilder(buf, field.Value, exprFrame{fn: true})
			if idx != len(v.Fields)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("}")
	case *ast.StringConcatOpExpr:
		exprBuilder(buf, v.Lhs, exprFrame{fn: true})
		buf.WriteString(" .. ")
		exprBuilder(buf, v.Rhs, exprFrame{fn: true})
	case *ast.LogicalOpExpr:
		exprBuilder(buf, v.Lhs, exprFrame{})
		buf.WriteString(" " + v.Operator + " ")
		exprBuilder(buf, v.Rhs, exprFrame{})
	case *ast.RelationalOpExpr:
		exprBuilder(buf, v.Lhs, exprFrame{})
		buf.WriteString(" " + v.Operator + " ")
		exprBuilder(buf, v.Rhs, exprFrame{})
	default:
		fmt.Println(reflect.TypeOf(v))
	}
}

func stmtBuilder(buf *bytes.Buffer, chunk []ast.Stmt, filter map[int]bool) {
	for idx, stmt := range chunk {
		if filter != nil && filter[idx] {
			continue
		}
		switch v := stmt.(type) {
		case *ast.AssignStmt:
			exprBuilder(buf, v.Lhs[0], exprFrame{})
			buf.WriteString("=")
			exprBuilder(buf, v.Rhs[0], exprFrame{})
			buf.WriteString("\n")
		case *ast.FuncCallStmt:
			exprBuilder(buf, v.Expr, exprFrame{})
		case *ast.LocalAssignStmt:
			buf.WriteString(fmt.Sprintf("local %s = \n", v.Names[0]))
			for _, expr := range v.Exprs {
				exprBuilder(buf, expr, exprFrame{})
			}
		case *ast.ReturnStmt:
			buf.WriteString("return ")
			for idx, expr := range v.Exprs {
				exprBuilder(buf, expr, exprFrame{})
				if idx != len(v.Exprs)-1 {
					buf.WriteString(", ")
				}
			}
			buf.WriteString("\n")
		case *ast.NumberForStmt:
			buf.WriteString(fmt.Sprintf(`for %s = `, v.Name))
			exprBuilder(buf, v.Init, exprFrame{})
			buf.WriteString(", ")
			exprBuilder(buf, v.Limit, exprFrame{})
			if v.Step != nil {
				buf.WriteString(", ")
				exprBuilder(buf, v.Step, exprFrame{})
			}
			buf.WriteString(" do\n")
			stmtBuilder(buf, v.Stmts, nil)
			buf.WriteString("end\n")
		case *ast.IfStmt:
			buf.WriteString("if ")
			exprBuilder(buf, v.Condition, exprFrame{})
			buf.WriteString(" then\n")
			stmtBuilder(buf, v.Then, nil)
			if len(v.Else) != 0 {
				buf.WriteString("else ")
				stmtBuilder(buf, v.Else, nil)
			}
			buf.WriteString("end\n")
		default:
			fmt.Println(v)
		}
	}
}

func BuildLuaScript(chunk []ast.Stmt, filter map[int]bool) string {
	var buf *bytes.Buffer = new(bytes.Buffer)
	stmtBuilder(buf, chunk, filter)
	return buf.String()
}

func buildBootScript(file string, tpl string, modes []string) {
	tmpl := build.NewTemplate()
	type mode struct {
		Name string
	}
	text, err := utils.ReadStraemFromFile(tpl)
	if err != nil {
		panic(err)
	}
	tmpl.Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})
	var ms []mode
	for _, m := range modes {
		ms = append(ms, mode{Name: m})
	}
	out, err := tmpl.OnceParse(string(text), ms)
	if err != nil {
		panic(err)
	}
	if !strings.HasSuffix(file, ".lua") {
		file = file + ".lua"
	}
	utils.WriteFile(file, []byte(out))
}