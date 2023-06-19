// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	vec := NewVector[int](10)
	vec.PushBack(10)
	vec.PushFront(9)
	fmt.Println(vec.Front(), vec.Back())
	vec.PushFront(8)
	vec.PushFront(7)
	vec.Remove(vec.Find(1))
	fmt.Println(vec.Find(1))
}
