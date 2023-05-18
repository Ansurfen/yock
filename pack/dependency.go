package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/yuin/gopher-lua/parse"
	"os"
	"path/filepath"
	"strings"

	"github.com/ansurfen/cushion/utils"
	"github.com/yuin/gopher-lua/ast"
)

type luaDependencyAnalyzer struct {
	Includes map[string][]LuaMethod `json:"includes"`
}

func NewLuaDependencyAnalyzer() *luaDependencyAnalyzer {
	return &luaDependencyAnalyzer{
		Includes: make(map[string][]LuaMethod),
	}
}

func (analyzer *luaDependencyAnalyzer) Tidy(file string) ([]string, map[string]bool) {
	ast := ParserASTFromFile(file)
	dec, cal := parseFuncStmt(file, ast)
	deps := make(map[string]bool)
	libs := []string{}
	for name := range cal {
		if _, ok := dec[name]; ok {
			continue
		}
		if m, ok := analyzer.Includes[name]; !ok {
			libs = append(libs, name)
		} else {
			deps[m[0].Pkg] = true
		}
	}
	return libs, deps
}

// Lookup returns files of declare of methods according to specify name
func (analyzer *luaDependencyAnalyzer) Lookup(name string) []LuaMethod {
	return analyzer.Includes[name]
}

// Load to anaylse declare of method in specify lua script
func (analyzer *luaDependencyAnalyzer) Load(str string) {
	var ast []ast.Stmt
	scope := "g"
	if filepath.Ext(str) == ".lua" {
		ast = ParserASTFromFile(str)
		scope = str
	} else {
		ast = ParserASTFromString(str)
	}
	dec, _ := parseFuncStmt(scope, ast)
	for name, method := range dec {
		analyzer.Includes[name] = append(analyzer.Includes[name], method)
	}
}

func (analyzer *luaDependencyAnalyzer) LoadG(str string) {
	var ast []ast.Stmt
	scope := "g"
	if filepath.Ext(str) == ".lua" {
		ast = ParserASTFromFile(str)
	} else {
		ast = ParserASTFromString(str)
	}
	dec, _ := parseFuncStmt(scope, ast)
	for name, method := range dec {
		analyzer.Includes[name] = append(analyzer.Includes[name], method)
	}
}

func (analyzer *luaDependencyAnalyzer) Preload(name string, method LuaMethod) {
	analyzer.Includes[name] = append(analyzer.Includes[name], method)
}

func (analyzer *luaDependencyAnalyzer) Export(file string) {
	out, err := json.Marshal(analyzer)
	if err != nil {
		panic(err)
	}
	utils.WriteFile(file, out)
}

type LuaMethod struct {
	Argc int      `json:"argc"`
	Argv []string `json:"argv"`
	Pkg  string   `json:"pkg"`
}

func parseFuncStmt(scope string, stmts []ast.Stmt) (declares, calls map[string]LuaMethod) {
	declares = make(map[string]LuaMethod)
	calls = make(map[string]LuaMethod)
	for _, stmt := range stmts {
		switch v := stmt.(type) {
		case *ast.FuncCallStmt:
			name := parseFuncExpr(v.Expr)
			calls[name] = LuaMethod{
				Pkg: scope,
			}
		case *ast.FuncDefStmt:
			name := ""
			if v.Name.Func != nil {
				name = fmt.Sprintf("%s()", parseFuncExpr(v.Name.Func))
			} else if v.Name.Receiver != nil {
				name = fmt.Sprintf("%s:%s()", parseFuncExpr(v.Name.Receiver), v.Name.Method)
			} else {
				continue
			}
			declares[name] = LuaMethod{
				Argv: v.Func.ParList.Names,
				Argc: len(v.Func.ParList.Names),
				Pkg:  scope,
			}
		}
	}
	return
}

func parseFuncExpr(expr ast.Expr) (ret string) {
	switch v := expr.(type) {
	case *ast.FuncCallExpr:
		if v.Func != nil {
			ret += fmt.Sprintf("%s()", parseFuncExpr(v.Func))
		} else if v.Receiver != nil {
			ret += fmt.Sprintf("%s:%s()", parseFuncExpr(v.Receiver), v.Method)
		}
	case *ast.AttrGetExpr:
		ret += fmt.Sprintf("%s.%s", parseFuncExpr(v.Object), parseFuncExpr(v.Key))
	case *ast.StringExpr:
		ret += v.Value
	case *ast.IdentExpr:
		ret += v.Value
	}
	return
}

func ParserASTFromFile(file string) []ast.Stmt {
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	chunk, err := parse.Parse(reader, file)
	if err != nil {
		panic(err)
	}
	return chunk
}

func ParserASTFromString(str string) []ast.Stmt {
	reader := bufio.NewReader(strings.NewReader(str))
	chunk, err := parse.Parse(reader, "<string>")
	if err != nil {
		panic(err)
	}
	return chunk
}
