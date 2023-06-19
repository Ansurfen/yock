// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"net"

	"github.com/ansurfen/yock/daemon/proto"
	"github.com/ansurfen/yock/daemon/server/conf"
	"github.com/ansurfen/yock/daemon/server/gateway"
	"github.com/ansurfen/yock/daemon/server/kernel"
	"google.golang.org/grpc"
)

type YockDaemon struct {
	proto.UnimplementedYockDaemonServer
	gate   *gateway.YockdGateWay
	kernel *kernel.YockKernel
	conf   *conf.YockdConf
}

func New() *YockDaemon {
	return &YockDaemon{
		gate: gateway.New(),
	}
}

func (yockd *YockDaemon) Close() {
}

func (yockd *YockDaemon) Run() {
	listen, err := net.ListenTCP("tcp", yockd.conf.Grpc.Addr.LocalV4TCPAddr())
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer(
		yockd.gate.GuardTransport(
			yockd.conf.TLS.Cert,
			yockd.conf.TLS.Key,
			yockd.conf.TLS.Ca...),
		yockd.gate.GuardUnary(),
		yockd.gate.GuardStream())

	go yockd.kernel.NetworkManager.MakeBridge()

	if err := srv.Serve(listen); err != nil {
		panic(err)
	}
}

// Ping is used to detect whether the connection is available
func (yockd *YockDaemon) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{}, nil
}

// Info can obtain the meta information of the target node,
// including CPU, DISK, MEM and so on.
// You can specify it by InfoRequest, and by default only basic parameters
// (the name of the node, the file uploaded, and the connection information) are returned.
func (yockd *YockDaemon) Info(ctx context.Context, req *proto.InfoRequest) (*proto.InfoResponse, error) {
	return &proto.InfoResponse{Name: yockd.conf.Name}, nil
}
