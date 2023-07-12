-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class sshClient
local sshClient = {}

---@param cmds string
function sshClient:Exec(cmds)
end

function sshClient:Shell()
end

---@param src string
---@param dst string
---@return err
function sshClient:Get(src, dst)
end

---@param src string
---@param dst string
---@return err
function sshClient:Put(src, dst)
end

---@class sshOpt
---@field user string
---@field pwd string
---@field ip string
---@field network string
---@field redirect boolean
local sshOpt = {}

---@param opt sshOpt
---@param cb fun(client: sshClient)
---@return sshClient
function ssh(opt, cb)
    return {}
end
