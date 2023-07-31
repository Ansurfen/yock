// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	"net"
)

type TCPServer struct {
	listen   *net.TCPListener
	protocal Protocal
	addr     *NetAddr
}

type TCPClient struct {
	conn     net.Conn
	protocal Protocal
}

func ListenTCP(addr *NetAddr, protocal Protocal) *TCPServer {
	listen, err := net.ListenTCP("tcp", addr.TCPAddr())
	if err != nil {
		panic(err)
	}
	return &TCPServer{
		listen:   listen,
		addr:     addr,
		protocal: protocal,
	}
}
