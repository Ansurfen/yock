// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import lua "github.com/yuin/gopher-lua"

type yocksDB struct {
	datas map[string]lua.LValue
}

func newYocksDB() *yocksDB {
	return &yocksDB{
		datas: make(map[string]lua.LValue),
	}
}

func (db *yocksDB) Get(k string) lua.LValue {
	return db.datas[k]
}

func (db *yocksDB) Put(k string, v lua.LValue) {
	db.datas[k] = v
}
