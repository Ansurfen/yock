package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	_ "github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/cushion/utils/build"
	"github.com/ansurfen/yock"
	"github.com/gocolly/colly"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	file := os.Args[1]
	if ext := filepath.Ext(file); ext == ".lua" {
		vm := runtime.NewVirtualMachine()
		vm.SetGlobalFn(runtime.LuaFuncs{
			"mv": func(l *lua.LState) int {
				mv := yock.NewMoveCmd()
				mv.Exec(fmt.Sprintf("%s %s", l.CheckString(1), l.CheckString(2)))
				return 0
			},
			"cp": func(l *lua.LState) int {
				cp := yock.NewCpCmd()
				cp.Exec(fmt.Sprintf("-r %s %s", l.CheckString(1), l.CheckString(2)))
				return 0
			},
			"rm": func(l *lua.LState) int {
				rm := yock.NewRmCmd()
				opts := []string{}
				for i := 1; i <= l.GetTop(); i++ {
					if i == 1 && l.CheckAny(i).Type().String() == "table" {
						l.CheckTable(i).ForEach(func(_, opt lua.LValue) {
							opts = append(opts, opt.String())
						})
					} else {
						tmp := append(opts, l.CheckString(i))
						rm.Exec(strings.Join(tmp, " "))
					}
				}
				return 0
			},
			"curl": func(l *lua.LState) int {
				c := colly.NewCollector()
				c.OnResponse(func(r *colly.Response) {
					utils.WriteFile(l.CheckString(2), r.Body)
				})
				fmt.Println("curl: ", l.CheckString(1), l.CheckString(2))
				c.Visit(l.CheckString(1))
				return 0
			},
			"file_replace": func(l *lua.LState) int {
				src, err := utils.ReadStraemFromFile(l.CheckString(1))
				if err != nil {
					panic(err)
				}
				dst := strings.ReplaceAll(string(src), l.CheckString(2), l.CheckString(3))
				utils.WriteFile(l.CheckString(1), []byte(dst))
				return 0
			},
			"exec": func(l *lua.LState) int {
				l.CheckTable(1).ForEach(func(_, cmd lua.LValue) {
					if out, err := utils.ExecStr(cmd.String()); err != nil {
						fmt.Println(string(out))
						return
					}
				})
				return 0
			},
			"compose": func(l *lua.LState) int { return 0 },
			"loadpkg": func(l *lua.LState) int {
				vm.Eval(fmt.Sprintf(`package.path = package.path .. ";%s.lua";require("%s")`, l.CheckString(1), l.CheckString(1)))
				return 0
			},
			"gsub": func(l *lua.LState) int {
				tmp := []string{}
				for i := 0; i <= l.GetTop(); i++ {
					switch l.CheckAny(i).Type() {
					case lua.LTNumber:
						tmp = append(tmp, l.CheckNumber(i).String())
					case lua.LTString:
						tmp = append(tmp, l.CheckString(i))
					}
				}
				l.Push(lua.LString(strings.Join(tmp, " ")))
				return 1
			},
		})
		if len(os.Args) > 2 {
			if err := vm.Eval(parse(file, os.Args[2])); err != nil {
				panic(err)
			}
		} else {
			if err := vm.Eval(parse(file, "")); err != nil {
				panic(err)
			}
		}
	} else if ext == ".starlark" {
		os.Exit(0)
	}
}

type LuaFile struct {
	vars   map[string]bool
	scopes []string
	script [][]string
	scope  int
	eid    string
}

func (fp *LuaFile) insert(line string) {
	fp.script[fp.scope] = append(fp.script[fp.scope], line)
}

func (fp *LuaFile) newScope(name string) {
	fp.scope++
	fp.script = append(fp.script, []string{})
	fp.scopes = append(fp.scopes, name)
}

type yockFrame struct {
	p          *yockFrame
	dump       bool
	decl       bool
	scope      string
	file       string
	prtivate   map[string]string
	isPrivate  bool
	curPrivate string
	internal   bool
	scripts    string
}

type Comparable interface {
	string
}

type Stack[T Comparable] struct {
	data []T
	top  int
	cap  int
}

func NewStack[T Comparable](data []T) *Stack[T] {
	return &Stack[T]{
		data: data,
		top:  len(data) - 1,
		cap:  cap(data),
	}
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
	s.top = -1
	s.cap = 1
}

func (s *Stack[T]) Empty() bool {
	return s.top == -1
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
	s.flashParam()
}

func (s *Stack[T]) Pop() error {
	if s.Empty() {
		return errors.New("out of range")
	}
	s.data = s.data[:s.top]
	s.flashParam()
	return nil
}

func (s *Stack[T]) Top() (T, error) {
	if s.Empty() {
		return "", errors.New("out of range")
	}
	return s.data[s.top], nil
}

func (s *Stack[T]) flashParam() {
	s.top = len(s.data) - 1
	s.cap = cap(s.data)
}

func parse(file, mode string) string {
	fp := &LuaFile{
		scope:  0,
		script: make([][]string, 1),
		scopes: []string{"global"},
		eid:    utils.RandString(8),
		vars:   make(map[string]bool),
	}
	if len(mode) == 0 {
		mode = fp.eid
	} else {
		mode = strings.ReplaceAll(mode, ".", fp.eid)
	}
	frame := &yockFrame{prtivate: make(map[string]string)}
	parseBuilder(file, fp, frame)
	script := ""
	s := NewStack([]string{})
	for i := 0; i < len(fp.script); i++ {
		scope := fp.scopes[i]
		tmp := ""
		for j := 0; j < len(fp.script[i]); j++ {
			tmp += fp.script[i][j] + "\n"
		}
		if i+1 < len(fp.scopes) && strings.HasPrefix(fp.scopes[i+1], scope) {
			script += tmp
			s.Push(scope)
		} else {
			if scope == "global" {
				script += tmp + runtime.LuaGoto("{{.}}").Value() + "\n"
				continue
			}
			if s.Empty() {
				script += tmp + runtime.LuaGoto(fp.eid).Value() + "\n"
				continue
			} else {
				e, _ := s.Top()
				if strings.HasPrefix(scope, e) {
					script += tmp + runtime.LuaIf(
						[]string{fmt.Sprintf(`"" == "%s"`, scope)},
						runtime.LuaGoto(fp.eid)).Value() + "\n"
					continue
				} else {
					s.Pop()
					script += runtime.LuaIf(
						[]string{fmt.Sprintf(`"" == "%s"`, e)},
						runtime.LuaGoto(fp.eid)).Value() + "\n" +
						tmp + runtime.LuaIf(
						[]string{fmt.Sprintf(`"" == "%s"`, scope)},
						runtime.LuaGoto(fp.eid)).Value() + "\n"
				}
			}
		}
	}
	script += fmt.Sprintf("::%s::\n", fp.eid)
	tmpl := build.NewTemplate()
	script, _ = tmpl.OnceParse(script, mode)
	if frame.dump {
		fmt.Println(script)
	}
	return script
}

var (
	isLocalVar       = regexp.MustCompile(`local *(\w*)`)
	isMeta           = regexp.MustCompile(`--- *@meta:(.*)`)
	isLink           = regexp.MustCompile(`--- *@link (.*)`)
	isLabel          = regexp.MustCompile(`::([A-Za-z0-9_]+)::`)
	isSyntaxStart    = regexp.MustCompile(`--\[\[(.*)`)
	isSyntaxEnd      = regexp.MustCompile(`(.*)]]$`)
	isSyntaxCompress = regexp.MustCompile(`--\[\[(.*)]]`)
)

func readLine(s string, cb func(s string) string) ([]byte, error) {
	if path.Ext(s) == ".lua" {
		return utils.ReadLineFromFile(s, cb)
	}
	return utils.ReadLineFromString(s, cb)
}

func parseBuilder(file string, fp *LuaFile, frame *yockFrame) {
	if file == frame.file {
		scope := ""
		if frame.p != nil {
			scope = frame.p.scope
		}
		fmt.Printf("cycle import, at file %s, scope %s\n", file, scope)
		os.Exit(1)
	}
	readLine(file, func(s string) string {
		if frame.internal {
			parseScript(s, fp, frame)
			return ""
		}
		parseMeta(s, frame)
		if parseScope(s, fp, frame) {
			return ""
		}
		if frame.isPrivate {
			frame.prtivate[frame.curPrivate] += s + "\n"
			return ""
		}
		if m := isLink.FindStringSubmatch(s); len(m) > 1 {
			attrs := strings.Split(m[1], " ")
			if path.Ext(attrs[0]) == ".lua" {
				var nframe *yockFrame
				if fp.scopes[fp.scope] == "global" {
					nframe = &yockFrame{file: file, p: nil, decl: frame.decl, prtivate: make(map[string]string)}
				} else {
					nframe = &yockFrame{p: frame, file: file, decl: frame.decl, prtivate: make(map[string]string)}
				}
				parseBuilder(attrs[0], fp, nframe)
			} else {
				if p, ok := frame.prtivate[attrs[0]]; ok {
					parseBuilder(p, fp, frame)
				}
			}
			return ""
		}
		if parseLocalVar(s, fp, frame) {
			return ""
		}
		if parseScript(s, fp, frame) {
			return ""
		}
		fp.insert(s)
		return ""
	})
}

func parseInsertLine(s string, fp *LuaFile, frame *yockFrame) {
	idx := len(fp.scopes) - 1
	for i := 0; i < len(fp.scopes); i++ {
		if fp.scopes[i] == frame.scope {
			idx = i
			break
		}
	}
	if idx != len(fp.scopes) {
	} else {
		fp.insert(s)
	}
}

func parseScript(s string, fp *LuaFile, frame *yockFrame) bool {
	if m := isSyntaxCompress.FindStringSubmatch(s); len(m) > 1 {
		return true
	}
	if m := isSyntaxStart.FindStringSubmatch(s); len(m) > 1 {
		frame.internal = true
		frame.scripts += m[1]
		return true
	}
	if m := isSyntaxEnd.FindStringSubmatch(s); len(m) > 1 {
		frame.internal = false
		frame.scripts = ""
		return true
	}
	if frame.internal {
		frame.scripts += s
		return true
	}
	return false
}

func parseScope(s string, fp *LuaFile, frame *yockFrame) bool {
	if m := isLabel.FindStringSubmatch(s); len(m) > 1 {
		label := m[1]
		if strings.HasPrefix(label, "__") {
			frame.isPrivate = true
			frame.curPrivate = label
			return true
		}
		if frame.p != nil {
			s = frame.p.scope
			fp.newScope(s + "." + label)
			frame.scope = frame.p.scope + "." + label
			fp.insert(fmt.Sprintf("::%s::", strings.ReplaceAll(s+fp.eid+label, ".", fp.eid)))
		} else {
			fp.newScope(label)
			frame.scope = label
			fp.insert(fmt.Sprintf(`::%s::`, label))
		}
		frame.isPrivate = false
		return true
	}
	return false
}

func parseLocalVar(s string, fp *LuaFile, fram *yockFrame) bool {
	if m := isLocalVar.FindStringSubmatch(s); len(m) > 1 {
		name := strings.TrimSpace(m[1])
		if _, ok := fp.vars[name]; ok && fram.decl {
			fp.insert(name + " " + isLocalVar.ReplaceAllString(s, ""))
		} else {
			fp.vars[name] = true
			fp.insert(s)
		}
		return true
	}
	return false
}

func parseMeta(s string, frame *yockFrame) {
	if m := isMeta.FindStringSubmatch(s); len(m) > 1 {
		for _, v := range strings.Split(m[1], ",") {
			switch strings.TrimSpace(v) {
			case "dump":
				frame.dump = true
			case "decl":
				frame.decl = true
			default:
			}
		}
	}
}
