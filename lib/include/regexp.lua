-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

regexp = {}

---@param str string
---@return reg
function regexp.MustCompile(str)
    return {}
end

---@param str string
---@return reg, err
function regexp.Compile(str)
    return {}
end

---@class reg
local reg = {}

---@param s string
---@return userdata
function reg:FindStringSubmatch(s)
end
