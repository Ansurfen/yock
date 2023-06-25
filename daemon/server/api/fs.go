// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"

	"github.com/ansurfen/yock/daemon/proto"
	"github.com/ansurfen/yock/daemon/server/fs"
	"github.com/ansurfen/yock/util"
)

// Upload pushes file information to peers so that peers can download files
func (yockd *YockDaemon) Upload(ctx context.Context, req *proto.UploadRequest) (*proto.UploadResponse, error) {
	volume := ""
	yockd.kernel.FileSystem.CreateFile(volume, fs.FileInfo{
		Owner:    req.Owner,
		Size:     req.Size,
		Hash:     req.Hash,
		CreateAt: req.CreateAt,
	})
	return &proto.UploadResponse{}, nil
}

// Download file in other peer
func (yockd *YockDaemon) Download(stream proto.YockDaemon_DownloadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	file, ok := yockd.kernel.FileSystem.FindFile("", req.Filename)
	if !ok {
		return util.ErrFileNotExist
	}
	if file.Owner == yockd.conf.Name {
		if req.Sender == file.Owner {
			return nil
		} else {
			raw, err := util.ReadStraemFromFile(util.Pathf("@/tmp/" + req.Filename))
			if err != nil {
				return err
			}
			for i := 0; i < len(raw); i++ {
				chunk := raw[i : i+yockd.conf.Fs.MTL]
				if err = stream.Send(&proto.DownloadResponse{Data: chunk}); err != nil {
					return err
				}
			}
		}
	}
	// if node, ok := daemon.nodeManager.Nodes()[file.Owner]; ok {
	// 	return node.cli.Download(req.Filename)
	// } else { // boardcast to every node
	// 	for _, n := range daemon.nodeManager.Nodes() {
	// 		if n.cli.Download(req.Filename) == nil {
	// 			break
	// 		}
	// 	}
	// }
	return nil
}
