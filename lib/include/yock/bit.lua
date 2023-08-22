-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---because of the sake of lua, we don't use operator (such as &、|、^、~、>>、<<)
---to implement easily bitwise operation. so, here yock provide a method based on
---tabel's field to make it.
---
---It just like traditional operator:
---```lua
---local a = bit.And(1, 2) -- equal a = 1 & 2
---local b = bit.Or(1, 2) -- equal a = 1 | 2
---```
---@class bit
---@field And fun(a: integer, b: integer): integer
---@field Or fun(a: integer, b:integer): integer
bit = {}
