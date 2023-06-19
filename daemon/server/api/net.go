// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"

	"github.com/ansurfen/yock/daemon/proto"
)

// Register tells the daemon the address of the peer.
func (yockd *YockDaemon) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	// for _, addr := range req.Addrs {
	// 	daemon.nodeManager.AddNode(addr)
	// }
	return &proto.RegisterResponse{}, nil
}

// Unregister tells the daemon to remove the peer according to addrs.
func (yockd *YockDaemon) Unregister(ctx context.Context, req *proto.UnregisterRequest) (*proto.UnregisterResponse, error) {
	// for _, addr := range req.Addrs {
	// 	daemon.nodeManager.DelNode(addr)
	// }
	return &proto.UnregisterResponse{}, nil
}
