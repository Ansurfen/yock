package util

import "sync"

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

func (m *SafeMap[T]) SafeGet(k string) (T, bool) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	return m.Get(k)
}

func (m *SafeMap[T]) Get(k string) (T, bool) {
	if v, ok := m.data[k]; ok {
		return v, ok
	} else {
		return v, false
	}
}

func (m *SafeMap[T]) SafeSet(k string, v T) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	m.data[k] = v
}

func (m *SafeMap[T]) Set(k string, v T) {
	m.data[k] = v
}

func (m *SafeMap[T]) SafeRange(handle func(k string, v T) bool) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	for k, v := range m.data {
		if !handle(k, v) {
			return
		}
	}
}

func (m *SafeMap[T]) Range(handle func(k string, v T) bool) {
	for k, v := range m.data {
		if !handle(k, v) {
			return
		}
	}
}
