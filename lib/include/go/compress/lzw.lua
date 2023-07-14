-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class lzw
---@field LSB any
---@field MSB any
lzw = {}

--- NewWriter creates a new io.WriteCloser.
--- Writes to the returned io.WriteCloser are compressed and written to w.
--- It is the caller's responsibility to call Close on the WriteCloser when
--- finished writing.
--- The number of bits to use for literal codes, litWidth, must be in the
--- range [2,8] and is typically 8. Input bytes must be less than 1<<litWidth.
---
--- It is guaranteed that the underlying type of the returned io.WriteCloser
--- is a *Writer.
---@param w ioWriter
---@param order lzwOrder
---@param litWidth number
---@return any
function lzw.NewWriter(w, order, litWidth) end

--- NewReader creates a new io.ReadCloser.
--- Reads from the returned io.ReadCloser read and decompress data from r.
--- If r does not also implement io.ByteReader,
--- the decompressor may read more data than necessary from r.
--- It is the caller's responsibility to call Close on the ReadCloser when
--- finished reading.
--- The number of bits to use for literal codes, litWidth, must be in the
--- range [2,8] and is typically 8. It must equal the litWidth
--- used during compression.
---
--- It is guaranteed that the underlying type of the returned io.ReadCloser
--- is a *Reader.
---@param r ioReader
---@param order lzwOrder
---@param litWidth number
---@return ioReadCloser
function lzw.NewReader(r, order, litWidth) end

--- Order specifies the bit ordering in an LZW data stream.
---@class lzwOrder
local lzwOrder = {}

--- Reader is an io.Reader which can be used to read compressed data in the
--- LZW format.
---@class flateReader
local flateReader = {}

--- Read implements io.Reader, reading uncompressed bytes from its underlying Reader.
---@param b byte[]
---@return number, err
function flateReader:Read(b) end

--- Close closes the Reader and returns an error for any future read operation.
--- It does not close the underlying io.Reader.
---@return err
function flateReader:Close() end

--- Reset clears the Reader's state and allows it to be reused again
--- as a new Reader.
---@param src ioReader
---@param order lzwOrder
---@param litWidth number
function flateReader:Reset(src, order, litWidth) end

--- Writer is an LZW compressor. It writes the compressed form of the data
--- to an underlying writer (see NewWriter).
---@class flateWriter
local flateWriter = {}

--- Write writes a compressed representation of p to w's underlying writer.
---@param p byte[]
---@return number, err
function flateWriter:Write(p) end

--- Close closes the Writer, flushing any pending output. It does not close
--- w's underlying writer.
---@return err
function flateWriter:Close() end

--- Reset clears the Writer's state and allows it to be reused again
--- as a new Writer.
---@param dst ioWriter
---@param order lzwOrder
---@param litWidth number
function flateWriter:Reset(dst, order, litWidth) end
