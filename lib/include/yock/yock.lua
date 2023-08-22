-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class byte: integer

---@param e any
---@param success_handle? function
function yassert(e, success_handle) end

---@class yocki
---@field connect fun(name: string, ip: string, port: integer) # dial server to be specified by ip and port, and name it for call.
---@field call fun(name: string, fn: string, arg: string): string, err # calls function on specified server
---@field list fun(): string[] # returns a list of services that have been connected
yocki = {}

---@param name string
---@param ip string
---@param port integer
function yocki.connect(name, ip, port) end

---@param name string
---@param fn string
---@param arg string
---@return string, err
function yocki.call(name, fn, arg) end

---@return string[]
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

---@class strf_opt: table

---strf wraps template engine to format string.
---
---`NOTE`: Because of the golang sake, string to be
---replaced must allows capital case in the first letter.
---Having a `Charset` field saved in internal, strf will
---converts string encoding.
---### Example:
---```lua
---print(strf("{{.A}} {{.B}}", { A = "Hello", B = "World" }))
---print(strf("你好", { Charset = "GB18030" }))
---```
---@param format string
---@param opt strf_opt
---@return string
function strf(format, opt) end

---strf is just like string.format(),
---and you can think it's short for
---string format. Using it isn't different
---with string.format().
---### Example:
---```lua
---print(string.format("%s %s", "Hello", "World"))
---print(strf("%s %s", "Hello", "World"))
---```
---@param format string
---@vararg any
---@return string
function strf(format, ...) end

---@param title string[]
---@param rows string[][]
function printf(title, rows) end

---pathf joins any number of path elements into a single path,
---separating them with slashes. Empty elements are ignored.
---The result is Cleaned. However, if the argument list is empty
---or all its elements are empty, Join returns an empty string.
---
---Except joining any number of path, pathf's first element can
---alias regular path according to different parameter.
---### Mapping rule:
---* `#(integer)` returns the real path of function from stack
---* `$` returns process's worksapce
---* `~` returns the path of executable file
---* `@` returns yock's worksapce
---
---### Example:
---```lua
---# assumes working directory is D:/tmp/main.lua
---pathf("@/", "a", "b") -- mapping: {HomeDir}/.yock/a/b
---pathf("#1", "../a") -- mapping: #1 returns D:/tmp/main.lua path, then joins "../a" to output "D:/tmp/a"
---```
---@vararg string
---@return string
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

---wrapzip returns string that wrapped platform zip format in default
---### Example:
---```lua
---wrapzip("a") -- on windows, result is a.zip
---
---wrapzip("a") -- on linux, darwin and etc, result is a.tar.gz
---```
---@param s string
---@return string
function wrapzip(s) end

---wrapexf returns string that wrapped platform executable filename extension in default
---### Example:
---```lua
---wrapexf("a") -- on windows, result is a.exe
---
---wrapexf("a") -- on linux, darwin and etc, result is a
---```
---@param s string
---@return string
function wrapexf(s) end

---wrapscript returns string that wrapped platform script format in default
---### Example:
---```lua
---wrapscript("a") -- on windows, result is a.bat
---
---wrapscript("a") -- on linux, darwin and etc, result is a.sh
---```
---@param s string
---@return string
function wrapscript(s) end

---@param path string
---@return Viper, err
function open_conf(path) end

---@class yockd_fs
local yockd_fs = {}

function yockd_fs.put(src, dst) end

function yockd_fs.get(src, dst, scope) end

function yockd_fs.ls(dir) end

function yockd_fs.rmdir(dir) end

---@param file string
---@return string[]
function yockd_fs.info(file) end

---@class file_descriptor: string

---@param file string
---@return file_descriptor[]
function yockd_fs.open(file) end

---@param file string
---@param owner string
---@return file_descriptor
function yockd_fs.open(file, owner) end

---@param fd file_descriptor
---@return string
function yockd_fs.read(fd) end

---@param file string
---@return string
function yockd_fs.read(file) end

---@param name? string
function yockd_fs.volume(name) end

---@class yockd_signal
local yockd_signal = {}

function yockd_signal.list() end

---@param sig string
---@return boolean exist, boolean status, err
function yockd_signal.info(sig) end

---@vararg string
function yockd_signal.clear(...) end

---@class yockd_net
local yockd_net = {}

---@param fromName string
---@param fromIP string
---@param fromPort integer
---@param fromPublic boolean
---@param toName string
---@param toIP string
---@param toPort integer
---@param toPublic boolean
---@return err
function yockd_net.dial(fromName, fromIP, fromPort, fromPublic,
                        toName, toIP, toPort, toPublic)
end

---@param node string
---@param method string
---@vararg string
function yockd_net.call(node, method, ...) end

---@class yockd
---@field fs yockd_fs
---@field signal yockd_signal
---@field net yockd_net
---@field process yockd_process
yockd = {}

---@class yockd_process
local yockd_process = {}

---@param type string|"cron"|"fs"|"script"
---@param sepc string
---@param cmd string
---@return integer pid, err
function yockd_process.spawn(type, sepc, cmd) end

---@param id integer
---@return process[], err
function yockd_process.find(id) end

---@param cmd string
---@return process[], err
function yockd_process.find(cmd) end

---@param id integer
function yockd_process.kill(id) end

---@class process
---@field pid integer
---@field state string|'create'|'ready'|'suspend'|'running'|'destory'
---@field spec string
---@field cmd string
---@field type string
local process = {}

---@return process[]
function yockd_process.list() end

---@param name string
---@return err
function yockd.ping(name) end

---@param name string
---@param ip string
---@param port integer
function yockd.dial(name, ip, port) end

---@param src string
---@param dst string
---@param perm? string
function yockd.upload(src, dst, perm) end

---@param src string
---@param dst string
---@param user? string
function yockd.download(src, dst, user) end

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

---@return userdata
function Viper:AllSettings() end

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

---@class FlagValue

---@class pflagFlag

---@class pflagFlagSet
