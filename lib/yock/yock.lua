--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global

-- require layer
--
-- yock overloads package.path to ensure the validity
-- of relative and absolute path search modules.
package.path = "?.lua"

env.yock_tmp = path.join(env.yock_path, "tmp")
env.yock_bin = path.join(env.yock_path, "bin")

---@param var string
---@return string
read = function(var)
    local line = io.read("*l")
    alias(var, line)
    return line
end

---@param file string
---@param data string
write = function(file, data)
    local _, err = write_file(file, data)
    -- local _, err = echo({
    --     fd = { file },
    --     mode = "c|t|rw"
    -- }, data)
    yassert(err)
end

---@param fileType string
---@param want table<string, string>
---@param opt table<string, string>
---@return string
multi_fetch = function(fileType, want, opt)
    local file, err
    local elements = {}
    for _, value in pairs(opt) do
        table.insert(elements, function(v)
            file, err = fetch.file(strf(value, want), fileType)
            if err ~= nil then
                return err
            end
            return ""
        end)
    end
    loadbalance({
        maxRetry = #elements + 1
    }, elements)
    return file
end

---@param todo table
---@param handle function
---@return table
multi_bind = function(todo, handle)
    local ret = {}
    for _, v in ipairs(todo) do
        table.insert(ret, function()
            local err = handle(v)
            if err ~= nil then
                return err
            end
            return ""
        end)
    end
    return ret
end

---@param s string
---@return string
wrapzip = function(s)
    return s .. env.platform:Zip()
end

---@param s string
---@return string
wrapexf = function(s)
    return s .. env.platform:Exf()
end

---@param s string
---@return string
wrapscript = function(s)
    return s .. env.platform:Script()
end

---@param src string
---@param dst string
rename = function(src, dst)
    mv(src, dst)
end
