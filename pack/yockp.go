// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockp

import (
	"bufio"
	"encoding/json"
	"fmt"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/ast"
	"github.com/yuin/gopher-lua/parse"
	"os"
	"path/filepath"
	"strings"
)

// YockPack serves as yock's preprocessing tool for decomposing Lua source code
// and for dependency analysis when YPM is not introduced.
type YockPack[T any] struct{}

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

// ParseStr parses the given string into a lua statement structure
func (*YockPack[T]) ParseStr(str string) []ast.Stmt {
	reader := bufio.NewReader(strings.NewReader(str))
	chunk, err := parse.Parse(reader, "<string>")
	if err != nil {
		ycho.Fatal(err)
	}
	return chunk
}

// ParseFile parses the given file content into a lua statement structure
func (*YockPack[T]) ParseFile(file string) []ast.Stmt {
	fp, err := os.Open(file)
	if err != nil {
		ycho.Fatal(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	chunk, err := parse.Parse(reader, file)
	if err != nil {
		ycho.Fatal(err)
	}
	return chunk
}

// DumpStr prints out a syntax tree based on the given source code string
func (yockpack *YockPack[T]) DumpStr(str string) string {
	stmts := yockpack.ParseStr(str)
	return parse.Dump(stmts)
}

// DumpFile prints out a syntax tree based on the given file
func (yockpack *YockPack[T]) DumpFile(file string) string {
	stmts := yockpack.ParseFile(file)
	return parse.Dump(stmts)
}

type CompileOpt struct {
	DisableAnalyse bool
	VM             yocki.YockRuntime
}

// Compile compiles the contents of the given file into functions that can be executed by the virtual machine.
func (yockpack *YockPack[T]) Compile(opt CompileOpt, file string) *lua.LFunction {
	fp, err := os.Open(file)
	if err != nil {
		ycho.Fatal(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	chunk, err := parse.Parse(reader, file)

	if opt.DisableAnalyse {
		anlyzer := NewLuaDependencyAnalyzer()
		out, err := util.ReadStraemFromFile(util.Pathf("~/lib/dep/stdlib.json"))
		if err != nil {
			ycho.Fatal(err)
		}
		if err = json.Unmarshal(out, anlyzer); err != nil {
			ycho.Fatal(err)
		}
		files, err := os.ReadDir(util.Pathf("~/lib"))
		if err != nil {
			ycho.Fatal(err)
		}
		for _, file := range files {
			if fn := file.Name(); filepath.Ext(fn) == ".lua" {
				anlyzer.Load(util.Pathf("~/lib/") + fn)
			}
		}
		undefines, _ := anlyzer.Completion(file)
		for _, undefine := range undefines {
			undefine = strings.TrimSuffix(undefine, "()")
			opt.VM.Eval(fmt.Sprintf(`%s = uninit_driver("%s")`, undefine, undefine))
		}
	}

	if err != nil {
		ycho.Fatal(err)
	}
	proto, err := lua.Compile(chunk, file)
	if err != nil {
		ycho.Fatal(err)
	}
	return opt.VM.State().LState().NewFunctionFromProto(proto)
}
