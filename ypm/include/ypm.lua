-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

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
function load_module(target)
end

---@param opt table
function yock_todo_loader(opt)
end

---@param url string
---@return fun(opt: table)
function github_loader(url)
end

---@class module
---@field name string
---@field version string
---@field load fun(opt: table)
local module = {}
