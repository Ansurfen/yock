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
    return exec({
        debug = opt["debug"] or false,
        redirect = opt["redirect"] or false,
    }, self:build())
end
