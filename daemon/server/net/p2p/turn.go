// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package p2p

import (
	. "github.com/ansurfen/yock/daemon/server/net"
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

type Turn struct {
}

type TurnServer struct {
	info     *NetAddr
	listener *net.TCPListener
}

type OptionTurn func(*Turn) error

func NewTurn(opts ...OptionTurn) *Turn {
	t := &Turn{}
	for _, opt := range opts {
		if err := opt(t); err != nil {
			panic(err)
		}
	}
	return t
}

func newTurn() *TurnServer {
	return &TurnServer{}
}

func (turn *TurnServer) Relay() error {
	var err error
	turn.listener, err = net.ListenTCP("tcp", turn.info.LocalV4TCPAddr())
	return err
}

func (turn *TurnServer) Close() {
	turn.listener.Close()
}

func (turn *TurnServer) Read() {
	turn.listener.Accept()
}

func (turn *TurnServer) Write() {

}
