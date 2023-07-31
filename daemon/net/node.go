// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import yocki "github.com/ansurfen/yock/interface"

type Node yocki.YockdClient

type NetworkManager struct {
	nodes map[string]Node
}

func NewNetworkManager() *NetworkManager {
	return &NetworkManager{
		nodes: make(map[string]Node),
	}
}

func (m *NetworkManager) MakeBridge() {}

func (m *NetworkManager) Node(name string) Node {
	if n := m.nodes[name]; n != nil {
		return n
	}
	return nil
}

func (m *NetworkManager) SetNode(name string, node Node) {
	m.nodes[name] = node
}

func (m *NetworkManager) Nodes() map[string]Node {
	return m.nodes
}
