// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockp

import (
	"fmt"
	"testing"
)

func TestVisitStmt(t *testing.T) {
	yockPack := YockPack[NilFrame]{}
	fmt.Println(yockPack.DumpStr(`local a = 10`))
	yockPack.VisitStr(`local a = 10`, NilFrame{}, VisitStmtHandle[NilFrame]{
		StmtAssign: func(idx int, stmt yockStmt, frame NilFrame) {
			fmt.Println(stmt.(AssignStmt))
		},
		StmtLocalAssign: func(idx int, stmt yockStmt, frame NilFrame) {
			s := stmt.(LocalAssignStmt)
			yockPack.VisitExpr(s.Exprs, frame, VisitExprHandle[NilFrame]{
				ExprNumber: func(idx int, expr yockExpr, frame NilFrame) {
					fmt.Println(expr.(NumberExpr))
				},
			})
		},
		HandleDefault: func(idx int, stmt yockStmt, frame NilFrame) {
			fmt.Println(stmt)
		},
	})
}
