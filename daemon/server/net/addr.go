// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ynet

import (
	"net"

	"github.com/ansurfen/cushion/utils"
)

type NetAddr struct {
	Family uint16
	IP     string
	Port   uint16
	Zone   string
}

func (info *NetAddr) String() string {
	return utils.JsonStr(utils.NewJsonObject(map[string]utils.JsonValue{
		"IP":     utils.NewJsonString(info.IP),
		"Port":   utils.NewJsonNumber(int64(info.Port)),
		"Zone":   utils.NewJsonString(info.Zone),
		"Family": utils.NewJsonNumber(int64(info.Family)),
	}))
}

func (info *NetAddr) TCPAddr() *net.TCPAddr {
	return &net.TCPAddr{IP: net.IP(info.IP).To4(), Port: int(info.Port), Zone: info.Zone}
}

func (info *NetAddr) UDPAddr() *net.UDPAddr {
	return &net.UDPAddr{IP: net.IP(info.IP).To4(), Port: int(info.Port), Zone: info.Zone}
}

func (info *NetAddr) LocalV4UDPAddr() *net.UDPAddr {
	return &net.UDPAddr{IP: net.IPv4zero, Port: int(info.Port), Zone: info.Zone}
}

func (info *NetAddr) LocalV4TCPAddr() *net.TCPAddr {
	return &net.TCPAddr{IP: net.IPv4zero, Port: int(info.Port), Zone: info.Zone}
}

func UDPAddr2NetAddr(addr *net.UDPAddr) *NetAddr {
	return &NetAddr{
		Port: uint16(addr.Port),
		IP:   string(addr.IP),
		Zone: addr.Zone,
	}
}

func TCPAddr2NetAddr(addr *net.TCPAddr) *NetAddr {
	return &NetAddr{
		Port: uint16(addr.Port),
		IP:   string(addr.IP),
		Zone: addr.Zone,
	}
}
