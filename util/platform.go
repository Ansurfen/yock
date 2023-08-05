// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

var CurPlatform Platform

type Platform struct {
	OS     string
	Ver    string
	Arch   string
	Locale string
	Lang   string
}

func (pf Platform) Exf() string {
	switch pf.OS {
	case "windows":
		return ".exe"
	default:
		return ""
	}
}

func (pf Platform) Script() string {
	switch pf.OS {
	case "windows":
		return ".bat"
	default:
		return ".sh"
	}
}

func (pf Platform) Zip() string {
	switch pf.OS {
	case "windows":
		return ".zip"
	default:
		return ".tar.gz"
	}
}
