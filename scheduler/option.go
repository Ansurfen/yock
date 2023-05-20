package scheduler

type YockSchedulerOption func(*YockScheduler) error

func OptionUpgradeSingalStream() YockSchedulerOption {
	return func(ys *YockScheduler) error {
		upgradeSingalStream(ys.signals.(*SingleSignalStream))
		return nil
	}
}
