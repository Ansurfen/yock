// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rpclib

import (
	"net/rpc"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadRpc(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("rpc")
	lib.SetField(map[string]any{
		// functions
		"NewServer":          rpc.NewServer,
		"ServeRequest":       rpc.ServeRequest,
		"ServeConn":          rpc.ServeConn,
		"NewClientWithCodec": rpc.NewClientWithCodec,
		"DialHTTP":           rpc.DialHTTP,
		"HandleHTTP":         rpc.HandleHTTP,
		"RegisterName":       rpc.RegisterName,
		"ServeCodec":         rpc.ServeCodec,
		"Accept":             rpc.Accept,
		"Register":           rpc.Register,
		"DialHTTPPath":       rpc.DialHTTPPath,
		"Dial":               rpc.Dial,
		"NewClient":          rpc.NewClient,
		// constants
		"DefaultRPCPath":   rpc.DefaultRPCPath,
		"DefaultDebugPath": rpc.DefaultDebugPath,
		// variable
		"ErrShutdown":   rpc.ErrShutdown,
		"DefaultServer": rpc.DefaultServer,
	})
}
