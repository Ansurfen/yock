// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ansurfen/yock/util"
)

type LsofInfo struct {
	Protocal string `json:"protocal"`
	Local    string `json:"local"`
	Foreign  string `json:"foregin"`
	State    string `json:"state"`
	Pid      int32  `json:"pid"`
}

func Lsof() ([]LsofInfo, error) {
	switch util.CurPlatform.OS {
	case "windows":
		str, err := Exec(ExecOpt{Quiet: true}, "netstat -ano")
		if err != nil {
			return nil, fmt.Errorf("%s%s", str, err)
		}
		return lsofWindows(str), nil
	case "linux":
		str, err := Exec(ExecOpt{Quiet: true}, "netstat -tlnp")
		if err != nil {
			return nil, fmt.Errorf("%s%s", str, err)
		}
		return lsofLinux(str), nil
	case "darwin":
	}
	return nil, fmt.Errorf("no support")
}

func lsofWindows(str string) (infos []LsofInfo) {
	level := 0
	re := regexp.MustCompile(`\s+`)
	util.ReadLineFromString(str, func(s string) string {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			return ""
		}
		if level < 2 {
			level++
			return ""
		}
		s = re.ReplaceAllString(s, " ")
		res := strings.SplitN(s, " ", 5)
		if len(res) == 5 {
			pid, _ := strconv.Atoi(res[4])
			infos = append(infos, LsofInfo{
				Protocal: res[0],
				Local:    res[1],
				Foreign:  res[2],
				State:    res[3],
				Pid:      int32(pid),
			})
		}
		return ""
	})
	return
}

func lsofLinux(str string) (infos []LsofInfo) {
	flag := false
	re := regexp.MustCompile(`\s+`)
	util.ReadLineFromString(str, func(s string) string {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			return ""
		}
		if strings.HasPrefix(s, "Proto") {
			flag = true
			return ""
		}
		if flag {
			s = re.ReplaceAllString(s, " ")
			res := strings.SplitN(s, " ", 7)
			if len(res) == 7 {
				if strings.Contains(res[6], "/") {
					res[6] = strings.Split(res[6], "/")[0]
				}
				pid, _ := strconv.Atoi(res[6])
				infos = append(infos, LsofInfo{
					Protocal: res[0],
					Local:    res[3],
					Foreign:  res[4],
					State:    res[5],
					Pid:      int32(pid),
				})
			}
		}
		return ""
	})
	return
}
