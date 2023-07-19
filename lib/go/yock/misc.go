// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	yockc "github.com/ansurfen/yock/cmd"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

func LoadMisc(yocks yocki.YockScheduler) {
	yocks.RegYockFn(yocki.YockFuns{
		"sh":           osSh,
		"cmdf":         osCmdf,
		"strf":         osStrf,
		"new_command":  osNewCommand,
		"is_url":       netIsURL,
		"is_localhost": netIsLocalhost,
		"safe_write":   safe_write,
		"write_file":   write_file,
		"printf":       printf,
		"pathf":        pathf,
		"open_conf":    openConf,
		"eval":         eval,
		"loadbalance":  loadbalance,
		"yassert":      yassert,
	})
	yocks.RegYocksFn(yocki.YocksFuncs{
		"curl": netCurl,
	})
	lib := yocks.CreateLib("ycho")
	lib.SetYFunction(map[string]yocki.YGFunction{
		"info": ychoInfo,
		"warn": ychoWarn,
	})
	yocks.SetGlobalFn(map[string]lua.LGFunction{
		"print": func(l *lua.LState) int {
			top := l.GetTop()
			for i := 1; i <= top; i++ {
				ycho.Print(l.ToStringMeta(l.Get(i)).String())
				if i != top {
					ycho.Print("\t")
				}
			}
			ycho.Println("")
			return 0
		},
	})
}

func ychoInfo(s yocki.YockState) int {
	ycho.Infof("%s%s", s.Stacktrace(), s.CheckString(1))
	return 0
}

func ychoWarn(s yocki.YockState) int {
	ycho.Warnf("%s%s", s.Stacktrace(), s.CheckString(1))
	return 0
}

func yassert(s yocki.YockState) int {
	if s.IsBool(1) {
		ok := s.CheckBool(1)
		if !ok {
			s.LState().RaiseError(s.LState().OptString(2, "assertion failed!"))
		}
	} else if s.IsString(1) {
		s.LState().RaiseError(s.CheckString(1))
	} else if s.IsUserData(1) {
		s.LState().RaiseError(s.LState().OptString(2, "assertion failed!"))
	}
	if s.Argc() > 1 && s.IsFunction(2) {
		s.Call(yocki.YockFuncInfo{
			Fn: s.CheckFunction(2),
		})
	}
	return 0
}

type loadbalanceOpt struct {
	MaxRetry int
}

type handle func() error

func loadbalance(s yocki.YockState) int {
	elements := []handle{}
	opt := loadbalanceOpt{}
	if err := s.CheckTable(1).Bind(&opt); err != nil {
		s.Throw(err)
		return 1
	}
	s.CheckTable(2).Value().ForEach(func(_, l2 lua.LValue) {
		elements = append(elements, func() error {
			err := s.Call(yocki.YockFuncInfo{
				Fn:   l2,
				NRet: 1,
			})
			if err != nil {
				return err
			}
			if msg := s.CheckString(1); len(msg) > 0 {
				err = errors.New(msg)
			}
			s.PopTop()
			return err
		})
	})
	s.PopTop()
	s.PopTop()
	if opt.MaxRetry == 0 {
		opt.MaxRetry = len(elements)/2 + 1
	}
	ld := util.NewWeightedRandom(elements)
	for i := 0; i < opt.MaxRetry; i++ {
		call, idx := ld.Next()
		if err := call(); err != nil {
			ld.Down(idx)
		} else {
			break
		}
	}
	s.PushNil()
	return 1
}

func eval(s yocki.YockState) int {
	err := s.LState().DoString(s.CheckString(1))
	s.PushError(err)
	return 1
}

// @param path string
//
// @return userdata, err
func openConf(s yocki.YockState) int {
	conf, err := util.OpenConf(s.CheckString(1))
	s.Pusha(conf).PushError(err)
	return 2
}

/*
* @param str string
* @param charset string
* @return string
 */
func osStrf(s yocki.YockState) int {
	out := ""
	format := s.CheckString(1)
	if s.IsTable(2) {
		opt := make(map[string]any)
		if err := s.CheckTable(2).Bind(&opt); err != nil {
			s.PushString(out)
			return 1
		}
		tmpl := util.NewTemplate()
		out, _ = tmpl.OnceParse(format, opt)
		if v := opt["Charset"]; v != nil {
			if charset, ok := v.(string); ok {
				out = util.ConvertByte2String([]byte(out), util.Charset(charset))
			}
		}
	} else {
		a := []any{}
		for i := 2; i <= s.Argc(); i++ {
			a = append(a, s.CheckAny(i))
		}
		out = fmt.Sprintf(format, a...)
	}
	s.PushString(out)
	return 1
}

// @param opt? table
//
// @param cmds string
//
// @return table, err
func osSh(s yocki.YockState) int {
	cmds := []string{}
	opt := yockc.ExecOpt{Quiet: true}
	if s.IsTable(1) {
		tbl := s.CheckTable(1)
		if err := tbl.Bind(&opt); err != nil {
			return s.PushNil().Throw(err).Exit()
		}
		for i := 2; i <= s.Argc(); i++ {
			cmds = append(cmds, s.CheckString(i))
		}
	} else {
		opt.Redirect = true
		for i := 1; i <= s.Argc(); i++ {
			cmds = append(cmds, s.CheckString(i))
		}
	}
	outs := &lua.LTable{}
	var g_err error
	for _, cmd := range cmds {
		util.ReadLineFromString(cmd, func(str string) string {
			if len(str) > 0 {
				out, err := yockc.Exec(opt, str)
				ychoLogger(err, "%ssh %s", s.Stacktrace(), str)
				outs.Append(lua.LString(out))
				if err != nil {
					g_err = err
				}
			}
			return ""
		})
	}
	return s.Push(outs).PushError(g_err).Exit()
}

// @param cmd ...string
//
// @return string
func osCmdf(s yocki.YockState) int {
	tmp := []string{}
	for i := 0; i <= s.Argc(); i++ {
		switch s.LState().CheckAny(i).Type() {
		case lua.LTNumber:
			tmp = append(tmp, s.LState().CheckNumber(i).String())
		case lua.LTString:
			tmp = append(tmp, s.LState().CheckString(i))
		}
	}
	s.PushString(strings.Join(tmp, " "))
	return 1
}

// @return userdata
func osNewCommand(l yocki.YockState) int {
	l.Pusha(&cobra.Command{})
	return 1
}

// netCurl is capable of sending HTTP requests, which defaults to the GET method
/*
* @param opt table
* @param urls ...string
* @retrun err
 */
func netCurl(yocks yocki.YockScheduler, s yocki.YockState) int {
	opt := yockc.CurlOpt{Method: "GET"}
	urls := []string{}
	if s.IsTable(1) {
		tbl := s.CheckTable(1)

		if fn := tbl.Value().RawGetString("filename"); fn.Type() == lua.LTFunction {
			opt.FilenameHandle = func(s string) string {
				tmp, _ := yocks.NewState()
				if err := tmp.Call(yocki.YockFuncInfo{
					NRet: 1,
					Fn:   fn,
				}, s); err != nil {
					panic(err)
				}
				return tmp.CheckString(1)
			}
		}

		if err := tbl.Bind(&opt); err != nil {
			s.Throw(err)
			return s.Exit()
		}
		for i := 2; i <= s.Argc(); i++ {
			urls = append(urls, s.CheckString(i))
		}
	} else {
		for i := 1; i <= s.Argc(); i++ {
			urls = append(urls, s.CheckString(i))
		}
	}
	str, err := yockc.Curl(opt, urls)
	ychoLogger(err, "%scurl %s", s.Stacktrace(), strings.Join(urls, ","))
	s.PushString(string(str)).PushError(err)
	return s.Exit()
}

// @param url string
//
// @return bool
func netIsURL(s yocki.YockState) int {
	s.PushBool(util.IsURL(s.CheckString(1)))
	return 1
}

// @param url string
//
// @return bool
func netIsLocalhost(s yocki.YockState) int {
	url := s.CheckString(1)
	if url == "localhost" {
		return s.PushBool(true).Exit()
	}
	addrs, err := net.LookupHost("localhost")
	if err != nil {
		return s.Throw(err).Exit()
	}
	return s.PushBool(len(addrs) > 1 && addrs[1] == url).Exit()
}

/*
* @param file string
* @param data string
* @return err
 */
func safe_write(s yocki.YockState) int {
	err := util.SafeWriteFile(s.CheckString(1), []byte(s.CheckString(2)))
	s.PushError(err)
	return 1
}

/*
* @param file string
* @param data string
* @return err
 */
func write_file(s yocki.YockState) int {
	err := util.WriteFile(s.CheckString(1), []byte(s.CheckString(2)))
	s.PushError(err)
	return 1
}

// @param title string
//
// @param rows string[][]
func printf(s yocki.YockState) int {
	title := []string{}
	rows := [][]string{}
	s.CheckTable(1).Value().ForEach(func(idx, el lua.LValue) {
		title = append(title, el.String())
	})
	s.CheckTable(2).Value().ForEach(func(ri, row lua.LValue) {
		tmp := []string{}
		row.(*lua.LTable).ForEach(func(fi, field lua.LValue) {
			tmp = append(tmp, field.String())
		})
		rows = append(rows, tmp)
	})
	util.Prinf(util.PrintfOpt{MaxLen: 30}, title, rows)
	return 0
}

// pathf formats path
//
// @/abc => {WorkSpace}/abc (WorkSpace = UserHome + .yock)
//
// ~/abc => {YockPath}/abc (YockPath = executable file path)
//
// @varag string
//
// @return string
func pathf(s yocki.YockState) int {
	path := s.CheckString(1)
	if len(path) > 0 && path[0] == '#' {
		i, err := strconv.Atoi(path[1:])
		if err != nil {
			s.PushString("")
			return 1
		}
		dbg, ok := s.Stack(i)
		if !ok {
			s.PushString("")
			return 1
		}
		path = dbg.Source
	} else {
		path = util.Pathf(path)
	}
	elem := []string{path}
	for i := 2; i <= s.Argc(); i++ {
		elem = append(elem, s.CheckString(i))
	}
	s.PushString(filepath.Join(elem...))
	return 1
}

func pathfV2(s yocki.YockState) int {
	path := s.CheckString(1)
	if len(path) == 0 {
		s.PushString("")
		return 1
	}
	switch path[0] {
	case '#':
		i, err := strconv.Atoi(path[1:])
		if err != nil {
			s.PushString("")
			return 1
		}
		dbg, ok := s.Stack(i)
		if !ok {
			s.PushString("")
			return 1
		}
		path = dbg.Source
	case '~':
		wd, err := os.UserHomeDir()
		if err != nil {
			s.PushString("")
			return 1
		}
		path = wd
	case '!':
		abs, err := filepath.Abs(path[1:])
		if err != nil {
			s.PushString("")
			return 1
		}
		path = abs
	case '?':

	case '@':

	case '$':
		wd, err := os.Getwd()
		if err != nil {
			s.PushString("")
			return 1
		}
		path = wd
	}
	elem := []string{path}
	for i := 2; i <= s.Argc(); i++ {
		elem = append(elem, s.CheckString(i))
	}
	s.PushString(filepath.Join(elem...))
	return 1
}
