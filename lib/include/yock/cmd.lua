-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

-- Example:
-- ```lua
-- sh({redirect = true}, "echo Hello World")
-- ```
---@param opt table
---@vararg string
---@return table<string>, err
function sh(opt, ...) end

-- Example:
-- ```lua
-- sh("echo a", "echo b")
-- ```
---@vararg string
---@return table<string>, err
function sh(...) end

---@class promptoptdesc
---@field use string
---@field short string
local promptoptdesc = {}

---@class promptoptflag
---@field default boolean|string|table<string>
---@field type flag_type
---@field name string
---@field shorthand string
---@field usage string
local promptoptflag = {}

---@class promptopt
---@field desc promptoptdesc
---@field sub table<promptopt>
---@field flags table<promptoptflag>
---@field run fun(cmd: userdata, args: table<string>)
local promptopt = {}

---{{.prompt}}
---@param opt promptopt
function prompt(opt) end

---@class command
---@field Use string
---@field Short string
---@field Long string
---@field Run fun(cmd: command, args: table)
local command = {}

---{{.new_command}}
---@return command
function new_command() end

---{.new_command_AddCommand}
---@vararg command
function command:AddCommand(...) end

---@return any
function command:PersistentFlags() end

---@return err
function command:Execute() end

---@vararg string
---@return string
function cmdf(...) end

---@class flag_type
---@field str number
---@field number_type number
---@field array_type number
---@field bool_type number
flag_type = {}

-- ```lua
-- local res = {}
-- argsparse(res, {
--     a = flag_type.str,
--     b = flag_type.bool,
-- })
-- table.dump(res)
-- ```
---@param env table
---@param todo table
function argsparse(env, todo) end

---@class argBuilder
---@field params table
argBuilder = {}

---@return argBuilder
function argBuilder:new() end

---@param cmd string
---@return argBuilder
function argBuilder:add(cmd) end

---@param cmd string
---@param v boolean|nil
---@return argBuilder
function argBuilder:add_bool(cmd, v) end

---@param cmd string
---@param v string|nil
---@return argBuilder
function argBuilder:add_str(cmd, v) end

---@param format string
---@param v string|nil
---@return argBuilder
function argBuilder:add_strf(format, v) end

---@param v any[]
---@return argBuilder
function argBuilder:add_arr(v) end

---@param ok boolean
---@param format string
---@vararg any
---@return argBuilder
function argBuilder:add_format(ok, format, ...) end

---@return string
function argBuilder:build() end

---@param opt table
---@return table, err
function argBuilder:exec(opt) end
