-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class conf
conf = {}

---@param file string
---@return conf
function conf.open(file) end

---@param file string
---@param tmpl string
---@return conf
function conf.create(file, tmpl) end

---@param k string
---@return table|nil
function conf:read(k) end

---@param k string
---@param v any
function conf:write(k, v) end

function conf:save() end
