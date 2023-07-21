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

---@class option_todo_ycho
---@field stdout boolean
local option_todo_ycho = {}

---@class option_todo
---@field ycho option_todo_ycho
---@field strict boolean
---@field sync boolean
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
