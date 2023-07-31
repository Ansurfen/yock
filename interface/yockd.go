// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocki

import (
	"context"
	"time"

	pb "github.com/ansurfen/yock/daemon/proto"
)

type Protocal interface {
	Type() pb.ProtocalType
	String() string
}

type Promise interface {
	Load(id int64) (any, bool)
	LoadWithTimeout(id int64, timeout time.Duration) (any, bool)
	Store(id int64, v any)
	// NextID returns the unique snowflake id to ensure concurrency security
	NextID() int64
}

type PromiseEvent interface {
	Id() int64
	Proto() Protocal
}

type YockdClient interface {
	YockdClientSignal
	YockdClientFS
	YockdClientGateway
	YockdClientMeta
	YockdClientNet
	YockdClientProcess
}

type YockdClientFS interface {
	FileSystemPut(src, dst string) error
	FileSystemGet(src, dst string) error
}

type YockdClientProcess interface {
	ProcessList() ([]*pb.Process, error)
	ProcessKill(pid int64) error
	ProcessFind(pid int64, cmd string) ([]*pb.Process, error)
	ProcessSpawn(t pb.ProcessSpawnType, spec, cmd string) (int64, error)
}

type YockdClientGateway interface{}

type YockdClientMeta interface {
	IsPublic() bool
	Name() string
	Info() (string, error)
	Status()
	Close()
}

type YockdClientNet interface {
	Ping() error
	Mark(name, addr string) error
	Dial(form, to *pb.NodeInfo) error
	Call(node, method string, args ...string) (string, error)
	MakeTunnel(name string, ctx context.Context, p Promise, event chan PromiseEvent) error
}

type YockdClientSignal interface {
	SignalNotify(sig string) error
	SignalWait(sig string) (bool, error)
	SignalInfo(sig string) (bool, bool, error)
	SignalList() ([]string, error)
	SignalClear(sigs ...string) error
}
