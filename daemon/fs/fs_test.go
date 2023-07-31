// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFileSystem(t *testing.T) {
	fs := NewFileSystem()
	fs.Put("./testdata", "D:/")
	fmt.Println(fs.List("D:/"))
	fs.Put("./testdata", "D:/b")
	fmt.Println(fs.List("D:/"))
	path, _ := filepath.Abs("./tmp")
	fmt.Println(fs.Get("D:/", path))
}
