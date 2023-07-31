// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/ansurfen/yock/daemon/conf"
	"github.com/ansurfen/yock/util"
	"github.com/ccding/go-stun/stun"
)

const defaultBufferSize = 1024

type UDPServer struct {
	conn     *net.UDPConn
	addr     *NetAddr
	protocal Protocal
}

func ListenUDP(addr *NetAddr, protocal Protocal) *UDPServer {
	conn, err := net.ListenUDP("udp", addr.LocalV4UDPAddr())
	if err != nil {
		panic(err)
	}
	return &UDPServer{
		conn:     conn,
		addr:     addr,
		protocal: protocal,
	}
}

func (s *UDPServer) Read() (Protocal, Context, error) {
	data := make([]byte, 1024)
	n, addr, err := s.conn.ReadFromUDP(data)
	if err != nil {
		return nil, Context{}, err
	}
	return s.protocal.Parse(data[:n]),
		Context{addr: UDPAddr2NetAddr(addr)}, nil
}

func (s *UDPServer) Write(p Protocal, addr *NetAddr) error {
	return s.WriteRaw(p.Bytes(), addr)
}

func (s *UDPServer) WriteRaw(b []byte, addr *NetAddr) error {
	_, err := s.conn.WriteToUDP(b, addr.UDPAddr())
	return err
}

type UDPClient struct {
	conn     *net.UDPConn
	addr     *NetAddr
	protocal Protocal
}

func DialUDP(laddr, raddr *NetAddr, protocal Protocal) *UDPClient {
	conn, err := net.DialUDP("udp", laddr.LocalV4UDPAddr(), raddr.UDPAddr())
	if err != nil {
		panic(err)
	}
	return &UDPClient{
		conn:     conn,
		addr:     laddr,
		protocal: protocal,
	}
}

func (cli *UDPClient) Write(p Protocal) error {
	return cli.WriteRaw(p.Bytes())
}

func (cli *UDPClient) WriteRaw(b []byte) error {
	_, err := cli.conn.Write(b)
	return err
}

func (cli *UDPClient) Read() (Protocal, Context, error) {
	data := make([]byte, defaultBufferSize)
	n, addr, err := cli.conn.ReadFromUDP(data)
	if err != nil {
		return nil, Context{}, err
	}
	return cli.protocal.Parse(data[:n]),
		Context{addr: UDPAddr2NetAddr(addr)}, nil
}

type StunProtocal struct {
	Hash  string      `json:"hash"`
	Ver   uint32      `json:"version"`
	Peers []*PeerAddr `json:"peer"`
	ID    string      `json:"id"`
}

func (p *StunProtocal) Parse(message []byte) Protocal {
	protocal := &StunProtocal{}
	if err := json.Unmarshal(message, protocal); err != nil {
		panic(err)
	}
	return protocal
}

func (p *StunProtocal) Version() string {
	return fmt.Sprintf("v%d", p.Ver)
}

func (p *StunProtocal) Bytes() []byte {
	raw, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return raw
}

const defaultStunVersion = 1

type Stun struct {
	server  *UDPServer
	client  *UDPClient
	stunCli *stun.Client
	ssp     []*PeerAddr
	csp     []*PeerAddr
	token   string
	urls    util.LoadBalancer[string]
	opt     conf.YockConfNetStun
}

func (s *Stun) WriteToClient() {
	p := &StunProtocal{
		Hash:  s.token,
		Ver:   defaultStunVersion,
		Peers: s.ssp,
	}
	raw := p.Bytes()
	for _, peer := range s.ssp {
		s.server.WriteRaw(raw, &peer.NetAddr)
	}
}

func (s *Stun) ReadFromClient() error {
	p, ctx, err := s.server.Read()
	if err != nil {
		return err
	}
	sp := p.(*StunProtocal)
	if sp.Hash != s.token {
		return nil
	}
	s.ssp = append(s.ssp, &PeerAddr{
		NetAddr: *ctx.addr,
		ID:      sp.ID,
	})
	return nil
}

func (s *Stun) ReadFromServer() {
	p, _, err := s.client.Read()
	if err != nil {
		panic(err)
	}
	sp := p.(*StunProtocal)
	s.csp = append(s.csp, sp.Peers...)
}

func (s *Stun) WriteToServer() {
	p := &StunProtocal{
		Hash: s.token,
		Ver:  defaultStunVersion,
	}
	if err := s.client.Write(p); err != nil {
		panic(err)
	}
}

func (s *Stun) Peer() []*PeerAddr {
	return s.csp
}

type StunOption func(*Stun) error

func OptionEnableStunServer(addr *NetAddr) StunOption {
	return func(s *Stun) error {
		s.server = ListenUDP(addr, &StunProtocal{})
		return nil
	}
}

func OptionEnableStunClient(laddr, raddr *NetAddr) StunOption {
	return func(s *Stun) error {
		s.client = DialUDP(laddr, raddr, &StunProtocal{})
		return nil
	}
}

func OptionSetToken(token string) StunOption {
	return func(s *Stun) error {
		s.token = token
		return nil
	}
}

func NewStun(opts ...StunOption) *Stun {
	var stunService []any
	raw, err := util.ReadStraemFromFile("stun.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(raw, &stunService)
	if err != nil {
		panic(err)
	}
	stuns := []string{}
	for _, s := range stunService {
		stuns = append(stuns, s.(string))
	}
	s := &Stun{
		urls:    util.NewWeightedRandom(stuns),
		stunCli: stun.NewClient(),
	}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			panic(err)
		}
	}
	return s
}

func (s *Stun) Discover() (addr PeerAddr, e error) {
	for cnt := 0; cnt < s.opt.RetryCnt; cnt++ {
		url, idx := s.urls.Next()
		s.stunCli.SetServerAddr(url)
		natType, host, err := s.stunCli.Discover()
		fmt.Println(url, natType, host, err)
		if err != nil {
			e = err
			s.urls.Down(idx)
			if err.Error() == stun.NATBlocked.String() ||
				err.Error() == stun.NATError.String() {
				continue
			} else {
				return
			}
		} else {
			if natType == stun.NATBlocked ||
				natType == stun.NATError {
				s.urls.Down(idx)
				continue
			}
			addr.natType = natType
			addr.Family = host.Family()
			addr.IP = host.IP()
			addr.Port = host.Port()
			s.urls.Up(idx)
		}
	}
	return
}

func (s *Stun) Run(context context.Context) {
	if s.server != nil {
		go func() {
			for {
				select {
				case <-context.Done():
					return
				default:
					err := s.ReadFromClient()
					if err != nil {
						continue
					}
					s.WriteToClient()
				}
			}
		}()
	}
	if s.client != nil {
		go func() {
			time.Sleep(3 * time.Second)
			s.WriteToServer()
			s.ReadFromServer()
		}()
	}
	select {
	case <-context.Done():
		return
	default:
		time.Sleep(1 * time.Second)
	}
}
