-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

local a = file("a.txt")
local b = file("b.txt")
local c = a + b
local d = c - b

_       = c < stream("Hello World\n")
_       = c <= stream("Hello World!!")
_       = d < stream(pwd())
