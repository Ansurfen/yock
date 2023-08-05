// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"os"
	"strings"

	"github.com/ansurfen/yock/util"
)

const (
	exportOverrideWindows = `Set wshShell = CreateObject("WScript.Shell")
Set env = wshShell.Environment("User")
env("%s") = "%s"`

	exportExpandWindows = `Set wshShell = CreateObject("WScript.Shell")
Set env = wshShell.Environment("User")
old = env("%s")
newVar = "%s"
if InStr(1, old, newVar, vbTextCompare) = 0 Then
	old = newVar & ";" & old
End If
env("%s") = old`

	exportOverridePosix = `#!/bin/bash
sed -i '/^export %s=%s/d' ~/.bashrc
echo 'export %s=%s' >> ~/.bashrc
. ~/.bashrc`
	exportExpandPosix = `#!/bin/bash
sed -i '/^export %s=%s/d' ~/.bashrc
echo 'export %s=$%s:%s' >> ~/.bashrc
. ~/.bashrc`

	unsetWindows = `
Set objShell = CreateObject("WScript.Shell")
objShell.Environment("User").Remove("%s")`
	unsetExpandWindows = `
Set wshShell = CreateObject("WScript.Shell")
Set env = wshShell.Environment("{{ .Target }}")
old = env("{{ .Key }}")
newVar = "{{ .Value }}"
if InStr(1, old, newVar, vbTextCompare) = 1 Then
	old = Replace(old, newVar+";", "")
End If
env("{{ .Key }}") = old`
	unsetPosix = `#!/bin/bash
sed -i '/^export %s=/d' ~/.bashrc
. ~/.bashrc`
	UnsetExpandPosix = `#!/bin/bash
sed -i '/^export %s=%s/d' ~/.bashrc
. ~/.bashrc`
)

type ExportOpt struct {
	Expand bool
	System bool
}

func Export(opt ExportOpt, k, v string) error {
	script := ""
	if strings.ToUpper(k) == "PATH" {
		if util.CurPlatform.OS == "windows" {
			k = "Path"
		} else {
			k = "PATH"
		}
	}
	if util.CurPlatform.OS == "windows" {
		if opt.Expand {
			script = fmt.Sprintf(exportExpandWindows, k, v, k)
		} else {
			script = fmt.Sprintf(exportOverrideWindows, k, v)
		}
	} else {
		if opt.Expand {
			script = fmt.Sprintf(exportExpandPosix, k, v, k, k, v)
		} else {
			script = fmt.Sprintf(exportOverridePosix, k, v, k, v)
		}
	}
	_, err := OnceScript(script)
	return err
}

type UnsetOpt struct {
	Expand bool
	System bool
}

func Unset(opt UnsetOpt, k, v string) error {
	script := ""
	target := "User"
	if opt.System {
		target = "System"
	}
	if strings.ToUpper(k) == "PATH" {
		if util.CurPlatform.OS == "windows" {
			k = "Path"
		} else {
			k = "PATH"
		}
	}
	if util.CurPlatform.OS == "windows" {
		if opt.Expand {
			script, _ = util.NewTemplate().OnceParse(unsetExpandWindows, map[string]string{
				"Target": target,
				"Key":    k,
				"Value":  v,
			})
		} else {
			script = fmt.Sprintf(unsetWindows, k)
		}
	} else {
		if opt.Expand {
			script = fmt.Sprintf(UnsetExpandPosix, k, v)
		} else {
			script = fmt.Sprintf(unsetPosix, k)
		}
	}
	_, err := OnceScript(script)
	return err
}

func ExportL(k, v string) error {
	return os.Setenv(k, v)
}

func Environ(k string) string {
	return os.Getenv(k)
}

func UnsetL(k string) error {
	return os.Unsetenv(k)
}
