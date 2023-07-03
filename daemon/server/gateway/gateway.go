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

	"github.com/ansurfen/yock/daemon/server/gateway/agent"
	"github.com/ansurfen/yock/daemon/server/gateway/rule"
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
	yockdGateway := &YockdGateWay{
		agents: map[string]agent.RuleAgent{
			"jwt": agent.NewJWTAgent(),
			// "pwd": agent.NewPwdAgent(),
		},
	}
	return yockdGateway
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

func (gate *YockdGateWay) AddToken(str []byte) error {
	_, err := gate.parseRule(str)
	return err
}

func (gate *YockdGateWay) DelToken(str []byte) error {
	var t map[string]any
	err := json.Unmarshal(str, &t)
	if err != nil {
		panic(err)
	}
	return nil
}

func (gate *YockdGateWay) parseRule(str []byte) (rule.Rule, error) {
	var t map[string]any
	err := json.Unmarshal(str, &t)
	if err != nil {
		return nil, err
	}
	ruleType := ""
	if rt, ok := t["type"].(string); ok {
		ruleType = rt
	} else {
		return nil, errors.New("rule not found")
	}
	if agent, ok := gate.agents[ruleType]; ok {
		name := ""
		if n, ok := t["name"].(string); ok {
			name = n
		}
		if r := agent.Get(name); r != nil {
			return r, nil
		}
		return agent.Release(name, t), nil
	}
	return nil, errors.New("agent not found")
}

func (gate *YockdGateWay) SetRule(key string, rules ...string) {
	rs := []rule.Rule{}
	for _, rule := range rules {
		r, err := gate.parseRule([]byte(rule))
		if err == nil {
			rs = append(rs, r)
		}
	}
	gate.policy.SetRule(key, rs...)
}

func (gate *YockdGateWay) UnsetRule(key string, rules ...string) {

}
