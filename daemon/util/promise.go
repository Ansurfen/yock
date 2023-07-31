// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"time"

	yocki "github.com/ansurfen/yock/interface"
	"github.com/bwmarrin/snowflake"
)

var _ yocki.Promise = (*Promise)(nil)

type Promise struct {
	ctx  map[int64]any
	node *snowflake.Node
}

func NewPromise() *Promise {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	return &Promise{
		ctx:  make(map[int64]any),
		node: node,
	}
}

func (p *Promise) LoadWithTimeout(id int64, timeout time.Duration) (any, bool) {
	t := 100 * time.Microsecond
	time.Sleep(t)
	v, ok := p.Load(id)
	if ok {
		return v, true
	}
	for t < time.Duration(timeout) {
		t = t << 2
		time.Sleep(t)
		v, ok = p.Load(id)
		if ok {
			return v, true
		}
	}
	return "", false
}

func (p *Promise) Store(id int64, v any) {
	p.ctx[id] = v
}

func (p *Promise) Load(id int64) (any, bool) {
	if v, ok := p.ctx[id]; ok {
		delete(p.ctx, id)
		return v, true
	}
	return nil, false
}

func (p *Promise) NextID() int64 {
	return p.node.Generate().Int64()
}

var _ yocki.PromiseEvent = (*PromiseEvent)(nil)

type PromiseEvent struct {
	id    int64
	proto yocki.Protocal
}

func PostPromiseEvent(id int64, proto yocki.Protocal) PromiseEvent {
	return PromiseEvent{id: id, proto: proto}
}

func (e PromiseEvent) Id() int64 {
	return e.id
}

func (e PromiseEvent) Proto() yocki.Protocal {
	return e.proto
}
