// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"github.com/ansurfen/yock/daemon/net"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
)

var (
	_ yocki.SignalStream = (*SingleSignalStream)(nil)
	_ yocki.SignalStream = (*CooperationSingalStream)(nil)
)

// SingleSignalStream is a single-process implementation of SignalStream,
// where signals only flow in the process.
type SingleSignalStream struct {
	sigs *util.SafeMap[bool]
}

func NewSingleSignalStream() *SingleSignalStream {
	return &SingleSignalStream{
		sigs: util.NewSafeMap[bool](),
	}
}

// Load returns the value of the specified singal.
// If the singal isn't exist, the second parameter returns false, and vice versa.
func (stream *SingleSignalStream) Load(sig string) (any, bool) {
	return stream.sigs.Get(sig)
}

// Store settings specify the value of the singal, similar to map's kv storage.
func (stream *SingleSignalStream) Store(sig string, v bool) {
	stream.sigs.SafeSet(sig, v)
}

// CooperationSingalStream is a distributed implementation of SignalStream,
// using grpc + protobuf to transmit signals.
type CooperationSingalStream struct {
	*SingleSignalStream
	cli yocki.YockdClient
}

func NewCooperationSingalStream(c *net.DirectClient) *CooperationSingalStream {
	return &CooperationSingalStream{
		SingleSignalStream: NewSingleSignalStream(),
		cli:                c,
	}
}

// upgradeSingalStream upgrades SingleSignalStream to CooperationSingalStream to meet distributed needs.
func upgradeSingalStream(stream yocki.SignalStream, c yocki.YockdClient) yocki.SignalStream {
	return &CooperationSingalStream{
		cli:                c,
		SingleSignalStream: stream.(*SingleSignalStream),
	}
}

// Load returns the value of the specified singal.
// If the singal isn't exist, the second parameter returns false, and vice versa.
// In CooperationSingalStream, each load will send a request to daemon to ask for the signal status,
// and set the value if it exists.
func (stream *CooperationSingalStream) Load(sig string) (any, bool) {
	v, ok := stream.sigs.Get(sig)
	if !ok {
		v, _ = stream.cli.SignalWait(sig)
		if v {
			stream.sigs.SafeSet(sig, v)
		}
	}
	return v, true
}

// Store settings specify the value of the singal, similar to map's kv storage and send it to daemon.
func (stream *CooperationSingalStream) Store(sig string, v bool) {
	stream.SingleSignalStream.Store(sig, v)
	stream.cli.SignalNotify(sig)
}
