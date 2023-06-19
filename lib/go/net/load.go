// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/lib/go/net/http"
)

func LoadNet(yocks yocki.YockScheduler) {
	http.LoadNetHttp(yocks)
}
