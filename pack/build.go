package yockpack

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

func (*YockPack[T]) BuildScript(stmts []ast.Stmt, filter map[int]bool) string {
	var buf *bytes.Buffer = new(bytes.Buffer)
	buildStmt(buf, stmts, filter)
	return buf.String()
}

type exprFrame struct {
	fn bool
}

func buildStmt(buf *bytes.Buffer, chunk []ast.Stmt, filter map[int]bool) {
	for idx, stmt := range chunk {
		if filter != nil && filter[idx] {
			continue
		}
		switch v := stmt.(type) {
		case *ast.AssignStmt:
			buildExpr(buf, v.Lhs[0], exprFrame{})
			buf.WriteString("=")
			buildExpr(buf, v.Rhs[0], exprFrame{})
			buf.WriteString("\n")
		case *ast.FuncCallStmt:
			buildExpr(buf, v.Expr, exprFrame{})
		case *ast.LocalAssignStmt:
			buf.WriteString(fmt.Sprintf("local %s = \n", v.Names[0]))
			for _, expr := range v.Exprs {
				buildExpr(buf, expr, exprFrame{})
			}
		case *ast.ReturnStmt:
			buf.WriteString("return ")
			for idx, expr := range v.Exprs {
				buildExpr(buf, expr, exprFrame{})
				if idx != len(v.Exprs)-1 {
					buf.WriteString(", ")
				}
			}
			buf.WriteString("\n")
		case *ast.NumberForStmt:
			buf.WriteString(fmt.Sprintf(`for %s = `, v.Name))
			buildExpr(buf, v.Init, exprFrame{})
			buf.WriteString(", ")
			buildExpr(buf, v.Limit, exprFrame{})
			if v.Step != nil {
				buf.WriteString(", ")
				buildExpr(buf, v.Step, exprFrame{})
			}
			buf.WriteString(" do\n")
			buildStmt(buf, v.Stmts, nil)
			buf.WriteString("end\n")
		case *ast.IfStmt:
			buf.WriteString("if ")
			buildExpr(buf, v.Condition, exprFrame{})
			buf.WriteString(" then\n")
			buildStmt(buf, v.Then, nil)
			if len(v.Else) != 0 {
				buf.WriteString("else ")
				buildStmt(buf, v.Else, nil)
			}
			buf.WriteString("end\n")
		default:
			fmt.Println(v)
		}
	}
}

func buildExpr(buf *bytes.Buffer, exp ast.Expr, frame exprFrame) {
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
		buildStmt(buf, v.Stmts, nil)
		buf.WriteString("end\n")
	case *ast.FuncCallExpr:
		buildExpr(buf, v.Func, exprFrame{})
		buf.WriteString("(")
		for idx, arg := range v.Args {
			buildExpr(buf, arg, exprFrame{fn: true})
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
		buildExpr(buf, v.Object, exprFrame{})
		buf.Write([]byte{'.'})
		buildExpr(buf, v.Key, exprFrame{})
	case *ast.TrueExpr:
		buf.WriteString("true")
	case *ast.FalseExpr:
		buf.WriteString("false")
	case *ast.TableExpr:
		buf.WriteString("{")
		for idx, field := range v.Fields {
			if field.Key != nil {
				buildExpr(buf, field.Key, exprFrame{})
				buf.WriteString("=")
			}
			buildExpr(buf, field.Value, exprFrame{fn: true})
			if idx != len(v.Fields)-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("}")
	case *ast.StringConcatOpExpr:
		buildExpr(buf, v.Lhs, exprFrame{fn: true})
		buf.WriteString(" .. ")
		buildExpr(buf, v.Rhs, exprFrame{fn: true})
	case *ast.LogicalOpExpr:
		buildExpr(buf, v.Lhs, exprFrame{})
		buf.WriteString(" " + v.Operator + " ")
		buildExpr(buf, v.Rhs, exprFrame{})
	case *ast.RelationalOpExpr:
		buildExpr(buf, v.Lhs, exprFrame{})
		buf.WriteString(" " + v.Operator + " ")
		buildExpr(buf, v.Rhs, exprFrame{})
	default:
		fmt.Println(reflect.TypeOf(v))
	}
}

type buildBootOpt struct {
	file  string
	tpl   string
	modes []string
}

func (yockpack *YockPack[T]) buildBoot(opt buildBootOpt) {
	tmpl := build.NewTemplate()
	type mode struct {
		Name string
	}
	text, err := utils.ReadStraemFromFile(opt.tpl)
	if err != nil {
		panic(err)
	}
	tmpl.Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})
	var ms []mode
	for _, m := range opt.modes {
		ms = append(ms, mode{Name: m})
	}
	out, err := tmpl.OnceParse(string(text), ms)
	if err != nil {
		panic(err)
	}
	if !strings.HasSuffix(opt.file, ".lua") {
		opt.file += ".lua"
	}
	utils.WriteFile(opt.file, []byte(out))
}
