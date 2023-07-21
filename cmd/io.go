// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ansurfen/yock/util"
)

// RmOpt indicates configuration of rm
type RmOpt struct {
	Safe bool
	// Pattern delete file to be matched
	Pattern string
}

func Rm(opt RmOpt, target string) error {
	if len(opt.Pattern) != 0 {
		err := filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				matched, _ := regexp.MatchString(opt.Pattern, info.Name())
				if matched {
					err := os.Remove(path)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		return err
	} else {
		if opt.Safe {
			return os.Remove(target)
		} else {
			return os.RemoveAll(target)
		}
	}
}

// CpOpt indicates configuration of cp
type CpOpt struct {
	Recurse bool
	Force   bool
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
			if strings.HasSuffix(src, "*") || strings.HasSuffix(dst, "*") {
				term.AppendCmds("-r")
			}
		}
	}
	_, err := term.Exec(&ExecOpt{
		Quiet: true,
		Redirect: true,
	})
	return err
}

// MvOpt indicates configuration of mv
type MvOpt struct{}

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
		Quiet: true,
	})
	return err
}

func Rename(old, new string) {
	os.Rename(old, new)
}
