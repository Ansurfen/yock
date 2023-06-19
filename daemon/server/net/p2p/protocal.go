// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package p2p

import "github.com/ansurfen/yock/daemon/server/net"

type Protocal interface {
	Parse(message []byte) Protocal
	Version() string
	Bytes() []byte
}

type P2PProtocal interface {
	Protocal
	ParseP2P(message string) P2PProtocal
	Peer() []*ynet.NetAddr
}

type Context struct {
	addr *ynet.NetAddr
}
