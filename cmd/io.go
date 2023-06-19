// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ansurfen/cushion/utils"
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
						err := os.Remove(path)
						if err != nil {
							if opt.Debug {
								util.Ycho.Warn(err.Error())
							}
							if opt.Strict {
								return err
							}
						} else {
							if opt.Debug {
								util.Ycho.Info(fmt.Sprintf("delete %s", path))
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
				if err := os.Remove(t); err != nil && opt.Debug {
					util.Ycho.Warn(fmt.Sprintf("%s\t%s", opt.Caller, err.Error()))
				}
			}
		} else {
			for _, t := range targets {
				if err := os.RemoveAll(t); err != nil && opt.Debug {
					util.Ycho.Warn(fmt.Sprintf("%s\t%s", opt.Caller, err.Error()))
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

	Strict bool
}

func Cp(opt CpOpt, src, dst string) error {
	var term *Terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = WindowsTerm()
		if term.this == TermPowershell {
			if opt.Recurse {
				term.SetCmds("cp", "-r", src, dst)
			} else {
				term.SetCmds("cp", src, dst)
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
}

func Mv(opt MvOpt, src, dst string) error {
	var term *Terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = WindowsTerm()
		if term.this == TermPowershell {
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
	})
	return err
}
