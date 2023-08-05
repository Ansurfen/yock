-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@return boolean
function Windows() end

---@return boolean
function Darwin() end

---@return boolean
function Linux() end

---@param want_os string
---@param want_ver string
---@return boolean
function OS(want_os, want_ver) end

---@param want string
---@param got string
---@return boolean
function CheckVersion(want, got) end
