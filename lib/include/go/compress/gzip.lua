-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class gzip
---@field NoCompression any
---@field BestSpeed any
---@field BestCompression any
---@field DefaultCompression any
---@field HuffmanOnly any
---@field ErrChecksum any
---@field ErrHeader any
gzip = {}

---{{.gzipNewReader}}
---@param r ioReader
---@return gzipReader, err
function gzip.NewReader(r)
end

---{{.gzipNewWriter}}
---@param w ioWriter
---@return gzipWriter
function gzip.NewWriter(w)
end

---{{.gzipNewWriterLevel}}
---@param w ioWriter
---@param level number
---@return gzipWriter, err
function gzip.NewWriterLevel(w, level)
end

---@class gzipHeader
---@field Comment string
---@field Extra any
---@field ModTime any
---@field Name string
---@field OS byte
local gzipHeader = {}

---@class gzipReader
local gzipReader = {}

---{{.gzipReaderRead}}
---@param p byte[]
---@return number, err
function gzipReader:Read(p)
end

---{{.gzipReaderClose}}
---@return err
function gzipReader:Close()
end

---{{.gzipReaderReset}}
---@param r ioReader
---@return err
function gzipReader:Reset(r)
end

---{{.gzipReaderMultistream}}
---@param ok boolean
function gzipReader:Multistream(ok)
end

---@class gzipWriter
local gzipWriter = {}

---{{.gzipWriterReset}}
---@param w ioWriter
function gzipWriter:Reset(w)
end

---{{.gzipWriterWrite}}
---@param p byte[]
---@return number, err
function gzipWriter:Write(p)
end

---{{.gzipWriterFlush}}
---@return err
function gzipWriter:Flush()
end

---{{.gzipWriterClose}}
---@return err
function gzipWriter:Close()
end
