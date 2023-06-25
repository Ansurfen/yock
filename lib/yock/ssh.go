// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package liby

import (
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
)

func LoadSSH(yocks yocki.YockScheduler) {
	yocks.RegYocksFn(yocki.YocksFuncs{
		"ssh": sshSSH,
	})
}

/*
* @param opt table
* @param cb function(*SSHClient)
* @return userdata (*SSHClient), err
 */
func sshSSH(yocks yocki.YockScheduler, state *yockr.YockState) int {
	opt := util.SSHOpt{}
	if state.IsTable(1) {
		state.CheckTable(1).Bind(&opt)
		cli, err := util.NewSSHClient(opt)
		if err != nil {
			state.PushNil().Throw(err)
			return 2
		}
		if state.GetTop() >= 2 && state.IsFunction(2) {
			fn := state.CheckFunction(2)
			state.Call(yockr.YockFuncInfo{
				Fn: fn,
			}, cli)
		}
		state.Pusha(cli)
	}
	state.PushNil()
	return 2
}