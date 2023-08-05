// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"testing"

	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/util/test"
)

const cronExprs = `0 * * * *
0 15 10 ? * * echo "Hello World"
0 15 10 * * ? ls -a
0 0 12 * * ?
0 0 10,14,16 * * ?
0 0/30 9-17 * * ?
0 * 14 * * ?
0 0-5 14 * * ?
0 0/5 14 * * ?
0 0/5 14,18 * * ?
0 0 12 ? * WED
0 15 10 15 * ?
0 15 10 L * ?
0 15 10 ? * 6L
0 15 10 ? * 6#3
0 10,44 14 ? 3 WED
0 15 10 ? * * 2022
0 15 10 ? * * *
0 0/5 14,18 * * ? 2022
0 15 10 ? * 6#3 2022,2023
0 0/30 9-17 * * ? 2022-2025
0 10,44 14 ? 3 WED 2022/2`

func TestCronParse(t *testing.T) {
	util.ReadLineFromString(cronExprs, func(s string) string {
		expr, cmd := CronParse(s)
		fmt.Println(expr.String(), cmd)
		return ""
	})
}

func TestCronToSchTasks(t *testing.T) {
	util.ReadLineFromString(cronExprs, func(s string) string {
		expr, _ := CronParse(s)
		fmt.Println(expr)
		fmt.Println(expr.toSchtasks())
		return ""
	})
}

func TestCrontabList(t *testing.T) {
	_, err := CrontabList("MyTest")
	test.Assert(err == nil)
}

func TestCrontabAdd(t *testing.T) {
	err := CrontabAdd("* * * * * echo a", "MyTest")
	test.Assert(err == nil)
}

func TestCrontabDel(t *testing.T) {
	err := CrontabDel("* * * * ? echo a", "MyTest")
	test.Assert(err == nil)
}

func TestCrontabListPosix(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/crontab/posix",
	}).Range(func(data string) error {
		fmt.Println(crontabListPosix(data))
		return nil
	})
}

func TestCrontabListWindowsV1(t *testing.T) {
	test.Batch(test.BatchOpt{
		Path: "./testdata/crontab/windows",
	}).Range(func(data string) error {
		fmt.Println(crontabListWindowsV1(data))
		return nil
	})
}

func TestCrontabListWindowsV2(t *testing.T) {
	if util.CurPlatform.OS == "windows" {
		str, err := Exec(ExecOpt{Quiet: true}, "schtasks /query /XML")
		if err != nil {
			panic(err)
		}
		fmt.Println(crontabListWindowsV2(str, true))
		str, err = Exec(ExecOpt{Quiet: true}, "schtasks /query /tn MyTest /XML")
		if err != nil {
			panic(err)
		}
		fmt.Println(crontabListWindowsV2(str, false))
	}
}

func TestCronAnchor(t *testing.T) {
	anchors := []cronAnchor{
		{nums: []cronTime{30, 1}, attr: A_STEP},
		{nums: []cronTime{14, 18}},
		{nums: []cronTime{9, 17}, attr: A_ROUND},
		{attr: A_INVALID},
		{attr: A_LAST},
		{nums: []cronTime{6, 1}, attr: A_LAST},
		{nums: []cronTime{6}, attr: A_LAST},
		{nums: []cronTime{6, 3}, attr: A_NO},
		{},
	}
	for _, anchor := range anchors {
		fmt.Println(anchor)
	}
}

func TestCronExpr(t *testing.T) {
	exprs := []cronExpr{
		{hour: cronAnchor{nums: []cronTime{1}, attr: A_STEP}},
		{
			min:  cronAnchor{nums: []cronTime{5}, attr: A_STEP},
			hour: cronAnchor{nums: []cronTime{2, 1, 10}, attr: A_STEP | A_ROUND},
			// hour: cronAnchor{nums: []cronTime{12, 14}, attr: A_ROUND},
			day: cronAnchor{nums: []cronTime{6}, attr: A_ENUM},
			mon: cronAnchor{nums: []cronTime{6}, attr: A_ENUM},
			// mon:  cronAnchor{nums: []cronTime{1, 2}, attr: A_ROUND},
			week: cronAnchor{nums: []cronTime{6, 3}, attr: A_NO},
			year: cronAnchor{nums: []cronTime{2006}},
		},
		{
			min:  cronAnchor{nums: []cronTime{15}, attr: A_ENUM},
			hour: cronAnchor{nums: []cronTime{10}, attr: A_ENUM},
		},
	}
	want := []string{
		"* * 0/1 * * * *",
		"* 0/5 1-10/2 6 6 6#3 2006",
		"* 15 10 * * * *",
	}
	schtask_wants := []int{1, 0, 1}
	for i, expr := range exprs {
		test.Assert(expr.String() == want[i])
		test.Assert(len(expr.toSchtasks()) == schtask_wants[i])
	}
}
