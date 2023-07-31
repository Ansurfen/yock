// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	"encoding/json"
	"fmt"

	pb "github.com/ansurfen/yock/daemon/proto"
	yocki "github.com/ansurfen/yock/interface"
)

type Protocal interface {
	Parse(message []byte) Protocal
	Version() string
	Bytes() []byte
}

type P2PProtocal interface {
	Protocal
	ParseP2P(message string) P2PProtocal
	Peer() []*NetAddr
}

type RelayProtocal interface {
	Type() pb.ProtocalType
	String() string
}

type EstablishProtocal struct {
	Name  string `json:"name"`
	Delay int    `json:"delay"`
}

func (p EstablishProtocal) Type() pb.ProtocalType {
	return pb.ProtocalType_Establish
}

func (p EstablishProtocal) String() string {
	return fmt.Sprintf(`{"name": "%s", "delay": %d}`, p.Name, p.Delay)
}

type MethodCallProtocal struct {
	Method string `json:"method"`
	Node   string `json:"node"`
}

func (MethodCallProtocal) Type() pb.ProtocalType {
	return pb.ProtocalType_MethodCall
}

func (p MethodCallProtocal) String() string {
	return fmt.Sprintf(`{"node": "%s", "method": "%s"}`, p.Node, p.Method)
}

func ParseProto[T yocki.Protocal](msg string) T {
	var v T
	json.Unmarshal([]byte(msg), &v)
	return v
}

type Context struct {
	addr *NetAddr
}
