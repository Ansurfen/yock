// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"fmt"
	"github.com/ansurfen/yock/daemon/net"
	"github.com/ansurfen/yock/daemon/proto"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/ycho"
	"google.golang.org/grpc/peer"
	"strconv"
	"strings"
	"time"
)

func (yockd *YockDaemon) Dial(ctx context.Context, req *proto.DialRequest) (*proto.DialResponse, error) {
	to := req.GetTo()
	form := req.GetFrom()
	ip := to.GetIp()
	port := int(to.GetPort())
	var err error
	if to.GetPublic() {
		remote := yockd.NetworkManager.Node(to.GetName())
		if remote != nil {
			err = remote.Mark(form.GetName(), fmt.Sprintf("%s:%d", form.GetIp(), form.GetPort()))
		} else {
			remote = net.NewDirect(&net.YockdClientOption{
				IP:     ip,
				Port:   port,
				Global: yockd.conf,
			})
			err = remote.Mark(form.GetName(), fmt.Sprintf("%s:%d", form.GetIp(), form.GetPort()))
			yockd.SetNode(form.GetName(), remote)
		}

		if err != nil {
			// 公网通不了， 可能需要去报告给其他peer
			ycho.Error(err)
		} else {
			remote := yockd.Node(to.GetName())
			// 这条是公网自己的连接
			yockd.gouroutines <- func(ctx context.Context) {
				remote.MakeTunnel(form.GetName(), ctx, yockd.SignalStream.System(), yockd.SignalStream.SystemEvent())
			}
		}
	} else {
		// 遍历一下公网的客户端列表，查看能做为中转的跳板服务器
		available := false
		for _, node := range yockd.Nodes() {
			if node.IsPublic() {
				available = true
				yockd.SetNode(to.GetName(), net.NewDelivery(to.GetName(), yockd.System(), yockd.SystemEvent()))
			}
		}
		if !available {
			err = fmt.Errorf("public server not found")
		}
	}
	return &proto.DialResponse{}, err
}

func proxyInvoke(promise yocki.Promise, stream proto.YockDaemon_TunnelServer) func(p yocki.Protocal) (any, error) {
	return func(p yocki.Protocal) (any, error) {
		id := promise.NextID()
		err := stream.Send(&proto.TunnelResponse{
			Type: p.Type(),
			Body: p.String(),
			Id:   id,
		})
		if err != nil {
			return nil, err
		}
		v, ok := promise.LoadWithTimeout(id, time.Second*10)
		if !ok {
			return nil, fmt.Errorf("context deadline exceeded")
		}
		return v, nil
	}
}

func (yockd *YockDaemon) Tunnel(stream proto.YockDaemon_TunnelServer) error {
	node, ok := peer.FromContext(stream.Context())
	if ok {
		ycho.Infof("new tunnel: %s", node.Addr.String())
	}
	proxySource := ""
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		switch req.GetType() {
		case proto.ProtocalType_Unknown:
		case proto.ProtocalType_Establish:
			p := net.ParseProto[net.EstablishProtocal](req.GetBody())
			proxySource = p.Name
			if node := yockd.NetworkManager.Node(proxySource); node != nil {
				yockd.SetNode(proxySource, &net.ProxyYockdClient{
					Invoke: proxyInvoke(yockd.SignalStream.System(), stream),
				})
				ycho.Infof("node register: %s -> %s", proxySource, "")
			} else {
				node.(*net.ProxyYockdClient).Invoke = proxyInvoke(yockd.SignalStream.System(), stream)
			}
		case proto.ProtocalType_MethodCall:
			p := net.ParseProto[net.MethodCallProtocal](req.GetBody())
			ycho.Infof("[%d] call %s.%s", req.GetId(), p.Node, p.Method)
			switch p.Method {
			case "info":
				res, err := yockd.Node(p.Node).Info()
				if err != nil {
					ycho.Error(err)
				}
				ycho.Infof("[%d] from %s to %s", req.GetId(), p.Node, proxySource)
				stream.Send(&proto.TunnelResponse{
					Id:   req.GetId(),
					Body: res,
					Type: proto.ProtocalType_MethodReturn,
				})
			}
		case proto.ProtocalType_MethodReturn:
			ycho.Infof("[%d] %s return ", req.GetId(), proxySource)
			yockd.SignalStream.System().Store(req.GetId(), req.GetBody())
		}
	}
}

func (yockd *YockDaemon) Mark(ctx context.Context, req *proto.MarkRequest) (*proto.MarkResponse, error) {
	if node := yockd.Node(req.GetName()); node == nil {
		yockd.SetNode(req.GetName(), &net.ProxyYockdClient{
			Addr: req.GetAddr(),
		})
		ycho.Infof("node register: %s -> %s", req.GetName(), req.GetAddr())
	}
	return &proto.MarkResponse{}, nil
}

func (yockd *YockDaemon) Call(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	var (
		ret   string
		g_err error = fmt.Errorf("fail to call")
		peer  net.Node
		args  []string = req.GetArgs()
	)
	if node := yockd.Node(req.GetNode()); node != nil {
		peer = node
	} else {
		return &proto.CallResponse{}, g_err
	}
	switch strings.ToLower(req.GetMethod()) {
	case "info":
		ret, g_err = peer.Info()
	case "signalwait":
		if len(args) > 0 {
			ok, err := peer.SignalWait(args[0])
			ret = strconv.FormatBool(ok)
			g_err = err
		} else {

		}
	case "signalnotify":
		if len(args) > 0 {
			g_err = peer.SignalNotify(args[0])
		}
	case "signallist":
		// res, err := peer.SignalList()

		// data, err := json.Marshal(res)
	default:
		g_err = fmt.Errorf("invalid method")
	}
	return &proto.CallResponse{
		Ret: ret,
	}, g_err
}
