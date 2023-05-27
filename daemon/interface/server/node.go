// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ansurfen/yock/daemon/interface/client"
	"github.com/ansurfen/yock/util"
)

type Node struct {
	name   string
	cli    *client.YockDaemonClient
	online bool
}

func newNode(opt *client.DaemonOption) *Node {
	return &Node{
		cli: client.New(opt),
	}
}

type NodeManager struct {
	nodes map[string]*Node
}

func NewNodeManager() *NodeManager {
	return &NodeManager{
		nodes: make(map[string]*Node),
	}
}

func (manager *NodeManager) Nodes() map[string]*Node {
	return manager.nodes
}

func (manager *NodeManager) AddNode(addr string) error {
	if !strings.Contains(addr, ":") {
		return util.ErrGeneral
	}
	ip, port, ok := strings.Cut(addr, ":")
	if !ok {
		return util.ErrGeneral
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return util.ErrGeneral
	}
	node := newNode(&client.DaemonOption{
		IP:   &ip,
		Port: &p,
	})
	info, err := node.cli.Info()
	if err != nil {
		return err
	}
	fmt.Println(info)
	return nil
}

func (manager *NodeManager) DelNode(addr string) {

}

func (manager *NodeManager) IsOnline(addr string) bool {
	return false
}

func (manager *NodeManager) Offline(addr string) {
	manager.nodes[addr].online = false
}
