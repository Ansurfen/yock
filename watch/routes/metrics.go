// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ansurfen/yock/watch/metrics"
	"github.com/ansurfen/yock/watch/models"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsRoutes struct{}

func (router *MetricsRoutes) InstallMetricsAPI(group *gin.RouterGroup) {
	metricsRouter := group.Group("/metrics")
	{
		counterRouter := metricsRouter.Group("/counter")
		{
			counterRouter.POST("new", metricsCounterNew)
			counterRouter.POST("add", metricsCounterAdd)
			counterRouter.POST("inc", metricsCounterInc)
			counterRouter.POST("rm", metricsCounterRm)
			counterRouter.GET("ls", metricsCounterLs)
		}
		gaugeRouter := metricsRouter.Group("/gauge")
		{
			gaugeRouter.POST("new", metricsGaugeNew)
			gaugeRouter.POST("add", metricsGaugeAdd)
			gaugeRouter.POST("sub", metricsGaugeSub)
			gaugeRouter.POST("inc", metricsGaugeInc)
			gaugeRouter.POST("dec", metricsGaugeDec)
			gaugeRouter.POST("set", metricsGaugeSet)
			gaugeRouter.POST("setToCurrentTime", metricsGaugeToCuurentTime)
			gaugeRouter.POST("rm", metricsGaugeRm)
			gaugeRouter.GET("ls", metricsGaugeLs)
		}
		histogramRouter := metricsRouter.Group("/histogram")
		{
			histogramRouter.POST("new", metricsHistogramNew)
			histogramRouter.POST("observe", metricsHistogramObserve)
			histogramRouter.POST("rm", metricsHistogramRm)
			histogramRouter.GET("ls", metricsHistogramLs)
		}
		summaryRouter := metricsRouter.Group("/summary")
		{
			summaryRouter.POST("new", metricsSummaryNew)
			summaryRouter.POST("observe", metricsSummaryObserve)
			summaryRouter.POST("rm", metricsSummaryRm)
			summaryRouter.GET("ls", metricsSummaryLs)
		}
		counterVecRouter := metricsRouter.Group("/counterVec")
		{
			counterVecRouter.POST("new", metricsCounterVecNew)
			counterVecRouter.POST("add", metricsCounterVecAdd)
			counterVecRouter.POST("inc", metricsCounterVecInc)
			counterVecRouter.POST("rm", metricsCounterVecRm)
			counterVecRouter.GET("ls", metricsCounterVecLs)
		}
		gaugeVecRouter := metricsRouter.Group("/gaugeVec")
		{
			gaugeVecRouter.POST("new", metricsGaugeVecNew)
			gaugeVecRouter.POST("add", metricsGaugeVecAdd)
			gaugeVecRouter.POST("sub", metricsGaugeVecSub)
			gaugeVecRouter.POST("inc", metricsGaugeVecInc)
			gaugeVecRouter.POST("dec", metricsGaugeVecDec)
			gaugeVecRouter.POST("set", metricsGaugeVecSet)
			gaugeVecRouter.POST("setToCurrentTime", metricsGaugeVecToCuurentTime)
			gaugeVecRouter.POST("rm", metricsGaugeVecRm)
			gaugeVecRouter.GET("ls", metricsGaugeVecLs)
		}
		histogramVecRouter := metricsRouter.Group("/histogramVec")
		{
			histogramVecRouter.POST("new", metricsHistogramVecNew)
			histogramVecRouter.POST("observe", metricsHistogramVecObserve)
			histogramVecRouter.POST("rm", metricsHistogramVecRm)
			histogramVecRouter.GET("ls", metricsHistogramVecLs)
		}
		summaryVecRouter := metricsRouter.Group("/summaryVec")
		{
			summaryVecRouter.POST("new", metricsSummaryVecNew)
			summaryVecRouter.POST("observe", metricsSummaryVecObserve)
			summaryVecRouter.POST("rm", metricsSummaryVecRm)
			summaryVecRouter.GET("ls", metricsSummaryVecLs)
		}
		metricsRouter.GET("snapshot")
	}
}

// @Summary Create Counter Metrics
// @Description Create Counter Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.MetricsOpts true "Counter Option"
// @Success 200 {string} Success
// @Router /metrics/counter/new [post]
func metricsCounterNew(ctx *gin.Context) {
	opt := prometheus.CounterOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewCounter(opt)
	if c != nil {
		ctx.String(http.StatusOK, "success to new counter")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new counter")
}

// @Summary Add Counter Metrics
// @Description Add Counter Metrics
// @Tags Metrics
// @Param name formData string true "Counter Name"
// @Param f formData number true "Index" minimum(0) default(1)
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/counter/add [post]
func metricsCounterAdd(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	err = metrics.DefaultMetricsWatch.AddCounter(name, f)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to add counter",
		})
		return
	}
	ctx.String(http.StatusOK, "success to add counter")
}

// @Summary Inc Counter Metrics
// @Description Inc Counter Metrics
// @Tags Metrics
// @Param name formData string true "Counter Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/counter/inc [post]
func metricsCounterInc(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.IncCounter(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to add counter",
		})
		return
	}
	ctx.String(http.StatusOK, "success to add counter")
}

// @Summary Remove Counter Metrics
// @Description Remove Counter Metrics
// @Tags Metrics
// @Param name formData string true "Counter Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/counter/rm [post]
func metricsCounterRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmCounter(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove counter",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove counter")
}

// @Summary List Counter Metrics
// @Description List Counter Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/counter/ls [get]
func metricsCounterLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsCounter())
}

// @Summary Create Gauge Metrics
// @Description Create Gauge Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.MetricsOpts true "Gauge Option"
// @Success 200 {string} Success
// @Router /metrics/gauge/new [post]
func metricsGaugeNew(ctx *gin.Context) {
	opt := prometheus.GaugeOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewGauge(opt)
	if c != nil {
		ctx.String(http.StatusOK, "success to new counter")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new counter")
}

// @Summary Create Gauge Metrics
// @Description Create Gauge Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param name formData string true "Name"
// @Param f formData number true "Index"
// @Success 200 {string} Success
// @Router /metrics/gauge/add [post]
func metricsGaugeAdd(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	err = metrics.DefaultMetricsWatch.AddGauge(name, f)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to add gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to add gauge")
}

// @Summary Create Gauge Metrics
// @Description Create Gauge Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param name formData string true "Name"
// @Param f formData number true "Index"
// @Success 200 {string} Success
// @Router /metrics/gauge/sub [post]
func metricsGaugeSub(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	err = metrics.DefaultMetricsWatch.SubGauge(name, f)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to sub gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to sub gauge")
}

// @Summary Inc Gauge Metrics
// @Description Inc Gauge Metrics
// @Tags Metrics
// @Param name formData string true "Gauge Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gauge/inc [post]
func metricsGaugeInc(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.IncGauge(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to inc gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to inc gauge")
}

// @Summary Dec Gauge Metrics
// @Description Dec Gauge Metrics
// @Tags Metrics
// @Param name formData string true "Gauge Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gauge/dec [post]
func metricsGaugeDec(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.DecGauge(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to dec gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to dec gauge")
}

// @Summary Set Gauge Metrics
// @Description Set Gauge Metrics
// @Tags Metrics
// @Param name formData string true "Gauge Name"
// @Param f formData number true "Index"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gauge/set [post]
func metricsGaugeSet(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	err = metrics.DefaultMetricsWatch.SetGauge(name, f)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to set gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to set gauge")
}

// @Summary Remove Gauge Metrics
// @Description Remove Gauge Metrics
// @Tags Metrics
// @Param name formData string true "Gauge Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gauge/rm [post]
func metricsGaugeRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmGauge(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove gauge")
}

// @Summary Set Gauge Metrics To Current Time
// @Description Set Gauge Metrics To Current Time
// @Tags Metrics
// @Param name formData string true "Gauge Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gauge/setToCurrentTime [post]
func metricsGaugeToCuurentTime(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.SetGaugeToCurrentTime(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to set gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to set gauge")
}

// @Summary List Gauge Metrics
// @Description List Gauge Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/gauge/ls [get]
func metricsGaugeLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsGauge())
}

// @Summary Create Histogram Metrics
// @Description Create Histogram Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.HistogramOpts true "Histogram Option"
// @Success 200 {string} Success
// @Router /metrics/histogram/new [post]
func metricsHistogramNew(ctx *gin.Context) {
	opt := prometheus.HistogramOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewHistogram(opt)
	if c != nil {
		ctx.String(http.StatusOK, "success to new histogram")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new histogram")
}

// @Summary Observe Histogram Metrics
// @Description Observe Histogram Metrics
// @Tags Metrics
// @Param name formData string true "Histogram Name"
// @Param f formData number true "Index"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/histogram/observe [post]
func metricsHistogramObserve(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	err = metrics.DefaultMetricsWatch.ObserveHistogram(name, f)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to observe histogram",
		})
		return
	}
	ctx.String(http.StatusOK, "success to observe histogram")
}

// @Summary Remove Histogram Metrics
// @Description Remove Histogram Metrics
// @Tags Metrics
// @Param name formData string true "Histogram Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/histogram/rm [post]
func metricsHistogramRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmHistogram(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove histogram",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove histogram")
}

// @Summary List Histogram Metrics
// @Description List Histogram Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/histogram/ls [get]
func metricsHistogramLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsHistogram())
}

// @Summary Create Summary Metrics
// @Description Create Summary Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.SummaryOpts true "Summary Option"
// @Success 200 {string} Success
// @Router /metrics/summary/new [post]
func metricsSummaryNew(ctx *gin.Context) {
	opt := models.SummaryOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewSummaries(opt.Adapter())
	if c != nil {
		ctx.String(http.StatusOK, "success to new summary")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new summary")
}

// @Summary Observe Summary Metrics
// @Description Observe Summary Metrics
// @Tags Metrics
// @Param name formData string true "Summary Name"
// @Param f formData number true "Index"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/summary/observe [post]
func metricsSummaryObserve(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	err = metrics.DefaultMetricsWatch.ObserveSummaries(name, f)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to observe summary",
		})
		return
	}
	ctx.String(http.StatusOK, "success to observe summary")
}

// @Summary Remove Summary Metrics
// @Description Remove Summary Metrics
// @Tags Metrics
// @Param name formData string true "Summary Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/summary/rm [post]
func metricsSummaryRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmSummary(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove summary",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove summary")
}

// @Summary List Summary Metrics
// @Description List Summary Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/summary/ls [get]
func metricsSummaryLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsSummaries())
}

// @Summary Create CounterVec Metrics
// @Description Create CounterVec Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.MetricsVecOpts true "CounterVec Option"
// @Success 200 {string} Success
// @Router /metrics/counterVec/new [post]
func metricsCounterVecNew(ctx *gin.Context) {
	opt := models.MetricsVecOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewCounterVec(opt)
	if c != nil {
		ctx.String(http.StatusOK, "success to new counterVec")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new counterVec")
}

func extractLabels(ctx *gin.Context) (any, error) {
	var lvs any
	label := ctx.PostForm("label")
	err := json.Unmarshal([]byte(label), &lvs)
	return lvs, err
}

// @Summary Add CounterVec Metrics
// @Description Add CounterVec Metrics
// @Tags Metrics
// @Param name formData string true "CounterVec Name"
// @Param f formData number true "Index" minimum(0) default(1)
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/counterVec/add [post]
func metricsCounterVecAdd(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.AddCounterVec(name, f, lvs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err,
			"msg": "fail to add counterVec",
		})
		return
	}
	ctx.String(http.StatusOK, "success to add counterVec")
}

// @Summary Inc CounterVec Metrics
// @Description Inc CounterVec Metrics
// @Tags Metrics
// @Param name formData string true "Counter Name"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/counterVec/inc [post]
func metricsCounterVecInc(ctx *gin.Context) {
	name := ctx.PostForm("name")
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.IncCounterVec(name, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to add counter",
		})
		return
	}
	ctx.String(http.StatusOK, "success to add counter")
}

// @Summary Remove CounterVec Metrics
// @Description Remove CounterVec Metrics
// @Tags Metrics
// @Param name formData string true "Counter Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/counterVec/rm [post]
func metricsCounterVecRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmCounterVec(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove counterVec",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove counterVec")
}

// @Summary List CounterVec Metrics
// @Description List CounterVec Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/counterVec/ls [get]
func metricsCounterVecLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsCounterVec())
}

// @Summary Create GaugeVec Metrics
// @Description Create GaugeVec Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.MetricsVecOpts true "GaugeVec Option"
// @Success 200 {string} Success
// @Router /metrics/gaugeVec/new [post]
func metricsGaugeVecNew(ctx *gin.Context) {
	opt := models.MetricsVecOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewGaugeVec(opt)
	if c != nil {
		ctx.String(http.StatusOK, "success to new gaugeVec")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new gaugeVec")
}

// @Summary Create GaugeVec Metrics
// @Description Create GaugeVec Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param name formData string true "Name"
// @Param f formData number true "Index"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @Router /metrics/gaugeVec/add [post]
func metricsGaugeVecAdd(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.AddGaugeVec(name, f, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to add gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to add gauge")
}

// @Summary Create GaugeVec Metrics
// @Description Create GaugeVec Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param name formData string true "Name"
// @Param f formData number true "Index"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @Router /metrics/gaugeVec/sub [post]
func metricsGaugeVecSub(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.SubGaugeVec(name, f, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to sub gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to sub gauge")
}

// @Summary Inc GaugeVec Metrics
// @Description Inc GaugeVec Metrics
// @Tags Metrics
// @Param name formData string true "GaugeVec Name"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gaugeVec/inc [post]
func metricsGaugeVecInc(ctx *gin.Context) {
	name := ctx.PostForm("name")
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.IncGaugeVec(name, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to inc gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to inc gauge")
}

// @Summary Dec GaugeVec Metrics
// @Description Dec GaugeVec Metrics
// @Tags Metrics
// @Param name formData string true "GaugeVec Name"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gaugeVec/dec [post]
func metricsGaugeVecDec(ctx *gin.Context) {
	name := ctx.PostForm("name")
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.DecGaugeVec(name, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to dec gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to dec gauge")
}

// @Summary Set GaugeVec Metrics
// @Description Set GaugeVec Metrics
// @Tags Metrics
// @Param name formData string true "GaugeVec Name"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gaugeVec/set [post]
func metricsGaugeVecSet(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.SetGaugeVec(name, f, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to set gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to set gauge")
}

// @Summary Remove GaugeVec Metrics
// @Description Remove GaugeVec Metrics
// @Tags Metrics
// @Param name formData string true "GaugeVec Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gaugeVec/rm [post]
func metricsGaugeVecRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmGauge(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove gauge",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove gauge")
}

// @Summary Set GaugeVec Metrics To Current Time
// @Description Set GaugeVec Metrics To Current Time
// @Tags Metrics
// @Param name formData string true "GaugeVec Name"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/gaugeVec/setToCurrentTime [post]
func metricsGaugeVecToCuurentTime(ctx *gin.Context) {
	name := ctx.PostForm("name")
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.SetGaugeVecToCurrentTime(name, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to set gaugeVec",
		})
		return
	}
	ctx.String(http.StatusOK, "success to set gaugeVec")
}

// @Summary List GaugeVec Metrics
// @Description List GaugeVec Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/gaugeVec/ls [get]
func metricsGaugeVecLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsGaugeVec())
}

// @Summary Create HistogramVec Metrics
// @Description Create HistogramVec Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.HistogramVecOpts true "HistogramVec Option"
// @Success 200 {string} Success
// @Router /metrics/histogramVec/new [post]
func metricsHistogramVecNew(ctx *gin.Context) {
	opt := models.HistogramVecOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewHistogramVec(opt)
	if c != nil {
		ctx.String(http.StatusOK, "success to new histogramVec")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new histogramVec")
}

// @Summary Observe HistogramVec Metrics
// @Description Observe HistogramVec Metrics
// @Tags Metrics
// @Param name formData string true "HistogramVec Name"
// @Param f formData number true "Index"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/histogramVec/observe [post]
func metricsHistogramVecObserve(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.ObserveHistogramVec(name, f, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to observe histogramVec",
		})
		return
	}
	ctx.String(http.StatusOK, "success to observe histogramVec")
}

// @Summary Remove Histogram Metrics
// @Description Remove Histogram Metrics
// @Tags Metrics
// @Param name formData string true "Histogram Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/histogram/rm [post]
func metricsHistogramVecRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmHistogramVec(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove histogramVec",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove histogramVec")
}

// @Summary List HistogramVec Metrics
// @Description List HistogramVec Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/histogramVec/ls [get]
func metricsHistogramVecLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsHistogramVec())
}

// @Summary Create SummaryVec Metrics
// @Description Create SummaryVec Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Param data body models.SummaryVecOpts true "SummaryVec Option"
// @Success 200 {string} Success
// @Router /metrics/summaryVec/new [post]
func metricsSummaryVecNew(ctx *gin.Context) {
	opt := models.SummaryVecOpts{}
	if err := ctx.BindJSON(&opt); err != nil {
		ctx.String(http.StatusInternalServerError, "server internal error")
		return
	}
	c := metrics.DefaultMetricsWatch.NewSummaryVec(opt)
	if c != nil {
		ctx.String(http.StatusOK, "success to new summaryVec")
		return
	}
	ctx.String(http.StatusNotModified, "fail to new summaryVec")
}

// @Summary Observe SummaryVec Metrics
// @Description Observe SummaryVec Metrics
// @Tags Metrics
// @Param name formData string true "SummaryVec Name"
// @Param f formData number true "Index"
// @Param label formData string true "Label"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/summaryVec/observe [post]
func metricsSummaryVecObserve(ctx *gin.Context) {
	name := ctx.PostForm("name")
	f, err := strconv.ParseFloat(ctx.PostForm("f"), 64)
	if err != nil {
		ctx.String(http.StatusNotModified, "illegal the range of float")
		return
	}
	lvs, err := extractLabels(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	err = metrics.DefaultMetricsWatch.ObserveSummaryVec(name, f, lvs)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to observe summaryVec",
		})
		return
	}
	ctx.String(http.StatusOK, "success to observe summaryVec")
}

// @Summary Remove SummaryVec Metrics
// @Description Remove SummaryVec Metrics
// @Tags Metrics
// @Param name formData string true "SummaryVec Name"
// @Success 200 {string} Success
// @failure 304 {object} string
// @Router /metrics/summaryVec/rm [post]
func metricsSummaryVecRm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	err := metrics.DefaultMetricsWatch.RmSummaryVec(name)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"err": err,
			"msg": "fail to remove summaryVec",
		})
		return
	}
	ctx.String(http.StatusOK, "success to remove summaryVec")
}

// @Summary List SummaryVec Metrics
// @Description List SummaryVec Metrics
// @Tags Metrics
// @Success 200 {string} Success
// @Router /metrics/summaryVec/ls [get]
func metricsSummaryVecLs(ctx *gin.Context) {
	ctx.String(http.StatusOK, metrics.DefaultMetricsWatch.LsSummaryVec())
}
