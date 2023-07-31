// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"fmt"
	"strconv"
	"strings"

	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/ctl/conf"
	"github.com/ansurfen/yock/daemon/proto"
	yocke "github.com/ansurfen/yock/env"
	yocki "github.com/ansurfen/yock/interface"
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ycho"
	lua "github.com/yuin/gopher-lua"
)

// yockJob (Job) is the smallest component of a task,
// and each task is freely combined through a job.
type yockJob struct {
	name string
	fn   *lua.LFunction
}

func (job *yockJob) Name() string {
	return job.name
}

func (job *yockJob) Func() *lua.LFunction {
	return job.fn
}

func loadTask(yocks yocki.YockScheduler) {
	yocks.RegYocksFn(yocki.YocksFuncs{
		"job":    taskJob,
		"jobs":   taskJobs,
		"option": yocksOption,
	})
}

// taskJob packages the job as a task and registers it with the scheduler
//
// NOTE: the name of the job and jobs cannot be duplicated
//
// @param jobName string
//
// @param jobFn function
func taskJob(yocks yocki.YockScheduler, l yocki.YockState) int {
	jobName := l.LState().CheckString(1)
	jobFn := l.LState().CheckFunction(2)
	if yocks.GetTask(jobName) {
		ycho.Fatal(util.ErrDumplicateJobName)
	} else {
		yocks.AppendTask(jobName, &yockJob{
			name: jobName,
			fn:   jobFn,
		})
	}
	return 0
}

// taskJob packages multiple jobs as a task and registers it with the scheduler
//
// NOTE: the name of the job and jobs cannot be duplicated
//
// @param name string
//
// @param jobs ...string
func taskJobs(ys yocki.YockScheduler, l yocki.YockState) int {
	yocks := ys.(*YockScheduler)
	groups := []string{}
	for i := 1; i <= l.LState().GetTop(); i++ {
		groups = append(groups, l.LState().CheckString(i))
	}
	if len(groups) <= 1 {
		return 0
	}
	name := groups[0]
	if yocks.GetTask(name) {
		ycho.Fatal(util.ErrDumplicateJobName)
	}
	for _, n := range groups[1:] {
		if job, ok := yocks.task[n]; ok {
			yocks.task[name] = append(yocks.task[name], job...)
		}
	}
	return 0
}

// yocksOption gets local environment declared in the script
// and stores them in the scheduler's opt field.
//
// @param opt table
func yocksOption(yocks yocki.YockScheduler, l yocki.YockState) int {
	opt := l.CheckTable(1)
	yocks.SetOpt(opt)
	cfg := conf.YockConf{}
	if err := opt.Bind(&cfg); err != nil {
		return 0
	}
	env := yocke.GetEnv[*conf.YockConf]()
	if m, ok := opt.GetBool("sync"); ok && m {
		env.SetValue("strict", cfg.Strict)
		env.SetValue("ycho.stdout", cfg.Ycho.Stdout)
		env.Save()
	}
	if cfg.Strict {
		yocki.Y_MODE.SetMode(yocki.Y_STRICT)
	} else {
		yocki.Y_MODE.UnsetMode(yocki.Y_STRICT)
	}
	if cfg.Ycho.Stdout {
		yocki.Y_MODE.SetMode(yocki.Y_DEBUG)
	} else {
		yocki.Y_MODE.UnsetMode(yocki.Y_DEBUG)
	}
	ys := yocks.(*YockScheduler)
	if cfg.Yockd.Port > 0 {
		env.Conf().Yockd.Port = cfg.Yockd.Port
	}
	if cfg.Yockd.SelfBoot {
		infos, err := yockc.Lsof()
		if err != nil {
			ycho.Error(err)
		}
		found := false
		for _, info := range infos {
			if strings.Contains(info.Local, strconv.Itoa(env.Conf().Yockd.Port)) {
				found = true
				break
			}
		}
		if !found {
			err = yockc.Nohup(fmt.Sprintf("yockd%s -p %d", util.CurPlatform.Exf(), env.Conf().Yockd.Port))
			if err != nil {
				ycho.Error(err)
			}
		}
	}
	hostName := util.Title(cfg.Yockd.Name)
	if host, ok := cfg.Yockd.Peer[hostName]; ok {
		for peerName, peer := range cfg.Yockd.Peer {
			if peerName != hostName {
				ys.defaultYockd().Dial(&proto.NodeInfo{
					Name:   hostName,
					Ip:     host.IP,
					Port:   host.Port,
					Public: host.Public,
				}, &proto.NodeInfo{
					Name:   peerName,
					Ip:     peer.IP,
					Port:   peer.Port,
					Public: peer.Public,
				})
			}
		}
	}
	if cfg.Yockw.Port > 0 {
		env.Conf().Yockw.Port = cfg.Yockw.Port
	}
	if cfg.Yockw.SelfBoot {
		env.Conf().Yockw.SelfBoot = true
		err := yockc.Nohup(fmt.Sprintf("%s -p %d",
			util.Pathf("~/yockw"+util.CurPlatform.Exf()), env.Conf().Yockw.Port))
		if err != nil {
			ycho.Error(err)
		}
	}
	arr2Table := func(arr []string) string {
		res := []string{}
		for _, e := range arr {
			res = append(res, fmt.Sprintf(`"%s"`, e))
		}
		return fmt.Sprintf("{%s}", strings.Join(res, ","))
	}
	for _, c := range cfg.Yockw.Metrics.Counter {
		if len(c.Lables) > 0 {
			yocks.Eval(fmt.Sprintf(`yockw.metrics.counter_vec.new(
				{ name = "%s", help = "%s", labels = %s })`, c.Name, c.Help, arr2Table(c.Lables)))
		} else {
			yocks.Eval(fmt.Sprintf(`yockw.metrics.counter.new(
				{ name = "%s", help = "%s"})`, c.Name, c.Help))
		}
	}
	for _, c := range cfg.Yockw.Metrics.Gauge {
		if len(c.Lables) > 0 {
			yocks.Eval(fmt.Sprintf(`yockw.metrics.gauge_vec.new(
				{ name = "%s", help = "%s", labels = %s })`, c.Name, c.Help, arr2Table(c.Lables)))
		} else {
			yocks.Eval(fmt.Sprintf(`yockw.metrics.gauge.new(
				{ name = "%s", help = "%s"})`, c.Name, c.Help))
		}
	}
	for _, c := range cfg.Yockw.Metrics.Histogram {
		if len(c.Lables) > 0 {
			// yocks.Eval(fmt.Sprintf(`yockw.metrics.histogram_vec.new(
			// 	{ name = "%s", help = "%s", labels = %s, buckets = { %s } })`,
			// 	c.Name, c.Help, arr2Table(c.Lables), strings.Join(c.Buckets, ",")))
		} else {
			yocks.Eval(fmt.Sprintf(`yockw.metrics.histogram.new(
				{ name = "%s", help = "%s"})`, c.Name, c.Help))
		}
	}
	for _, c := range cfg.Yockw.Metrics.Summary {
		if len(c.Lables) > 0 {
			yocks.Eval(fmt.Sprintf(`yockw.metrics.summary.new(
				{ name = "%s", help = "%s", labels = %s })`, c.Name, c.Help, arr2Table(c.Lables)))
		} else {
			yocks.Eval(fmt.Sprintf(`yockw.metrics.summary.new(
				{ name = "%s", help = "%s"})`, c.Name, c.Help))
		}
	}
	return 0
}
