// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reflect

import (
	"reflect"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadReflect(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("reflect")
	lib.SetField(map[string]any{
		"TypeOf":  reflect.TypeOf,
		"ValueOf": reflect.ValueOf,
	})
}
