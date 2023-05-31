-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---{{.sh}}
---@param opt? table
---@vararg string
---@return table, err
function sh(opt, ...)
end

---{{.prompt}}
---@param tbl table
function prompt(tbl)
end

---@class command
---@field Use string
---@field Short string
---@field Long string
---@field Run fun(cmd: command, args: table)
local command = {}

---{{.new_command}}
---@return command
function new_command()
    return {}
end

---{.new_command_AddCommand}
---@vararg command
function command:AddCommand(...)
end

---@return any
function command:PersistentFlags()
    return {}
end

---@return err
function command:Execute()
end

--
---{{.cmdf}}
--
---@vararg string
---@return string
---
function cmdf(...)
    return ""
end

---@class flag_type
---@field str number
---@field number_type number
---@field array_type number
---@field bool_type number
flag_type = {}

---{{.argsparse}}
---@param env env
---@param todo table
function argsparse(env, todo)
end
