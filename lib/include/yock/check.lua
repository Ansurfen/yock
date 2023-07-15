-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---{{.Windows}}
---@return boolean
function Windows()
    return false
end

---{{.Darwin}}
---@return boolean
function Darwin()
    return false
end

---{{.Linux}}
---@return boolean
function Linux()
    return false
end

---{{.OS}}
---@param want_os string
---@param want_ver string
---@return boolean
function OS(want_os, want_ver)
    return false
end

---@param want string
---@param got string
---@return boolean
function CheckVersion(want, got)
    return false
end
