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
---@return string, err
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

---@vararg string
---@return err
function mkdir(...)
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

---@param cmd table|string
function sudo(cmd)
end

---@param opt table
function grep(opt)
end

---@param opt table
function awk(opt)
end

---@param opt table
function sed(opt)
end

function find()
end

---@param k string
---@return string
function whereis(k)
end

---@param k string
---@param v string
function alias(k, v)
end

---@vararg string
function unalias(...)
end
