// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/ansurfen/yock/scheduler"
)

func main() {
	var wg sync.WaitGroup
	ys := yocks.New()
	lib := ys.CreateLib("log")
	lib.SetField(map[string]any{
		"Info": func(msg string) {
			log.Println(msg)
		},
		"Infof": func(msg string, a ...any) {
			log.Printf(msg, a...)
		},
	})
	if err := ys.Eval(`log.Info("Hello World!")`); err != nil {
		panic(err)
	}
	go ys.EventLoop()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ys.Do(func() {
				ys.Eval(fmt.Sprintf(`log.Info("%d")`, i))
				wg.Done()
			})
		}(i)
	}
	wg.Wait()
}
