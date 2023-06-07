// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import "github.com/ansurfen/yock/daemon/client"

type DaemonOption struct {
	*client.DaemonOption
}

var Gopt = &DaemonOption{client.Gopt}
