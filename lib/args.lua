--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global
argBuilder = {}

function argBuilder:new()
    local obj = {
        params = {}
    }
    setmetatable(obj, self)
    self.__index = self
    return obj
end

function argBuilder:add(cmd)
    table.insert(self.params, cmd)
    return self
end

function argBuilder:add_bool(cmd, v)
    if v then
        table.insert(self.params, cmd)
    end
    return self
end

function argBuilder:add_str(cmd, v)
    if v ~= nil then
        table.insert(self.params, cmd)
    end
    return self
end

function argBuilder:build()
    local arg = ""
    for _, v in ipairs(self.params) do
        arg = arg .. v .. " "
    end
    return arg
end

function argBuilder:exec(opt)
    return sh({
        debug = opt["debug"] or false,
        redirect = opt["redirect"] or false,
    }, self:build())
end
