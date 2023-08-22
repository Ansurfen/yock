-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class sh_opt
---@field redirect? boolean # redirects stdin, stdout, stderr to this terminal
---@field quiet? boolean # prints the result of execution
---@field sandbox? boolean # launches subprocess without coping environment variables

---sh is designed to execute raw command, according to string provided.
---It's hardly any handling except using alias to mapping variable.
---### Option:
---* redirect, boolean (default false), redirects stdin, stdout, stderr to this terminal
---* quiet, boolean (default true), prints the result of execution
---* sandbox, boolean (default false), launches subprocess without coping environment variables
---### Example:
---```lua
---sh({ redirect = true }, "echo 'Hello World'") -- single command
---
---sh({ redirect = true }, "echo Hello", "echo World") -- multiple commands
---```
---@param opt sh_opt
---@vararg string
---@return string[], err
function sh(opt, ...) end

---sh is designed to execute raw command, according to string provided.
---It's hardly any handling except using alias to mapping variable.
---Example:
---```lua
---# multiple line commands
---sh([[
---echo Hello
---echo World
---]])
---
---sh("echo Hello", "echo World") -- multiple argument commands
---```
---@vararg string
---@return string[], err
function sh(...) end

---@class prompt_opt_desc
---@field use string # Use is the one-line usage message.
---@field short? string # Short is the short description shown in the 'help' output.
---@field long? string # Long is the long message shown in the 'help <this-command>' output.

---@class prompt_opt_flag
---@field default boolean|string|string[]
---@field type flag_type
---@field name string
---@field shorthand string
---@field usage string

---@class prompt_opt
---@field desc prompt_opt_desc # description of current command
---@field sub? prompt_opt[] # sub commands
---@field flags? prompt_opt_flag[] # current command with flags, such as -f, --flag
---@field run? fun(cmd: userdata, args: string[]) # callback when command is triggered

---prompt allow you to build command line tool with fast
---### Option:
---* desc, prompt_opt_desc, description of current command
---* sub, prompt_opt[], sub commands
---* flags, prompt_opt_flag[], current command's flags, such as -f, --flag
---* run, fun(cmd: userdata, args: string[]), callback when command is triggered
---
---### Example:
---```lua
---# ctl.lua
---
---prompt({
---    desc = {
---        use = "mycmd",
---        short = "mycmd is an terminal application"
---    },
---    sub = {
---        {
---            desc = { use = "echo" },
---            run = function(cmd, args)
---                for i = 1, #args, 1 do
---                    print(args[i])
---                end
---            end,
---        }
---    },
---    flags = {
---        {
---            type = flag_type.bool,
---            default = true,
---            name = "flag",
---            shorthand = "f",
---            usage = "flag test"
---        }
---    }
---})
---# use `yock run ctl.lua -- mycmd echo Hello World` to test on terminal when inputs above code.
---# meanwhile, you also could mount it to environment variable, use `yock mount mycmd ./ctl.lua`
---# if done, you'll can directly use `mycmd echo Hello World` to run it.
---```
---@param opt prompt_opt
function prompt(opt) end

---@class command
---@field Use string
---@field Short string
---@field Long string
---@field Run fun(cmd: command, args: string[])
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

---@vararg string
---@return argBuilder
function argBuilder:add(...) end

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

---@param opt table|nil
---@return table, err
function argBuilder:exec(opt) end

---@class args_builder
---@field combine table<string, boolean>
---@field available boolean
---@field params string[]
---@field args_builder fun(): args_builder
args_builder = {}

---@param header? string
---@return args_builder
function args_builder.new(header) end

---@vararg string
---@return args_builder
function args_builder:reg_combine(...) end

---@param ok boolean
---@param k string
---@return args_builder
function args_builder:add_combine(ok, k) end

---@param ok boolean
---@param k string
---@return args_builder
function args_builder:add_combine_must(ok, k) end

---@vararg string
---@return args_builder
function args_builder:add(...) end

---@param ok boolean
---@param format string
---@param ... any
---@return args_builder
function args_builder:add_str(ok, format, ...) end

---@param ok boolean
---@param format string
---@param ... any
---@return args_builder
function args_builder:add_str_must(ok, format, ...) end

---@param ok boolean
---@param format string
---@param arr any[]
---@return args_builder
function args_builder:add_arr(ok, format, arr) end

---@param ok boolean
---@param format string
---@param arr any[]
---@return args_builder
function args_builder:add_arr_must(ok, format, arr) end

---@param ok boolean
---@param v string
---@return args_builder
function args_builder:add_bool(ok, v) end

---@param ok boolean
---@param v string
---@return args_builder
function args_builder:add_bool_must(ok, v) end

---@param ok boolean
---@param format string
---@param map table
---@return args_builder
function args_builder:add_map(ok, format, map) end

---@param ok boolean
---@param format string
---@param map table
---@return args_builder
function args_builder:add_map_must(ok, format, map) end

---@param v string
---@return args_builder
function args_builder:set_header(v) end

---@return string
function args_builder:build() end

---@param opt sh_opt
---@return string[]|nil, err
function args_builder:exec(opt) end
