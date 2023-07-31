// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import "github.com/ansurfen/yock/watch/models"

type yockWatch struct {
	SelfBoot bool             `yaml:"self_boot"`
	Port     int              `yaml:"port"`
	Metrics  yockWatchMetrics `yaml:"metrics"`
}

type yockWatchMetrics struct {
	Resolved  []string                  `yaml:"resolved"`
	Counter   []models.MetricsVecOpts   `yaml:"counter"`
	Gauge     []models.MetricsVecOpts   `yaml:"gauge"`
	Histogram []models.HistogramVecOpts `yaml:"histogram"`
	Summary   []models.SummaryVecOpts   `yaml:"summary"`
}
