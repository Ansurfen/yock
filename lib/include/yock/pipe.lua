-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field

---@meta _

---@class pipe
---@field type integer
---@field payload any
local pipe = {}

---@return pipe
function pipe:clone() end

---@vararg string
---@return pipe
function file(...) end

---@param str string
---@return pipe
function stream(str) end
