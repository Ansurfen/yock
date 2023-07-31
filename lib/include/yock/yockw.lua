-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class yockw
---@field metrics yockw_metrics
yockw = {}

---@class yockw_metrics
---@field counter yockw_metrics_counter
---@field counter_vec yockw_metrics_counter_vec
local metrics = {}

---@class yockw_metrics_counter
local metrics_counter = {}

---@param opt option_yockw_metrics_counter
function metrics_counter.new(opt) end

---@param name string
---@param f number
function metrics_counter.add(name, f) end

---@param name string
function metrics_counter.inc(name) end

---@param name string
function metrics_counter.rm(name) end

---@return string[]
function metrics_counter.ls() end

---@class yockw_metrics_counter_vec
local metrics_counter_vec = {}

---@param name string
---@param f number
---@vararg string
function metrics_counter_vec.add(name, f, ...) end

---@param name string
---@param f number
---@param label table<string, string>
function metrics_counter_vec.add(name, f, label) end

---@param name string
---@vararg string
function metrics_counter_vec.inc(name, ...) end

---@param name string
---@param label table<string, string>
function metrics_counter_vec.inc(name, label) end

---@class yockw_metrics_gauge
local metrics_gauge = {}

---@class yockw_metrics_gauge_vec
local metrics_gauge_vec = {}

---@class yockw_metrics_hisogram
local metrics_histogram = {}

---@class yockw_metrics_hisogram_vec
local metrics_histogram_vec = {}

---@class yockw_metrics_summary
local metrics_summary = {}

---@class yockw_metrics_summary_vec
local metrics_summary_vec = {}
