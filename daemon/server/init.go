// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ansurfen/yock/daemon/conf"
	"github.com/ansurfen/yock/daemon/gateway"
	"github.com/ansurfen/yock/daemon/gateway/agent"
	"github.com/ansurfen/yock/daemon/kernel"
	pb "github.com/ansurfen/yock/daemon/proto"
	"github.com/ansurfen/yock/daemon/util"
	yocke "github.com/ansurfen/yock/env"
	"github.com/ansurfen/yock/ycho"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var _ pb.YockDaemonServer = (*YockDaemon)(nil)

type YockDaemon struct {
	pb.UnimplementedYockDaemonServer
	gate *gateway.YockdGateWay
	*kernel.YockKernel
	conf        *conf.YockdConf
	gouroutines chan func(context.Context)
}

func New() *YockDaemon {
	yockd := &YockDaemon{
		gate:        gateway.New(),
		conf:        yocke.GetEnv[*conf.YockdConf]().Conf(),
		YockKernel:  kernel.NewKernel(),
		gouroutines: make(chan func(context.Context)),
	}
	for name, opt := range yockd.conf.Gateway.Agent {
		if opt.Enable {
			switch name {
			case "jwt":
				yockd.gate.SetAgent(name, agent.NewJWTAgent(opt.Path...))
			case "pwd":
				yockd.gate.SetAgent(name, agent.NewPwdAgent(opt.Path...))
			case "crt":
			default: // diy
			}
			for _, rule := range opt.Rule {
				err := yockd.gate.AddRule([]byte(rule))
				if err != nil {
					ycho.Error(err)
				}
			}
		}
	}
	switch yockd.conf.Gateway.Policy {
	case "user":
		yockd.gate.SetPolicy(gateway.NewUserPolicy())
		// yockd.gate.SetRule("root", "jwt.default")
		ycho.Info("enable user policy")
	case "router":
		yockd.gate.SetPolicy(gateway.NewRouterPolicy())
		ycho.Info("enable router policy")
	default:
		yockd.gate.SetPolicy(gateway.NullPolicy{})
		ycho.Warnf("insecure, lack of rule")
	}
	return yockd
}

func (yockd *YockDaemon) Close() {}

func (yockd *YockDaemon) Run() {
	if len(yockd.conf.Grpc.Addr.IP) == 0 {
		yockd.conf.Grpc.Addr.IP = ""
	}
	if yockd.conf.Grpc.Addr.Port == 0 {
		ycho.Fatalf("invalid port")
	}
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", yockd.conf.Grpc.Addr.IP, yockd.conf.Grpc.Addr.Port))
	// listen, err := net.ListenTCP("tcp", yockd.conf.Grpc.Addr.LocalV4TCPAddr())
	if err != nil {
		ycho.Fatal(err)
	}
	opts := []grpc.ServerOption{yockd.gate.GuardUnary(), yockd.gate.GuardStream(), grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 30 * time.Minute,
	})}

	if yockd.conf.Gateway.TLS.Enable {
		opts = append(opts, yockd.gate.GuardTransport(
			yockd.conf.Gateway.TLS.Cert,
			yockd.conf.Gateway.TLS.Key,
			yockd.conf.Gateway.TLS.Ca...))
	}

	srv := grpc.NewServer(opts...)

	go yockd.NetworkManager.MakeBridge()
	ctx := context.Background()
	go func() {
		for {
			select {
			case fn := <-yockd.gouroutines:
				go fn(ctx)
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	go yockd.YockKernel.Init()
	
	ycho.Infof("start at %s:%d", yockd.conf.Grpc.Addr.IP, yockd.conf.Grpc.Addr.Port)
	pb.RegisterYockDaemonServer(srv, yockd)
	if err := srv.Serve(listen); err != nil {
		panic(err)
	}
}

// Ping is used to detect whether the connection is available
func (yockd *YockDaemon) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}

// Info can obtain the meta information of the target node,
// including CPU, DISK, MEM and so on.
// You can specify it by InfoRequest, and by default only basic parameters
// (the name of the node, the file uploaded, and the connection information) are returned.
func (yockd *YockDaemon) Info(ctx context.Context, req *pb.InfoRequest) (*pb.InfoResponse, error) {
	return &pb.InfoResponse{Name: util.ID}, nil
}
