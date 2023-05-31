-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@param file string
---@return string, err
function read_file(file)
    return ""
end

---@param file string
---@param data string
function safe_write(file, data)
end

---@param zipPath string
---@vararg string
---@return err
function zip(zipPath, ...)
end

---@param src string
---@param dst string
---@return err
function unzip(src, dst)
end

---@param file string
---@param data string
---@return err
function write_file(file, data)
end

---@param path string
---@return boolean
function is_exist(path)
    return false
end
