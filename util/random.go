// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"math/rand"
)

type LoadBalancer[T any] interface {
	// Next returns an element and its index.
	// If error, returns -1.
	Next() (T, int)
	// Up increases the probability of element to be specified
	Up(idx int)
	// Down decreases the probability of element to be specified
	Down(idx int)
	Put(e T)
	Del(idx int)
	Weights() []float64
}

var _ LoadBalancer[int] = (*WeightedRandom[int])(nil)

type WeightedRandom[T any] struct {
	weights  []float64
	elements []T
	policy   func(weights []float64) float64
}

func defaultRandomWeight(weights []float64) float64 {
	return 0.1
}

func NewWeightedRandom[T any](
	elements []T,
	policy ...func(weights []float64) float64) *WeightedRandom[T] {
	wr := &WeightedRandom[T]{elements: elements}
	if len(policy) > 0 {
		wr.policy = policy[0]
	} else {
		wr.policy = defaultRandomWeight
	}
	for i := 0; i < len(elements); i++ {
		wr.weights = append(wr.weights, 0.5)
	}
	return wr
}

// Up increases the probability of element to be specified
func (wr *WeightedRandom[T]) Up(idx int) {
	if idx >= 0 && idx < len(wr.elements) {
		wr.weights[idx] += wr.policy(wr.weights)
	}
}

// Down decreases the probability of element to be specified
func (wr *WeightedRandom[T]) Down(idx int) {
	if idx >= 0 && idx < len(wr.elements) {
		wr.weights[idx] -= wr.policy(wr.weights)
	}
}

func (wr *WeightedRandom[T]) Put(e T) {
	wr.elements = append(wr.elements, e)
	var totalWeight float64
	for _, w := range wr.weights {
		totalWeight += w
	}
	wr.weights = append(wr.weights, totalWeight/float64(len(wr.elements)))
}

func (wr *WeightedRandom[T]) Del(idx int) {
	wr.elements = append(wr.elements[:idx], wr.elements[idx+1:]...)
}

// Next returns an element and its index.
// If error, returns -1.
func (wr *WeightedRandom[T]) Next() (T, int) {
	var (
		totalWeight      float64
		cumulativeWeight float64
	)
	for _, w := range wr.weights {
		totalWeight += w
	}
	randomWeight := rand.Float64() * totalWeight
	for i, e := range wr.elements {
		cumulativeWeight += wr.weights[i]
		if randomWeight < cumulativeWeight {
			return e, i
		}
	}
	var v T
	return v, -1
}

func (wr *WeightedRandom[T]) Weights() []float64 {
	return wr.weights
}
