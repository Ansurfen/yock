-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class sshClient
local sshClient = {}

---@param cmds string
---@return string, err
function sshClient:Exec(cmds) end

---@param file string
---@vararg any
---@return string, err
function sshClient:Sh(file, ...) end

---@return string
function sshClient:OS() end

function sshClient:Shell() end

---@param src string
---@param dst string
---@return err
function sshClient:Get(src, dst) end

---@param src string
---@param dst string
---@return err
function sshClient:Put(src, dst) end

---@class sshOpt
---@field user string
---@field pwd string
---@field ip string
---@field port integer
---@field network string
---@field redirect boolean
local sshOpt = {}

---@param opt sshOpt
---@param cb fun(client: sshClient)
---@return sshClient
function ssh(opt, cb) end
