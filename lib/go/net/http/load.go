// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadNetHttp(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("http")
	lib.SetField(map[string]any{
		"HandleFunc":        http.HandleFunc,
		"ListenAndServe":    http.ListenAndServe,
		"ListenAndServeTLS": http.ListenAndServeTLS,
		"Client":            netHttpClient,
	})
}

func netHttpClient() *http.Client {
	return &http.Client{}
}
