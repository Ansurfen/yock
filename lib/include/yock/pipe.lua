-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field

---@meta _

---pipe object is designed to simulate the pipe operation of terminal
---on lua.
---@class pipe
---@field type integer
---@field payload any
---@field clone fun(self: pipe): pipe
---@operator add(pipe):pipe
---@operator sub(pipe):pipe

---file saves file descriptor based-on filename and creates
---empty file when given file not exist.
---### Example:
---```lua
---local a = file("1.txt") -- single file stream
---local b = file("2.txt", "3.txt") -- multiple file stream
---
--- # operator reload to reset file stream
---local c = a + b -- converge file stream to handle at the same time
---local d = c - file("2.txt") -- remove file stream based-on the second parameter
---```
---@vararg string
---@return pipe
function file(...) end

---stream converts string into pipe object, which allow you
---use operator to handle file stream, just like the pipe operation
---of terminal.
---### Example:
---```lua
---local a = stream("Hello World") # create pipe object
---_ = file("test.txt") < a -- write with truncation to test.txt
---
---_ = file("test.txt") <= a -- write with append to test.txt
---```
---@param str string
---@return pipe
function stream(str) end
