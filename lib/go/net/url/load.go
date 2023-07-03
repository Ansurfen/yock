// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package urllib

import (
	"net/url"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadUrl(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("url")
	lib.SetField(map[string]any{
		// functions
		"QueryEscape":     url.QueryEscape,
		"JoinPath":        url.JoinPath,
		"ParseQuery":      url.ParseQuery,
		"PathUnescape":    url.PathUnescape,
		"PathEscape":      url.PathEscape,
		"User":            url.User,
		"Parse":           url.Parse,
		"QueryUnescape":   url.QueryUnescape,
		"ParseRequestURI": url.ParseRequestURI,
		"UserPassword":    url.UserPassword,
		// constants
		// variable
	})
}
