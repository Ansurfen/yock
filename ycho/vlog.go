// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

import (
	"io"

	yocki "github.com/ansurfen/yock/interface"
)

var (
	_ io.Writer  = (*vWriter)(nil)
	_ yocki.Ycho = (*Vlog)(nil)
)

type vWriter struct{}

func (vw *vWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

type Vlog struct{}

func (v *Vlog) Eventloop() {}

func (v *Vlog) Progress(total int64, r io.Reader) io.Writer {
	return &vWriter{}
}

func (v *Vlog) Info(msg string) {}

func (v *Vlog) Infof(msg string, a ...any) {}

func (v *Vlog) Fatal(msg string) {}

func (v *Vlog) Fatalf(msg string, a ...any) {}

func (v *Vlog) Debug(msg string) {}

func (v *Vlog) Debugf(msg string, a ...any) {}

func (v *Vlog) Warn(msg string) {}

func (v *Vlog) Warnf(msg string, a ...any) {}

func (v *Vlog) Error(msg string) {}

func (v *Vlog) Errorf(msg string, a ...any) {}
