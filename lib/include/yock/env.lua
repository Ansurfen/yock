-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class yockConf
---@field Ycho ychoOpt
---@field Lang string
---@field Yockd yockd_opt
---@field Yockw yockw_opt

---@class yockw_opt
---@field SelfBoot boolean
---@field Port integer

---@class ychoOpt
---@field Level string
---@field Format string
---@field Path string
---@field FileName string
---@field FileMaxSize number
---@field FileMaxBackups number
---@field MaxAge number
---@field Compress boolean
---@field Stdout boolean

---@class yockd_opt
---@field IP string
---@field Port integer
---@field Name string
---@field SelfBoot boolean

---@class env
---@field args string[]
---@field platform platform
---@field flags table
---@field job string
---@field workdir string
---@field yock_path string
---@field conf yockConf
---@field yock_tmp string
---@field yock_bin string
---@field params table<string, table<string, starType>>?
env = {}

-- ---@param path string
-- ---@return err
-- function env.set_path(path) end

-- ---@param k string
-- ---@param v any
-- ---@return err
-- function env.set(k, v) end

-- ---@param k string
-- ---@param v any
-- ---@return err
-- function env.safe_set(k, v) end

-- ---@param k string
-- ---@param v any
-- ---@return err
-- function env.setl(k, v) end

-- ---@param k string
-- ---@param v any
-- ---@return err
-- function env.safe_setl(k, v) end

-- ---@param k string
-- ---@return err
-- function env.unset(k) end

-- ---@param file string
-- ---@return err
-- function env.export(file)
-- end

-- function env.print()
-- end

---@param args table
function env.set_args(args) end

---@class platform
---@field OS string|"aix"|"android"|"darwin"|"dragonfly"|"freebsd"|"hurd"|"illumos"|"ios"|"js"|"linux"|"nacl"|"netbsd"|"openbsd"|"plan9"|"solaris"|"windows"|"zos"
---@field Ver string
---@field Arch string|"386"|"amd64"|"amd64p32"|"arm"|"armbe"|"arm64"|"arm64be"|"loong64"|"mips"|"mipsle"|"mips64"|"mips64le"|"mips64p32"|"mips64p32le"|"ppc"|"ppc64"|"ppc64le"|"riscv"|"riscv64"|"s390"|"s390x"|"sparc"|"sparc64"|"wasm"
local platform = {}

---Exf returns executable file extension name in default.
---
---windows: .exe
---
---posix:
---@return string
function platform:Exf() end

---Script returns script extension name in default.
---
---windows: .bat
---
---posix: .sh
---@return string
function platform:Script() end

---Zip returns zip extension name in default.
---
---windows: .zip
---
---posix: .tar.gz
---@return string
function platform:Zip() end
