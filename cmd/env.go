// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
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
sed -i '/^export %s/d' ~/.bashrc
echo 'export %s=%s' >> ~/.bashrc
. ~/.bashrc`
	exportExpandPosix = `#!/bin/bash
sed -i '/^export %s/d' ~/.bashrc
echo 'export %s=$%s:%s' >> ~/.bashrc
. ~/.bashrc`

	unsetWindows = `
Set objShell = CreateObject("WScript.Shell")
objShell.Environment("User").Remove("%s")`
	unsetPosix = `unset %s`
)

type ExportOpt struct {
	Expand bool
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
			script = fmt.Sprintf(exportExpandPosix, k, k, k, v)
		} else {
			script = fmt.Sprintf(exportOverridePosix, k, k, v)
		}
	}
	_, err := OnceScript(script)
	return err
}

func Unset(k string) error {
	script := ""
	if util.CurPlatform.OS == "windows" {
		script = fmt.Sprintf(unsetWindows, k)
	} else {
		script = fmt.Sprintf(unsetPosix, k)
	}
	_, err := OnceScript(script)
	return err
}
