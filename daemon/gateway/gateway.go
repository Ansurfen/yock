// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gateway

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/ansurfen/yock/daemon/gateway/agent"
	"github.com/ansurfen/yock/daemon/gateway/rule"
	"github.com/ansurfen/yock/ycho"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type YockdGateWay struct {
	agents map[string]agent.RuleAgent
	policy PermPolicy
}

func New() *YockdGateWay {
	return &YockdGateWay{
		agents: make(map[string]agent.RuleAgent),
	}
}

func (gate *YockdGateWay) SetAgent(name string, agent agent.RuleAgent) {
	gate.agents[name] = agent
}

func (gate *YockdGateWay) SetPolicy(p PermPolicy) {
	gate.policy = p
}

func (gate *YockdGateWay) GuardUnary() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			err = errors.New("invalid context")
			return
		}
		err = gate.policy.Auth(&md, info.FullMethod)
		if err != nil {
			return
		}
		return handler(ctx, req)
	})
}

func (gate *YockdGateWay) GuardStream() grpc.ServerOption {
	return grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			err = errors.New("invalid context")
			return
		}
		err = gate.policy.Auth(&md, info.FullMethod)
		if err != nil {
			return
		}
		return handler(srv, ss)
	})
}

func (gate *YockdGateWay) GuardTransport(c, k string, cas ...string) grpc.ServerOption {
	cert, err := tls.LoadX509KeyPair(c, k)
	if err != nil {
		ycho.Fatalf("failed to load key pair: %s", err)
	}

	caCertPool := x509.NewCertPool()
	for _, ca := range cas {
		caCert, err := ioutil.ReadFile(ca)
		if err != nil {
			ycho.Fatalf("failed to read CA certificate: %s", err)
		}
		if !caCertPool.AppendCertsFromPEM(caCert) {
			ycho.Fatalf("failed to append CA certificate")
		}
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	})
	return grpc.Creds(creds)
}

func (gate *YockdGateWay) AddRule(str []byte) error {
	_, err := gate.parseRule(str)
	return err
}

func (gate *YockdGateWay) DelRule(str []byte) error {
	var t map[string]any
	err := json.Unmarshal(str, &t)
	if err != nil {
		panic(err)
	}
	return nil
}

type Rule struct {
	Type  string         `json:"type"`
	Name  string         `json:"name"`
	Token map[string]any `json:"token"`
}

func (gate *YockdGateWay) parseRule(str []byte) (rule.Rule, error) {
	var r Rule
	err := json.Unmarshal(str, &r)
	if err != nil {
		return nil, err
	}

	if agent, ok := gate.agents[r.Type]; ok {
		if r := agent.Get(r.Name); r != nil {
			return r, nil
		}
		ycho.Infof("release %s", str)
		return agent.Release(r.Name, r.Token), nil
	}
	return nil, errors.New("agent not found")
}

func (gate *YockdGateWay) SetRule(key string, rules ...string) {
	defer func() {
		msg := recover()
		switch v := msg.(type) {
		case error:
			ycho.Errorf("error happen %s, when mount %s.%s", v, gate.policy.Policy(), key)
		}
	}()
	rs := []rule.Rule{}
	names := ""
	for _, rule := range rules {
		kv := strings.Split(rule, ".")
		r := gate.agents[kv[0]].Get(kv[1])
		if r != nil {
			rs = append(rs, r)
			names += rule + " "
		}
	}
	gate.policy.Policy()
	ycho.Infof("%s.%s mount %s", gate.policy.Policy(), key, names)
	gate.policy.SetRule(key, rs...)
}

func (gate *YockdGateWay) UnsetRule(key string, rules ...string) {

}
