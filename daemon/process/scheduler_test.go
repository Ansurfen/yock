// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package process

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScheduler(t *testing.T) {
	s := NewScheduler()
	s.CreateCronTask("*/1 * * * *", "rmdir tmp")
	pwd, _ := os.Getwd()
	s.CreateFSListenTask([]string{filepath.Join(pwd, "testdata")}, "mkdir tmp")
	s.Run()
}
