package yocki

import (
	"os"
	"testing"
	"time"

	yocki "github.com/ansurfen/yock/interface/go"
)

func TestYocki(t *testing.T) {
	go func() {
		time.Sleep(10 * time.Second)
		os.Exit(0)
	}()
	s := New()
	s.Register("SayHello", func(req *yocki.CallRequest) (*yocki.CallResponse, error) {
		return &yocki.CallResponse{Buf: "I'm Golang"}, nil
	})
	s.Run(9090)
}
