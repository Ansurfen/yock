// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/ansurfen/yock/util"
)

// RmOpt indicates configuration of rm
type RmOpt struct {
	Safe bool
	// Debug prints output when it's true
	Debug bool
	// Caller is used to mark parent caller of HTTP function
	//
	// It'll printed on console when debug is true
	Caller string
	// Strict will exit at once when error occur
	Strict bool
	// Pattern delete file to be matched
	Pattern string

	Info func(path string)

	Error func(err error) error
}

func Rm(opt RmOpt, targets []string) error {
	if len(opt.Pattern) != 0 {
		for _, t := range targets {
			filepath.Walk(t, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					matched, _ := regexp.MatchString(opt.Pattern, info.Name())
					if matched {
						if opt.Info != nil {
							opt.Info(path)
						}
						err := os.Remove(path)
						if err != nil && opt.Error != nil {
							if err = opt.Error(err); err != nil {
								return err
							}
						}
					}
				}
				return nil
			})
		}
	} else {
		if opt.Safe {
			for _, t := range targets {
				if opt.Info != nil {
					opt.Info(t)
				}
				if err := os.Remove(t); err != nil && opt.Error != nil {
					opt.Error(err)
				}
			}
		} else {
			for _, t := range targets {
				if opt.Info != nil {
					opt.Info(t)
				}
				if err := os.RemoveAll(t); err != nil && opt.Error != nil {
					opt.Error(err)
				}
			}
		}
	}
	return nil
}

// CpOpt indicates configuration of cp
type CpOpt struct {
	Recurse bool
	// Debug prints output when it's true
	Debug bool
	// Caller is used to mark parent caller of HTTP function
	//
	// It'll printed on console when debug is true
	Caller string
	Force  bool

	Info func(name, args string)
}

func Cp(opt CpOpt, src, dst string) error {
	var term *Terminal
	switch util.CurPlatform.OS {
	case "windows":
		term = WindowsTerm()
		if term.kind == TermPowershell {
			if opt.Recurse {
				term.SetCmds("cp", "-r", src, dst)
			} else {
				term.SetCmds("cp", src, dst)
			}
			if opt.Force {
				term.AppendCmds("-Force")
			}
		} else {
			term.SetCmds("copy", src, dst)
		}
	default:
		term = PosixTerm()
		if opt.Recurse {
			term.SetCmds("cp", "-r", src, dst)
		} else {
			term.SetCmds("cp", src, dst)
		}
	}
	_, err := term.Exec(&ExecOpt{
		Debug:  opt.Debug,
		Caller: opt.Caller,
		Quiet:  true,
		Info:   opt.Info,
	})
	return err
}

// MvOpt indicates configuration of mv
type MvOpt struct {
	// Debug prints output when it's true
	Debug bool
	// Caller is used to mark parent caller of HTTP function
	//
	// It'll printed on console when debug is true
	Caller string

	Info func(name, args string)
}

func Mv(opt MvOpt, src, dst string) error {
	var term *Terminal
	switch util.CurPlatform.OS {
	case "windows":
		term = WindowsTerm()
		if term.kind == TermPowershell {
			term.SetCmds("mv", src, dst)
		} else {
			term.SetCmds("move", src, dst)
		}
	default:
		term = PosixTerm("mv", src, dst)
	}
	_, err := term.Exec(&ExecOpt{
		Debug:  opt.Debug,
		Caller: opt.Caller,
		Quiet:  true,
		Info:   opt.Info,
	})
	return err
}

func Rename(old, new string) {
	os.Rename(old, new)
}
