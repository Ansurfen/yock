// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import (
	"fmt"
	"testing"

	"github.com/ansurfen/yock/util/test"
)

func TestFormatPath(t *testing.T) {
	testset := map[string]string{
		".":       "%",
		"/a":      "%a%",
		"./a":     "%a%",
		"/a/b":    "%a%b%",
		"./a/a/b": "%a%a%b%",
		".\\a/b":  "%a%b%",
	}
	for got, want := range testset {
		test.Assert(FormatPath(got) == want)
	}
}

func TestResolvePath(t *testing.T) {
	testset := map[string]string{
		"%":         "/",
		"%a":        "/a",
		"%a%":       "/a/",
		"%a%%b.txt": "/a/%b.txt",
		"a%%a%%b":   "a/%a/%b",
		"%%%%a%%b":  "/%%%a/%b",
	}
	for got, want := range testset {
		test.Assert(ResolvePath(got) == want)
	}
}

func TestSplitPath(t *testing.T) {
	fmt.Println(SplitPath("D:./a"))
}
