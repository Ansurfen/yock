// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockr

import (
	"testing"

	"github.com/ansurfen/yock/util/test"
	lua "github.com/yuin/gopher-lua"
)

func TestTable(t *testing.T) {
	l := lua.NewState()
	tbl := &lua.LTable{}

	old := UpgradeTable(tbl)
	test.Assert(tbl == old.Value())

	new := old.Clone(l)
	test.Assert(old != new)

	old.SetString("first", "1")
	old.SetBool("second", true)
	old.SetNil("thrid")
	old.SetInt("fourth", 4)
	old.SetTable("fifth", new)
	old.SetLTable("sixth", tbl)
	old.SetField(l, "seven", "7")
	old.SetFields(l, map[string]any{})
	// old.SetDo("", func(ys yocki.YockState) lua.LValue {

	// })

	_, ok := old.GetString("first")
	test.Assert(ok)

	_, ok = old.GetBool("second")
	test.Assert(ok)

	_, ok = old.GetFloat("fourth")
	test.Assert(ok)

	new_t, ok := old.GetTable("fifth")
	test.Assert(ok)
	test.Assert(new_t.Value() == new.Value())

	tt, ok := old.GetTable("sixth")
	if ok {
		test.Assert(tt.Value() == tbl)
		test.Assert(old.MustGetTable("sixth").Value() == tbl)
	}

	_, ok = old.GetString("seven")
	test.Assert(ok)
}
