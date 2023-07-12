-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field

---@meta _

json = {}

---@param v any
---@vararg string
---@return string
function json.encode(v, ...)
    return ""
end

---@param str string
---@return table
function json.decode(str)
    return {}
end

---@class jsonobject
---@field buf table
---@field file string
jsonobject = {}

---@param file string
---@param str? string
---@return jsonobject
function json.create(file, str) end

---@param file string
---@return jsonobject
function json.open(file) end

---@param k string
---@return boolean
function jsonobject:getbool(k) end

---@param k string
---@return number
function jsonobject:getnumber(k) end

---@param k string
---@return string
function jsonobject:getstr(k) end

---@param k string
---@return table
function jsonobject:gettable(k) end

---@param k string
---@return any
function jsonobject:get(k) end

---@param k string
---@param v any
function jsonobject:set(k, v) end

---@return string
function jsonobject:string() end

---@param pretty? boolean
function jsonobject:save(pretty) end
