// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestPlist(t *testing.T) {
	fp, err := CreatePlistFile("myplist2.plist")
	if err != nil {
		panic(err)
	}
	defer fp.Free()
	fp.v = []any{10, "6", map[string]any{
		"a": map[string]any{
			"b": 10,
		},
	}}
	fmt.Println(fp.Write())
	fp2, err := OpenPlistFile("myplist2.plist")
	if err != nil {
		panic(err)
	}
	if err := fp2.Backup(); err != nil {
		panic(err)
	}
	defer fp2.Free()
}

func TestCFValue(t *testing.T) {
	fp, err := CreatePlistFile("demo.plist")
	if err != nil {
		panic(err)
	}
	defer fp.Free()
	fp.Set(cfDictionary{
		"a": 6,
		"b": "basdaad",
		"c": cfDictionary{
			"c1": cfDictionary{
				"c11": cfArray{"6", cfDictionary{"6": 60}},
			},
			"c2": "10",
		},
		"d": cfArray{"10", 6},
	})

	fmt.Println(fp.GetArrByField("c.c1.c11"))

	arr := fp.GetArrByField("d").Copy()
	arr.SetByIdx(1, true)
	fp.SetByField("c.c1.c11", arr)
	fp.SetByField("d", CFArray{})
	fmt.Println(fp.GetDict())

	fp.Set(cfString("aa"))
	fmt.Println(fp.GetBaseValue())

	fp.Set(cfArray{10, "6", 7.8})
	fp.SetByIdx(3, true)
	fp.SetByIdx(2, true)
	fp.GetArr().Foreach(func(idx int, v CFValue) bool {
		fmt.Println(idx, v)
		return true
	})
	fp.Backup()
	fp.Write()
}
