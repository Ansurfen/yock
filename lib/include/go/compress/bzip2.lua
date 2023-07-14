-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

bzip2 = {}

--- NewReader returns an io.Reader which decompresses bzip2 data from r.
--- If r does not also implement io.ByteReader,
--- the decompressor may read more data than necessary from r.
---@param r ioReader
---@return ioReader
function bzip2.NewReader(r) end

--- A StructuralError is returned when the bzip2 data is found to be
--- syntactically invalid.
---@class bzip2StructuralError
local bzip2StructuralError = {}


---@return string
function bzip2StructuralError:Error() end
