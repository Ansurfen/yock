package client

import (
	"context"
	"fmt"
	"time"

	yocki "github.com/ansurfen/yock/interface/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type YockInterface struct {
	conn *grpc.ClientConn
	cli  yocki.YockInterfaceClient
}

func New(ip string, port int) *YockInterface {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &YockInterface{
		conn: conn,
		cli:  yocki.NewYockInterfaceClient(conn),
	}
}

func (yock *YockInterface) Close() {
	yock.conn.Close()
}

func (yock *YockInterface) Call(fn, arg string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := yock.cli.Call(ctx, &yocki.CallRequest{Fn: fn, Arg: arg})
	return res.Buf, err
}
