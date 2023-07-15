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
function argBuilder:add_bool(cmd, v)
    if v then
        table.insert(self.params, cmd)
    end
    return self
end

---@param cmd string
---@param v string|nil
function argBuilder:add_str(cmd, v)
    if v ~= nil then
        table.insert(self.params, cmd)
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
        debug = opt["debug"] or false,
        redirect = opt["redirect"] or false,
    }, self:build())
end
