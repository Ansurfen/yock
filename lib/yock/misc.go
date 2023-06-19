// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	"net"
	"strings"

	"github.com/ansurfen/cushion/utils"
	yockc "github.com/ansurfen/yock/cmd"
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
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
		"safe_write": safe_write,
		"zip":        zip,
		"unzip":      unzip,
		"write_file": write_file,
		"is_exist":   is_exist,
		"printf":     printf,
		"pathf":      ioPathf,
		"read_file":  ioReadFile,
	})
	yocks.RegYocksFn(yocki.YocksFuncs{
		"http": netHTTP,
	})
}

/*
* @param str string
* @param charset string
* @return string
 */
func osStrf(l *yockr.YockState) int {
	out := utils.ConvertByte2String([]byte(l.CheckString(1)), utils.Charset(l.CheckString(2)))
	l.PushString(out)
	return 1
}

/*
* @param opt? table
* @param cmds string
* @return table, err
 */
func osSh(l *yockr.YockState) int {
	cmds := []string{}
	opt := yockc.ExecOpt{Quiet: true}
	if l.IsTable(1) {
		tbl := l.CheckTable(1)
		if err := tbl.Bind(&opt); err != nil {
			return l.PushNil().Throw(err).Exit()
		}
		for i := 2; i <= l.GetTop(); i++ {
			cmds = append(cmds, l.CheckString(i))
		}
	} else {
		opt.Redirect = true
		for i := 1; i <= l.GetTop(); i++ {
			cmds = append(cmds, l.CheckString(i))
		}
	}
	outs := &lua.LTable{}
	var g_err error
	for _, cmd := range cmds {
		utils.ReadLineFromString(cmd, func(s string) string {
			if len(s) > 0 {
				out, err := yockc.Exec(opt, s)
				outs.Append(lua.LString(out))
				if err != nil {
					if opt.Debug {
						util.Ycho.Warn(err.Error())
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

	return l.Push(outs).PushError(g_err).Exit()
}

// @param cmd ...string
//
// @return string
func osCmdf(l *yockr.YockState) int {
	tmp := []string{}
	for i := 0; i <= l.GetTop(); i++ {
		switch l.CheckAny(i).Type() {
		case lua.LTNumber:
			tmp = append(tmp, l.CheckNumber(i).String())
		case lua.LTString:
			tmp = append(tmp, l.CheckString(i))
		}
	}
	l.PushString(strings.Join(tmp, " "))
	return 1
}

// @return userdata
func osNewCommand(l *yockr.YockState) int {
	l.Pusha(&cobra.Command{})
	return 1
}

// netHTTP is capable of sending HTTP requests, which defaults to the GET method
/*
* @param opt table
* @param urls ...string
* @retrun err
 */
func netHTTP(yocks yocki.YockScheduler, s *yockr.YockState) int {
	opt := yockc.HttpOpt{Method: "GET"}
	urls := []string{}
	if s.IsTable(1) {
		tbl := s.CheckTable(1)

		if fn := tbl.RawGetString("filename"); fn.Type() == lua.LTFunction {
			opt.Filename = func(s string) string {
				lvm, _ := yocks.State().NewThread()
				if err := lvm.CallByParam(lua.P{
					NRet: 1,
					Fn:   fn.(*lua.LFunction),
				}, lua.LString(s)); err != nil {
					panic(err)
				}
				return lvm.CheckString(1)
			}
		}

		if err := tbl.Bind(&opt); err != nil {
			s.Throw(err)
			return s.Exit()
		}
		for i := 2; i <= s.GetTop(); i++ {
			urls = append(urls, s.CheckString(i))
		}
	} else {
		for i := 1; i < s.GetTop(); i++ {
			urls = append(urls, s.CheckString(i))
		}
	}
	s.PushError(yockc.HTTP(opt, urls))
	return s.Exit()
}

// @param url string
//
// @return bool
func netIsURL(l *yockr.YockState) int {
	l.PushBool(utils.IsURL(l.CheckString(1)))
	return 1
}

// @param url string
//
// @return bool
func netIsLocalhost(s *yockr.YockState) int {
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
func safe_write(l *yockr.YockState) int {
	err := utils.SafeWriteFile(l.CheckString(1), []byte(l.CheckString(2)))
	l.PushError(err)
	return 1
}

/*
* @param src string
* @param dst string
* @return err
 */
func zip(l *yockr.YockState) int {
	zipPath := l.CheckString(1)
	paths := []string{}
	for i := 2; i <= l.GetTop(); i++ {
		paths = append(paths, l.CheckString(i))
	}
	err := utils.Zip(zipPath, paths...)
	l.PushError(err)
	return 1
}

/*
* @param src string
* @param dst string
* @return err
 */
func unzip(l *yockr.YockState) int {
	err := utils.Unzip(l.CheckString(1), l.CheckString(2))
	l.PushError(err)
	return 1
}

/*
* @param file string
* @param data string
* @return err
 */
func write_file(l *yockr.YockState) int {
	err := utils.WriteFile(l.CheckString(1), []byte(l.CheckString(2)))
	l.PushError(err)
	return 1
}

// @param path string
//
// @return bool
func is_exist(l *yockr.YockState) int {
	ok := utils.IsExist(l.CheckString(1))
	l.PushBool(ok)
	return 1
}

// @param title string
//
// @param rows string[][]
func printf(l *yockr.YockState) int {
	title := []string{}
	rows := [][]string{}
	l.CheckTable(1).ForEach(func(idx, el lua.LValue) {
		title = append(title, el.String())
	})
	l.CheckTable(2).ForEach(func(ri, row lua.LValue) {
		tmp := []string{}
		row.(*lua.LTable).ForEach(func(fi, field lua.LValue) {
			tmp = append(tmp, field.String())
		})
		rows = append(rows, tmp)
	})
	utils.Prinf(utils.PrintfOpt{MaxLen: 30}, title, rows)
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
func ioPathf(l *yockr.YockState) int {
	l.PushString(util.Pathf(l.CheckString(1)))
	return 1
}

/*
* @param file string
* @return string, err
 */
func ioReadFile(l *yockr.YockState) int {
	out, err := utils.ReadStraemFromFile(l.CheckString(1))
	l.PushString(string(out)).PushError(err)
	return 2
}
