// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"net"
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
		"is_exist":     is_exist,
		"printf":       printf,
		"pathf":        ioPathf,
		"open_conf":    openConf,
	})
	yocks.RegYocksFn(yocki.YocksFuncs{
		"curl": netCurl,
	})
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
	out := util.ConvertByte2String([]byte(s.CheckString(1)), util.Charset(s.CheckString(2)))
	s.PushString(out)
	return 1
}

func aliasHandle(cmd string) string {
	for k, v := range aliases {
		cmd = strings.ReplaceAll(cmd, "$"+k, v)
	}
	return cmd
}

/*
* @param opt? table
* @param cmds string
* @return table, err
 */
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
		util.ReadLineFromString(cmd, func(s string) string {
			if len(s) > 0 {
				out, err := yockc.Exec(opt, aliasHandle(s))
				outs.Append(lua.LString(out))
				if err != nil {
					if opt.Debug {
						ycho.Warn(err)
					}
					if opt.Strict {
						return ""
					} else {
						g_err = util.ErrGeneral
					}
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
				}, lua.LString(s)); err != nil {
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
		for i := 1; i < s.Argc(); i++ {
			urls = append(urls, s.CheckString(i))
		}
	}
	s.PushError(yockc.Curl(opt, urls))
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

// @param path string
//
// @return bool
func is_exist(s yocki.YockState) int {
	ok := util.IsExist(s.CheckString(1))
	s.PushBool(ok)
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

// ioPathf formats path
//
// @/abc => {WorkSpace}/abc (WorkSpace = UserHome + .yock)
//
// ~/abc => {YockPath}/abc (YockPath = executable file path)
/*
* @param path string
* @return string
 */
func ioPathf(s yocki.YockState) int {
	s.PushString(util.Pathf(s.CheckString(1)))
	return 1
}
