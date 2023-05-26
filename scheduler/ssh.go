package scheduler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
	"github.com/pkg/sftp"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"golang.org/x/crypto/ssh"
	luar "layeh.com/gopher-luar"
)

func init() {
	center = &SSHCenter{}
}

func loadSSH(yocks *YockScheduler) lua.LValue {
	return yocks.Interp().NewClosure(func(l *lua.LState) int {
		opt := SSHOpt{}
		mode := l.CheckAny(1)
		if mode.Type() == lua.LTTable {
			gluamapper.Map(l.CheckTable(1), &opt)
			cli, err := NewSSHClient(opt)
			if err != nil {
				l.Push(lua.LNil)
				l.Push(lua.LString(err.Error()))
				return 2
			}
			if l.GetTop() >= 2 && l.CheckAny(2).Type() == lua.LTFunction {
				fn := l.CheckFunction(2)
				yocks.FastEvalFunc(fn, []lua.LValue{luar.New(l, cli)})
			}
			l.Push(luar.New(l, cli))
		}
		l.Push(lua.LNil)
		return 2
	})
}

type SSHReader struct {
	channel chan string
}

func NewSSHReader() *SSHReader {
	return &SSHReader{channel: make(chan string, 2)}
}

type SSHWriter struct {
	channel chan string
}

func NewSSHWriter() *SSHWriter {
	return &SSHWriter{channel: make(chan string, 2)}
}

func (r *SSHReader) Read(p []byte) (n int, err error) {
	cmd := <-r.channel
	tmpl := []byte(cmd + "\n")
	copy(p, tmpl)
	return len(tmpl), err
}

func (w *SSHWriter) Write(p []byte) (n int, err error) {
	w.channel <- string(p)
	return len(p), err
}

type SSHOpt struct {
	User     string
	Pwd      string
	IP       string
	Network  string
	Redirect bool
}

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

func (cli *SSHClient) Put(src, dst string) {
	sftpClient, err := sftp.NewClient(cli.Client)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	defer sftpClient.Close()
	remoteFile, err := sftpClient.Create(dst)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	defer remoteFile.Close()
	out, err := utils.ReadStraemFromFile(src)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	_, err = remoteFile.Write(out)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
}

func (cli *SSHClient) Get(src, dst string) {
	sftpClient, err := sftp.NewClient(cli.Client)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	defer sftpClient.Close()
	remoteFile, err := sftpClient.Open(src)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	defer remoteFile.Close()
	fp, err := os.Create(dst)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
	defer fp.Close()
	_, err = remoteFile.WriteTo(fp)
	if err != nil {
		util.Ycho.Fatal(err.Error())
	}
}

func (cli *SSHClient) Exec(cmds []string) {
	for _, cmd := range cmds {
		session, err := cli.NewSession()
		if err != nil {
			fmt.Printf("fail to create session: %s\n", err.Error())
			continue
		}
		defer session.Close()
		output, err := session.CombinedOutput(cmd)
		if err != nil {
			fmt.Printf("fail to execute command: %s\n", err.Error())
			continue
		}
		fmt.Println(string(output))
	}
}

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
		fmt.Println("fail to allocate term")
		return
	}
	if err := session.Shell(); err != nil {
		fmt.Println("fail to start shell")
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

var center *SSHCenter
