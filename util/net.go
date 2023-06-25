// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"net"
	"net/url"
)

// RandomPort return an available port with random.
func RandomPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// IsURL determine whether urlStr is a URL
func IsURL(urlStr string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	} else if u.Scheme == "" || u.Host == "" {
		return false
	} else {
		return true
	}
}
