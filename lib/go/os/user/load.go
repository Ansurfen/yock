// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package userlib

import (
	"os/user"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadUser(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("user")
	lib.SetField(map[string]any{
		// functions
		"LookupGroupId": user.LookupGroupId,
		"Current":       user.Current,
		"Lookup":        user.Lookup,
		"LookupId":      user.LookupId,
		"LookupGroup":   user.LookupGroup,
		// constants
		// variable
	})
}
