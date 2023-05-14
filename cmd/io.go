package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ansurfen/cushion/utils"
)

// RmOpt indicates configuration of rm
type RmOpt struct {
	Safe bool
	// Debug prints output when it's true
	Debug bool
	// Caller is used to mark parent caller of HTTP function
	//
	// It'll printed on console when debug is true
	Caller  string
	Pattern string
}

func Rm(opt RmOpt, targets []string) int {
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
					fmt.Println(err)
				}
			}
		} else {
			for _, t := range targets {
				if err := os.RemoveAll(t); err != nil && opt.Debug {
					fmt.Println(err)
				}
			}
		}
	}
	return 0
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
	var term *terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = windowsTerm()
		if term.this == TermPowershell {
			if opt.Recurse {
				term.setCmds("cp", "-r", src, dst)
			} else {
				term.setCmds("cp", src, dst)
			}
		} else {
			term.setCmds("copy", src, dst)
		}
	default:
		term = posixTerm()
		if opt.Recurse {
			term.setCmds("cp", "-r", src, dst)
		} else {
			term.setCmds("cp", src, dst)
		}
	}
	return term.exec(&ExecOpt{
		Debug:  opt.Debug,
		Caller: opt.Caller,
		Quiet:  true,
	})
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
	var term *terminal
	switch utils.CurPlatform.OS {
	case "windows":
		term = windowsTerm()
		if term.this == TermPowershell {
			term.setCmds("mv", src, dst)
		} else {
			term.setCmds("move", src, dst)
		}
	default:
		term = posixTerm("mv", src, dst)
	}
	return term.exec(&ExecOpt{
		Debug:  opt.Debug,
		Caller: opt.Caller,
		Quiet:  true,
	})
}
