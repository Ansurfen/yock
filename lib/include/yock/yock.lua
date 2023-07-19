-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class byte: integer

---@param e any
---@param success_handle? function
function yassert(e, success_handle) end

yocki = {}

---@param name string
---@param ip string
---@param port number
function yocki.connect(name, ip, port) end

---@param name string
---@param fn string
---@param arg string
---@return string, err
function yocki.call(name, fn, arg) end

---@return table<string>
function yocki.list() end

cachetable = {}

---@param name string
---@param level integer
---@param expire timeTime
---@param lock string
---@param attr integer
---@return ycache
function cachetable:create(name, level, expire, lock, attr) end

---@param name string
---@param lock string
---@return ycache|nil
function cachetable:get(name, lock) end

---@param level integer
---@param expire? boolean
function cachetable:free(level, expire) end

---@class ycache
local ycache = {}

---@param index string
---@param dir string
---@return ycache
function ycache:new(index, dir) end

---@param k string
---@param v any
function ycache:put(k, v) end

---@param k string
---@return any
function ycache:get(k) end

function ycache:free() end

function ycache:save() end

---@alias charset "UTF-8" | "GB18030"

---@class strfopt
local strfopt = {}

---@param format string
---@param opt strfopt
---@return string
function strf(format, opt) end

---@param format string
---@vararg any
---@return string
function strf(format, ...) end

---@param title string[]
---@param rows string[][]
function printf(title, rows) end

---`#(integer)` returns the real path of function from stack
---
---`$` returns process's worksapce
---
---`~` returns the path of executable file
---
---`@` returns yock's worksapce
---
---example:
---```lua
---pathf("@/", "a", "b")
---pathf("#1")
---```
---
---@vararg string
---@return string
---
function pathf(...) end

---@param opt table
---@param handles table<fun(): string>
function loadbalance(opt, handles) end

---@param fileType string
---@param want table<string, string>
---@param opt table<string, string>
---@return string
function multi_fetch(fileType, want, opt) end

---@param todo table
---@param handle function
---@return table
function multi_bind(todo, handle) end

---@param s string
---@return string
function wrapzip(s) end

---@param s string
---@return string
function wrapexf(s) end

---@param s string
---@return string
function wrapscript(s) end

---@param path string
---@return Viper, err
function open_conf(path) end

---@class Viper
local Viper = {}

---@param p string
function Viper:AddConfigPath(p)
end

---@param provider string
---@param endpoint string
---@param path string
---@return err
function Viper:AddRemoteProvider(provider, endpoint, path)
end

---@param provider string
---@param endpoint string
---@param path string
---@param secretkeyring string
---@return err
function Viper:AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
end

---@return string[]
function Viper:AllKeys()
end

---@return table
function Viper:AllSettings()
end

---@param allowEmptyEnv boolean
function Viper:AllowEmptyEnv(allowEmptyEnv)
end

function Viper:AutomaticEnv()
end

---@vararg string
---@return err
function Viper:BindEnv(...)
end

---@param key string
---@param flag FlagValue
---@return err
function Viper:BindFlagValue(key, flag)
end

---@param flags FlagValueSet
---@return err
function Viper:BindFlagValues(flags)
end

---@param key string
---@param flag pflagFlag
---@return err
function Viper:BindPFlag(key, flag)
end

---@param key string
---@param flags pflagFlagSet
---@return err
function Viper:BindPFlags(key, flags)
end

---@return string
function Viper:ConfigFileUsed()
end

function Viper:Debug()
end

---@param w ioWriter
function Viper:DebugTo(w)
end

---@param key string
---@return userdata
function Viper:Get(key)
end

---@param key string
---@return boolean
function Viper:GetBool(key)
end

---@param key string
---@return number
function Viper:GetFloat64(key)
end

---@param key string
---@return integer
function Viper:GetInt(key)
end

---@param key string
---@return integer
function Viper:GetInt32(key)
end

---@param key string
---@return integer
function Viper:GetInt64(key)
end

---@param key string
---@return integer[]
function Viper:GetIntSlice(key)
end

---@param key string
---@return integer
function Viper:GetSizeInBytes(key)
end

---@param key string
---@return string
function Viper:GetString(key)
end

---@param key string
---@return table
function Viper:GetStringMap(key)
end

---@param key string
---@return table
function Viper:GetStringMapString(key)
end

---@param key string
---@return table
function Viper:GetStringMapStringSlice(key)
end

---@param key string
---@return timeDuration
function Viper:GetDuration(key)
end

---@param key string
---@param value any
function Viper:Set(key, value)
end

---@return err
function Viper:WriteConfig()
end

---@return err
function Viper:SafeWriteConfig()
end

---@param filename string
---@return err
function Viper:WriteConfigAs(filename)
end

---@param filename string
---@return err
function Viper:SafeWriteConfigAs(filename)
end

---@class FlagValueSet
local FlagValueSet = {}

---@class FlagValue
local FlagValue = {}

---@class pflagFlag
local pflagFlag = {}

---@class pflagFlagSet
local pflagFlagSet = {}
