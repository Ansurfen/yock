-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global
---@diagnostic disable: duplicate-doc-field
---@diagnostic disable: missing-fields
---@diagnostic disable: duplicate-set-field

---@type pipe
local pipe = {}

local pipe_type = {
    stream = 1,
    file = 2,
}

---@vararg string
---@return pipe
function file(...)
    local files = { ... }
    local obj = {
        type = pipe_type.file,
        payload = files
    }
    for _, filename in ipairs(files) do
        if not find(filename) then
            write(filename, "")
        end
    end
    setmetatable(obj, pipe)
    obj.clone = pipe.clone
    return obj
end

---@param str string
---@return pipe
function stream(str)
    local obj = {
        payload = str,
        type = pipe_type.stream
    }
    setmetatable(obj, pipe)
    return obj
end

---@return pipe
function pipe:clone()
    if self.type ~= pipe_type.file then
        yassert("invalid file pipe")
    end
    local tmp = {}
    for _, value in ipairs(self.payload) do
        table.insert(tmp, value)
    end
    local obj = {
        type = pipe_type.file,
        payload = tmp
    }
    setmetatable(obj, pipe)
    obj.clone = pipe.clone
    return obj
end

---@param a pipe
---@param b pipe
---@return err
pipe.__lt  = function(a, b)
    if not (a.type == pipe_type.file and b.type == pipe_type.stream) then
        yassert("pipe type not matched")
    end
    for _, file in ipairs(a.payload) do
        write(file, b.payload)
    end
    return nil
end

---@param a pipe
---@param b pipe
---@return err
pipe.__le  = function(a, b)
    if not (a.type == pipe_type.file and b.type == pipe_type.stream) then
        yassert("pipe type not matched")
    end
    for _, file in ipairs(a.payload) do
        write(file, cat(file) .. b.payload)
    end
    return nil
end

---@param a pipe
---@param b pipe
---@return pipe
pipe.__add = function(a, b)
    if not (a.type == pipe_type.file and b.type == pipe_type.file) then
        yassert("invalid file pipe")
    end
    local c = a:clone()
    for _, v in ipairs(b.payload) do
        table.insert(c.payload, v)
    end
    return c
end

---@param a pipe
---@param b pipe
---@return pipe
pipe.__sub = function(a, b)
    if not (a.type == pipe_type.file and b.type == pipe_type.file) then
        yassert("invalid file pipe")
    end
    local c = a:clone()
    for _, bf in ipairs(b.payload) do
        for i, af in ipairs(c.payload) do
            if af == bf then
                table.remove(c.payload, i)
            end
        end
    end
    return c
end
