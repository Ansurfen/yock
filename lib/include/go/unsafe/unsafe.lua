-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

unsafe = {}

---@param x unsafeArbitraryType
---@return any
function unsafe.Offsetof(x)
end

---@param x unsafeArbitraryType
---@return any
function unsafe.Alignof(x)
end

---@param ptr unsafePointer
---@param len unsafeIntegerType
---@return unsafePointer
function unsafe.Add(ptr, len)
end

---@param ptr unsafeArbitraryType
---@param len unsafeIntegerType
---@return any
function unsafe.Slice(ptr, len)
end

---@param slice any
---@return unsafeArbitraryType
function unsafe.SliceData(slice)
end

---@param ptr byte
---@param len unsafeIntegerType
---@return string
function unsafe.String(ptr, len)
end

---@param str string
---@return byte
function unsafe.StringData(str)
end

---@param x unsafeArbitraryType
---@return any
function unsafe.Sizeof(x)
end

---@class unsafeArbitraryType
local unsafeArbitraryType = {}

---@class unsafeIntegerType
local unsafeIntegerType = {}

---@class unsafePointer
local unsafePointer = {}
