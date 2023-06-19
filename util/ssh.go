// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/ansurfen/cushion/utils"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SSHReader implements io.Reader,
// which reads string stream by channel.
type SSHReader struct {
	channel chan string
}

func NewSSHReader() *SSHReader {
	return &SSHReader{channel: make(chan string, 2)}
}

func (r *SSHReader) Read(p []byte) (n int, err error) {
	cmd := <-r.channel
	tmpl := []byte(cmd + "\n")
	copy(p, tmpl)
	return len(tmpl), err
}

// SSHWriter implements io.Writer,
// which writes to a string stream by channel.
type SSHWriter struct {
	channel chan string
}

func NewSSHWriter() *SSHWriter {
	return &SSHWriter{channel: make(chan string, 2)}
}

func (w *SSHWriter) Write(p []byte) (n int, err error) {
	w.channel <- string(p)
	return len(p), err
}

// SSHOpt indicates configuration of newSSHClient
type SSHOpt struct {
	User string
	// password
	Pwd string
	IP  string
	// tcp, udp, etc.
	Network  string
	Redirect bool
}

// SSHClient packs the SSH connection
type SSHClient struct {
	*ssh.Client
}

func newSSHClient(opt SSHOpt) (*SSHClient, error) {
	conf := &ssh.ClientConfig{
		User: opt.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(opt.Pwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial(opt.Network, opt.IP+":22", conf)
	if err != nil {
		return nil, err
	}
	return &SSHClient{
		Client: conn,
	}, nil
}

// Put uploads local files to a remote server
func (cli *SSHClient) Put(src, dst string) {
	sftpClient, err := sftp.NewClient(cli.Client)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
	defer sftpClient.Close()
	remoteFile, err := sftpClient.Create(dst)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
	defer remoteFile.Close()
	out, err := utils.ReadStraemFromFile(src)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
	_, err = remoteFile.Write(out)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
}

// Get download remote file to localhost from remote server
func (cli *SSHClient) Get(src, dst string) {
	sftpClient, err := sftp.NewClient(cli.Client)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
	defer sftpClient.Close()
	remoteFile, err := sftpClient.Open(src)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
	defer remoteFile.Close()
	fp, err := os.Create(dst)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
	defer fp.Close()
	_, err = remoteFile.WriteTo(fp)
	if err != nil {
		Ycho.Fatal(err.Error())
	}
}

// Exec creates a temporary session to execute commands
func (cli *SSHClient) Exec(cmds []string) {
	for _, cmd := range cmds {
		session, err := cli.NewSession()
		if err != nil {
			Ycho.Warn(fmt.Sprintf("%s: %s", ErrCreateSession.Error(), err.Error()))
			continue
		}
		defer session.Close()
		output, err := session.CombinedOutput(cmd)
		if err != nil {
			Ycho.Warn(fmt.Sprintf("%s: %s", ErrExecuteCommand.Error(), err.Error()))
			continue
		}
		fmt.Println(string(output))
	}
}

// Shell assigns a terminal to the user while redrecting stdout, stderr, stdin.
// Input exit to release the terminal to close the session.
func (cli *SSHClient) Shell() {
	session, _ := cli.NewSession()
	defer session.Close()
	w := NewSSHWriter()
	r := NewSSHReader()
	session.Stdout = w
	session.Stdin = r
	session.Stderr = w
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm", 25, 80, modes); err != nil {
		Ycho.Warn(ErrAllocTerm.Error())
		return
	}
	if err := session.Shell(); err != nil {
		Ycho.Warn(ErrAllocShell.Error())
		return
	}
	go func() {
		for {
			select {
			case res := <-w.channel:
				fmt.Print(res)
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	buf := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := buf.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		if cmd == "exit" {
			session.Close()
			break
		}
		r.channel <- cmd
	}
}

// SSHCenter manages the SSHClient
type SSHCenter struct {
	clients []*SSHClient
}

func NewSSHClient(opt SSHOpt) (*SSHClient, error) {
	cli, err := newSSHClient(opt)
	if err != nil {
		return nil, err
	}
	center.clients = append(center.clients, cli)
	return cli, nil
}

var (
	_      io.Writer = (*SSHWriter)(nil)
	_      io.Reader = (*SSHReader)(nil)
	center *SSHCenter
)
