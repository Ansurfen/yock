-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

reglib = {}

---@param patterns table
---@return table
function reglib:new(patterns)
end

---@param p string
---@param s string
---@return string|nil
function reglib:find_str(p, s)
end

---@param p string
---@param s string
---@return boolean
function reglib:match_str(p, s)
end
