-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.
---@diagnostic disable: lowercase-global
---@diagnostic disable: duplicate-set-field

---@class conf
---@field buf table
---@field viper Viper
conf = {}

---@param file string
---@param tmpl string
---@return conf
function conf.create(file, tmpl)
    if not find(file) then
        write(file, tmpl)
        yassert(string.format("please complete context in %s", file))
    end
    return conf.open(file)
end

---@param file string
---@return conf
function conf.open(file)
    local viper, err = open_conf(file)
    yassert(err)
    local obj = {
        viper = viper
    }
    obj.buf = map2Table(obj.viper:AllSettings())
    setmetatable(obj, { __index = conf })
    return obj
end

---@param k string
---@return table|nil
function conf:read(k)
    local keys = strings.Split(k, ".")
    local x = self.buf
    for _, key in ipairs(keys) do
        if x == nil then
            return nil
        end
        x = x[key]
    end
    return x
end

---@param k string
---@param v any
function conf:write(k, v)
    self.viper:Set(k, v)
end

function conf:save()
    self.viper:WriteConfig()
end
