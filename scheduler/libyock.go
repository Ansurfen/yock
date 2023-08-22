// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import liby "github.com/ansurfen/yock/lib/go/yock"

var libyock = []loader{
	liby.LoadCheck,
	liby.LoadGoroutine,
	liby.LoadXML,
	liby.LoadType,
	liby.LoadGNU,
	liby.LoadJSON,
	liby.LoadWatch,
	liby.LoadMisc,
	liby.LoadRandom,
	liby.LoadCrypto,
	// liby.LoadTea,
	liby.LoadBit,
	loadEnv,
	loadTask,
	liby.LoadGin,
}
