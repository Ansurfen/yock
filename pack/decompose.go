// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockp

import (
	"strconv"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/yuin/gopher-lua/ast"
)

// DecomposeOpt indicates configuration of Decompose
type DecomposeOpt struct {
	// file to be decomposed
	File string
	// divide file into modes to be specified
	Modes []string
	// template file, is used to generate control script.
	Tpl    string
	Prefix string
}

type modeBlock struct {
	limit  int
	filter map[int]bool
}

const defaultPrefix = "host"

func (yockpack *YockPack[T]) Decompose(opt DecomposeOpt, stmts []ast.Stmt) {
	var frame T
	tasks := make(map[string][]string)
	records := make(map[string]int)
	yockpack.VisitStmt(stmts, frame, VisitStmtHandle[T]{
		StmtFuncCall: func(si int, stmt yockStmt, frame T) {
			v := stmt.(FuncCallStmt)
			yockpack.VisitExpr([]yockExpr{v.Expr}, frame, VisitExprHandle[T]{
				ExprFuncCall: func(ei int, expr yockExpr, frame T) {
					v := expr.(FuncCallExpr)
					if vv, ok := v.Func.(IdentExpr); ok && vv.Value == "jobs" {
						if len(v.Args) < 2 {
							return
						}
						taskName := ""
						jobs := []string{}
						for idx, arg := range v.Args {
							if str, ok := arg.(StringExpr); ok {
								if idx == 0 {
									taskName = str.Value
								} else {
									jobs = append(jobs, str.Value)
								}
							}
						}
						if len(taskName) > 0 {
							tasks[taskName] = jobs
						}
					} else if ok && vv.Value == "job" {
						if len(v.Args) < 2 {
							return
						}
						if str, ok := v.Args[0].(StringExpr); ok {
							records[str.Value] = si
							tasks[str.Value] = append(tasks[str.Value], str.Value)
						}
					}
				},
			})
		},
	})
	modeBlocks := make([]modeBlock, len(opt.Modes))
	for i, mode := range opt.Modes {
		if modeBlocks[i].filter == nil {
			modeBlocks[i].filter = make(map[int]bool)
		}
		max := -1
		for _, record := range records {
			modeBlocks[i].filter[record] = true
		}
		for _, job := range tasks[mode] {
			jobPos := records[job]
			modeBlocks[i].filter[jobPos] = false
			if max < jobPos {
				max = jobPos
			}
		}
		modeBlocks[i].limit = max
	}
	prefix := defaultPrefix
	if len(opt.Prefix) > 0 {
		prefix = opt.Prefix
	}
	unique := util.RandString(3)
	for idx, mb := range modeBlocks {
		if mb.limit == -1 {
			ycho.Warnf("invalid mode block")
			continue
		}
		util.WriteFile(unique+prefix+strconv.Itoa(idx)+".lua", []byte(yockpack.BuildScript(stmts[:mb.limit+1], mb.filter)))
	}
	if len(opt.File) == 0 {
		opt.File = unique + prefix
	}
	yockpack.buildBoot(buildBootOpt{
		file:  opt.File,
		tpl:   opt.Tpl,
		modes: opt.Modes,
	})
}
