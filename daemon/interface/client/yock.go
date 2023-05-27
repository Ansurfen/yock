// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/ansurfen/cushion/utils"
	yocki "github.com/ansurfen/yock/daemon/interface"
	"github.com/ansurfen/yock/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type YockDaemonClient struct {
	conn *grpc.ClientConn
	cli  yocki.YockInterfaceClient
	opt  *DaemonOption
}

func New(opt *DaemonOption) *YockDaemonClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *opt.IP, *opt.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &YockDaemonClient{
		conn: conn,
		cli:  yocki.NewYockInterfaceClient(conn),
		opt:  opt,
	}
}

func (client *YockDaemonClient) Close() {
	client.conn.Close()
}

// Ping is used to detect whether the connection is available
func (c *YockDaemonClient) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.cli.Ping(ctx, &yocki.PingRequest{})
	return err
}

// Wait is used to request signal from the daemon
func (c *YockDaemonClient) Wait(sig string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.cli.Wait(ctx, &yocki.WaitRequest{Sig: sig})
	return res.Ok, err
}

// Notify pushes signal to Daemon
func (c *YockDaemonClient) Notify(sig string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.cli.Notify(ctx, &yocki.NotifyRequest{Sig: sig})
	return err
}

// Upload pushes file information to peers so that peers can download files
func (c *YockDaemonClient) Upload(file string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
	_, err = c.cli.Upload(ctx, &yocki.UploadRequest{
		Filename: file,
		Hash:     utils.SHA256(string(out)),
		Owner:    *c.opt.Name,
		Size:     info.Size(),
		CreateAt: info.ModTime().Format(time.RFC3339),
	})
	return err
}

// Download file in other peer
func (c *YockDaemonClient) Download(file string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.cli.Download(ctx)
	if err != nil {
		return err
	}
	if err = stream.Send(&yocki.DownloadRequest{Filename: file, Sender: *c.opt.Name}); err != nil {
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
			utils.WriteFile(util.Pathf("@/tmp/"+file), data)
			return err
		}
		data = append(data, chunk.Data...)
	}
	return nil
}

// Register tells the daemon the address of the peer.
func (c *YockDaemonClient) Register(addrs ...string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.cli.Register(ctx, &yocki.RegisterRequest{Addrs: addrs})
	return res.GetAddrs(), err
}

// Unregister tells the daemon to remove the peer according to addrs.
func (c *YockDaemonClient) Unregister(addrs ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.cli.Unregister(ctx, &yocki.UnregisterRequest{Addrs: addrs})
	return err
}

// Info can obtain the meta information of the target node,
// including CPU, DISK, MEM and so on.
// You can specify it by InfoRequest, and by default only basic parameters
// (the name of the node, the file uploaded, and the connection information) are returned.
func (c *YockDaemonClient) Info() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.cli.Info(ctx, &yocki.InfoRequest{})
	return res.GetPayload(), err
}
