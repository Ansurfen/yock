// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package maillib

import (
	"net/mail"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadMail(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("mail")
	lib.SetField(map[string]any{
		// functions
		"ReadMessage":      mail.ReadMessage,
		"ParseAddress":     mail.ParseAddress,
		"ParseAddressList": mail.ParseAddressList,
		"ParseDate":        mail.ParseDate,
		// constants
		// variable
		"ErrHeaderNotPresent": mail.ErrHeaderNotPresent,
	})
}
