// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"

	pb "github.com/ansurfen/yock/daemon/proto"
	"github.com/ansurfen/yock/ycho"
)

// Wait is used to request signal from the daemon
func (yockd *YockDaemon) SignalWait(ctx context.Context, req *pb.WaitRequest) (*pb.WaitResponse, error) {
	ok := yockd.SignalStream.Wait(req.GetSig())
	ycho.Infof("waiting for %s, status: %v", req.GetSig(), ok)
	return &pb.WaitResponse{Ok: ok}, nil
}

// Notify pushes signal to Daemon
func (yockd *YockDaemon) SignalNotify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	yockd.SignalStream.Notify(req.GetSig())
	ycho.Infof("notify %s signal", req.GetSig())
	return &pb.NotifyResponse{}, nil
}

func (yockd *YockDaemon) SignalList(context.Context, *pb.SignalListRequest) (*pb.SignalListResponse, error) {
	return &pb.SignalListResponse{
		Sigs: yockd.SignalStream.List(),
	}, nil
}

func (yockd *YockDaemon) SingalClear(ctx context.Context, req *pb.SignalClearRequest) (*pb.SignalClearResponse, error) {
	ycho.Infof("clear signals: %v", req.GetSigs())
	yockd.SignalStream.Clear(req.GetSigs()...)
	return &pb.SignalClearResponse{}, nil
}

func (yockd *YockDaemon) SingalInfo(ctx context.Context, req *pb.SignalInfoRequest) (*pb.SignalInfoResponse, error) {
	status, exist := yockd.SignalStream.Info(req.GetSig())
	return &pb.SignalInfoResponse{
		Status: status,
		Exist:  exist,
	}, nil
}
