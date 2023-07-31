// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package test

func Assert(ok bool, msg ...string) {
	if !ok {
		v := "fail asserted!"
		if len(msg) > 0 {
			v = msg[0]
		}
		panic(v)
	}
}
