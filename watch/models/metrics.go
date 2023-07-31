// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricsOpts struct {
	Namespace string `json:"nameSpace"`
	Subsystem string `json:"subSystem"`
	Name      string `json:"name"`
	Help      string `json:"help"`
}

type HistogramOpts struct {
	MetricsOpts
	Buckets []float64 `json:"buckets"`
}

type SummaryOpts struct {
	MetricsOpts
	Objectives map[string]float64 `json:"objectives"`
	MaxAge     time.Duration      `json:"maxAge"`
	AgeBuckets uint32             `json:"ageBuckets"`
	BufCap     uint32             `json:"bufCap"`
}

func (opt SummaryOpts) Adapter() prometheus.SummaryOpts {
	obj := make(map[float64]float64)
	for k, v := range opt.Objectives {
		f, err := strconv.ParseFloat(k, 64)
		if err != nil {
			panic(err)
		}
		obj[f] = v
	}
	return prometheus.SummaryOpts{
		Namespace:  opt.Namespace,
		Subsystem:  opt.Subsystem,
		Name:       opt.Name,
		Help:       opt.Help,
		Objectives: obj,
		MaxAge:     opt.MaxAge,
		AgeBuckets: opt.AgeBuckets,
		BufCap:     opt.BufCap,
	}
}

type MetricsVecOpts struct {
	prometheus.Opts
	Lables []string `json:"labels"`
}

type HistogramVecOpts struct {
	prometheus.HistogramOpts
	Lables []string `json:"labels"`
}

type SummaryVecOpts struct {
	SummaryOpts
	Lables []string `json:"labels"`
}
