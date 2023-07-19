-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@param callback fun()
---@async
function go(callback) end

---@param sig string
---@param timeout? time
function wait(sig, timeout) end

---@param ... string|time
function waits(...) end

---@param sig string
function notify(sig) end
