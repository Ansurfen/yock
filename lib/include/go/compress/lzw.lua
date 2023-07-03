-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class lzw
---@field LSB any
---@field MSB any
lzw = {}

---{{.lzwNewReader}}
---@param r ioReader
---@param order lzwOrder
---@param litWidth number
---@return any
function lzw.NewReader(r, order, litWidth)
end

---{{.lzwNewWriter}}
---@param w ioWriter
---@param order lzwOrder
---@param litWidth number
---@return any
function lzw.NewWriter(w, order, litWidth)
end

---@class lzwOrder
local lzwOrder = {}

---@class flateReader
local flateReader = {}

---{{.flateReaderRead}}
---@param b byte[]
---@return number, err
function flateReader:Read(b)
end

---{{.flateReaderClose}}
---@return err
function flateReader:Close()
end

---{{.flateReaderReset}}
---@param src ioReader
---@param order lzwOrder
---@param litWidth number
function flateReader:Reset(src, order, litWidth)
end

---@class flateWriter
local flateWriter = {}

---{{.flateWriterWrite}}
---@param p byte[]
---@return number, err
function flateWriter:Write(p)
end

---{{.flateWriterClose}}
---@return err
function flateWriter:Close()
end

---{{.flateWriterReset}}
---@param dst ioWriter
---@param order lzwOrder
---@param litWidth number
function flateWriter:Reset(dst, order, litWidth)
end
