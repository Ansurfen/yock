// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"net"

	"github.com/denisbrodbeck/machineid"
)

var ID string

func init() {
	id, err := machineid.ID()
	if err != nil {
		interfaces, err := net.Interfaces()
		if err != nil || len(interfaces) == 0 {
			panic("fail to init machine id")
		}
		id = interfaces[0].HardwareAddr.String()
	}
	ID = id
}
