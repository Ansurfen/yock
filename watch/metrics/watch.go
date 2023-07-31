// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package metrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ansurfen/yock/watch/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var DefaultMetricsWatch *MetricsWatch

type MetricsWatch struct {
	counters      map[string]prometheus.Counter
	gauges        map[string]prometheus.Gauge
	histograms    map[string]prometheus.Histogram
	summaries     map[string]prometheus.Summary
	counterVecs   map[string]*prometheus.CounterVec
	gaugeVecs     map[string]*prometheus.GaugeVec
	histogramVecs map[string]*prometheus.HistogramVec
	summaryVecs   map[string]*prometheus.SummaryVec
}

func New() *MetricsWatch {
	return &MetricsWatch{
		counters:      make(map[string]prometheus.Counter),
		gauges:        make(map[string]prometheus.Gauge),
		histograms:    make(map[string]prometheus.Histogram),
		summaries:     make(map[string]prometheus.Summary),
		counterVecs:   make(map[string]*prometheus.CounterVec),
		gaugeVecs:     make(map[string]*prometheus.GaugeVec),
		histogramVecs: make(map[string]*prometheus.HistogramVec),
		summaryVecs:   make(map[string]*prometheus.SummaryVec),
	}
}

func (mw *MetricsWatch) NewCounter(opt prometheus.CounterOpts) prometheus.Counter {
	if c, ok := mw.counters[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewCounter(opt)
		mw.counters[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) AddCounter(name string, f float64) error {
	if f < 0 {
		return errors.New("the value is must more than zero")
	}
	if c, ok := mw.counters[name]; ok {
		c.Add(f)
		return nil
	}
	return errors.New("counter not found")
}

func (mw *MetricsWatch) IncCounter(name string) error {
	if c, ok := mw.counters[name]; ok {
		c.Inc()
		return nil
	}
	return errors.New("counter not found")
}

func (mw *MetricsWatch) RmCounter(name string) error {
	if c, ok := mw.counters[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.counters, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) LsCounter() string {
	data, err := json.Marshal(mw.counters)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) NewGauge(opt prometheus.GaugeOpts) prometheus.Gauge {
	if c, ok := mw.gauges[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewGauge(opt)
		mw.gauges[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) AddGauge(name string, f float64) error {
	if c, ok := mw.gauges[name]; ok {
		c.Add(f)
		return nil
	}
	return errors.New("gauge not found")
}

func (mw *MetricsWatch) SubGauge(name string, f float64) error {
	if c, ok := mw.gauges[name]; ok {
		c.Sub(f)
		return nil
	}
	return errors.New("gauge not found")
}

func (mw *MetricsWatch) IncGauge(name string) error {
	if c, ok := mw.gauges[name]; ok {
		c.Inc()
		return nil
	}
	return errors.New("gauge not found")
}

func (mw *MetricsWatch) DecGauge(name string) error {
	if c, ok := mw.gauges[name]; ok {
		c.Dec()
		return nil
	}
	return errors.New("gauge not found")
}

func (mw *MetricsWatch) SetGauge(name string, f float64) error {
	if c, ok := mw.gauges[name]; ok {
		c.Set(f)
		return nil
	}
	return errors.New("gauge not found")
}

func (mw *MetricsWatch) RmGauge(name string) error {
	if c, ok := mw.gauges[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.gauges, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) SetGaugeToCurrentTime(name string) error {
	if c, ok := mw.gauges[name]; ok {
		c.SetToCurrentTime()
		return nil
	}
	return errors.New("gauge not found")
}

func (mw *MetricsWatch) LsGauge() string {
	data, err := json.Marshal(mw.gauges)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) NewHistogram(opt prometheus.HistogramOpts) prometheus.Histogram {
	if c, ok := mw.histograms[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewHistogram(opt)
		mw.histograms[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) ObserveHistogram(name string, f float64) error {
	if c, ok := mw.histograms[name]; ok {
		c.Observe(f)
		return nil
	}
	return errors.New("histogram not found")
}

func (mw *MetricsWatch) RmHistogram(name string) error {
	if c, ok := mw.histograms[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.histograms, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) LsHistogram() string {
	data, err := json.Marshal(mw.histograms)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) NewSummaries(opt prometheus.SummaryOpts) prometheus.Summary {
	if c, ok := mw.summaries[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewSummary(opt)
		mw.summaries[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) ObserveSummaries(name string, f float64) error {
	if c, ok := mw.summaries[name]; ok {
		c.Observe(f)
		return nil
	}
	return errors.New("summary not found")
}

func (mw *MetricsWatch) RmSummary(name string) error {
	if c, ok := mw.summaries[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.summaries, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) LsSummaries() string {
	data, err := json.Marshal(mw.summaries)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) NewCounterVec(opt models.MetricsVecOpts) *prometheus.CounterVec {
	if c, ok := mw.counterVecs[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewCounterVec(prometheus.CounterOpts(opt.Opts), opt.Lables)
		mw.counterVecs[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) parseCounterVec(name string, labels any) (prometheus.Counter, error) {
	if c, ok := mw.counterVecs[name]; ok {
		switch v := labels.(type) {
		case []any:
			lvs := []string{}
			for _, label := range v {
				if s, ok := label.(string); ok {
					lvs = append(lvs, s)
				}
			}
			return c.WithLabelValues(lvs...), nil
		case map[string]any:
			lvs := make(map[string]string)
			for k, vv := range v {
				if s, ok := vv.(string); ok {
					lvs[k] = s
				}
			}
			return c.With(lvs), nil
		default:
			return nil, errors.New("invalid label type")
		}
	}
	return nil, errors.New("gaugeVec not found")
}

func (mw *MetricsWatch) AddCounterVec(name string, f float64, labels any) error {
	if f < 0 {
		return errors.New("the value is must more than zero")
	}
	counter, err := mw.parseCounterVec(name, labels)
	if err != nil {
		return err
	}
	counter.Add(f)
	return nil
}

func (mw *MetricsWatch) IncCounterVec(name string, labels any) error {
	counter, err := mw.parseCounterVec(name, labels)
	if err != nil {
		return err
	}
	counter.Inc()
	return nil
}

func (mw *MetricsWatch) RmCounterVec(name string) error {
	if c, ok := mw.counterVecs[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.counterVecs, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) LsCounterVec() string {
	data, err := json.Marshal(mw.counterVecs)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) parseGaugeVec(name string, labels any) (prometheus.Gauge, error) {
	if c, ok := mw.gaugeVecs[name]; ok {
		switch v := labels.(type) {
		case []any:
			lvs := []string{}
			for _, label := range v {
				if s, ok := label.(string); ok {
					lvs = append(lvs, s)
				}
			}
			return c.WithLabelValues(lvs...), nil
		case map[string]any:
			lvs := make(map[string]string)
			for k, vv := range v {
				if s, ok := vv.(string); ok {
					lvs[k] = s
				}
			}
			return c.With(lvs), nil
		default:
			return nil, errors.New("invalid label type")
		}
	}
	return nil, errors.New("gaugeVec not found")
}

func (mw *MetricsWatch) NewGaugeVec(opt models.MetricsVecOpts) *prometheus.GaugeVec {
	if c, ok := mw.gaugeVecs[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewGaugeVec(prometheus.GaugeOpts(opt.Opts), opt.Lables)
		mw.gaugeVecs[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) AddGaugeVec(name string, f float64, labels any) error {
	gauge, err := mw.parseGaugeVec(name, labels)
	if err != nil {
		return err
	}
	gauge.Add(f)
	return nil
}

func (mw *MetricsWatch) SubGaugeVec(name string, f float64, labels any) error {
	gauge, err := mw.parseGaugeVec(name, labels)
	if err != nil {
		return err
	}
	gauge.Sub(f)
	return nil
}

func (mw *MetricsWatch) IncGaugeVec(name string, labels any) error {
	gauge, err := mw.parseGaugeVec(name, labels)
	if err != nil {
		return err
	}
	gauge.Inc()
	return nil
}

func (mw *MetricsWatch) DecGaugeVec(name string, labels any) error {
	gauge, err := mw.parseGaugeVec(name, labels)
	if err != nil {
		return err
	}
	gauge.Dec()
	return nil
}

func (mw *MetricsWatch) SetGaugeVec(name string, f float64, labels any) error {
	gauge, err := mw.parseGaugeVec(name, labels)
	if err != nil {
		return err
	}
	gauge.Set(f)
	return nil
}

func (mw *MetricsWatch) SetGaugeVecToCurrentTime(name string, labels any) error {
	gauge, err := mw.parseGaugeVec(name, labels)
	if err != nil {
		return err
	}
	gauge.SetToCurrentTime()
	return nil
}

func (mw *MetricsWatch) RmGaugeVec(name string) error {
	if c, ok := mw.gaugeVecs[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.gaugeVecs, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) LsGaugeVec() string {
	data, err := json.Marshal(mw.gaugeVecs)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) NewHistogramVec(opt models.HistogramVecOpts) *prometheus.HistogramVec {
	if c, ok := mw.histogramVecs[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewHistogramVec(opt.HistogramOpts, opt.Lables)
		mw.histogramVecs[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) parseHistogramVec(name string, labels any) (prometheus.Observer, error) {
	if c, ok := mw.histogramVecs[name]; ok {
		switch v := labels.(type) {
		case []any:
			lvs := []string{}
			for _, label := range v {
				if s, ok := label.(string); ok {
					lvs = append(lvs, s)
				}
			}
			return c.WithLabelValues(lvs...), nil
		case map[string]any:
			lvs := make(map[string]string)
			for k, vv := range v {
				if s, ok := vv.(string); ok {
					lvs[k] = s
				}
			}
			return c.With(lvs), nil
		default:
			return nil, errors.New("invalid label type")
		}
	}
	return nil, errors.New("histogramVec not found")
}

func (mw *MetricsWatch) ObserveHistogramVec(name string, f float64, label any) error {
	histogram, err := mw.parseHistogramVec(name, label)
	if err != nil {
		return err
	}
	histogram.Observe(f)
	return nil
}

func (mw *MetricsWatch) RmHistogramVec(name string) error {
	if c, ok := mw.histogramVecs[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.histogramVecs, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) LsHistogramVec() string {
	data, err := json.Marshal(mw.histogramVecs)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) NewSummaryVec(opt models.SummaryVecOpts) *prometheus.SummaryVec {
	if c, ok := mw.summaryVecs[opt.Name]; ok {
		return c
	} else {
		c := prometheus.NewSummaryVec(opt.Adapter(), opt.Lables)
		mw.summaryVecs[opt.Name] = c
		prometheus.Register(c)
		return c
	}
}

func (mw *MetricsWatch) parseSummaryVec(name string, labels any) (prometheus.Observer, error) {
	if c, ok := mw.summaryVecs[name]; ok {
		switch v := labels.(type) {
		case []any:
			lvs := []string{}
			for _, label := range v {
				if s, ok := label.(string); ok {
					lvs = append(lvs, s)
				}
			}
			return c.WithLabelValues(lvs...), nil
		case map[string]any:
			lvs := make(map[string]string)
			for k, vv := range v {
				if s, ok := vv.(string); ok {
					lvs[k] = s
				}
			}
			return c.With(lvs), nil
		default:
			return nil, errors.New("invalid label type")
		}
	}
	return nil, errors.New("summaryVec not found")
}

func (mw *MetricsWatch) ObserveSummaryVec(name string, f float64, label any) error {
	summary, err := mw.parseSummaryVec(name, label)
	if err != nil {
		return err
	}
	summary.Observe(f)
	return nil
}

func (mw *MetricsWatch) RmSummaryVec(name string) error {
	if c, ok := mw.summaryVecs[name]; ok {
		if !prometheus.Unregister(c) {
			return errors.New("internal error")
		}
		delete(mw.summaryVecs, name)
		return nil
	}
	return nil
}

func (mw *MetricsWatch) LsSummaryVec() string {
	data, err := json.Marshal(mw.summaryVecs)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (mw *MetricsWatch) Snapshot(file ...string) error {
	metrics, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return err
	}
	filename := "tmp.txt"
	if len(file) > 0 {
		filename = file[0]
	}
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	for _, m := range metrics {
		data, err := json.Marshal(m)
		if err != nil {
			return err
		}
		_, err = fp.WriteString(fmt.Sprintf("%s\n", data))
		if err != nil {
			return err
		}
	}
	return nil
}

func (mw *MetricsWatch) Document() http.Handler {
	return promhttp.Handler()
}
