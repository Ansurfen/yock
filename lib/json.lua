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
