// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package netlib

import (
	"net"

	yocki "github.com/ansurfen/yock/interface"
	httplib "github.com/ansurfen/yock/lib/go/net/http"
	maillib "github.com/ansurfen/yock/lib/go/net/mail"
	netiplib "github.com/ansurfen/yock/lib/go/net/netip"
	rpclib "github.com/ansurfen/yock/lib/go/net/rpc"
	stmplib "github.com/ansurfen/yock/lib/go/net/stmp"
	textprotolib "github.com/ansurfen/yock/lib/go/net/textproto"
	urllib "github.com/ansurfen/yock/lib/go/net/url"
)

func LoadNet(yocks yocki.YockScheduler) {
	urllib.LoadUrl(yocks)
	httplib.LoadHttp(yocks)
	maillib.LoadMail(yocks)
	netiplib.LoadNetip(yocks)
	rpclib.LoadRpc(yocks)
	textprotolib.LoadTextproto(yocks)
	stmplib.LoadSmtp(yocks)
	lib := yocks.CreateLib("net")
	lib.SetField(map[string]any{
		// functions
		"IPv4":                net.IPv4,
		"LookupNS":            net.LookupNS,
		"LookupTXT":           net.LookupTXT,
		"Pipe":                net.Pipe,
		"IPv4Mask":            net.IPv4Mask,
		"ResolveIPAddr":       net.ResolveIPAddr,
		"LookupPort":          net.LookupPort,
		"DialTCP":             net.DialTCP,
		"DialUDP":             net.DialUDP,
		"ListenPacket":        net.ListenPacket,
		"FilePacketConn":      net.FilePacketConn,
		"ResolveUDPAddr":      net.ResolveUDPAddr,
		"InterfaceAddrs":      net.InterfaceAddrs,
		"ParseIP":             net.ParseIP,
		"TCPAddrFromAddrPort": net.TCPAddrFromAddrPort,
		"ListenMulticastUDP":  net.ListenMulticastUDP,
		"DialUnix":            net.DialUnix,
		"InterfaceByIndex":    net.InterfaceByIndex,
		"InterfaceByName":     net.InterfaceByName,
		"CIDRMask":            net.CIDRMask,
		"LookupIP":            net.LookupIP,
		"LookupCNAME":         net.LookupCNAME,
		"ListenUnixgram":      net.ListenUnixgram,
		"FileListener":        net.FileListener,
		"ParseCIDR":           net.ParseCIDR,
		"ListenIP":            net.ListenIP,
		"LookupAddr":          net.LookupAddr,
		"ListenTCP":           net.ListenTCP,
		"ListenUnix":          net.ListenUnix,
		"LookupMX":            net.LookupMX,
		"LookupHost":          net.LookupHost,
		"DialTimeout":         net.DialTimeout,
		"Listen":              net.Listen,
		"FileConn":            net.FileConn,
		"JoinHostPort":        net.JoinHostPort,
		"SplitHostPort":       net.SplitHostPort,
		"LookupSRV":           net.LookupSRV,
		"ResolveTCPAddr":      net.ResolveTCPAddr,
		"ListenUDP":           net.ListenUDP,
		"UDPAddrFromAddrPort": net.UDPAddrFromAddrPort,
		"Dial":                net.Dial,
		"Interfaces":          net.Interfaces,
		"DialIP":              net.DialIP,
		"ParseMAC":            net.ParseMAC,
		"ResolveUnixAddr":     net.ResolveUnixAddr,
		// constants
		"FlagUp":           net.FlagUp,
		"FlagBroadcast":    net.FlagBroadcast,
		"FlagLoopback":     net.FlagLoopback,
		"FlagPointToPoint": net.FlagPointToPoint,
		"FlagMulticast":    net.FlagMulticast,
		"FlagRunning":      net.FlagRunning,
		"IPv4len":          net.IPv4len,
		"IPv6len":          net.IPv6len,
		// variable
		"IPv4bcast":                  net.IPv4bcast,
		"IPv4allsys":                 net.IPv4allsys,
		"IPv4allrouter":              net.IPv4allrouter,
		"IPv4zero":                   net.IPv4zero,
		"IPv6zero":                   net.IPv6zero,
		"IPv6unspecified":            net.IPv6unspecified,
		"IPv6loopback":               net.IPv6loopback,
		"IPv6interfacelocalallnodes": net.IPv6interfacelocalallnodes,
		"IPv6linklocalallnodes":      net.IPv6linklocalallnodes,
		"IPv6linklocalallrouters":    net.IPv6linklocalallrouters,
		"DefaultResolver":            net.DefaultResolver,
		"ErrWriteToConnected":        net.ErrWriteToConnected,
		"ErrClosed":                  net.ErrClosed,
	})
}
