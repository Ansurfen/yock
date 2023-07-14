-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class context
local context = {}

---@param msg string
function context.info(msg)
end

---@param timeout? timeTime
function context.yeild(timeout)
end

---@param error string
function context.throw(error)
end


---@param name string
---@param callback fun(ctx: table)
function job(name, callback)
end

---
---{{.jobs}}
---
---@param name string
---@vararg string
---
function jobs(name, ...)
end

---
---{{.job_option}}
---
---@param opt table
---
function job_option(opt)
end
