-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class random
---@field str fun(n:number): string # str returns the string of given length n
---@field port fun():integer # port returns an idle port
random = {}