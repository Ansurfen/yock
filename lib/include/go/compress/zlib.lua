-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class zlib
---@field NoCompression any
---@field BestSpeed any
---@field BestCompression any
---@field DefaultCompression any
---@field HuffmanOnly any
---@field ErrChecksum any
---@field ErrDictionary any
---@field ErrHeader any
zlib = {}

---{{.zlibNewReader}}
---@param r ioReader
---@return any, err
function zlib.NewReader(r)
end

---{{.zlibNewReaderDict}}
---@param r ioReader
---@param dict byte[]
---@return any, err
function zlib.NewReaderDict(r, dict)
end

---{{.zlibNewWriter}}
---@param w ioWriter
---@return flateWriter
function zlib.NewWriter(w)
end

---{{.zlibNewWriterLevel}}
---@param w ioWriter
---@param level number
---@return flateWriter, err
function zlib.NewWriterLevel(w, level)
end

---{{.zlibNewWriterLevelDict}}
---@param w ioWriter
---@param level number
---@param dict byte[]
---@return flateWriter, err
function zlib.NewWriterLevelDict(w, level, dict)
end

---@class flateResetter
local flateResetter = {}

---@class flateWriter
local flateWriter = {}

---{{.flateWriterWrite}}
---@param p byte[]
---@return number, err
function flateWriter:Write(p)
end

---{{.flateWriterFlush}}
---@return err
function flateWriter:Flush()
end

---{{.flateWriterClose}}
---@return err
function flateWriter:Close()
end

---{{.flateWriterReset}}
---@param w ioWriter
function flateWriter:Reset(w)
end
