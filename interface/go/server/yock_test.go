package yocki

import (
	"testing"

	yocki "github.com/ansurfen/yock/interface/go"
)

func TestYocki(t *testing.T) {
	s := New()
	s.Register("SayHello", func(req *yocki.CallRequest) (*yocki.CallResponse, error) {
		return &yocki.CallResponse{Buf: "I'm Golang"}, nil
	})
	s.Run(9090)
}
