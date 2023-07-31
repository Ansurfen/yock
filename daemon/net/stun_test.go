// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestStunServer(t *testing.T) {
	s := NewStun(
		OptionSetToken("root"),
		OptionEnableStunServer(&NetAddr{Port: 9090}))
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
		OptionEnableStunClient(&NetAddr{Port: 9091}, &NetAddr{IP: "localhost", Port: 9090}))
	s.WriteToServer()
	s.ReadFromServer()
	fmt.Println(s.Peer())
}

func TestStun(t *testing.T) {
	server := NewStun(
		OptionSetToken("root"),
		OptionEnableStunServer(&NetAddr{Port: 9090}))
	go server.Run(context.Background())

	client1 := NewStun(
		OptionSetToken("root"),
		OptionEnableStunClient(&NetAddr{Port: 9091}, &NetAddr{IP: "127.0.0.1", Port: 9090}))
	go client1.Run(context.Background())

	client2 := NewStun(
		OptionSetToken("root"),
		OptionEnableStunClient(&NetAddr{Port: 9092}, &NetAddr{IP: "127.0.0.1", Port: 9090}))
	go client2.Run(context.Background())

	time.Sleep(20 * time.Second)
	fmt.Println(client1.Peer(), client2.Peer())
}

func TestStunDiscover(t *testing.T) {
	s := NewStun()
	s.opt.RetryCnt = 10
	addr, err := s.Discover()
	if err != nil {
		panic(err)
	}
	fmt.Println(addr.CanMakeHole(), addr.natType)
}
