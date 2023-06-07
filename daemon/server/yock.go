// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"fmt"
	"net"

	"github.com/ansurfen/cushion/utils"
	yockd "github.com/ansurfen/yock/daemon/proto"
	"github.com/ansurfen/yock/util"
	"google.golang.org/grpc"
)

type YockDaemon struct {
	yockd.UnimplementedYockDaemonServer
	signals     *util.SafeMap[bool]
	fs          map[string]FileInfo
	nodeManager *NodeManager
	opt         *DaemonOption
}

func New(opt *DaemonOption) *YockDaemon {
	return &YockDaemon{
		signals:     util.NewSafeMap[bool](),
		opt:         opt,
		fs:          make(map[string]FileInfo),
		nodeManager: NewNodeManager(),
	}
}

func (daemon *YockDaemon) Close() {
	for _, node := range daemon.nodeManager.Nodes() {
		node.cli.Close()
	}
}

// Ping is used to detect whether the connection is available
func (daemon *YockDaemon) Ping(ctx context.Context, req *yockd.PingRequest) (*yockd.PingResponse, error) {
	return &yockd.PingResponse{}, nil
}

// Wait is used to request signal from the daemon
func (daemon *YockDaemon) Wait(ctx context.Context, req *yockd.WaitRequest) (*yockd.WaitResponse, error) {
	if v, ok := daemon.signals.Get(req.Sig); !ok {
		daemon.signals.SafeSet(req.Sig, false)
		return &yockd.WaitResponse{Ok: false}, nil
	} else if ok && v {
		return &yockd.WaitResponse{Ok: true}, nil
	}
	return &yockd.WaitResponse{Ok: false}, nil
}

// Notify pushes signal to Daemon
func (daemon *YockDaemon) Notify(ctx context.Context, req *yockd.NotifyRequest) (*yockd.NotifyResponse, error) {
	daemon.signals.SafeSet(req.Sig, true)
	return &yockd.NotifyResponse{}, nil
}

// Upload pushes file information to peers so that peers can download files
func (daemon *YockDaemon) Upload(ctx context.Context, req *yockd.UploadRequest) (*yockd.UploadResponse, error) {
	daemon.fs[req.Filename] = FileInfo{
		owner:    req.Owner,
		size:     req.Size,
		hash:     req.Hash,
		createAt: req.CreateAt,
	}
	return &yockd.UploadResponse{}, nil
}

// Download file in other peer
func (daemon *YockDaemon) Download(stream yockd.YockDaemon_DownloadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	file, ok := daemon.fs[req.Filename]
	if !ok {
		return util.ErrFileNotExist
	}
	if file.owner == *daemon.opt.Name {
		if req.Sender == file.owner {
			return nil
		} else {
			raw, err := utils.ReadStraemFromFile(util.Pathf("@/tmp/" + req.Filename))
			if err != nil {
				return err
			}
			for i := 0; i < len(raw); i++ {
				chunk := raw[i : i+*daemon.opt.MTL]
				if err = stream.Send(&yockd.DownloadResponse{Data: chunk}); err != nil {
					return err
				}
			}
		}
	}
	if node, ok := daemon.nodeManager.Nodes()[file.owner]; ok {
		return node.cli.Download(req.Filename)
	} else { // boardcast to every node
		for _, n := range daemon.nodeManager.Nodes() {
			if n.cli.Download(req.Filename) == nil {
				break
			}
		}
	}
	return nil
}

// Register tells the daemon the address of the peer.
func (daemon *YockDaemon) Register(ctx context.Context, req *yockd.RegisterRequest) (*yockd.RegisterResponse, error) {
	for _, addr := range req.Addrs {
		daemon.nodeManager.AddNode(addr)
	}
	return &yockd.RegisterResponse{}, nil
}

// Unregister tells the daemon to remove the peer according to addrs.
func (daemon *YockDaemon) Unregister(ctx context.Context, req *yockd.UnregisterRequest) (*yockd.UnregisterResponse, error) {
	for _, addr := range req.Addrs {
		daemon.nodeManager.DelNode(addr)
	}
	return &yockd.UnregisterResponse{}, nil
}

// Info can obtain the meta information of the target node,
// including CPU, DISK, MEM and so on.
// You can specify it by InfoRequest, and by default only basic parameters
// (the name of the node, the file uploaded, and the connection information) are returned.
func (daemon *YockDaemon) Info(ctx context.Context, req *yockd.InfoRequest) (*yockd.InfoResponse, error) {
	return &yockd.InfoResponse{Name: *daemon.opt.Name}, nil
}

func (daemon *YockDaemon) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *daemon.opt.Port))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer()
	yockd.RegisterYockDaemonServer(gsrv, daemon)
	if err := gsrv.Serve(listen); err != nil {
		panic(err)
	}
}
