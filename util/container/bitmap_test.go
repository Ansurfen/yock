// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"fmt"
	"testing"
)

func TestBitmap(t *testing.T) {
	bits := NewBitmap(10)
	bits.Set(10)
	fmt.Println(bits.Chcek(9), bits.Chcek(10))
}
