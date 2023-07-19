// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unsafe"
)

var CurPlatform Platform

type Platform struct {
	OS     string
	Ver    string
	Arch   string
	Locale string
	Lang   string
}

func (pf Platform) Exf() string {
	switch pf.OS {
	case "windows":
		return ".exe"
	default:
		return ""
	}
}

func (pf Platform) Script() string {
	switch pf.OS {
	case "windows":
		return ".bat"
	default:
		return ".sh"
	}
}

func (pf Platform) Zip() string {
	switch pf.OS {
	case "windows":
		return ".zip"
	default:
		return ".tar.gz"
	}
}

func init() {
	CurPlatform = Platform{
		Arch:   runtime.GOARCH,
		OS:     runtime.GOOS,
		Lang:   "en",
		Locale: "US",
	}
	switch CurPlatform.OS {
	case "windows":
		CurPlatform.Ver = windowsVersion()
		// cmd := wmic os get MUILanguages
		cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "-")
			CurPlatform.Lang = langLoc[0]
			CurPlatform.Locale = langLoc[1]
		}
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

func windowsVersion() string {
	plugin, err := NewPlugin("ntdll.dll")
	if err != nil {
		return ""
	}
	procRtlGetVersion, err := plugin.Func("RtlGetVersion")
	if err != nil {
		return ""
	}
	type RTL_OSVERSIONINFOEX struct {
		dwOSVersionInfoSize uint32
		dwMajorVersion      uint32
		dwMinorVersion      uint32
		dwBuildNumber       uint32
		dwPlatformId        uint32
		szCSDVersion        [128]uint16
	}
	var info RTL_OSVERSIONINFOEX
	info.dwOSVersionInfoSize = uint32(unsafe.Sizeof(info))
	r, _ := procRtlGetVersion.Call(uintptr(unsafe.Pointer(&info)))
	if r != 0 {
		return ""
	}
	switch {
	case info.dwMajorVersion == 5 && info.dwMinorVersion == 1:
		return "XP"
	case info.dwMajorVersion == 6 && info.dwMinorVersion == 0:
		return "Vista"
	case info.dwMajorVersion == 6 && info.dwMinorVersion == 1:
		return "7"
	case info.dwMajorVersion == 6 && info.dwMinorVersion == 2 && info.dwBuildNumber == 9200:
		return "8"
	case info.dwMajorVersion == 6 && info.dwMinorVersion == 3 && info.dwBuildNumber == 9600:
		return "8.1"
	case info.dwMajorVersion == 10 && info.dwMinorVersion == 0 && info.dwBuildNumber >= 22000:
		return "11"
	case info.dwMajorVersion == 10 && info.dwMinorVersion == 0:
		return "10"
	default:
		return ""
	}
}
