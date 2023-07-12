// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"os"
	"testing"
	"time"
)

func TestSSH(t *testing.T) {
	go func ()  {
		time.Sleep(20 * time.Second)
		os.Exit(0)	
	}()
	sh, _ := NewSSHClient(SSHOpt{
		User:     "ubuntu",
		Pwd:      "root",
		IP:       "192.168.127.128",
		Redirect: true,
		Network:  "tcp",
	})

	sh.Put("../release.tar.gz", "yock.tar")
	// sh.Get("myfile.txt", "../myfile.txt")
}
