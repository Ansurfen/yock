// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"testing"
)

func TestSSH(t *testing.T) {
	sh, _ := NewSSHClient(SSHOpt{
		User:     "ubuntu",
		Pwd:      "root",
		IP:       "192.168.127.128",
		Redirect: true,
		Network:  "tcp",
	})

	sh.Put("../yock.tar", "release.tar")
	sh.Get("myfile.txt", "../myfile.txt")
}
