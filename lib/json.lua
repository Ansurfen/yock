--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field
---@diagnostic disable: duplicate-set-field

---@class jsonfile
---@field fp? file*
---@field buf table
---@field filename string
jsonfile = {}

---@param filename string
---@param no_strict? boolean
---@return jsonfile
function jsonfile:open(filename, no_strict)
    local obj = {}
    local fp, err = io.open(filename, "r+")
    if no_strict == nil then
        no_strict = false
    end
    if err ~= nil and not no_strict then
        yassert(err)
    end
    obj.fp = fp
    obj.filename = filename
    if type(fp) ~= "nil" then
        obj.buf = json.decode(fp:read("*a"))
    elseif type(no_strict) == "boolean" and no_strict then
    else
        yassert("invalid file")
    end
    setmetatable(obj, { __index = self })
    return obj
end

---@param filename string
---@return jsonfile
function jsonfile:create(filename)
    local obj = {
        filename = filename,
        buf = {}
    }
    setmetatable(obj, { __index = self })
    return obj
end

---@param k string
---@return any
function jsonfile:read(k)
    local keys = strings.Split(k, ".")
    local v = self.buf
    for _, kk in ipairs(keys) do
        if v == nil then
            return nil
        end
        v = v[kk]
    end
    return v
end

---@param pretty? boolean
function jsonfile:write(pretty)
    if pretty then
        yassert(write_file(self.filename, json.encode(self.buf, "", "    ")))
        return
    end
    yassert(write_file(self.filename, json.encode(self.buf)))
end

function jsonfile:close()
    self.fp:close()
end

---@class jsonobject
---@field buf table
---@field file string
jsonobject = {}

---@param file string
---@param str? string
---@return jsonobject
function json.create(file, str)
    local obj = {
        file = file
    }
    if not find(file) then
        write(file, str or "{}")
    end
    local stream, err = cat(file)
    yassert(err)
    obj.buf = json.decode(stream)
    setmetatable(obj, { __index = jsonobject })
    return obj
end

---@param file string
---@return jsonobject
function json.open(file)
    local obj = {
        file = file
    }
    local stream, err = cat(file)
    yassert(err)
    obj.buf = json.decode(stream)
    setmetatable(obj, { __index = jsonobject })
    return obj
end

---@param k string
---@return boolean
function jsonobject:getbool(k)
    local v = self:get(k)
    if type(v) == "string" then
        return v
    end
    return false
end

---@param k string
---@return number
function jsonobject:getnumber(k)
    local v = self:get(k)
    if type(v) == "number" then
        return v
    end
    return 0
end

---@param k string
---@return string
function jsonobject:getstr(k)
    local v = self:get(k)
    if type(v) == "string" then
        return v
    end
    return ""
end

---@param k string
---@return table
function jsonobject:gettable(k)
    local v = self:get(k)
    if type(v) == "table" then
        return v
    end
    return {}
end

---@param k string
---@return any
function jsonobject:get(k)
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
function jsonobject:set(k, v)
    local fields = strings.Split(k, ".")
    local currentTable = self.buf
    for i = 1, #fields - 1 do
        local field = fields[i]
        if not currentTable[field] then
            currentTable[field] = {}
        end
        currentTable = currentTable[field]
    end
    currentTable[fields[#fields]] = v
end

---@return string
function jsonobject:string()
    return json.encode(self.buf)
end

---@param pretty? boolean
function jsonobject:save(pretty)
    if pretty then
        write(self.file, json.encode(self.buf, "", "    "))
        return
    end
    write(self.file, json.encode(self.buf))
end
