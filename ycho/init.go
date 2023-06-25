// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import yocki "github.com/ansurfen/yock/interface"

var ycho yocki.Ycho

func init() {}

func Info(msg string) {}

func Infof(msg string, v ...any) {}

func Fatal(msg string) {}

func Fatalf(msg string, v ...any) {}

func Debug(msg string) {}

func Debugf(msg string, v ...any) {}

func Warn(msg string) {}

func Warnf(msg string, v ...any) {}
