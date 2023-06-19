// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ynet

import (
	"github.com/ansurfen/yock/daemon/client"
	"github.com/ansurfen/yock/daemon/server/conf"
)

type NetworkManager struct {
	nm  *NodeManager
	opt conf.YockdConfNet
}

func (m *NetworkManager) MakeBridge() {
	
}

type NodeManager struct {
	nodes map[string]*Node
}

func NewNodeManager() *NodeManager {
	return &NodeManager{
		nodes: make(map[string]*Node),
	}
}

type Node struct {
	name string
	cli  *client.YockDaemonClient
}

func newNode(opt *conf.YockdConf) *Node {
	return &Node{
		cli: client.New(nil),
	}
}
