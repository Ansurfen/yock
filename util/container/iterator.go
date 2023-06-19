// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

type Iterator[T any] interface {
	Front() T
	Back() T
}
