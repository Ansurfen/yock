// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package metrics

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/ansurfen/yock/util"
	"github.com/prometheus/client_golang/prometheus"
)

func TestMetricsWatch(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	DefaultMetricsWatch = New()

	go DefaultMetricsWatch.Snapshot()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			name := util.RandString(32)
			c := DefaultMetricsWatch.NewCounter(prometheus.CounterOpts{
				Name: name,
			})
			go func() {
				for {
					c.Inc()
					time.Sleep(2 * time.Second)
				}
			}()
		}
	}()

	http.Handle("/metrics", DefaultMetricsWatch.Document())

	go http.ListenAndServe(":2112", nil)
	<-c
}
