-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class sh_opt
---@field redirect? boolean
---@field quiet? strfopt
---@field sandbox? boolean
local sh_opt = {}

-- Example:
-- ```lua
-- sh({redirect = true}, "echo Hello World")
-- ```
---@param opt sh_opt
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
---@field desc? promptoptdesc
---@field sub? table<promptopt>
---@field flags? table<promptoptflag>
---@field run? fun(cmd: userdata, args: table<string>)
local promptopt = {}


---@param opt promptopt
function prompt(opt) end

---@class command
---@field Use string
---@field Short string
---@field Long string
---@field Run fun(cmd: command, args: table)
local command = {}

---@return command
function new_command() end

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
---@field num number
---@field arr number
---@field bool number
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
