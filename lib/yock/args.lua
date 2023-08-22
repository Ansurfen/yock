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

---@vararg string
---@return argBuilder
function argBuilder:add(...)
    for _, value in ipairs({ ... }) do
        table.insert(self.params, value)
    end
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
    return table.concat(self.params, " ")
end

---@param opt table|nil
---@return table, err
function argBuilder:exec(opt)
    if opt == nil then
        opt = {}
    end
    return sh({
        redirect = opt["redirect"] or false,
    }, self:build())
end

---@type args_builder
---@diagnostic disable-next-line: missing-fields
args_builder = {}

---@param header? string
---@return args_builder
function args_builder.new(header)
    local obj = {
        available = true,
        params = {},
        combine = {},
        header = header
    }
    setmetatable(obj, { __index = args_builder })
    return obj
end

---@vararg string
function args_builder:reg_combine(...)
    for _, v in ipairs({ ... }) do
        self.combine[v] = false
    end
    return self
end

---@param ok boolean
---@param k string
function args_builder:add_combine(ok, k)
    if ok then
        if self.combine[k] ~= nil then
            self.combine[k] = true
        end
    end
    return self
end

---@param ok boolean
---@param k string
function args_builder:add_combine_must(ok, k)
    if not ok then
        self.available = false
        return self
    end
    return self:add_combine(ok, k)
end

---@vararg string
function args_builder:add(...)
    for _, v in ipairs({ ... }) do
        table.insert(self.params, v)
    end
    return self
end

---@param ok boolean
---@param format string
---@vararg any
function args_builder:add_str(ok, format, ...)
    if ok then
        table.insert(self.params, string.format(format, ...))
    end
    return self
end

---@param ok boolean
---@param format string
---@param ... any
---@return args_builder
function args_builder:add_str_must(ok, format, ...)
    if ok then
        table.insert(self.params, string.format(format, ...))
    else
        self.available = false
    end
    return self
end

---@param ok boolean
---@param format string
---@param arr any[]
function args_builder:add_arr(ok, format, arr)
    if ok and type(arr) == "table" then
        for _, v in ipairs(arr) do
            table.insert(self.params, string.format(format, v))
        end
    end
    return self
end

---@param ok boolean
---@param format string
---@param arr any[]
function args_builder:add_arr_must(ok, format, arr)
    if not ok then
        self.available = false
        return self
    end
    return self:add_arr(ok, format, arr)
end

---@param ok boolean
---@param v string
function args_builder:add_bool(ok, v)
    if ok then
        table.insert(self.params, v)
    end
    return self
end

---@param ok boolean
---@param v string
function args_builder:add_bool_must(ok, v)
    if not ok then
        self.available = false
        return self
    end
    return self:add_bool(ok, v)
end

---@param ok boolean
---@param format string
---@param map table
function args_builder:add_map(ok, format, map)
    if ok then
        local tmp = {}
        for k, v in pairs(map) do
            if type(v) == "boolean" and v then
                table.insert(tmp, k)
            elseif type(v) == "table" then
                for _, vv in ipairs(v) do
                    table.insert(tmp, string.format("%s %s", k, vv))
                end
            end
        end
        table.insert(self.params, string.format(format, table.concat(tmp, " ")))
    end
    return self
end

---@param ok boolean
---@param format string
---@param map table
function args_builder:add_map_must(ok, format, map)
    if not ok then
        self.available = false
        return self
    end
    return self:add_map(ok, format, map)
end

---@param v string
function args_builder:set_header(v)
    self.header = v
    return self
end

---@return string
function args_builder:build()
    if not self.available then
        return ""
    end
    local combine = ""
    for k, v in pairs(self.combine) do
        if v then
            combine = combine .. k
        end
    end
    if #combine > 0 then
        table.insert(self.params, 1, combine)
    end
    if self.header ~= nil then
        table.insert(self.params, 1, self.header)
    end
    return table.concat(self.params, " ")
end

---@param opt sh_opt
---@return string[]|nil, err
function args_builder:exec(opt)
    local script = self:build()
    if #script == 0 then
        return nil, "invalid command"
    end
    return sh(opt or { redirect = false, quiet = true }, self:build())
end
