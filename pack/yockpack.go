package yockpack

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/ast"
	"github.com/yuin/gopher-lua/parse"
)

type YockPack[T any] struct{}

type VisitOption[T any] func(*YockPack[T])

func New() YockPack[NilFrame] {
	return YockPack[NilFrame]{}
}

const (
	StmtAssign = iota
	StmtLocalAssign
	StmtFuncCall
	StmtDoBlock
	StmtWhile
	StmtRepeat
	StmtIf
	StmtNumbderFor
	StmtGenericFor
	StmtFuncDef
	StmtReturn
	StmtBreak
	StmtLabel
	StmtGoto

	ExprTrue
	ExprFalse
	ExprNil
	ExprNumber
	ExprString
	ExprComma3
	ExprIdent
	ExprAttrGet
	ExprTable
	ExprFuncCall
	ExprLogicalOp
	ExprRelationalOp
	ExprStringConcatOp
	ExprArithmeticOp
	ExprUnaryMinus
	ExprUnaryNotOp
	ExprUnaryLenOp
	ExprFunciton

	HandleDefault
)

type (
	NilFrame struct{}
	yockStmt ast.Stmt
	yockExpr = ast.Expr

	// Statment

	AssignStmt      = *ast.AssignStmt
	LocalAssignStmt = *ast.LocalAssignStmt
	FuncCallStmt    = *ast.FuncCallStmt
	DoBlockStmt     = *ast.DoBlockStmt
	WhileStmt       = *ast.WhileStmt
	RepeatStmt      = *ast.RepeatStmt
	IfStmt          = *ast.IfStmt
	NumberForStmt   = *ast.NumberForStmt
	GenericForStmt  = *ast.GenericForStmt
	FuncDefStmt     = *ast.FuncDefStmt
	ReturnStmt      = *ast.ReturnStmt
	BreakStmt       = *ast.BreakStmt
	LabelStmt       = *ast.LabelStmt
	GotoStmt        = *ast.GotoStmt

	// Expression

	TrueExpr           = *ast.TrueExpr
	FalseExpr          = *ast.FalseExpr
	NilExpr            = *ast.NilExpr
	NumberExpr         = *ast.NumberExpr
	StringExpr         = *ast.StringExpr
	Comma3Expr         = *ast.Comma3Expr
	IdentExpr          = *ast.IdentExpr
	AttrGetExpr        = *ast.AttrGetExpr
	TableExpr          = *ast.TableExpr
	FuncCallExpr       = *ast.FuncCallExpr
	LogicalOpExpr      = *ast.LogicalOpExpr
	RelationalOpExpr   = *ast.RelationalOpExpr
	StringConcatOpExpr = *ast.StringConcatOpExpr
	ArithmeticOpExpr   = *ast.ArithmeticOpExpr
	UnaryMinusOpExpr   = *ast.UnaryMinusOpExpr
	UnaryNotOpExpr     = *ast.UnaryNotOpExpr
	UnaryLenOpExpr     = *ast.UnaryLenOpExpr
	FunctionExpr       = *ast.FunctionExpr
)

func (*YockPack[T]) ParseStr(str string) []ast.Stmt {
	reader := bufio.NewReader(strings.NewReader(str))
	chunk, err := parse.Parse(reader, "<string>")
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	return chunk
}

func (*YockPack[T]) ParseFile(file string) []ast.Stmt {
	fp, err := os.Open(file)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	chunk, err := parse.Parse(reader, file)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	return chunk
}

func (yockpack *YockPack[T]) DumpStr(str string) string {
	stmts := yockpack.ParseStr(str)
	return parse.Dump(stmts)
}

func (yockpack *YockPack[T]) DumpFile(file string) string {
	stmts := yockpack.ParseFile(file)
	return parse.Dump(stmts)
}

type CompileOpt struct {
	DisableAnalyse bool
	VM             runtime.VirtualMachine
}

func (yockpack *YockPack[T]) Compile(opt CompileOpt, file string) *lua.LFunction {
	fp, err := os.Open(file)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	chunk, err := parse.Parse(reader, file)

	if opt.DisableAnalyse {
		anlyzer := NewLuaDependencyAnalyzer()
		out, err := utils.ReadStraemFromFile(util.Pathf("~/lib/dep/stdlib.json"))
		if err != nil {
			util.Ycho.Fatal(err.Error())
		}
		if err = json.Unmarshal(out, anlyzer); err != nil {
			util.Ycho.Fatal(err.Error())
		}
		files, err := os.ReadDir(util.Pathf("~/lib"))
		if err != nil {
			util.Ycho.Fatal(err.Error())
		}
		for _, file := range files {
			if fn := file.Name(); filepath.Ext(fn) == ".lua" {
				anlyzer.Load(util.Pathf("~/lib/") + fn)
			}
		}
		undefines, _ := anlyzer.Tidy(file)
		for _, undefine := range undefines {
			undefine = strings.TrimSuffix(undefine, "()")
			opt.VM.Eval(fmt.Sprintf(`%s = uninit_driver("%s")`, undefine, undefine))
		}
	}

	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	proto, err := lua.Compile(chunk, file)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	return opt.VM.Interp().NewFunctionFromProto(proto)
}
