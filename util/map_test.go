// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := NewSafeMap[bool]()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		idx := 0
		for {
			str := strconv.Itoa(idx) + " => "
			m.Range(func(k string, v bool) bool {
				str += fmt.Sprintf("%s:%v", k, v)
				return true
			})
			fmt.Println(str)
			idx++
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			m.SafeSet(strconv.Itoa(i), false)
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	go func() {
		for i := 5; i >= 0; i-- {
			m.SafeSet(strconv.Itoa(i), true)
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
}
