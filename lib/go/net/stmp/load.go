// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package smtplib

import (
	"net/smtp"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadSmtp(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("smtp")
	lib.SetField(map[string]any{
		// functions
		"PlainAuth":   smtp.PlainAuth,
		"CRAMMD5Auth": smtp.CRAMMD5Auth,
		"Dial":        smtp.Dial,
		"NewClient":   smtp.NewClient,
		"SendMail":    smtp.SendMail,
		// constants
		// variable
	})
}
