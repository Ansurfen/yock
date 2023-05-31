-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---
---{{.pathf}}
---
---@param path string
---@return string
---
function pathf(path)
    return ""
end

---
---{{.path}}
---
---@class path
---@field Separator string
---
path = {}

---
---{{.path_filename}}
---
---@param filepath string
---@return string
---
function path.filename(filepath)
    return ""
end

---
---{{.path_exist}}
---
---@param filepath string
---@return boolean
---
function path.exist(filepath)
    return false
end

---@vararg string
---@return string
function path.join(...)
    return ""
end

---@param path string
---@return string
function path.dir(path)
    return ""
end

---@param path string
---@return string
function path.base(path)
    return ""
end

---@param path string
---@return string
function path.clean(path)
    return ""
end

---@param path string
---@return string
function path.ext(path)
    return ""
end

---@param path string
---@return string, string
function path.abs(path)
    return "", ""
end
