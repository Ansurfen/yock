// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"strings"

	pb "github.com/ansurfen/yock/daemon/proto"
)

func (daemon *YockDaemon) ProcessList(ctx context.Context, req *pb.ProcessListRequest) (*pb.ProcessListResponse, error) {
	res := []*pb.Process{}
	for _, p := range daemon.Scheduler.ProcessState() {
		res = append(res, &pb.Process{
			Pid:   p.Pid(),
			State: int32(p.State()),
			Spec:  p.Spec(),
			Cmd:   p.Cmd(),
		})
	}
	return &pb.ProcessListResponse{
		Res: res,
	}, nil
}

func (daemon *YockDaemon) ProcessKill(ctx context.Context, req *pb.ProcessKillRequest) (*pb.ProcessKillResponse, error) {
	daemon.Scheduler.Kill(req.GetPid())
	return &pb.ProcessKillResponse{}, nil
}

func (daemon *YockDaemon) ProcessFind(ctx context.Context, req *pb.ProcessFindRequest) (*pb.ProcessFindResponse, error) {
	res := []*pb.Process{}
	if id := req.GetPid(); id != 0 {
		p := daemon.Scheduler.FindByPID(id)
		res = append(res, &pb.Process{
			Pid:   p.Pid(),
			State: int32(p.State()),
			Spec:  p.Spec(),
			Cmd:   p.Cmd(),
		})
	} else if len(req.GetCmd()) > 0 {
		ps := daemon.Scheduler.FindByCmd(req.GetCmd())
		for _, p := range ps {
			res = append(res, &pb.Process{
				Pid:   p.Pid(),
				State: int32(p.State()),
				Spec:  p.Spec(),
				Cmd:   p.Cmd(),
			})
		}
	}
	return &pb.ProcessFindResponse{
		Res: res,
	}, nil
}

func (daemon *YockDaemon) ProcessSpawn(ctx context.Context, req *pb.ProcessSpawnRequest) (*pb.ProcessSpawnResponse, error) {
	var (
		pid int64
		err error
	)
	switch req.GetType() {
	case pb.ProcessSpawnType_Cron:
		pid, err = daemon.Scheduler.CreateCronTask(req.GetSpec(), req.GetCmd())

	case pb.ProcessSpawnType_FS:
		pid, err = daemon.Scheduler.CreateFSListenTask(strings.Split(req.GetSpec(), ";"), req.GetCmd())
	case pb.ProcessSpawnType_Script:

	}
	if err != nil {
		return &pb.ProcessSpawnResponse{
			Pid: 0,
		}, err
	}
	return &pb.ProcessSpawnResponse{
		Pid: pid,
	}, nil
}
