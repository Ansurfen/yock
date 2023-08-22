-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field

---@meta _

---@class env
---@field yock_modules string
env = {}

---import imports module based on module's
---name and filepath, and wraps load_module
---and load_file functions. Because using load_module
---is based on module's name, therefore it's not supported
---for import the relative path is without explict dot and separator sign.
---### Example:
---```lua
---# locates opencmd, if not found opencmd.lua, and search opencmd/index.lua
---# details see load_module function
---import("opencmd@1.0.0") -- call load_module
---import("opencmd") -- if not given version explictly, searchs local module.json in default
---
---# if you want to import file through the relative path, there are the following examples.
---import("test") -- error, import will call load_module to search test module.
---import("./test") -- correct
---
---import("D:/test")
---```
---@param target string
---@return unknown
---@return unknown loaderdata
function import(target) end

---load_module imports file based-on module's name.
---
---It's worth noting if target not found, it'll seen target
---as directory, and joining path with `index.lua` to search
---file in joined path.
---### Example:
---```lua
---# locates opencmd, if not found opencmd.lua, and search opencmd/index.lua
---import("opencmd@1.0.0")
---import("opencmd") -- if not given version explictly, searchs local module.json in default
---```
---@param target string
---@return any
function load_module(target) end

---load_file imports file based-on absoulte or relative path,
---and wraps require() function. Directly using require() isn't recommended.
---
---It's worth noting filename extension isn't required, and of
---course, explictly declaration also is supported.
---### Example:
---```lua
---load_file("./test.lua")
---load_file("./test") -- same with the above 
---load_file("D:/test")
---```
---@param path string
---@return any
function load_file(path) end

---@class module
---@field name string # module name
---@field version string
---@field url? string # module address
---@field license? string
---@field author? string
---@field load? fun(opt: table) # will be called when `ypm install [module]` command executed successfully.
---@field unload? fun(opt: table) # will be called when `ypm uninstall [module]` command is executing.

---register_service registers service with given name,
---and has a callback that receives avaiable port
---to user for making command and returns for nohup executing
---in `init` function.
---### Example:
---```lua
---register_service("python-eval", function (port)
---    return string.format("python main.py -p %d", port)
---end)
---``` 
---@param name string
---@param cmd fun(port: integer): string
---@return integer
function register_service(name, cmd) end

---unregister_service unregisters service with given name
---@param name string
function unregister_service(name) end

---init initializes specified target, which is equal
---to module's name. Just like import, you also can 
---specify version, e.g. `python-eval@1.0.0`. Then 
---it'll try to search and load `{target}/init.lua`.
---@param target string
---@return any
function init(target) end
