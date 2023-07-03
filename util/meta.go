// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

var (
	// YockBuild has two values, dev and release,
	// which correspond to two different modes.
	//
	// The value YOCK_PATH under dev is taken from main.go under the /ctl.
	// If it is in release mode, the actual location of the executable file shall prevail.
	//
	// This value will be changed by -ldflag at compile time. Details to see /auto/build.lua.
	YockBuild   = "dev"
	YockVersion = "0.0.16"
)
