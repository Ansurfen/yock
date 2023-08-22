// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"strings"

	"github.com/ansurfen/yock/util"
)

func Nohup(cmds string) error {
	if util.CurPlatform.OS == "windows" {
		name := ""
		args := ""
		idx := strings.Index(cmds, " ")
		if idx == -1 {
			name = cmds
			args = " "
		} else {
			name = cmds[:idx]
			args = cmds[idx:]
		}
		cmds = fmt.Sprintf(`Start-Process -FilePath %s -ArgumentList "%s" -WindowStyle Hidden`, name, args)
	} else {
		cmds += " &"
	}
	_, err := Exec(ExecOpt{}, cmds)
	return err
}
