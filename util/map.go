// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "sync"

// SafeMap is a simple alternative version of sync.map,
// designed specifically for SignalStream.
// Provides unlocked (unsafe) and locked (safe) to operate map
// to meet special scenarios to improve performance.
type SafeMap[T any] struct {
	mutex *sync.Mutex
	data  map[string]T
}

func NewSafeMap[T any]() *SafeMap[T] {
	return &SafeMap[T]{
		mutex: &sync.Mutex{},
		data:  make(map[string]T),
	}
}

// SafeGet locks to get the value of the specified k.
// If the value doesn't exist, the second parameter returns false, and vice versa.
func (m *SafeMap[T]) SafeGet(k string) (T, bool) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	return m.Get(k)
}

// Get directly obtains the value of the specified k without locking.
// If the value doesn't exist, the second parameter returns false, and vice versa.
func (m *SafeMap[T]) Get(k string) (T, bool) {
	if v, ok := m.data[k]; ok {
		return v, ok
	} else {
		return v, false
	}
}

// SafeSet locks to set value of key to be specified
func (m *SafeMap[T]) SafeSet(k string, v T) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	m.data[k] = v
}

// SafeSet sets value of key to be specified with locking
func (m *SafeMap[T]) Set(k string, v T) {
	m.data[k] = v
}

// SafeRange locks to range map. You can set callback to implement demand.
func (m *SafeMap[T]) SafeRange(handle func(k string, v T) bool) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	for k, v := range m.data {
		if !handle(k, v) {
			return
		}
	}
}

// Range ranges map without locking. You can set callback to implement demand.
func (m *SafeMap[T]) Range(handle func(k string, v T) bool) {
	for k, v := range m.data {
		if !handle(k, v) {
			return
		}
	}
}
