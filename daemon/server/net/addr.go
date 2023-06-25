// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ynet

import (
	"net"

	"github.com/ansurfen/yock/util"
)

type NetAddr struct {
	Family uint16
	IP     string
	Port   uint16
	Zone   string
}

func (info *NetAddr) String() string {
	return util.JsonStr(util.NewJsonObject(map[string]util.JsonValue{
		"IP":     util.NewJsonString(info.IP),
		"Port":   util.NewJsonNumber(int64(info.Port)),
		"Zone":   util.NewJsonString(info.Zone),
		"Family": util.NewJsonNumber(int64(info.Family)),
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
