-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class waitGroup
local waitGroup = {}

---@param delta number
function waitGroup:Add(delta) end

function waitGroup:Done()
end

function waitGroup:Wait() end

sync = {}

---@return waitGroup
function sync.new() end
