package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lyock
#include "yock.h"
*/
import "C"
import (
	context "context"
	"encoding/json"
	"fmt"
	"net"
	"unsafe"

	"github.com/ansurfen/yock/util"
	grpc "google.golang.org/grpc"
)

type YockInterface struct {
	UnimplementedYockInterfaceServer
	dict map[string]func(string) string
}

func newYockInterface() *YockInterface {
	return &YockInterface{
		dict: make(map[string]func(string) string),
	}
}

func (yock *YockInterface) Callback(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	if cb, ok := yock.dict[req.Fn]; ok {
		res := cb(util.JsonStr(util.NewJsonObject(map[string]util.JsonValue{
			"Fn":  util.NewJsonString(req.Fn),
			"Arg": util.NewJsonString(req.Arg),
		})))
		var ret CallResponse
		err := json.Unmarshal([]byte(res), &ret)
		return &ret, err
	}
	return &CallResponse{Buf: "unknown"}, nil
}

func (yock *YockInterface) Register(name string, cb func(string) string) {
	yock.dict[name] = cb
}

//export newYock
func newYock() *C.Yock {
	s := newYockInterface()
	ret := (*C.Yock)(C.malloc(C.size_t(unsafe.Sizeof(C.Yock{}))))
	ret.ptr = unsafe.Pointer(s)
	return ret
}

//export yockRegisterCall
func yockRegisterCall(yock *C.Yock, name *C.char, cb C.Call) {
	s := CastPtr[YockInterface](yock.ptr)
	s.Register(C.GoString(name), func(s string) string {
		str := C.yockCall(cb, C.CString(s))
		return C.GoString(str)
	})
}

//export yockRun
func yockRun(yock *C.Yock, port *C.char) {
	s := CastPtr[YockInterface](yock.ptr)
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", C.GoString(port)))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer()
	RegisterYockInterfaceServer(gsrv, s)
	if err := gsrv.Serve(listen); err != nil {
		panic(err)
	}
}

func CastPtr[T any](ptr unsafe.Pointer) *T {
	return (*T)(ptr)
}

func main() {}
