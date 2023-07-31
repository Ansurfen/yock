// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/ansurfen/yock/util"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func init() {
	center = &SSHCenter{}
}

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
	Pwd  string
	IP   string
	Port int
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
	conn, err := ssh.Dial(opt.Network, fmt.Sprintf("%s:%d", opt.IP, opt.Port), conf)
	if err != nil {
		return nil, err
	}
	return &SSHClient{
		Client: conn,
	}, nil
}

// Put uploads local files to a remote server
func (cli *SSHClient) Put(src, dst string) error {
	sftpClient, err := sftp.NewClient(cli.Client)
	if err != nil {
		return err
	}
	defer sftpClient.Close()
	remoteFile, err := sftpClient.Create(dst)
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	out, err := util.ReadStraemFromFile(src)
	if err != nil {
		return err
	}
	_, err = remoteFile.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// Get download remote file to localhost from remote server
func (cli *SSHClient) Get(src, dst string) error {
	sftpClient, err := sftp.NewClient(cli.Client)
	if err != nil {
		return err
	}
	defer sftpClient.Close()
	remoteFile, err := sftpClient.Open(src)
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	fp, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = remoteFile.WriteTo(fp)
	if err != nil {
		return err
	}
	return nil
}

// Exec creates a temporary session to execute commands
func (cli *SSHClient) Exec(cmd string) (string, error) {
	session, err := cli.NewSession()
	if err != nil {
		return "", fmt.Errorf("%s: %s", util.ErrCreateSession.Error(), err.Error())
	}
	defer session.Close()
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", fmt.Errorf("%s: %s", util.ErrExecuteCommand.Error(), err.Error())
	}
	return string(output), nil
}

func (cli *SSHClient) Sh(file string, args ...string) (string, error) {
	var (
		out string
		arg = strings.Join(args, " ")
	)
	switch filepath.Ext(file) {
	case ".sh":
		sh := util.RandString(32) + ".sh"
		err := cli.Put(file, sh)
		if err != nil {
			return "", err
		}
		out, err = cli.Exec(fmt.Sprintf("chmod +x %s && sh %s %s && rm %s", sh, sh, arg, sh))
		if err != nil {
			return out, err
		}
	case ".bat":
		bat := util.RandString(32) + ".bat"
		err := cli.Put(file, bat)
		if err != nil {
			return "", err
		}
		out, err = cli.Exec(fmt.Sprintf("%s %s & del %s", bat, arg, bat))
		if err != nil {
			return out, err
		}
	}
	return out, nil
}

func (cli *SSHClient) OS() string {
	raw, err := cli.Exec("echo $OSTYPE")
	raw = strings.TrimRight(raw, "\n")
	if err != nil {
		return "unknown"
	}
	if raw != "$OSTYPE" {
		return raw
	}
	raw, err = cli.Exec("echo %OS%")
	raw = strings.TrimRight(raw, "\n")
	if err != nil {
		return "unknown"
	}
	if raw != "%OS%" {
		return raw
	}
	return "unknown"
}

// Shell assigns a terminal to the user while redrecting stdout, stderr, stdin.
// Input exit to release the terminal to close the session.
func (cli *SSHClient) Shell() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	session, err := cli.NewSession()
	if err != nil {
		return err
	}
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
		return util.ErrAllocTerm
	}
	if err := session.Shell(); err != nil {
		return util.ErrAllocShell
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
	go func() {
		for {
			select {
			case <-c:
				session.Close()
				return
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	go func() {
		buf := bufio.NewReader(os.Stdin)
		for {
			cmd, err := buf.ReadString('\n')
			if err != nil {
				continue
			}
			r.channel <- strings.TrimSpace(cmd)
		}
	}()
	if err := session.Wait(); err != nil {
		return err
	}
	return nil
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
