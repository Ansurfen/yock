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

---Boolean convert lua.boolean to golang.bool type
---@param b boolean
---@return BooleanType
function Boolean(b) end

---Ptr returns value's pointer address
---@return userdata
function BooleanType:Ptr() end

---Var returns concrete value
---@return boolean
function BooleanType:Var() end

---@class StringType
local StringType = {}

---String convert lua.string to golang.string type
---@param str string
---@return StringType
function String(str) end

---Ptr returns value's pointer address
---@return userdata
function StringType:Ptr() end

---Var returns concrete value
---@return string
function StringType:Var() end

---@class StringArrayType
local StringArrayType = {}

---StringArray convert lua.string[] to golang.string[] type
---@vararg string
---@return StringArrayType
function StringArray(...) end

---Ptr returns value's pointer address
---@return userdata
function StringArrayType:Ptr() end

---Var returns concrete value
---@return table
function StringArrayType:Var() end

---@param m userdata
---@return table
function map2Table(m) end
