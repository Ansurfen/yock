// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	filter := NewBloomFilter(10000, 5)
	filter.Set("abc")
	fmt.Println(filter.Check("abc"), filter.Check("ab"))
}
