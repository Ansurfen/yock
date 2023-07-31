--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.
---@diagnostic disable: duplicate-set-field
---@diagnostic disable: lowercase-global
---@class assign
assign = {}

---@param a string
---@param b string
---@return string
-- a = b when b exist
function assign.string(a, b)
    if type(b) == "string" then
        return b
    end
    return a
end

---@param a number
---@param b number
---@return number
-- a = b when b exist
function assign.number(a, b)
    if type(b) == "number" then
        return b
    end
    return a
end

---@param a boolean
---@param b boolean
---@return boolean
-- a = b when b exist
function assign.bool(a, b)
    if type(b) == "boolean" then
        return b
    end
    return a
end

---@param a table
---@param b table
---@return table
-- a = b when b exist
function assign.table(a, b)
    if type(b) == "table" then
        return b
    end
    return a
end

---@param a function
---@param b function
---@return function
-- a = b when b exist
function assign.func(a, b)
    if type(b) == "function" then
        return b
    end
    return a
end
