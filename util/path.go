// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "os"

var (
	// WorkSpace is the .yock path in the UserHome
	//
	// You can think of it as yock's global workspace
	// for storing user's information.
	WorkSpace  string
	PluginPath string
	DriverPath string
	// executable file path
	YockPath string
)

// Pathf to format path
//
// @/abc => {WorkSpace}/abc (WorkSpace = UserHome + .yock)
//
// ~/abc => {YockPath}/abc (YockPath = executable file path)
func Pathf(path string) string {
	if len(path) > 0 {
		switch path[0] {
		case '@':
			path = WorkSpace + path[1:]
		case '~':
			path = YockPath + path[1:]
		case '$':
			wd, _ := os.Getwd()
			path = wd + path[1:]
		}
	}
	return path
}
