//go:build !windows
// +build !windows

// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func init() {
	CurPlatform = Platform{
		Arch:   runtime.GOARCH,
		OS:     runtime.GOOS,
		Lang:   "en",
		Locale: "US",
	}
	switch CurPlatform.OS {
	case "linux":
		envlang, ok := os.LookupEnv("LANG")
		if ok {
			langLocRaw := strings.Split(strings.TrimSpace(envlang), ".")[0]
			langLoc := strings.Split(langLocRaw, "_")
			if len(langLoc) >= 2 {
				CurPlatform.Lang = langLoc[0]
				CurPlatform.Locale = langLoc[1]
			}
		}
	case "darwin":
		cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "_")
			if len(langLoc) >= 2 {
				CurPlatform.Lang = langLoc[0]
				CurPlatform.Locale = langLoc[1]
			}
		}
	}
}
