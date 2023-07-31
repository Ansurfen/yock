// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocki

import (
	"os"
	"testing"
	"time"

	yocki "github.com/ansurfen/yock/interface/go"
)

func TestYocki(t *testing.T) {
	go func() {
		time.Sleep(100 * time.Second)
		os.Exit(0)
	}()
	s := New()
	s.Register("SayHello", func(req *yocki.CallRequest) (*yocki.CallResponse, error) {
		return &yocki.CallResponse{Buf: "I'm Golang"}, nil
	})
	s.Run(9090)
}
