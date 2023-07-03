// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package signallib

import (
	"os/signal"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadSignal(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("signal")
	lib.SetField(map[string]any{
		// functions
		"Notify":        signal.Notify,
		"Reset":         signal.Reset,
		"Stop":          signal.Stop,
		"Ignore":        signal.Ignore,
		"Ignored":       signal.Ignored,
		"NotifyContext": signal.NotifyContext,
		// constants
		// variable
	})
}
