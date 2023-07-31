// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package net

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/ansurfen/yock/daemon/conf"
	"github.com/ansurfen/yock/ycho"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type YockdClientOption struct {
	IP     string
	Port   int
	Global *conf.YockdConf
}

func OptionTransportCreds(c, k string, cas ...string) grpc.DialOption {
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
	return grpc.WithTransportCredentials(creds)
}
