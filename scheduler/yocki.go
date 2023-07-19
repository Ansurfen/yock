// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"github.com/ansurfen/yock/interface/go/client"
)

type yockiClient = *client.YockInterface

type yockInterfaces struct {
	clients map[string]yockiClient
}

func newYockInterface() *yockInterfaces {
	return &yockInterfaces{
		clients: make(map[string]yockiClient),
	}
}

func (yi *yockInterfaces) Connect(name, ip string, port int) {
	if _, ok := yi.clients[name]; !ok {
		yi.clients[name] = client.New(ip, port)
	}
}

func (yi *yockInterfaces) Close(name string) {
	if client, ok := yi.clients[name]; ok {
		client.Close()
		delete(yi.clients, name)
	}
}

func (yi *yockInterfaces) Call(name, fn, arg string) (string, error) {
	if client, ok := yi.clients[name]; ok {
		return client.Call(fn, arg)
	}
	return "", nil
}

func (yi *yockInterfaces) Shutdown() {
	for _, client := range yi.clients {
		client.Close()
	}
}
