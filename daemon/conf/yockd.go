// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"net"

	"github.com/ansurfen/yock/ycho"
)

type YockdConf struct {
	Name    string           `yaml:"name"`
	Fs      yockdConfFS      `yaml:"fs"`
	Grpc    yockdConfGrpc    `yaml:"grpc"`
	Gateway yockdConfGateway `yaml:"gateway"`
	Net     YockdConfNet     `yaml:"net"`
	Ycho    ycho.YchoOpt     `yaml:"ycho"`
}

type yockdConfGrpc struct {
	Addr yockdConfGrpcAddr `yaml:"addr"`
}

type yockdConfGrpcAddr struct {
	IP   string `yaml:"ip"`
	Port uint16 `yaml:"port"`
}

func (addr yockdConfGrpcAddr) LocalV4TCPAddr() *net.TCPAddr {
	return &net.TCPAddr{IP: net.IPv4zero, Port: int(addr.Port)}
}

func OpenYockdConf() *YockdConf {
	return &YockdConf{}
}

func CreateYockdConf() *YockdConf {
	return &YockdConf{}
}
