// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	yocki "github.com/ansurfen/yock/interface"
	yockr "github.com/ansurfen/yock/runtime"
	"github.com/ansurfen/yock/util"
	lua "github.com/yuin/gopher-lua"
)

// yockJob (Job) is the smallest component of a task,
// and each task is freely combined through a job.
type yockJob struct {
	fn *lua.LFunction
}

func (job *yockJob) Func() *lua.LFunction {
	return job.fn
}

func loadTask(yocks yocki.YockScheduler) {
	yocks.RegYocksFn(yocki.YocksFuncs{
		"job":        taskJob,
		"jobs":       taskJobs,
		"job_option": taskJobOption,
	})
}

// taskJob packages the job as a task and registers it with the scheduler
//
// NOTE: the name of the job and jobs cannot be duplicated
/*
* @param jobName string
* @param jobFn function
 */
func taskJob(yocks yocki.YockScheduler, l *yockr.YockState) int {
	jobName := l.CheckString(1)
	jobFn := l.CheckFunction(2)
	if yocks.GetTask(jobName) {
		util.Ycho.Fatal(util.ErrDumplicateJobName.Error())
	} else {
		yocks.AppendTask(jobName, &yockJob{
			fn: jobFn,
		})
	}
	return 0
}

// taskJob packages multiple jobs as a task and registers it with the scheduler
//
// NOTE: the name of the job and jobs cannot be duplicated
/*
* @param name string
* @param jobs ...string
 */
func taskJobs(ys yocki.YockScheduler, l *yockr.YockState) int {
	yocks := ys.(*YockScheduler)
	groups := []string{}
	for i := 1; i <= l.GetTop(); i++ {
		groups = append(groups, l.CheckString(i))
	}
	if len(groups) <= 1 {
		return 0
	}
	name := groups[0]
	if yocks.GetTask(name) {
		util.Ycho.Fatal(util.ErrDumplicateJobName.Error())
	}
	for _, n := range groups[1:] {
		if job, ok := yocks.task[n]; ok {
			yocks.task[name] = append(yocks.task[name], job...)
		}
	}
	return 0
}

// taskJobOption gets local environment declared in the script
// and stores them in the scheduler's opt field.
//
// @param opt table
func taskJobOption(yocks yocki.YockScheduler, l *yockr.YockState) int {
	yocks.SetOpt(l.CheckTable(1))
	return 0
}
