// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"errors"

	"github.com/ansurfen/yock/daemon/net"
	pb "github.com/ansurfen/yock/daemon/proto"
	yockr "github.com/ansurfen/yock/runtime"
	lua "github.com/yuin/gopher-lua"
)

func (yocks *YockScheduler) LoadYockd() {
	lib := yocks.CreateLib("yockd")
	lib.SetField(map[string]any{
		"ping": func(name string) error {
			return yocks.daemon[name].Ping()
		},
		"dial": func(name, ip string, port int) {
			yocks.daemon[name] = net.NewDirect(&net.YockdClientOption{
				IP:   ip,
				Port: port,
			})
		},
		"upload": func(name, src, dst string) {
			// yocks.daemon[name].Upload(src)
		},
	})
	signal := yockr.NewTable()
	signal.SetFields(yocks.LState(), map[string]any{
		"list": func() *lua.LTable {
			tbl := &lua.LTable{}
			sigs, err := yocks.defaultYockd().SignalList()
			if err != nil {
				return tbl
			}
			for _, sig := range sigs {
				tbl.Append(lua.LString(sig))
			}
			return tbl
		},
		"clear": func(sigs ...string) error {
			return yocks.defaultYockd().SignalClear(sigs...)
		},
		"info": func(sig string) (bool, bool, error) {
			return yocks.defaultYockd().SignalInfo(sig)
		},
	})
	fs := yockr.NewTable()
	fs.SetFields(yocks.LState(), map[string]any{
		"put": func(src, dst string) error {
			return yocks.defaultYockd().FileSystemPut(src, dst)
		},
		"get": func(src, dst string) error {
			return yocks.defaultYockd().FileSystemGet(src, dst)
		},
	})
	net := yockr.NewTable()
	net.SetFields(yocks.LState(), map[string]any{
		"dial": func(fromName, fromIP string, fromPort int32, fromPublic bool,
			toName, toIP string, toPort int32, toPublic bool) error {
			return yocks.defaultYockd().Dial(&pb.NodeInfo{
				Name:   fromName,
				Ip:     fromIP,
				Port:   fromPort,
				Public: fromPublic,
			}, &pb.NodeInfo{
				Name:   toName,
				Ip:     toIP,
				Port:   toPort,
				Public: toPublic,
			})
		},
		"call": func(node, method string, args ...string) (string, error) {
			return yocks.defaultYockd().Call(node, method, args...)
		},
	})
	process := yockr.NewTable()
	process.SetFields(yocks.LState(), map[string]any{
		"spawn": func(t, spec, cmd string) (int64, error) {
			switch t {
			case "cron":
				return yocks.defaultYockd().ProcessSpawn(pb.ProcessSpawnType_Cron, spec, cmd)
			case "fs":
				return yocks.defaultYockd().ProcessSpawn(pb.ProcessSpawnType_FS, spec, cmd)
			case "script":
				return yocks.defaultYockd().ProcessSpawn(pb.ProcessSpawnType_Script, spec, cmd)
			default:
				return 0, errors.New("invalid type")
			}
		},
		"find": func(v lua.LValue) (ret []*pb.Process, err error) {
			switch v.Type() {
			case lua.LTString:
				ret, err = yocks.defaultYockd().ProcessFind(0, string(v.(lua.LString)))
			case lua.LTNumber:
				ret, err = yocks.defaultYockd().ProcessFind(int64(v.(lua.LNumber)), "")
			default:
				err = errors.New("invalid key")
			}
			return
		},
		"kill": func(pid int64) error {
			return yocks.defaultYockd().ProcessKill(pid)
		},
		"list": func() (*lua.LTable, error) {
			tbl := &lua.LTable{}
			res, err := yocks.defaultYockd().ProcessList()
			if err != nil {
				return nil, err
			}
			for _, p := range res {
				tmp := &lua.LTable{}
				tbl.Append(tmp)
				tmp.RawSetString("cmd", lua.LString(p.Cmd))
				tmp.RawSetString("pid", lua.LNumber(p.Pid))
				tmp.RawSetString("state", lua.LNumber(p.State))
				tmp.RawSetString("spec", lua.LString(p.Spec))
			}
			return tbl, nil
		},
	})
	lib.SetField(map[string]any{
		"signal":  signal.Value(),
		"fs":      fs.Value(),
		"net":     net.Value(),
		"process": process.Value(),
	})
}
