-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@alias echomode string
---|> "c" # create a new file if none exists.
---| "t" # truncate regular writable file when opened.
---| "r" # open the file read-only.
---| "w" # open the file write-only.
---| "rw" # open the file read-write.
---| "a" # append data to the file when writing.
---| "e" # used with `c`, file must not exist.
---| "s" # open for synchronous I/O.

---@class echoopt
---@field mode echomode
---@field fd table<string>
local echoopt = {}

---@vararg string
function echo(...) end

-- Example:
--
-- append write
-- ```lua
-- echo({ fd = { "stdout", "test.txt" }, mode = "c|a|rw" }, "Hello World!")
-- ```
--
-- truncate write
-- ```lua
-- echo({ fd = { "stdout", "test.txt" }, mode = "c|t|rw" }, "Hello World!")
-- ```
---@param opt echoopt
---@vararg string
---@return table<string>, err
function echo(opt, ...) end

---@return string, err
function whoami() end

function clear() end

---@param dir string
---@return err
function cd(dir) end

---@param file string
---@return err
function touch(file) end

---@param file string
---@return string, err
function cat(file) end

---@param opt table|string
---@return table<string>|string, err
function ls(opt) end

---@param name string
---@param mode number
---@return err
function chmod(name, mode) end

---@param name string
---@param uid number
---@param gid number
---@return err
function chown(name, uid, gid) end

---@vararg string
---@return err
function mkdir(...) end

-- Example:
-- ```lua
-- cp("a", "b")
-- ```
---@param src string
---@param dst string
function cp(src, dst) end

---@class cpopt
---@field recurse boolean
---@field force boolean
local cpopt = {}

-- Example:
-- ```lua
-- cp({ recurse = true }, {
--      ["a"] = "b",
--      ["c"] = "d",
-- }
-- ```
---@param opt cpopt
---@param path table<string, string>
function cp(opt, path) end

---@param src string
---@param dst string
function mv(src, dst) end

---@return string, err
function pwd() end

---@class rmOpt
---@field safe boolean
---@field pattern string
local rmOpt = {}

---Example:
--
-- like rmdir
-- ```lua
-- rm("a")
-- ```
--
-- delete file with recuse
-- ```lua
-- rm({ safe = true }, "/a", "/b")
-- ```
---@param opt rmOpt
---@vararg string
---@return err
---
function rm(opt, ...) end

---Example:
-- ```lua
-- rm("/a", "/b")
-- ```
---@vararg string
---@return err
function rm(...) end

---@param src string
---@param dst string
function rename(src, dst) end

---@param cmd table|string
function sudo(cmd) end

---@param opt table
function grep(opt) end

---@param opt table
function awk(opt) end

---@param opt table
function sed(opt) end

---@class findopt
---@field pattern string
---@field dir boolean
---@field file boolean
local findopt = {}

---
---Example:
---```lua
-- find({
--     dir = false,
--     pattern = "\\.lua"
-- }, "/script")
---```
---
---@param opt findopt
---@param path string
---@return table, err
function find(opt, path) end

---@param path string
---@return boolean
function find(path) end

---@param k string
---@return string, err
function whereis(k) end

-- Example:
-- ```lua
-- alias("CC", "go")
-- sh("$CC version")
-- ```
---@param k string
---@param v string
function alias(k, v) end

-- Example:
-- ```lua
-- alias("CC", "go")
-- unalias("CC")
-- sh("$CC version")
-- ```
---@vararg string
function unalias(...) end

---@param cmd string
---@return err
function nohup(cmd) end

---@class pgrepres
---@field name string
---@field pid integer
local pgrepres = {}

---@param name string
---@return pgrepres[]
function pgrep(name) end

---@class psinfo
---@field name string
---@field cmd string
---@field cpu? number
---@field start? number
---@field mem? any
---@field user? string
local psinfo = {}

---@class psopt
---@field user boolean
---@field cpu boolean
---@field time boolean
---@field mem boolean
local psopt = {}

---@param opt psopt|string|integer|nil
---@return table<integer, psinfo>
function ps(opt) end

---@param k integer|string
---@return err
function kill(k) end

---@param src string
---@param dst string
function tarc(src, dst) end

---@param src string
---@param dst string
function zipc(src, dst) end

---@param src string
---@param dst string
function untar(src, dst) end

---@param src string
---@param dst string
function unzip(src, dst) end

---@param src string
---@param dst string
function compress(src, dst) end

---@param src string
---@param dst string
function uncompress(src, dst) end

---Example:
---```lua
---export("PATH", "/bin/yock")
---```
---@param k string
---@param v string
---@return err
function export(k, v) end

---Example:
---```lua
---export("PATH:/bin/yock")
---```
---@param kv string
---@return err
function export(kv) end

---@param k string
function unset(k) end

---@class ifconfig_addr
---@field addr string
local ifconfig_addr = {}

---@alias ifconfig_flag string|"up"|"broadcast"|"multicast"|"loopback"

---@class ifconfig_result
---@field index integer
---@field mtu integer
---@field name string
---@field hardwareAddr string
---@field flags ifconfig_flag[]
---@field addrs ifconfig_addr[]
local ifconfig_result = {}

---@return ifconfig_result[]
function ifconfig() end

---@class systemctlopt
local systemctlopt = {}

-- ---@param opt systemctlopt
-- function systemctl(opt)
-- end

systemctl = {}

---@alias sysstate string
---|> "all"
---| "active"
---| "inactive"

---@alias systype string
---|> "target"
---|"service"

---@param t? systype
---@param s? sysstate
---@return sysservice[]
function systemctl.list(t, s) end

---@param name string
---@return err
function systemctl.restart(name)
end

---@param name string
---@return err
function systemctl.start(name) end

---@param name string
---@return err
function systemctl.stop(name) end

---@param name string
---@return err
function systemctl.delete(name) end

---@param name string
---@return err
function systemctl.disable(name) end

---@param name string
---@return err
function systemctl.enable(name) end

---@class sccreateoptunit
---@field description string
---@field before string
---@field after string
local sccreateoptunit = {}

---@class sccreateoptservice
---@field type "simple"|"exec"|"forking"|"oneshot"|"dbus"|"notify"|"idle"
---@field execStart string
---@field execStop string
---@field privateTmp boolean
---@field restartSec integer
---@field restart string
local sccreateoptservice = {}

---@class sccreateoptinstall
---@field wantedBy string
local sccreateoptinstall = {}

---@class sscreateoptspec
local sscreateoptspec = {}

---@class sccreateopt
---@field unit sccreateoptunit
---@field service sccreateoptservice
---@field install sccreateoptinstall
---@field spec sscreateoptspec
local sccreateopt = {}

---@param name string
---@param opt sccreateopt
---@return err
function systemctl.create(name, opt) end

---@alias servicestatus integer
---|> "running"
---| "stopped"
---| "unknown"

---@class sysservice
---@field pid integer
---@field name string
---@field status servicestatus
local sysservice = {}

---@param name string
---@return sysservice, err
function systemctl.status(name) end

iptables = {}

---@class iptableslistopt
---@field name string
---@field chain string
---@field legacy boolean
local iptableslistopt = {}

---@class firewarerule
---@field name string
---@field proto string
---@field src string
---@field dst string
---@field action string
local firewarerule = {}

-- legacy: determine to use iptables or iptables-legacy (except windows)
--
-- name: returns service to be specified and all services when the length of name is empty/zero.
---@param opt iptableslistopt
---@return firewarerule[]|firewarerule, err
function iptables.list(opt) end

---@class iptablesopopt
---@field chain string
---@field name string
---@field protocol string
---@field action string
---@field destination string
---@field legacy boolean
local iptablesopopt = {}

-- chain: INPUT(linux), in(windows)
--
-- action: drop(linux), block(windows)
---@param opt iptablesopopt
---@return err
function iptables.add(opt) end

---@param opt iptablesopopt
---@return err
function iptables.del(opt) end

---@class lsofinfo
---@field pid string
---@field state string
---@field proto string
---@field Local string
---@field foreign string
local lsofinfo = {}

---@param port? integer
---@return lsofinfo[]|lsofinfo
function lsof(port) end

---@class curlopt
---@field header? table<string, string>
---@field method? string
---@field data? string
---@field save? boolean
---@field dir? string
---@field filename? fun(s: string): string
---@field debug? boolean
---@field strict? boolean
---@field caller? string
---@field async? boolean
local curlopt = {}

---@param opt curlopt
---@vararg string
---@return string, err
function curl(opt, ...) end

---@vararg string
---@return string, err
function curl(...) end

---@param file string
---@param data string
---@return err
function write(file, data) end
