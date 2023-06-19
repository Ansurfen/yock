// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}
