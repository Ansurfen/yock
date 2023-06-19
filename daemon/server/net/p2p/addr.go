// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package p2p

import (
	"encoding/json"

	"github.com/ansurfen/yock/daemon/server/net"
	"github.com/ccding/go-stun/stun"
)

const (
	Center = 0b01
	Worker = 0b10
)

type PeerAddr struct {
	ynet.NetAddr
	natType  stun.NATType
	Peer     []string
	Attr     uint8
	ID       string
	Protocal P2PProtocal
}

func (info *PeerAddr) IsCenter() bool {
	return (info.Attr & 0b01) == Center
}

func (info *PeerAddr) IsWorker() bool {
	return (info.Attr & 0b10) == Worker
}

func (info *PeerAddr) String() string {
	raw, _ := json.Marshal(info)
	return string(raw)
}

func (info *PeerAddr) CanMakeHole() bool {
	return info.natType != stun.NATSymmetric && info.natType != stun.NATUnknown
}
