-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class ioReader
local ioReader = {}

---@param p string
---@return integer, err
function ioReader:Read(p)
end

---@class ioWriter
local ioWriter = {}

---@param p string
---@return integer, err
function ioWriter:Write(p)
end

---@param rd ioReader
---@return bufioReader
function bufio.NewReader(rd)
end

---@param rd ioReader
---@param size integer
---@return bufioReader
function bufio.NewReaderSize(rd, size)
end

---@class bufioReader
local bufioReader = {}

---@return integer
function bufioReader:Size()
end

---@param r ioReader
function bufioReader:Reset(r)
end

---@param n integer
---@return string, err
function bufioReader:Peek(n)
end

---@param n integer
---@return integer, err
function bufioReader:Discard(n)
end

---@param p string
---@return integer, err
function bufioReader:Read(p)
end

---@return integer, err
function bufioReader:ReadByte()
end

---@return err
function bufioReader:UnreadByte()
end

---@return integer, integer, err
function bufioReader:ReadRune(r, size, err)
end

---@return err
function bufioReader:UnreadRune()
end

---@return integer
function bufioReader:Buffered()
end

---@param delim integer
---@return integer[], err
function bufioReader:ReadSlice(delim)
end

---@return integer[], boolean, err
function bufioReader:ReadLine()
end

---@param delim integer
---@return integer[], err
function bufioReader:ReadBytes(delim)
end

---@param delim integer
---@return string, err
function bufioReader:ReadString(delim)
end

---@param w ioWriter
---@return integer, err
function bufioReader:WriteTo(w)
end

---@class bufioWriter
local bufioWriter = {}

---@param w ioWriter
---@param size integer
---@return bufioWriter
function bufio.NewWriterSize(w, size)
end

---@param w ioWriter
---@return bufioWriter
function bufio.NewWriter(w)
end

---@return integer
function bufioWriter:Size()
end

---@param w ioWriter
function bufioWriter:Reset(w)
end

---@return err
function bufioWriter:Flush()
end

---@return integer
function bufioWriter:Available()
end

---@return integer[]
function bufioWriter:AvailableBuffer()
end

---@return integer
function bufioWriter:Buffered()
end

---@param p string
---@return integer, err
function bufioWriter:Write(p)
end

---@param c integer
---@return err
function bufioWriter:WriteByte(c)
end

---@param r integer
---@return integer, err
function bufioWriter:WriteRune(r)
end

---@param s string
---@return integer, err
function bufioWriter:WriteString(s)
end

---@param r ioReader
---@return integer, err
function bufioWriter:ReadFrom(r)
end

---@class bufioReadWriter
local bufioReadWriter = {}

---@param r bufioReader
---@param w bufioWriter
---@return bufioReadWriter
function bufio.NewReadWriter(r, w)
end
