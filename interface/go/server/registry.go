package yocki

import yocki "github.com/ansurfen/yock/interface/go"

type YockCall func(req *yocki.CallRequest) (*yocki.CallResponse, error)

var UnimplementYockCall = func(*yocki.CallRequest) (*yocki.CallResponse, error) {
	return &yocki.CallResponse{Buf: "unimplement method or bad request"}, nil
}

type Registry struct {
	dict map[string]YockCall
}

func (reg *Registry) register(fn string, call YockCall) {
	reg.dict[fn] = call
}

func (reg *Registry) unregister(fn string) {
	delete(reg.dict, fn)
}

func (reg *Registry) Find(name string) YockCall {
	if call, ok := reg.dict[name]; ok {
		return call
	}
	return UnimplementYockCall
}
