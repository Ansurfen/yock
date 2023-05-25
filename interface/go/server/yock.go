package yocki

import (
	"context"
	"fmt"
	"net"

	yocki "github.com/ansurfen/yock/interface/go"
	"google.golang.org/grpc"
)

var _ yocki.YockInterfaceServer = &YockInterface{}

type YockInterface struct {
	yocki.UnimplementedYockInterfaceServer
	reg *Registry
}

func New() *YockInterface {
	return &YockInterface{
		reg: &Registry{
			dict: make(map[string]YockCall),
		},
	}
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

func (yock *YockInterface) Run(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer()
	yocki.RegisterYockInterfaceServer(gsrv, yock)
	if err := gsrv.Serve(listen); err != nil {
		panic(err)
	}
}
