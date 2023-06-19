// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/beevik/etree"
)

func LoadXML(yocks yocki.YockScheduler) {
	yocks.RegYockFn(yocki.YockFuns{
		"xml": xmlXML,
	})
}

func xmlXML(l *yockr.YockState) int {
	l.Pusha(etree.NewDocument())
	return 1
}
