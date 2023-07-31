// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	"net"
)

type UDPListener struct {
	conn *net.UDPConn
	addr *NetAddr
}

func NewUDPListener(addr *NetAddr) *UDPListener {
	conn, err := net.ListenUDP("udp", addr.LocalV4UDPAddr())
	if err != nil {
		panic(err)
	}
	return &UDPListener{
		conn: conn,
		addr: addr,
	}
}

func (listen *UDPListener) Write() {}

func (listen *UDPListener) Read() {}

func (listen *UDPListener) Dial(addr *NetAddr) {
	// listen.conn.WriteToUDP()
}

type UDPConn struct {}

func (conn *UDPConn) Write() {}

func (conn *UDPConn) Read() {}

type TCPListerner struct {
	listen *net.TCPListener
}

func NewTCPListerner(addr *NetAddr) *TCPListerner {
	listen, err := net.ListenTCP("tcp", addr.LocalV4TCPAddr())
	if err != nil {
		panic(err)
	}
	return &TCPListerner{listen: listen}
}

func (listen *TCPListerner) Dial(addr *NetAddr) {
	// net.DialTCP()
}

func (listen *TCPListerner) Accept() {
	listen.listen.Accept()
}

func New(opts ...NodeOption) {}

type NodeOption func()
