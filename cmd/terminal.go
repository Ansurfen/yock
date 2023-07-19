// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ansurfen/yock/util"
)

const (
	TermUndefined = iota
	TermPowershell
	TermCmd
	TermBash
)

var (
	cmdConf        = [2]string{"cmd", "/C"}
	powerShellConf = [1]string{"powershell"}
	bashConf       = [2]string{"/bin/sh", "-c"}
)

// Terminal is a struct to abstract different termnial
type Terminal struct {
	cmd  []string
	conf []string
	kind uint8
}

func (term *Terminal) Exec(opt *ExecOpt) ([]byte, error) {
	name := term.conf[0]
	args := []string{}
	if len(term.conf) > 1 {
		args = append(args, term.conf[1:]...)
	}
	args = append(args, term.cmd...)
	cmd := exec.Command(name, args...)

	if opt.Redirect {
		var out bytes.Buffer
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = io.MultiWriter(os.Stdout, &out)
		err := cmd.Run()
		if err != nil {
			return nil, err
		}
		return out.Bytes(), err
	}

	out, err := cmd.CombinedOutput()

	switch util.CurPlatform.Lang {
	case "zh":
		out = []byte(util.ConvertByte2String(out, util.GB18030))
	}

	if !opt.Quiet {
		fmt.Print(string(out))
	}
	return out, err
}

func (term *Terminal) SetCmds(cmds ...string) {
	term.cmd = cmds
}

func (term *Terminal) AppendCmds(cmds ...string) {
	term.cmd = append(term.cmd, cmds...)
}

func (term *Terminal) Clear() {
	term.cmd = []string{}
}

func (term *Terminal) Type() uint8 {
	return term.kind
}

func WindowsTerm(cmds ...string) *Terminal {
	switch util.CurPlatform.Ver {
	case "10", "11":
		return powershellTerm(cmds...)
	default:
		return cmdTerm(cmds...)
	}
}

func PosixTerm(cmds ...string) *Terminal {
	return bashTerm(cmds...)
}

func powershellTerm(cmds ...string) *Terminal {
	return &Terminal{conf: powerShellConf[:], cmd: cmds, kind: TermPowershell}
}

func cmdTerm(cmds ...string) *Terminal {
	return &Terminal{conf: cmdConf[:], cmd: cmds, kind: TermCmd}
}

func bashTerm(cmds ...string) *Terminal {
	return &Terminal{conf: bashConf[:], cmd: cmds, kind: TermBash}
}

type ArgsBuilder struct {
	cmds []string
}

func NewArgsBuilder(cmds ...string) *ArgsBuilder {
	return &ArgsBuilder{
		cmds: cmds,
	}
}

func (builder *ArgsBuilder) AddInt(k string, i int) *ArgsBuilder {
	if i != 0 {
		builder.cmds = append(builder.cmds, fmt.Sprintf(k, i))
	}
	return builder
}

func (builder *ArgsBuilder) AddString(k string, v string) *ArgsBuilder {
	if len(v) > 0 {
		builder.cmds = append(builder.cmds, fmt.Sprintf(k, v))
	}
	return builder
}

func (builder *ArgsBuilder) MustAddString() {

}

func (builder *ArgsBuilder) Add(cmds ...string) *ArgsBuilder {
	builder.cmds = append(builder.cmds, cmds...)
	return builder
}

func (builder *ArgsBuilder) Build() string {
	return strings.Join(builder.cmds, " ")
}

var (
	randomScript  string
	scriptStarter string
	scripts       map[string]string
)

func init() {
	scripts = make(map[string]string)
	if util.CurPlatform.OS == "windows" {
		randomScript = "*.vbs"
		scriptStarter = `cscript //Nologo "%s"`
	} else {
		randomScript = "*.sh"
		scriptStarter = "sh %s"
	}
}

func OnceScript(s string) (string, error) {
	script, err := os.CreateTemp("", randomScript)
	if err != nil {
		return "", err
	}
	_, err = script.Write([]byte(s))
	if err != nil {
		return "", err
	}
	script.Close() // release file lock
	return Exec(ExecOpt{Quiet: true}, fmt.Sprintf(scriptStarter, script.Name()))
}

func MultiScript(name, s string) (string, error) {
	if script, ok := scripts[name]; ok && util.IsExist(script) {
		_, err := Exec(ExecOpt{Quiet: true}, fmt.Sprintf(scriptStarter, script))
		return "", err
	}
	script, err := os.CreateTemp("", randomScript)
	if err != nil {
		return "", err
	}
	_, err = script.Write([]byte(s))
	if err != nil {
		return "", err
	}
	script.Close() // release file lock
	str, err := Exec(ExecOpt{Quiet: true}, fmt.Sprintf(scriptStarter, script.Name()))
	if err != nil {
		return str, err
	}
	scripts[name] = script.Name()
	return str, nil
}
