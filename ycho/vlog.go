// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import yocki "github.com/ansurfen/yock/interface"

var _ yocki.Ycho = (*vlog)(nil)

type vlog struct{}

func (v *vlog) Info(msg string) {}

func (v *vlog) Infof(msg string, a ...any) {}

func (v *vlog) Fatal(msg string) {}

func (v *vlog) Fatalf(msg string, a ...any) {}

func (v *vlog) Debug(msg string) {}

func (v *vlog) Debugf(msg string, a ...any) {}

func (v *vlog) Warn(msg string) {}

func (v *vlog) Warnf(msg string, a ...any) {}

func (v *vlog) Error(msg string) {}

func (v *vlog) Errorf(msg string, a ...any) {}
