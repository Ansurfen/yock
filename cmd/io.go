package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type RmOpt struct {
	Safe    bool
	Debug   bool
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
