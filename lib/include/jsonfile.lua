-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field

---@meta _

---@class jsonfile
---@field fp? file*
---@field buf table
---@field filename string
jsonfile = {}

---@param filename string
---@param no_strict? boolean
---@return jsonfile
function jsonfile:open(filename, no_strict)
    return {}
end

---@param filename string
---@return jsonfile
function jsonfile:create(filename)
end

---@param k string
---@return any
function jsonfile:read(k)
end

---@param pretty? boolean
function jsonfile:write(pretty)
end

function jsonfile:close()
end
