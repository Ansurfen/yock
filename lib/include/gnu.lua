-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@param str string
---@param debug? boolean
---@return string
function echo(str, debug)
    return ""
end

---@return string, err
function whoami()
    return ""
end

function clear()
end

---@param dir string
---@return err
function cd(dir)
end

---@param file string
---@return err
function touch(file)
end

---@param file string
---@return err
function cat(file)
end

---@param opt table
---@return table|string, err
function ls(opt)
    return {}
end

---@param name string
---@param mode number
---@return err
function chmod(name, mode)
end

---@param name string
---@param uid number
---@param gid number
---@return err
function chown(name, uid, gid)
end

---@param dir string
function mkdir(dir)
end

---@param src string
---@param dst string
function cp(src, dst)
end

---@param src string
---@param dst string
function mv(src, dst)
end

---@return string, err
function pwd()
    return ""
end

---
---@class rmOpt
---@field safe boolean
---@field pattern string
---
---{{.rm_opt}}
---
---
local rmOpt = {}

---
--- {{.rm}}
---
---@param opt rmOpt
---@vararg string
---
function rm(opt, ...)
end
