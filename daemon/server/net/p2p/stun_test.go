// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package p2p

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ansurfen/yock/daemon/server/net"
)

func TestStunServer(t *testing.T) {
	s := NewStun(
		OptionSetToken("root"),
		OptionEnableStunServer(&ynet.NetAddr{Port: 9090}))
	go func() {
		for {
			err := s.ReadFromClient()
			if err != nil {
				panic(err)
			}
			s.WriteToClient()
		}
	}()
	select {}
}

func TestStunClient(t *testing.T) {
	s := NewStun(
		OptionSetToken("root"),
		OptionEnableStunClient(&ynet.NetAddr{Port: 9091}, &ynet.NetAddr{IP: "localhost", Port: 9090}))
	s.WriteToServer()
	s.ReadFromServer()
	fmt.Println(s.Peer())
}

func TestStun(t *testing.T) {
	server := NewStun(
		OptionSetToken("root"),
		OptionEnableStunServer(&ynet.NetAddr{Port: 9090}))
	go server.Run(context.Background())

	client1 := NewStun(
		OptionSetToken("root"),
		OptionEnableStunClient(&ynet.NetAddr{Port: 9091}, &ynet.NetAddr{IP: "127.0.0.1", Port: 9090}))
	go client1.Run(context.Background())

	client2 := NewStun(
		OptionSetToken("root"),
		OptionEnableStunClient(&ynet.NetAddr{Port: 9092}, &ynet.NetAddr{IP: "127.0.0.1", Port: 9090}))
	go client2.Run(context.Background())

	time.Sleep(20 * time.Second)
	fmt.Println(client1.Peer(), client2.Peer())
	// todo 打通后，可以开启正常的tcp连接，那需要给个地址
}

func TestStunDiscover(t *testing.T) {

}
