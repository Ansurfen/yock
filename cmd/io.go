package cmd

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
						if opt.Debug {
							if err != nil {
								fmt.Printf("delete %s, err: %s\n", path, err)
							} else {
								fmt.Printf("Deleting file %s\n", path)
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
					util.YchoWarn(opt.Caller, err.Error())
				}
			}
		} else {
			for _, t := range targets {
				if err := os.RemoveAll(t); err != nil && opt.Debug {
					util.YchoWarn(opt.Caller, err.Error())
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
