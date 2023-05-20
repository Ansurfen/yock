package yocki

import (
	"context"

	yocki "github.com/ansurfen/yock/interface/go"
)

var _ yocki.YockInterfaceServer = &YockInterface{}

type YockInterface struct {
	yocki.UnimplementedYockInterfaceServer
	reg *Registry
}

func New() *YockInterface {
	return &YockInterface{}
}

func (yock *YockInterface) Ping(ctx context.Context, req *yocki.PingRequest) (*yocki.PingResponse, error) {
	return &yocki.PingResponse{}, nil
}

func (yock *YockInterface) Call(ctx context.Context, req *yocki.CallRequest) (*yocki.CallResponse, error) {
	return yock.reg.Find(req.Fn)(req)
}

func (yock *YockInterface) Register(fn string, call YockCall) {
	yock.reg.register(fn, call)
}

func (yock *YockInterface) Unregister(fn string) {
	yock.reg.unregister(fn)
}
