// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package textprotolib

import (
	"net/textproto"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadTextproto(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("textproto")
	lib.SetField(map[string]any{
		// functions
		"TrimBytes":              textproto.TrimBytes,
		"NewConn":                textproto.NewConn,
		"Dial":                   textproto.Dial,
		"TrimString":             textproto.TrimString,
		"NewWriter":              textproto.NewWriter,
		"NewReader":              textproto.NewReader,
		"CanonicalMIMEHeaderKey": textproto.CanonicalMIMEHeaderKey,
		// constants
		// variable
	})
}
