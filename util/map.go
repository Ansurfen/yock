package util

import "sync"

type SafeMap[T any] struct {
	mutex *sync.Mutex
}

func NewSafeMap[T any]() *SafeMap[T] {
	return &SafeMap[T]{
		mutex: &sync.Mutex{},
	}
}
