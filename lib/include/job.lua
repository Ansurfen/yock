-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---
---{{.job}}
---
---@param name string
---@param callback fun(cenv: table):boolean
---
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
