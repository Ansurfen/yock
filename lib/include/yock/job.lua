-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class context
---@field platform platform
---@field args table<string>
---@field task string
---@field flags? table<string, any>
local context = {}

---@param msg string
function context.info(msg) end

---@alias ec integer
---|> 0 # abort all peer jobs (default)
---| 1 # continue to run peer jobs
---| 2 # continue to run peer jobs with inherit

---@param code? ec
function context.exit(code) end

---@generic T
---@param ok? T
---@param msg? string
function context.assert(ok, msg) end

---@param error string
function context.throw(error) end

---@param timeout? integer
function context.yield(timeout) end

---@vararg string
function context.resume(...) end

---@param k string
---@param v any
function context.put(k, v) end

---@param k string
---@return any
function context.get(k) end

---@param os string
function context.set_os(os) end

---@param name string
---@param callback fun(ctx: context)
function job(name, callback) end

-- Example:
-- ```lua
-- job("j1", function(ctx) print("j1") end)
-- job("j2", function(ctx) print("j2") end)
-- jobs("all", "j1", "j2")
-- ```
---@param name string
---@vararg string
function jobs(name, ...) end

---@class option_ycho
---@field stdout? boolean
local option_ycho = {}

---@class option_yockd
---@field name string
---@field peer? table<string, option_yockd_peer>
---@field self_boot? boolean
---@field port? integer
local option_yockd = {}

---@class option_yockd_peer
---@field ip string
---@field port integer
---@field public boolean
local option_yockd_peer = {}

---@class option_yockw
---@field self_boot? boolean
---@field port? integer
---@field metrics?
---|>table<"couter", option_yockw_metrics_counter[]>
---|table<"gauge", option_yockw_metrics_gauge[]>
---|table<"histogram", option_yockw_metrics_hisogram[]>
---|table<"summary", option_yockw_metrics_summary[]>
---|table<"resolved", string[]>
local option_yockw = {}

---@class option_yockw_metrics_counter
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string
---@field label? string[]
local yockw_metrics_counter = {}

---@class option_yockw_metrics_gauge
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string
local yockw_metrics_gauge = {}

---@class option_yockw_metrics_hisogram
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string
---@field buckets number[]
local yockw_metrics_hisogram = {}

---@class option_yockw_metrics_summary
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string
---@field objectives table<string, number>
---@field max_age? integer
---@field buf_cap? integer
local yockw_metrics_summary = {}

---@class option_todo
---@field ycho? option_ycho
---@field yockd? option_yockd
---@field yockw? option_yockw
---@field strict? boolean
---@field sync? boolean
local option_todo = {}

-- Example:
-- ```lua
-- option({
--     ycho = { stdout = true },
--     strict = false,
-- })
-- ```
---@param opt option_todo
function option(opt) end
