// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ycho

type tlog struct{}

func (t *tlog) Info(msg string) {

}

func (t *tlog) Infof(msg string, v ...any) {

}

func NewTLog(conf YchoOpt) (*tlog, error) {
	return nil, nil
}
