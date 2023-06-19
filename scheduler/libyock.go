// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import liby "github.com/ansurfen/yock/lib/yock"

var libyock = []loader{
	liby.LoadCheck,
	liby.LoadGoroutine,
	liby.LoadXML,
	liby.LoadTemplate,
	liby.LoadType,
	liby.LoadGNU,
	liby.LoadJSON,
	liby.LoadWatch,
	liby.LoadSSH,
	liby.LoadMisc,
	loadEnv,
	loadTask,
}
