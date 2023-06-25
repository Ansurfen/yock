// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// YockDaemonClient is the client API for YockDaemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YockDaemonClient interface {
	// Ping is used to detect whether the connection is available
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Wait is used to request signal from the daemon
	Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*WaitResponse, error)
	// Notify pushes signal to Daemon
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	// Upload pushes file information to peers so that peers can download files
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	// Download file in other peer
	Download(ctx context.Context, opts ...grpc.CallOption) (YockDaemon_DownloadClient, error)
	// Register tells the daemon the address of the peer.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Unregister tells the daemon to remove the peer according to addrs.
	Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error)
	// Info can obtain the meta information of the target node,
	// including CPU, DISK, MEM and so on.
	// You can specify it by InfoRequest, and by default only basic parameters
	// (the name of the node, the file uploaded, and the connection information) are returned.
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type yockDaemonClient struct {
	cc grpc.ClientConnInterface
}

func NewYockDaemonClient(cc grpc.ClientConnInterface) YockDaemonClient {
	return &yockDaemonClient{cc}
}

func (c *yockDaemonClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/Yockd.YockDaemon/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yockDaemonClient) Wait(ctx context.Context, in *WaitRequest, opts ...grpc.CallOption) (*WaitResponse, error) {
	out := new(WaitResponse)
	err := c.cc.Invoke(ctx, "/Yockd.YockDaemon/Wait", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yockDaemonClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/Yockd.YockDaemon/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yockDaemonClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, "/Yockd.YockDaemon/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yockDaemonClient) Download(ctx context.Context, opts ...grpc.CallOption) (YockDaemon_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_YockDaemon_serviceDesc.Streams[0], "/Yockd.YockDaemon/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &yockDaemonDownloadClient{stream}
	return x, nil
}

type YockDaemon_DownloadClient interface {
	Send(*DownloadRequest) error
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type yockDaemonDownloadClient struct {
	grpc.ClientStream
}

func (x *yockDaemonDownloadClient) Send(m *DownloadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yockDaemonDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *yockDaemonClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/Yockd.YockDaemon/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yockDaemonClient) Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error) {
	out := new(UnregisterResponse)
	err := c.cc.Invoke(ctx, "/Yockd.YockDaemon/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yockDaemonClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/Yockd.YockDaemon/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YockDaemonServer is the server API for YockDaemon service.
// All implementations must embed UnimplementedYockDaemonServer
// for forward compatibility
type YockDaemonServer interface {
	// Ping is used to detect whether the connection is available
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Wait is used to request signal from the daemon
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
	// Notify pushes signal to Daemon
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
	// Upload pushes file information to peers so that peers can download files
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	// Download file in other peer
	Download(YockDaemon_DownloadServer) error
	// Register tells the daemon the address of the peer.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Unregister tells the daemon to remove the peer according to addrs.
	Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error)
	// Info can obtain the meta information of the target node,
	// including CPU, DISK, MEM and so on.
	// You can specify it by InfoRequest, and by default only basic parameters
	// (the name of the node, the file uploaded, and the connection information) are returned.
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	mustEmbedUnimplementedYockDaemonServer()
}

// UnimplementedYockDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedYockDaemonServer struct {
}

func (*UnimplementedYockDaemonServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedYockDaemonServer) Wait(context.Context, *WaitRequest) (*WaitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wait not implemented")
}
func (*UnimplementedYockDaemonServer) Notify(context.Context, *NotifyRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedYockDaemonServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedYockDaemonServer) Download(YockDaemon_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (*UnimplementedYockDaemonServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedYockDaemonServer) Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (*UnimplementedYockDaemonServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedYockDaemonServer) mustEmbedUnimplementedYockDaemonServer() {}

func RegisterYockDaemonServer(s *grpc.Server, srv YockDaemonServer) {
	s.RegisterService(&_YockDaemon_serviceDesc, srv)
}

func _YockDaemon_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YockDaemonServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Yockd.YockDaemon/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YockDaemonServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YockDaemon_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YockDaemonServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Yockd.YockDaemon/Wait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YockDaemonServer).Wait(ctx, req.(*WaitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YockDaemon_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YockDaemonServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Yockd.YockDaemon/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YockDaemonServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YockDaemon_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YockDaemonServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Yockd.YockDaemon/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YockDaemonServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YockDaemon_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YockDaemonServer).Download(&yockDaemonDownloadServer{stream})
}

type YockDaemon_DownloadServer interface {
	Send(*DownloadResponse) error
	Recv() (*DownloadRequest, error)
	grpc.ServerStream
}

type yockDaemonDownloadServer struct {
	grpc.ServerStream
}

func (x *yockDaemonDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yockDaemonDownloadServer) Recv() (*DownloadRequest, error) {
	m := new(DownloadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _YockDaemon_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YockDaemonServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Yockd.YockDaemon/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YockDaemonServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YockDaemon_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YockDaemonServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Yockd.YockDaemon/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YockDaemonServer).Unregister(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YockDaemon_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YockDaemonServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Yockd.YockDaemon/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YockDaemonServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _YockDaemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Yockd.YockDaemon",
	HandlerType: (*YockDaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _YockDaemon_Ping_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _YockDaemon_Wait_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _YockDaemon_Notify_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _YockDaemon_Upload_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _YockDaemon_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _YockDaemon_Unregister_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _YockDaemon_Info_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _YockDaemon_Download_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "yockd.proto",
}