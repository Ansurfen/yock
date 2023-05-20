package server

import (
	"context"
	"fmt"
	"net"

	"github.com/ansurfen/cushion/utils"
	yocki "github.com/ansurfen/yock/daemon/interface"
	"github.com/ansurfen/yock/util"
	"google.golang.org/grpc"
)

type YockDaemon struct {
	yocki.UnimplementedYockInterfaceServer
	signals     *util.SafeMap[bool]
	fs          map[string]FileInfo
	nodeManager *NodeManager
	opt         *DaemonOption
}

func New(opt *DaemonOption) *YockDaemon {
	return &YockDaemon{
		signals:     util.NewSafeMap[bool](),
		opt:         opt,
		fs:          make(map[string]FileInfo),
		nodeManager: NewNodeManager(),
	}
}

func (daemon *YockDaemon) Close() {
	for _, node := range daemon.nodeManager.Nodes() {
		node.cli.Close()
	}
}

// Ping is used to detect whether the connection is available
func (daemon *YockDaemon) Ping(ctx context.Context, req *yocki.PingRequest) (*yocki.PingResponse, error) {
	return &yocki.PingResponse{}, nil
}

// Wait is used to request signal from the daemon
func (daemon *YockDaemon) Wait(ctx context.Context, req *yocki.WaitRequest) (*yocki.WaitResponse, error) {
	if v, ok := daemon.signals.Get(req.Sig); !ok {
		daemon.signals.SafeSet(req.Sig, false)
		return &yocki.WaitResponse{Ok: false}, nil
	} else if ok && v {
		return &yocki.WaitResponse{Ok: true}, nil
	}
	return &yocki.WaitResponse{Ok: false}, nil
}

// Notify pushes signal to Daemon
func (daemon *YockDaemon) Notify(ctx context.Context, req *yocki.NotifyRequest) (*yocki.NotifyResponse, error) {
	daemon.signals.SafeSet(req.Sig, true)
	return &yocki.NotifyResponse{}, nil
}

// Upload pushes file information to peers so that peers can download files
func (daemon *YockDaemon) Upload(ctx context.Context, req *yocki.UploadRequest) (*yocki.UploadResponse, error) {
	daemon.fs[req.Filename] = FileInfo{
		owner:    req.Owner,
		size:     req.Size,
		hash:     req.Hash,
		createAt: req.CreateAt,
	}
	return &yocki.UploadResponse{}, nil
}

// Download file in other peer
func (daemon *YockDaemon) Download(stream yocki.YockInterface_DownloadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	file, ok := daemon.fs[req.Filename]
	if !ok {
		return util.ErrFileNotExist
	}
	if file.owner == *daemon.opt.Name {
		if req.Sender == file.owner {
			return nil
		} else {
			raw, err := utils.ReadStraemFromFile(util.Pathf("@/tmp/" + req.Filename))
			if err != nil {
				return err
			}
			for i := 0; i < len(raw); i++ {
				chunk := raw[i : i+*daemon.opt.MTL]
				if err = stream.Send(&yocki.DownloadResponse{Data: chunk}); err != nil {
					return err
				}
			}
		}
	}
	if node, ok := daemon.nodeManager.Nodes()[file.owner]; ok {
		return node.cli.Download(req.Filename)
	} else { // boardcast to every node
		for _, n := range daemon.nodeManager.Nodes() {
			if n.cli.Download(req.Filename) == nil {
				break
			}
		}
	}
	return nil
}

// Register tells the daemon the address of the peer.
func (daemon *YockDaemon) Register(ctx context.Context, req *yocki.RegisterRequest) (*yocki.RegisterResponse, error) {
	for _, addr := range req.Addrs {
		daemon.nodeManager.AddNode(addr)
	}
	return &yocki.RegisterResponse{}, nil
}

// Unregister tells the daemon to remove the peer according to addrs.
func (daemon *YockDaemon) Unregister(ctx context.Context, req *yocki.UnregisterRequest) (*yocki.UnregisterResponse, error) {
	for _, addr := range req.Addrs {
		daemon.nodeManager.DelNode(addr)
	}
	return &yocki.UnregisterResponse{}, nil
}

// Info can obtain the meta information of the target node,
// including CPU, DISK, MEM and so on.
// You can specify it by InfoRequest, and by default only basic parameters
// (the name of the node, the file uploaded, and the connection information) are returned.
func (daemon *YockDaemon) Info(ctx context.Context, req *yocki.InfoRequest) (*yocki.InfoResponse, error) {
	return &yocki.InfoResponse{Name: *daemon.opt.Name}, nil
}

func (daemon *YockDaemon) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *daemon.opt.Port))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer()
	yocki.RegisterYockInterfaceServer(gsrv, daemon)
	if err := gsrv.Serve(listen); err != nil {
		panic(err)
	}
}
