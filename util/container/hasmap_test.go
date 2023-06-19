// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	m := NewLinkedHashMap[int, string](func(key, cap int) int { return key })
	m.Put(0, "abc")
	fmt.Println(m.Get(0))
	fmt.Println(m.Get(10))
}
