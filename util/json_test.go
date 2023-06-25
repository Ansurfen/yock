// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestJsonStr(t *testing.T) {
	fmt.Println(JsonStr(NewJsonObject(map[string]JsonValue{
		"a": JsonTrue,
		"b": NewJsonObject(map[string]JsonValue{
			"b1": JsonNil,
			"b2": NewJsonArray([]JsonValue{JsonFalse, NewJsonObject(map[string]JsonValue{
				"b21": NewJsonArray([]JsonValue{}),
			})}),
		}),
	})))
}
