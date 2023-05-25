package scheduler

import (
	"github.com/ansurfen/yock/daemon/interface/client"
	"github.com/ansurfen/yock/util"
)

var (
	_ SignalStream = &SingleSignalStream{}
	_ SignalStream = &CooperationSingalStream{}
)

type SignalStream interface {
	Load(sig string) (any, bool)
	Store(sig string, v bool)
}

type SingleSignalStream struct {
	sigs *util.SafeMap[bool]
}

func NewSingleSignalStream() *SingleSignalStream {
	return &SingleSignalStream{
		sigs: util.NewSafeMap[bool](),
	}
}

func (stream *SingleSignalStream) Load(sig string) (any, bool) {
	return stream.sigs.Get(sig)
}

func (stream *SingleSignalStream) Store(sig string, v bool) {
	stream.sigs.SafeSet(sig, v)
}

type CooperationSingalStream struct {
	*SingleSignalStream
	cli *client.YockDaemonClient
}

func NewCooperationSingalStream() *CooperationSingalStream {
	return &CooperationSingalStream{
		SingleSignalStream: NewSingleSignalStream(),
		cli:                client.New(client.Gopt),
	}
}

func upgradeSingalStream(stream *SingleSignalStream) *CooperationSingalStream {
	return &CooperationSingalStream{
		SingleSignalStream: stream,
		cli:                client.New(client.Gopt),
	}
}

func (stream *CooperationSingalStream) Load(sig string) (any, bool) {
	v, ok := stream.sigs.Get(sig)
	if !ok {
		v, _ = stream.cli.Wait(sig)
		if v {
			stream.sigs.SafeSet(sig, v)
		}
	}
	return v, true
}

func (stream *CooperationSingalStream) Store(sig string, v bool) {
	stream.SingleSignalStream.Store(sig, v)
	stream.cli.Notify(sig)
}
