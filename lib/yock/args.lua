--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field
---@diagnostic disable: lowercase-global
argBuilder = {}

---@return argBuilder
function argBuilder:new()
    local obj = {
        params = {}
    }
    setmetatable(obj, self)
    self.__index = self
    return obj
end

---@param cmd string
---@return argBuilder
function argBuilder:add(cmd)
    table.insert(self.params, cmd)
    return self
end

---@param cmd string
---@param v boolean|nil
---@return argBuilder
function argBuilder:add_bool(cmd, v)
    if v then
        table.insert(self.params, cmd)
    end
    return self
end

---@param cmd string
---@param v string|nil
---@return argBuilder
function argBuilder:add_str(cmd, v)
    if v ~= nil then
        table.insert(self.params, cmd)
    end
    return self
end

---@param format string
---@param v string|nil
---@return argBuilder
function argBuilder:add_strf(format, v)
    if v ~= nil then
        table.insert(self.params, string.format(format, v))
    end
    return self
end

---@param v any[]
---@return argBuilder
function argBuilder:add_arr(v)
    for _, value in ipairs(v) do
        table.insert(self.params, value)
    end
    return self
end

---@param ok boolean
---@param format string
---@vararg any
---@return argBuilder
function argBuilder:add_format(ok, format, ...)
    if ok then
        table.insert(self.params, string.format(format, ...))
    end
    return self
end

---@return string
function argBuilder:build()
    local arg = ""
    for _, v in ipairs(self.params) do
        arg = arg .. v .. " "
    end
    return arg
end

---@param opt table
---@return table, err
function argBuilder:exec(opt)
    return sh({
        redirect = opt["redirect"] or false,
    }, self:build())
end
