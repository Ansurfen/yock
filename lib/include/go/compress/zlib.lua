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

--- NewWriterLevelDict is like NewWriterLevel but specifies a dictionary to
--- compress with.
---
--- The dictionary may be nil. If not, its contents should not be modified until
--- the Writer is closed.
---@param w ioWriter
---@param level number
---@param dict byte[]
---@return flateWriter, err
function zlib.NewWriterLevelDict(w, level, dict) end

--- NewReader creates a new ReadCloser.
--- Reads from the returned ReadCloser read and decompress data from r.
--- If r does not implement io.ByteReader, the decompressor may read more
--- data than necessary from r.
--- It is the caller's responsibility to call Close on the ReadCloser when done.
---
--- The ReadCloser returned by NewReader also implements Resetter.
---@param r ioReader
---@return any, err
function zlib.NewReader(r) end

--- NewReaderDict is like NewReader but uses a preset dictionary.
--- NewReaderDict ignores the dictionary if the compressed data does not refer to it.
--- If the compressed data refers to a different dictionary, NewReaderDict returns ErrDictionary.
---
--- The ReadCloser returned by NewReaderDict also implements Resetter.
---@param r ioReader
---@param dict byte[]
---@return any, err
function zlib.NewReaderDict(r, dict) end

--- NewWriter creates a new Writer.
--- Writes to the returned Writer are compressed and written to w.
---
--- It is the caller's responsibility to call Close on the Writer when done.
--- Writes may be buffered and not flushed until Close.
---@param w ioWriter
---@return flateWriter
function zlib.NewWriter(w) end

--- NewWriterLevel is like NewWriter but specifies the compression level instead
--- of assuming DefaultCompression.
---
--- The compression level can be DefaultCompression, NoCompression, HuffmanOnly
--- or any integer value between BestSpeed and BestCompression inclusive.
--- The error returned will be nil if the level is valid.
---@param w ioWriter
---@param level number
---@return flateWriter, err
function zlib.NewWriterLevel(w, level) end

--- A Writer takes data written to it and writes the compressed
--- form of that data to an underlying writer (see NewWriter).
---@class flateWriter
local flateWriter = {}

--- Write writes a compressed form of p to the underlying io.Writer. The
--- compressed bytes are not necessarily flushed until the Writer is closed or
--- explicitly flushed.
---@param p byte[]
---@return number, err
function flateWriter:Write(p) end

--- Flush flushes the Writer to its underlying io.Writer.
---@return err
function flateWriter:Flush() end

--- Close closes the Writer, flushing any unwritten data to the underlying
--- io.Writer, but does not close the underlying io.Writer.
---@return err
function flateWriter:Close() end

--- Reset clears the state of the Writer z such that it is equivalent to its
--- initial state from NewWriterLevel or NewWriterLevelDict, but instead writing
--- to w.
---@param w ioWriter
function flateWriter:Reset(w) end

--- Resetter resets a ReadCloser returned by NewReader or NewReaderDict
--- to switch to a new underlying Reader. This permits reusing a ReadCloser
--- instead of allocating a new one.
---@class flateResetter
local flateResetter = {}
