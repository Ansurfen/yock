// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterp(t *testing.T) {
	interp := New()
	fmt.Println(reflect.TypeOf(interp))
	interp2 := New(OptionEnableInterpPool())
	fmt.Println(reflect.TypeOf(interp2))
}
