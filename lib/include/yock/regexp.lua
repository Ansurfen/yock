-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class reglib
---@field new fun(self: reglib, patterns: string): table
---@field find_str fun(self: reglib, p: string, s: string): string|nil
---@field match_str fun(self: reglib, p: string, s: string): boolean
reglib = {}
