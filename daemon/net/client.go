// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	pb "github.com/ansurfen/yock/daemon/proto"
	du "github.com/ansurfen/yock/daemon/util"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

var (
	_ yocki.YockdClient = (*ProxyYockdClient)(nil)
	_ yocki.YockdClient = (*DeliveryClient)(nil)
	_ yocki.YockdClient = (*DirectClient)(nil)
)

type ProxyYockdClient struct {
	Addr   string
	Invoke func(p yocki.Protocal) (any, error)
	Node   string
}

func (c *ProxyYockdClient) IsPublic() bool {
	return false
}

func (c *ProxyYockdClient) Name() string {
	return c.Node
}

func (c *ProxyYockdClient) Info() (string, error) {
	a, err := c.Invoke(MethodCallProtocal{
		Method: "info",
		Node:   c.Node,
	})
	if err != nil {
		return "", err
	}
	return a.(string), nil
}

func (c *ProxyYockdClient) Ping() error {
	return nil
}

func (c *ProxyYockdClient) Mark(name, addr string) error {
	return nil
}

func (c *ProxyYockdClient) Dial(form, to *pb.NodeInfo) error {
	return nil
}

func (c *ProxyYockdClient) Close() {}

func (c *ProxyYockdClient) Download(file string) error {
	return nil
}

func (c *ProxyYockdClient) FileSystemPut(src, dst string) error {
	return nil
}

func (c *ProxyYockdClient) FileSystemGet(src, dst string) error {
	return nil
}

func (c *ProxyYockdClient) MakeTunnel(name string, ctx context.Context, p yocki.Promise, e chan yocki.PromiseEvent) error {
	return nil
}

func (c *ProxyYockdClient) SignalNotify(sig string) error {
	return nil
}

func (c *ProxyYockdClient) SignalWait(sig string) (bool, error) {
	return false, nil
}

func (c *ProxyYockdClient) SignalInfo(sig string) (bool, bool, error) {
	return false, false, nil
}

func (c *ProxyYockdClient) SignalList() ([]string, error) {
	return nil, nil
}

func (c *ProxyYockdClient) SignalClear(sigs ...string) error {
	return nil
}

func (c *ProxyYockdClient) ProcessList() ([]*pb.Process, error) {
	return nil, nil
}

func (c *ProxyYockdClient) ProcessKill(pid int64) error {
	return nil
}

func (c *ProxyYockdClient) ProcessFind(pid int64, cmd string) ([]*pb.Process, error) {
	return nil, nil
}

func (c *ProxyYockdClient) ProcessSpawn(t pb.ProcessSpawnType, spec, cmd string) (int64, error) {
	return 0, nil
}

func (c *ProxyYockdClient) Status() {}

func (c *ProxyYockdClient) Call(node, method string, args ...string) (string, error) {
	return "", nil
}

type DeliveryClient struct {
	event         chan yocki.PromiseEvent
	node          string
	maxTimeout    time.Duration
	maxRetryCount int
	promise       yocki.Promise
}

func NewDelivery(node string, promise yocki.Promise, e chan yocki.PromiseEvent) *DeliveryClient {
	return &DeliveryClient{
		event:         e,
		maxTimeout:    time.Second * 10,
		maxRetryCount: 0,
		promise:       promise,
		node:          node,
	}
}

func (c *DeliveryClient) Name() string {
	return c.node
}

func (DeliveryClient) IsPublic() bool {
	return false
}

func (c *DeliveryClient) Info() (string, error) {
	v, ok := c.invoke("info", 5*time.Second)
	if !ok {
		return "", fmt.Errorf("context deadline exceeded")
	}
	return v.(string), nil
}

func (c *DeliveryClient) Ping() error {
	return nil
}

func (c *DeliveryClient) Mark(name, addr string) error {
	return nil
}

func (c *DeliveryClient) Dial(form, to *pb.NodeInfo) error {
	return nil
}

func (c *DeliveryClient) Close() {}

func (c *DeliveryClient) Download(file string) error {
	return nil
}

func (c *DeliveryClient) FileSystemPut(src, dst string) error {
	return nil
}

func (c *DeliveryClient) FileSystemGet(src, dst string) error {
	return nil
}

func (c *DeliveryClient) MakeTunnel(name string, ctx context.Context, p yocki.Promise, e chan yocki.PromiseEvent) error {
	return nil
}

func (c *DeliveryClient) ProcessList() ([]*pb.Process, error) {
	v, ok := c.invoke("processlist", 5*time.Second)
	if !ok {
		return nil, fmt.Errorf("context deadline exceeded")
	}
	return v.([]*pb.Process), nil
}

func (c *DeliveryClient) ProcessKill(pid int64) error {
	_, ok := c.invoke("processkill", 5*time.Second)
	if !ok {
		return fmt.Errorf("context deadline exceeded")
	}
	return nil
}

func (c *DeliveryClient) ProcessFind(pid int64, cmd string) ([]*pb.Process, error) {
	v, ok := c.invoke("processfind", 5*time.Second)
	if !ok {
		return nil, fmt.Errorf("context deadline exceeded")
	}
	return v.([]*pb.Process), nil
}

func (c *DeliveryClient) ProcessSpawn(t pb.ProcessSpawnType, spec, cmd string) (int64, error) {
	v, ok := c.invoke("processspawn", 5*time.Second)
	if !ok {
		return 0, fmt.Errorf("context deadline exceeded")
	}
	return v.(int64), nil
}

func (c *DeliveryClient) SignalNotify(sig string) error {
	_, ok := c.invoke("notify", 5*time.Second)
	if !ok {
		return fmt.Errorf("context deadline exceeded")
	}
	return nil
}

func (c *DeliveryClient) SignalWait(sig string) (bool, error) {
	v, ok := c.invoke("wait", 5*time.Second)
	if !ok {
		return false, fmt.Errorf("context deadline exceeded")
	}
	return v.(bool), nil
}

func (c *DeliveryClient) SignalInfo(sig string) (bool, bool, error) {
	v, ok := c.invoke("signalinfo", 5*time.Second)
	if !ok {
		return false, false, fmt.Errorf("context deadline exceeded")
	}
	res := []bool{}
	json.Unmarshal([]byte(v.(string)), &res)
	return res[0], res[1], nil
}

func (c *DeliveryClient) SignalList() ([]string, error) {
	return nil, nil
}

func (c *DeliveryClient) SignalClear(sigs ...string) error {
	return nil
}

func (c *DeliveryClient) Status() {}

func (c *DeliveryClient) Call(node, method string, args ...string) (string, error) {
	return "", nil
}

func (c *DeliveryClient) invoke(method string, timeout ...time.Duration) (any, bool) {
	id := c.promise.NextID()
	c.event <- du.PostPromiseEvent(id, MethodCallProtocal{
		Node:   c.node,
		Method: method,
	})
	t := c.maxTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}
	return c.promise.LoadWithTimeout(id, t)
}

type DirectClient struct {
	conn *grpc.ClientConn
	cli  pb.YockDaemonClient
	opt  *YockdClientOption
	name string
}

func NewDirect(opt *YockdClientOption) yocki.YockdClient {
	kaParams := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}

	dialOption := grpc.WithKeepaliveParams(kaParams)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", opt.IP, opt.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()), dialOption)

	if err != nil {
		panic(err)
	}
	return &DirectClient{
		conn: conn,
		cli:  pb.NewYockDaemonClient(conn),
		opt:  opt,
	}
}

func (c *DirectClient) ProcessList() ([]*pb.Process, error) {
	res, err := c.cli.ProcessList(context.Background(), &pb.ProcessListRequest{})
	return res.GetRes(), err
}

func (c *DirectClient) ProcessKill(pid int64) error {
	_, err := c.cli.ProcessKill(context.Background(), &pb.ProcessKillRequest{
		Pid: pid,
	})
	return err
}

func (c *DirectClient) ProcessFind(pid int64, cmd string) ([]*pb.Process, error) {
	res, err := c.cli.ProcessFind(context.Background(), &pb.ProcessFindRequest{
		Pid: pid,
		Cmd: cmd,
	})
	if err != nil {
		return nil, err
	}
	return res.GetRes(), nil
}

func (c *DirectClient) ProcessSpawn(t pb.ProcessSpawnType, spec, cmd string) (int64, error) {
	res, err := c.cli.ProcessSpawn(context.Background(), &pb.ProcessSpawnRequest{
		Type: t,
		Spec: spec,
		Cmd:  cmd,
	})
	return res.GetPid(), err
}

func (client *DirectClient) Name() string {
	return client.name
}

func (client *DirectClient) Status() {
	fmt.Println(client.conn.GetState())
}

func (client *DirectClient) Close() {
	client.conn.Close()
}

func (c *DirectClient) Mark(name, addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := c.cli.Mark(ctx, &pb.MarkRequest{
		Name: name,
		Addr: addr,
	})
	return err
}

// Ping is used to detect whether the connection is available
func (c *DirectClient) Ping() error {
	md := metadata.New(map[string]string{
		// "user": "root",
		// "token-x" : "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJleHAiOjE2OTAwNDg0MzYsImlhdCI6IjIwMjMtMDctMjJUMTU6NTM6NTYuMTA1OTY0OSswODowMCIsImtleSI6InlvY2tkX2tleSIsInN1YiI6InlvY2sifQ",
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	_, err := c.cli.Ping(ctx, &pb.PingRequest{})
	return err
}

// Wait is used to request signal from the daemon
func (c *DirectClient) SignalWait(sig string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.cli.SignalWait(ctx, &pb.WaitRequest{Sig: sig})
	if err != nil {
		ycho.Error(err)
	}
	return res.GetOk(), err
}

// Notify pushes signal to Daemon
func (c *DirectClient) SignalNotify(sig string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := c.cli.SignalNotify(ctx, &pb.NotifyRequest{Sig: sig})
	return err
}

func (c *DirectClient) SignalInfo(sig string) (bool, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.cli.SignalInfo(ctx, &pb.SignalInfoRequest{Sig: sig})
	return res.GetStatus(), res.GetExist(), err
}

func (c *DirectClient) SignalList() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.cli.SignalList(ctx, &pb.SignalListRequest{})
	return res.GetSigs(), err
}

func (c *DirectClient) SignalClear(sigs ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := c.cli.SignalClear(ctx, &pb.SignalClearRequest{Sigs: sigs})
	return err
}

// Upload pushes file information to peers so that peers can download files
func (c *DirectClient) Upload(file string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	out, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}
	info, err := fp.Stat()
	if err != nil {
		return err
	}
	_, err = c.cli.Upload(ctx, &pb.UploadRequest{
		Filename: file,
		Hash:     util.SHA256(string(out)),
		Owner:    du.ID,
		Size:     info.Size(),
		CreateAt: info.ModTime().Format(time.RFC3339),
	})
	return err
}

// Download file in other peer
func (c *DirectClient) FileSystemDownload(file string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := c.cli.FileSystemDownload(ctx)
	if err != nil {
		return err
	}
	if err = stream.Send(&pb.FileSystemDownloadRequest{Filename: file, Sender: du.ID}); err != nil {
		return err
	}
	data := []byte{}
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// breakpoint
			util.WriteFile(util.Pathf("@/tmp/"+file), data)
			return err
		}
		data = append(data, chunk.Data...)
	}
	return nil
}

// Register tells the daemon the address of the peer.
func (c *DirectClient) Register(addrs ...string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.cli.Register(ctx, &pb.RegisterRequest{Addrs: addrs})
	return res.GetAddrs(), err
}

// Unregister tells the daemon to remove the peer according to addrs.
func (c *DirectClient) Unregister(addrs ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := c.cli.Unregister(ctx, &pb.UnregisterRequest{Addrs: addrs})
	return err
}

// Info can obtain the meta information of the target node,
// including CPU, DISK, MEM and so on.
// You can specify it by InfoRequest, and by default only basic parameters
// (the name of the node, the file uploaded, and the connection information) are returned.
func (c *DirectClient) Info() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	res, err := c.cli.Info(ctx, &pb.InfoRequest{})
	return res.GetName(), err
}

func (c *DirectClient) FileSystemPut(src, dst string) error {
	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cannel()
	path, err := filepath.Abs(src)
	if err != nil {
		return err
	}
	_, err = c.cli.FileSystemPut(ctx, &pb.FileSystemPutRequest{
		Src: path,
		Dst: dst,
	})
	return err
}

func (c *DirectClient) FileSystemGet(src, dst string) error {
	return nil
}

func (c *DirectClient) Dial(from, to *pb.NodeInfo) error {
	ctx, cannel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cannel()
	_, err := c.cli.Dial(ctx, &pb.DialRequest{
		From: from,
		To:   to,
	})
	return err
}

func (c *DirectClient) Call(node, method string, args ...string) (string, error) {
	ctx, cannel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cannel()
	res, err := c.cli.Call(ctx, &pb.CallRequest{
		Node:   node,
		Method: method,
		Args:   args,
	})
	return res.GetRet(), err
}

func tunProtocal(p RelayProtocal) *pb.TunnelRequest {
	return &pb.TunnelRequest{
		Type: p.Type(),
		Body: p.String(),
		// Id:  ,
	}
}

func (c *DirectClient) IsPublic() bool {
	return true
}

func (c *DirectClient) MakeTunnel(name string, ctx context.Context, promise yocki.Promise, event chan yocki.PromiseEvent) error {
	stream, err := c.cli.Tunnel(ctx)
	if err != nil {
		return err
	}
	defer stream.CloseSend()
	ycho.Infof("try to make tunnel via %s", name)
	localhost := NewDirect(&YockdClientOption{
		IP:   c.opt.Global.Grpc.Addr.IP,
		Port: int(c.opt.Global.Grpc.Addr.Port),
	})
	for {
		if err = stream.Send(tunProtocal(EstablishProtocal{
			Name: name,
		})); err != nil {
			return err
		}
		heartbeat := make(chan bool)
		go func() {
			for {
				select {
				case e := <-event:
					msg := tunProtocal(e.Proto())
					msg.Id = e.Id()
					ycho.Infof("[%d] call %s", e.Id(), msg.String())
					err := stream.Send(msg)
					if err != nil {
						ycho.Error(err)
					}
				case ok := <-heartbeat:
					if !ok {
						return
					}
				case <-ctx.Done():
					return
				default:
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					continue
				}
				ycho.Error(err)
				if yocki.Y_MODE.Strict() {
					heartbeat <- false
					return err
				}
				break
			}
			switch res.GetType() {
			case pb.ProtocalType_MethodCall:
				var p MethodCallProtocal
				json.Unmarshal([]byte(res.GetBody()), &p)
				ycho.Infof("[%d] call %s.%s", res.GetId(), p.Node, p.Method)
				switch p.Method {
				case "info":
					info, err := localhost.Info()
					if err != nil {
						ycho.Error(err)
					}
					stream.Send(&pb.TunnelRequest{
						Type: pb.ProtocalType_MethodReturn,
						Body: info,
						Id:   res.GetId(),
					})
				}
			case pb.ProtocalType_MethodReturn:
				ycho.Infof("[%d] return ", res.GetId())
				promise.Store(res.GetId(), res.GetBody())
			}
		}
	}
}
