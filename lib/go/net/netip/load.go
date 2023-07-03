// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package netiplib

import (
	"net/netip"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadNetip(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("netip")
	lib.SetField(map[string]any{
		// functions
		"IPv6LinkLocalAllRouters": netip.IPv6LinkLocalAllRouters,
		"PrefixFrom":              netip.PrefixFrom,
		"MustParseAddrPort":       netip.MustParseAddrPort,
		"IPv6Loopback":            netip.IPv6Loopback,
		"IPv6LinkLocalAllNodes":   netip.IPv6LinkLocalAllNodes,
		"AddrFrom16":              netip.AddrFrom16,
		"ParseAddr":               netip.ParseAddr,
		"AddrPortFrom":            netip.AddrPortFrom,
		"ParseAddrPort":           netip.ParseAddrPort,
		"AddrFrom4":               netip.AddrFrom4,
		"MustParseAddr":           netip.MustParseAddr,
		"IPv6Unspecified":         netip.IPv6Unspecified,
		"IPv4Unspecified":         netip.IPv4Unspecified,
		"AddrFromSlice":           netip.AddrFromSlice,
		"ParsePrefix":             netip.ParsePrefix,
		"MustParsePrefix":         netip.MustParsePrefix,
		// constants
		// variable
	})
}
