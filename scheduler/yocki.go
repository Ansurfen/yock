// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import yocki "github.com/ansurfen/yock/interface/go"

type yockclient yocki.YockInterfaceClient

// TODO implments sdk
type yockInterfaces struct {
	clients map[string]yockclient
}

func NewYockInterface() {

}
