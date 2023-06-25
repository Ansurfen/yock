-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class oslib
---@field Stdout userdata
---@field Stderr userdata
---@field Stdin userdata
---@field O_RDONLY integer
---@field O_WRONLY integer
---@field O_RDWR integer
---@field O_APPEND integer
---@field O_CREATE integer
---@field O_EXCL integer
---@field O_SYNC integer
---@field O_TRUNC integer
os = {}

---@return osFile, err
function os.Open()
end

---@param name string
---@return osFile, err
function os.Create(name)
end

---@param name string
---@param flag integer
---@param perm integer
---@return osFile, err
function os.OpenFile(name, flag, perm)
end

---@class osFile
local osFile = {}

---@return err
function osFile:Close()
end
