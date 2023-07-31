// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestLoggerWatch(t *testing.T) {
	lw := New()
	files, err := os.ReadDir("testdata")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		err = lw.Parse(filepath.Join("testdata", file.Name()))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(lw.Find("*", "*", "*", "*", "*"))
}
