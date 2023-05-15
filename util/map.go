package util

import "sync"

type SafeMap struct {
	mutex *sync.Mutex
}
