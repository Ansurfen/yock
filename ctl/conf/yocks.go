// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

type yockScheduler struct {
	Goroutine  yockGoroutine `yaml:"goroutine"`
	InterpPool bool          `yaml:"interPool"`
	MaxInterp  int           `yaml:"maxInterp"`
}

type yockGoroutine struct {
	MaxGoroutine int64 `yaml:"maxGoroutine"`
	MaxWaitRound int64 `yaml:"maxWaitRound"`
	RoundStep    int   `yaml:"roundStep"`
}
