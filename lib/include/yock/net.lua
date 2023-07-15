-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@param path string
---@return boolean
function is_url(path)
    return false
end

---@class curlopt
---@field header table<string, string>
---@field method string
---@field data string
---@field save boolean
---@field dir string
---@field filename fun(s: string): string
---@field debug boolean
---@field strict boolean
---@field caller string
---@field async boolean
local curlopt = {}

---@param opt curlopt
---@vararg string
---@return string, err
function curl(opt, ...)
end

---@param url string
---@return boolean
function is_localhost(url)
    return false
end
