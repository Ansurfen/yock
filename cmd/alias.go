// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import "strings"

var aliases map[string]string

func init() {
	aliases = make(map[string]string)
}

func Alias(k, v string) {
	aliases[k] = v
}

func Unalias(k string) {
	delete(aliases, k)
}

func aliasMap(cmd string) string {
	for k, v := range aliases {
		cmd = strings.ReplaceAll(cmd, "$"+k, v)
	}
	return cmd
}
