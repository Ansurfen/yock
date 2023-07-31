// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package process

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFSNotify(t *testing.T) {
	fsn := NewOSNotify()
	pwd, _ := os.Getwd()
	fsn.AddFunc([]string{filepath.Join(pwd, "testdata")}, func() {
		fmt.Println("ping")
	})
	fsn.Listen()
}
