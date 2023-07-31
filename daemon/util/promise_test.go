// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "testing"

func TestPromise(t *testing.T) {
	p := NewPromise()
	p.Store(10, "")
	p.LoadWithTimeout(10, 10)
}
