-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@param e err
---@param msg? any
function yassert(e, msg)
end

ycache = {}

---@param k string
---@param v string
function ycache:put(k, v)
end

---@param k string
---@return string|nil
function ycache:get(k)
end

function ycache:free()
end
