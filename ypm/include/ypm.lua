-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class env
---@field yock_modules string
env = {}

---@param opt table
---@return function
function ymodule(opt)
    return function()
    end
end

---@param target string
---@return unknown
---@return unknown loaderdata
function import(target)
end

---@param todo table
---@return function
function ymodule(todo)
    return function()
    end
end

---@param target string
---@return any
function load_module(target) end

---@class module
---@field name string
---@field version string
---@field load fun(opt: table)
local module = {}

---@param name string
---@param cmd fun(port: integer): string
---@return integer
function register_service(name, cmd) end

---@param name string
function unregister_service(name) end

---@param target string
function init(target) end
