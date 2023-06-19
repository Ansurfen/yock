// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fmt

import (
	"fmt"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadFmt(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("fmt")
	lib.SetField(map[string]any{
		"Print":    fmt.Print,
		"Printf":   fmt.Printf,
		"Println":  fmt.Println,
		"Fprint":   fmt.Fprint,
		"Fprintf":  fmt.Fprintf,
		"Fprintln": fmt.Fprintln,
	})
}
