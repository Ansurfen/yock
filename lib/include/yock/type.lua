-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@alias err string | nil | userdata

-----@class err: string

---@alias starType
---| BooleanType
---| StringType
---| StringArrayType

---@class BooleanType
local BooleanType = {}

---@param b boolean
---@return BooleanType
function Boolean(b) end

---@return userdata
function BooleanType:Ptr() end

---@return boolean
function BooleanType:Var() end

---@class StringType
local StringType = {}

---@param str string
---@return StringType
function String(str) end

---@return userdata
function StringType:Ptr() end

---@return string
function StringType:Var() end

---@class StringArrayType
local StringArrayType = {}

---@vararg string
---@return StringArrayType
function StringArray(...) end

---@return userdata
function StringArrayType:Ptr() end

---@return table
function StringArrayType:Var() end

---@param m userdata
---@return table
function map2Table(m) end
