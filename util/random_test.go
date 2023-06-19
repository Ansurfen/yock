// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

func TestWeightedRandom(t *testing.T) {
	testset := map[string]int{
		"A": 0,
		"B": 0,
		"C": 0,
		"D": 0,
		"E": 0,
		"F": 0,
	}
	elements := []string{}
	for k := range testset {
		elements = append(elements, k)
	}
	lb := NewWeightedRandom(elements)
	for i := 0; i < 100; i++ {
		e, idx := lb.Next()
		testset[e]++
		lb.Up(idx)
	}
	fmt.Println(testset, lb.weights)
}
