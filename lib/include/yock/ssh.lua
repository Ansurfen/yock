-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class sshClient
---@field Exec fun(self: sshClient, cmd: string): string, err # execute command on remote host.
---@field Sh fun(self: sshClient, file: string, ...): string, err # upload local script to remote, create temporary file and execute it on remote host. File-name extension is set automatically based-on remote's os (windows: .bat, posix: .sh).
---@field Shell fun(self: sshClient) # redirect stdio, stdin, stderr to this terminal and allocate a shell to handle.
---@field OS fun(self: sshClient): string # returns os for remote host. If not get, returns unknown. Its implement is execute `echo $OSTYPE` and `echo %OS%` to infer.
---@field Get fun(self: sshClient, src: string, dst: string): err # fetches src from remote host and saves it to dst on local. Just like ftp, but it's based-on sftp protocol.
---@field Put fun(self: sshClient, src: string, dst: string): err # upload src on local to dst on remote. Just like ftp, but it's based-on sftp protocol.
local sshClient = {}

---@param cmd string
---@return string, err
function sshClient:Exec(cmd) end

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

---@class ssh_opt
---@field user string # account what you want to login in on remote
---@field pwd string # user's password
---@field ip string # remote ip
---@field port integer # running port of ssh server
---@field network string # indicates network protocol (tcp, udp) to dial
---@field redirect boolean

---ssh dial remote host to be specified by ssh_opt. There
---are two different method to handle it, but it should be
---three to be exact. Either way, the effect is the same.
---### Option:
---* user, string, account what you want to login in on remote
---* pwd, string, user's password
---* ip, string, remote ip
---* port, integer, running port of ssh server
---* network, string, indicates network protocol (tcp, udp) to dial
---
---### Example:
---```lua
---# method 1, callback
---ssh({
---    user = "root",
---    pwd = "root",
---    ip = "localhost",
---    port = 22,
---    network = "tcp",
---}, function(c)
---    c:Exec("echo Hello World")
---end)
---
---# method 2, variable
---local c = ssh({
---    user = "root",
---    pwd = "root",
---    ip = "localhost",
---    port = 22,
---    network = "tcp",
---})
---c:Exec("echo Hello World")
---
---# method 3, combine above two method
---```
---@param opt ssh_opt
---@param cb fun(client: sshClient)
---@return sshClient
function ssh(opt, cb) end
