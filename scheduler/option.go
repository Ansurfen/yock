// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	yocke "github.com/ansurfen/yock/env"
)

type YockSchedulerOption func(*YockScheduler) error

// OptionUpgradeSingalStream upgrades SingleSignalStream to CooperationSingalStream to meet distributed needs.
func OptionUpgradeSingalStream() YockSchedulerOption {
	return func(ys *YockScheduler) error {
		upgradeSingalStream(ys.signals.(*SingleSignalStream))
		return nil
	}
}

// OptionEnableYockDriverMode enables dependency analysis pattern.
//
// NOTE: It was deprecated in latest version, please use ypm to instead.
func OptionEnableYockDriverMode() YockSchedulerOption {
	return func(ys *YockScheduler) error {
		ys.driverManager = newDriverManager()
		return nil
	}
}

// OptionEnableEnvVar can CRUD environment variable.
// In some systems, you need administrator privileges to start
func OptionEnableEnvVar() YockSchedulerOption {
	return func(ys *YockScheduler) error {
		ys.envVar = yocke.NewEnvVar()
		return nil
	}
}
