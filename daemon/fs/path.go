// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import (
	"path/filepath"
	"strings"
)

func FormatPath(path string) string {
	path = filepath.Join(path)
	if filepath.Separator == '\\' {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	if path == "." {
		return "%"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return strings.ReplaceAll(path, "/", "%")
}

func ResolvePath(path string) string {
	ret := ""
	pathLen := len(path)
	for i := 0; i < pathLen; i++ {
		ch := path[i]
		if ch == '%' {
			ret += "/"
			for {
				if i+1 < pathLen && path[i+1] == '%' {
					ret += "%"
					i++
				} else {
					break
				}
			}
		} else {
			ret += string(ch)
		}
	}
	return ret
}

func SplitPath(fullpath string) (volume, path string) {
	before, after, ok := strings.Cut(fullpath, ":")
	if ok {
		volume = before
		path = after
	}
	return
}
