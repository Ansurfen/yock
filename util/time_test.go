// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestTimestamp(t *testing.T) {
	fmt.Println(FmtTimestamp(NowTimestamp()))
}

func TestTime(t *testing.T) {
	// TODO tm := ParseTime("02/01/2006 07:00:06", "2020")
	// tm.AddMinute(70)
	tm := ParseTime("02/01/2006", "01/01/2022")
	tm2 := ParseTime("02/01/2006", "02/01/2022")
	fmt.Println(tm.Format("02/01/2006"))
	fmt.Println(tm.Diff(tm2))
}

func TestDuration(t *testing.T) {
	testset := []string{
		"7h0m50s",
		"18h30m",
		"7",
		"23:00:59",
	}
	for _, test := range testset {
		t := ParseDuration(test)
		t.AddMinute(10)
		fmt.Println(t.Format(":"))
	}
}
