// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockpack

import (
	"github.com/yuin/gopher-lua/ast"
)

type (
	VisitStmtHandle[T any] map[uint8]func(idx int, stmt yockStmt, frame T)
	VisitExprHandle[T any] map[uint8]func(idx int, expr yockExpr, frame T)
)

func (yockpack *YockPack[T]) VisitStr(str string, frame T, handle VisitStmtHandle[T]) {
	yockpack.VisitStmt(yockpack.ParseStr(str), frame, handle)
}

func (yockpack *YockPack[T]) VisitFile(file string, frame T, handle VisitStmtHandle[T]) {
	yockpack.VisitStmt(yockpack.ParseFile(file), frame, handle)
}

// VisitStmt recursively traverses the lua statement,
// and you can pass a callback function to handle the incoming statement.
func (yockpack *YockPack[T]) VisitStmt(stmts []ast.Stmt, frame T, handle VisitStmtHandle[T]) {
	for idx, stmt := range stmts {
		switch v := stmt.(type) {
		case AssignStmt:
			if fn, ok := handle[StmtAssign]; ok {
				fn(idx, v, frame)
			}
		case LocalAssignStmt:
			if fn, ok := handle[StmtLocalAssign]; ok {
				fn(idx, v, frame)
			}
		case FuncCallStmt:
			if fn, ok := handle[StmtFuncCall]; ok {
				fn(idx, v, frame)
			}
		case DoBlockStmt:
			if fn, ok := handle[StmtDoBlock]; ok {
				fn(idx, v, frame)
			}
		case WhileStmt:
			if fn, ok := handle[StmtWhile]; ok {
				fn(idx, v, frame)
			}
		case RepeatStmt:
			if fn, ok := handle[StmtRepeat]; ok {
				fn(idx, v, frame)
			}
		case IfStmt:
			if fn, ok := handle[StmtIf]; ok {
				fn(idx, v, frame)
			}
		case NumberForStmt:
			if fn, ok := handle[StmtNumbderFor]; ok {
				fn(idx, v, frame)
			}
		case GenericForStmt:
			if fn, ok := handle[StmtGenericFor]; ok {
				fn(idx, v, frame)
			}
		case FuncDefStmt:
			if fn, ok := handle[StmtFuncDef]; ok {
				fn(idx, v, frame)
			}
		case ReturnStmt:
			if fn, ok := handle[StmtReturn]; ok {
				fn(idx, v, frame)
			}
		case BreakStmt:
			if fn, ok := handle[StmtBreak]; ok {
				fn(idx, v, frame)
			}
		case LabelStmt:
			if fn, ok := handle[StmtLabel]; ok {
				fn(idx, v, frame)
			}
		case GotoStmt:
			if fn, ok := handle[StmtGoto]; ok {
				fn(idx, v, frame)
			}
		default:
			if fn, ok := handle[HandleDefault]; ok {
				fn(idx, v, frame)
			}
		}
	}
}

// VisitExpr recursively traverses the lua expression,
// and you can pass a callback function to handle the incoming expression.
func (yockpack *YockPack[T]) VisitExpr(exprs []ast.Expr, frame T, handle VisitExprHandle[T]) {
	for idx, expr := range exprs {
		switch v := expr.(type) {
		case TrueExpr:
			if fn, ok := handle[ExprTrue]; ok {
				fn(idx, v, frame)
			}
		case FalseExpr:
			if fn, ok := handle[ExprFalse]; ok {
				fn(idx, v, frame)
			}
		case NilExpr:
			if fn, ok := handle[ExprNil]; ok {
				fn(idx, v, frame)
			}
		case NumberExpr:
			if fn, ok := handle[ExprNumber]; ok {
				fn(idx, v, frame)
			}
		case StringExpr:
			if fn, ok := handle[ExprString]; ok {
				fn(idx, v, frame)
			}
		case Comma3Expr:
			if fn, ok := handle[ExprComma3]; ok {
				fn(idx, v, frame)
			}
		case IdentExpr:
			if fn, ok := handle[ExprIdent]; ok {
				fn(idx, v, frame)
			}
		case AttrGetExpr:
			if fn, ok := handle[ExprAttrGet]; ok {
				fn(idx, v, frame)
			}
		case TableExpr:
			if fn, ok := handle[ExprTable]; ok {
				fn(idx, v, frame)
			}
		case FuncCallExpr:
			if fn, ok := handle[ExprFuncCall]; ok {
				fn(idx, v, frame)
			}
		case LogicalOpExpr:
			if fn, ok := handle[ExprLogicalOp]; ok {
				fn(idx, v, frame)
			}
		case RelationalOpExpr:
			if fn, ok := handle[ExprRelationalOp]; ok {
				fn(idx, v, frame)
			}
		case StringConcatOpExpr:
			if fn, ok := handle[ExprStringConcatOp]; ok {
				fn(idx, v, frame)
			}
		case ArithmeticOpExpr:
			if fn, ok := handle[ExprArithmeticOp]; ok {
				fn(idx, v, frame)
			}
		case UnaryMinusOpExpr:
			if fn, ok := handle[ExprUnaryMinus]; ok {
				fn(idx, v, frame)
			}
		case UnaryNotOpExpr:
			if fn, ok := handle[ExprUnaryNotOp]; ok {
				fn(idx, v, frame)
			}
		case UnaryLenOpExpr:
			if fn, ok := handle[ExprUnaryLenOp]; ok {
				fn(idx, v, frame)
			}
		case FunctionExpr:
			if fn, ok := handle[ExprFunciton]; ok {
				fn(idx, v, frame)
			}
		default:
			if fn, ok := handle[HandleDefault]; ok {
				fn(idx, v, frame)
			}
		}
	}
}
