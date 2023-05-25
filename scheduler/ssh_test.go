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
