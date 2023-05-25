package scheduler

import "github.com/ansurfen/cushion/utils"

type YockSchedulerOption func(*YockScheduler) error

func OptionUpgradeSingalStream() YockSchedulerOption {
	return func(ys *YockScheduler) error {
		upgradeSingalStream(ys.signals.(*SingleSignalStream))
		return nil
	}
}

func OptionEnableYockDriverMode() YockSchedulerOption {
	return func(ys *YockScheduler) error {
		ys.driverManager = newDriverManager()
		return nil
	}
}

func OptionEnableEnvVar() YockSchedulerOption {
	return func(ys *YockScheduler) error {
		ys.envVar = utils.NewEnvVar()
		return nil
	}
}
