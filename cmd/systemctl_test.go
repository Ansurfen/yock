// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"testing"

	yocke "github.com/ansurfen/yock/env"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/test"
)

func TestSystemCtlCreate4Windows(t *testing.T) {
	fmt.Println(SystemCtlCreate("TestService", SCCreateOpt{
		Service: scCreateOptService{
			ExecStart: "test.exe -p 10000",
		},
	}))
}

func TestSytemCtlCreate4Darwin(t *testing.T) {
	systemCtlCreateDarwin("yockc.TestService", SCCreateOpt{
		Unit: scCreateOptUnit{
			Description: "TestService",
		},
		Service: scCreateOptService{
			ExecStart: "echo 'Hello World'",
		},
	})
}

func TestSystemCtlStatus(t *testing.T) {
	SystemCtlStatus(SystemCtlStatusOpt{})
}

func TestSystemCtlStatus4Linux(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/systemctl/linux",
	}).FRange(func(s *test.TestSet, file, data string) error {
		infos, err := systemCtlStatusLinux(data)
		s.Assert(err)
		file = util.Filename(file)
		if file == "a" {
			s.AssertEqual(len(infos), 1)
		} else if file == "b" {
			s.AssertNil(infos)
		}
		return nil
	})
}

func TestSystemCtlStatus4Windows(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/systemctl/windows",
	}).Range(func(data string) error {
		fmt.Println(systemCtlStatusWindows(data))
		return nil
	})
}

func TestSystemCtlStatus4Darwin(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/systemctl/darwin",
	}).Range(func(data string) error {
		fmt.Println(systemCtlStatusDarwin(data))
		return nil
	})
}

func TestServiceFile4Linux(t *testing.T) {
	opt := SCCreateOpt{
		Unit: scCreateOptUnit{
			Description: "TestService",
		},
		Service: scCreateOptService{
			Type:       "forking",
			ExecStart:  "./test",
			RestartSec: 10,
		},
	}
	fmt.Println(opt.String())
}

func TestServiceFile4Darwin(t *testing.T) {
	fp, err := yocke.CreatePlistFile("TestService.plist")
	if err != nil {
		panic(err)
	}
	fp.Set(yocke.MetaMap{
		"Description": "TestService",
		"KeepAlive": yocke.MetaMap{
			"SuccessfulExit": false,
		},
		"Label":             "yock.TestService",
		"ProgramArguments":  yocke.MetaArr{},
		"RunAtLoad":         false,
		"WorkingDirectory":  "/usr/local/var",
		"StandardErrorPath": "/usr/local/var/yock.log",
		"StandardOutPath":   "/usr/local/var/yock.log",
	})
	err = fp.Write()
	if err != nil {
		panic(err)
	}
}
