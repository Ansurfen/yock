// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/ansurfen/yock/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func WithTransportCreds(c, k string, cas ...string) grpc.DialOption {
	cert, err := tls.LoadX509KeyPair(c, k)
	if err != nil {
		util.Ycho.Fatal(fmt.Sprintf("failed to load key pair: %s", err))
	}

	caCertPool := x509.NewCertPool()
	for _, ca := range cas {
		caCert, err := ioutil.ReadFile(ca)
		if err != nil {
			util.Ycho.Fatal(fmt.Sprintf("failed to read CA certificate: %s", err))
		}
		if !caCertPool.AppendCertsFromPEM(caCert) {
			util.Ycho.Fatal("failed to append CA certificate")
		}
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	})
	return grpc.WithTransportCredentials(creds)
}
