// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ioutillib

import (
	"io/ioutil"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadIoutil(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("ioutil")
	lib.SetField(map[string]any{
		"ToString": func(str []byte) string {
			return string(str)
		},
		// functions
		"NopCloser": ioutil.NopCloser,
		"TempFile":  ioutil.TempFile,
		"TempDir":   ioutil.TempDir,
		"ReadAll":   ioutil.ReadAll,
		"ReadFile":  ioutil.ReadFile,
		"WriteFile": ioutil.WriteFile,
		"ReadDir":   ioutil.ReadDir,
		// constants
		// variable
		"Discard": ioutil.Discard,
	})
}
