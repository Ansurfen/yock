// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"

	"github.com/ansurfen/yock/daemon/proto"
)

// Wait is used to request signal from the daemon
func (yockd *YockDaemon) Wait(ctx context.Context, req *proto.WaitRequest) (*proto.WaitResponse, error) {
	ok := yockd.kernel.SingalStream.Wait(req.Sig)
	return &proto.WaitResponse{Ok: ok}, nil
}

// Notify pushes signal to Daemon
func (yockd *YockDaemon) Notify(ctx context.Context, req *proto.NotifyRequest) (*proto.NotifyResponse, error) {
	yockd.kernel.SingalStream.Notify(req.Sig)
	return &proto.NotifyResponse{}, nil
}
