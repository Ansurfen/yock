// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kernel

import (
	"sync"

	du "github.com/ansurfen/yock/daemon/util"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
)

type SignalStream struct {
	userSignals *util.SafeMap[bool]
	mut         *sync.Mutex

	sysSignals *du.Promise
	event      chan yocki.PromiseEvent
}

func newSingalStream() *SignalStream {
	return &SignalStream{
		userSignals: util.NewSafeMap[bool](),
		mut:         &sync.Mutex{},
		event:       make(chan yocki.PromiseEvent),
		sysSignals:  du.NewPromise(),
	}
}

func (stream *SignalStream) Wait(sig string) bool {
	v, ok := stream.userSignals.Get(sig)
	if !ok {
		stream.userSignals.SafeSet(sig, false)
	}
	return ok && v
}

func (stream *SignalStream) Notify(sig string) {
	stream.userSignals.SafeSet(sig, true)
}

func (stream *SignalStream) Info(sig string) (bool, bool) {
	return stream.userSignals.Get(sig)
}

func (stream *SignalStream) Clear(sigs ...string) {
	stream.mut.Lock()
	defer stream.mut.Unlock()
	for _, sig := range sigs {
		stream.userSignals.Delete(sig)
	}
}

func (stream *SignalStream) List() []string {
	return stream.userSignals.Keys()
}

func (stream *SignalStream) System() yocki.Promise {
	return stream.sysSignals
}

func (stream *SignalStream) User() *util.SafeMap[bool] {
	return stream.userSignals
}

func (stream *SignalStream) SystemEvent() chan yocki.PromiseEvent {
	return stream.event
}
