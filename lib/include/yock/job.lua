-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---context is designed for handling the lifetime of task
---and composing jobs
---@class context
---@field platform platform
---@field args string[]
---@field task string
---@field flags? table<string, any>
local context = {}

---@alias ec integer
---|> 0 # abort all peer jobs (default)
---| 1 # continue to run peer jobs
---| 2 # continue to run peer jobs with inherit

---exit aborts running for the context, 
---and decides whether to continue the next
---according to code's value. Its scope only
---is limited to task, not impacting other tasks.
---@param code? ec
function context.exit(code) end

---@generic T
---@param ok? T
---@param msg? string
function context.assert(ok, msg) end

---throw uses os.exit to abort entire program with
---force, and prints the info of error.
---@param error string
function context.throw(error) end

---@param timeout? integer
function context.yield(timeout) end

---@vararg string
function context.resume(...) end

---put sets keyed value in yock's memory database,
---which can be taken by `context.get()` and the lifetime
---is in total program, but any context.
---### Example:
---```lua
---# assumes yock is running in windows
---job("", function (ctx)
---    print(ctx.put("a", 10))
---    print(ctx.get("a") == 10)
---end)
---```
---@param k string
---@param v any
function context.put(k, v) end

---get takes value from yock's memory database
---### Example:
---```lua
---# assumes yock is running in windows
---job("", function (ctx)
---    print(ctx.put("a", 10))
---    print(ctx.get("a") == 10)
---end)
---```
---@param k string
---@return any
function context.get(k) end

---set_os sets the os of platform field in context,
---which returns corresponding platform information.
---### Example:
---```lua
---# assumes yock is running in windows
---job("", function (ctx)
---    print(ctx.platform:Script()) -- .bat
---    ctx.set_os("linux")
---    print(ctx.platform:Script()) -- .sh
---end)
---```
---@param os string
function context.set_os(os) end

---job is the smallest component of a task, a task can consist of one or more jobs.
---It is not difficult to see from the function signature that each job is a unit
---whose name is bound to the callback function. If a user defines a job with the same
---name in the same file, Yock will throw an error, so each job name must be unique.
---
---### Example:
---```lua
---# main.lua
---
---job("test", function(ctx)
---     ctx.info("Hello World!")
---end)
---
---# use `yock run main.lua test` to run test task for the above code.
---```
---@param name string
---@param callback fun(ctx: context)
function job(name, callback) end

---jobs composes multiple jobs to form a task and share the namespace with the job.
---This means that if jobs and job have the same name, yock will also throw an error directly.
---### Example:
---```lua
---# main.lua
---job("test", function(ctx)
---    print("start test...")
---end)
---
---job("build", function(ctx)
---    print("start build...")
---end)
---
---job("deploy", function(ctx)
---    print("start deploy...")
---end)
---
---jobs("all", "test", "build", "deploy")
---# just like scheduling job, use `yock run main.lua all` to run all task one by one.
---# if you want to run multiple task with cover at the same time, it's also supported
---# and use the form of `yock run main.lua all deploy` to make it.
---```
---@param name string
---@vararg string
function jobs(name, ...) end

---@class option_ycho
---@field stdout? boolean # allows logger print on terminal, if true

---@class option_yockd
---@field name string
---@field peer? table<string, option_yockd_peer>
---@field self_boot? boolean
---@field port? integer

---@class option_yockd_peer
---@field ip string
---@field port integer
---@field public boolean

---@class option_yockw
---@field self_boot? boolean
---@field port? integer
---@field metrics?
---|>table<"couter", option_yockw_metrics_counter[]>
---|table<"gauge", option_yockw_metrics_gauge[]>
---|table<"histogram", option_yockw_metrics_hisogram[]>
---|table<"summary", option_yockw_metrics_summary[]>
---|table<"resolved", string[]>

---@class option_yockw_metrics_counter
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string
---@field label? string[]

---@class option_yockw_metrics_gauge
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string

---@class option_yockw_metrics_hisogram
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string
---@field buckets number[]

---@class option_yockw_metrics_summary
---@field namespace? string
---@field subsystem? string
---@field name string
---@field help string
---@field objectives table<string, number>
---@field max_age? integer
---@field buf_cap? integer

---@class option_todo
---@field ycho? option_ycho
---@field yockd? option_yockd
---@field yockw? option_yockw
---@field strict? boolean # automatically panic when error occurs.
---@field sync? boolean # synchronize to configuration (yock.yaml), and it's not recommended.

---option can reset yock configuration (yock.yaml) at runtime,
---and you can think it's an local or temporary environment.
---General said, most setting is effective.
---You also can add sync field to synchronize configuration,
---but it's not recommended and possible to destroy global.
---### Option:
---* strict?, boolean, catches error and panic when error occurs.
---* sync?, boolean, synchronize to configuration (yock.yaml), and it's not recommended.
---
---### Example:
---```lua
---option({
---    ycho = { stdout = true },
---    strict = false,
---})
---```
---@param opt option_todo
function option(opt) end
