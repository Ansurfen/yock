// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"strconv"
	"strings"
)

type CheckedVersion struct {
	lower struct {
		subVersions []int
		mode        int
	}
	upper struct {
		subVersions []int
		mode        int
	}
	pass bool
}

func NewCheckedVersion(version string) *CheckedVersion {
	if version == "-" {
		return &CheckedVersion{pass: true}
	}
	ver := &CheckedVersion{
		lower: struct {
			subVersions []int
			mode        int
		}{
			subVersions: make([]int, 0),
			mode:        NONE,
		},
		upper: struct {
			subVersions []int
			mode        int
		}{
			subVersions: make([]int, 0),
			mode:        NONE,
		},
		pass: false,
	}
	if strings.Contains(version, ",") {
		before, after, ok := strings.Cut(version, ",")
		before = strings.TrimSpace(before)
		after = strings.TrimSpace(after)
		if ok {
			ch := before[len(before)-1]
			var val string
			if ch == '+' || ch == '-' {
				val = before[:len(before)-1]
			} else {
				val = before
			}
			vals := strings.Split(val, ".")
			for _, v := range vals {
				i, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				ver.lower.subVersions = append(ver.lower.subVersions, i)
			}
			switch ch {
			case '+':
				ver.lower.mode = UPPER
			case '-':
				ver.lower.mode = DOWN
			default:
				ver.lower.mode = NONE
			}
			ch = after[len(after)-1]
			if ch == '+' || ch == '-' {
				val = after[:len(after)-1]
			} else {
				val = after
			}
			vals = strings.Split(val, ".")
			for _, v := range vals {
				i, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				ver.upper.subVersions = append(ver.upper.subVersions, i)
			}
			switch ch {
			case '+':
				ver.upper.mode = UPPER
			case '-':
				ver.upper.mode = DOWN
			default:
				ver.upper.mode = NONE
			}
		}
	} else {
		ch := version[len(version)-1]
		var val string
		if ch == '+' || ch == '-' {
			val = version[:len(version)-1]
		} else {
			val = version
		}
		vals := strings.Split(val, ".")
		for _, v := range vals {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			ver.lower.subVersions = append(ver.lower.subVersions, i)
		}
		switch ch {
		case '+':
			ver.lower.mode = UPPER
		case '-':
			ver.lower.mode = DOWN
		default:
			ver.lower.mode = NONE
		}
	}
	return ver
}

const (
	NONE = iota
	UPPER
	DOWN
)

func (want *CheckedVersion) Compare(got *CheckedVersion) bool {
	if want.pass {
		return true
	}
	if len(want.lower.subVersions) > 0 {
		if !compare(want.lower.subVersions, got.lower.subVersions, want.lower.mode) {
			return false
		}
	}
	if len(want.upper.subVersions) > 0 {
		if !compare(want.upper.subVersions, got.lower.subVersions, want.upper.mode) {
			return false
		}
	}
	return true
}

func compare(targetVersion, curVersion []int, firstAction int) bool {
	firLen, secLen := len(targetVersion), len(curVersion)
	count := 0
	minVal := 0
	for i := 0; i < secLen-firLen; i++ {
		targetVersion = append(targetVersion, 0)
	}
	firLen, secLen = len(targetVersion), len(curVersion)
	if firLen > secLen {
		minVal = secLen
	} else {
		minVal = firLen
	}
	for count < minVal {
		if targetVersion[count] < curVersion[count] {
			if firstAction == UPPER {
				return true
			} else {
				return false
			}
		} else if targetVersion[count] > curVersion[count] {
			if firstAction == DOWN {
				return true
			} else {
				return false
			}
		}
		count++
	}
	if firLen == secLen && count == minVal {
		return true
	}
	return false
}
