-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

bzip2 = {}

---{{.bzip2NewReader}}
---@param r ioReader
---@return ioReader
function bzip2.NewReader(r)
end

---@class bzip2StructuralError
local bzip2StructuralError = {}

---{{.bzip2StructuralErrorError}}
---@return string
function bzip2StructuralError:Error()
end
