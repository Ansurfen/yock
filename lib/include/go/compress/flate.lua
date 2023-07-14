-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class flate
---@field NoCompression any
---@field BestSpeed any
---@field BestCompression any
---@field DefaultCompression any
---@field HuffmanOnly any
flate = {}

--- NewWriter returns a new Writer compressing data at the given level.
--- Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression);
--- higher levels typically run slower but compress more. Level 0
--- (NoCompression) does not attempt any compression; it only adds the
--- necessary DEFLATE framing.
--- Level -1 (DefaultCompression) uses the default compression level.
--- Level -2 (HuffmanOnly) will use Huffman compression only, giving
--- a very fast compression for all types of input, but sacrificing considerable
--- compression efficiency.
---
--- If level is in the range [-2, 9] then the error returned will be nil.
--- Otherwise the error returned will be non-nil.
---@param w ioWriter
---@param level number
---@return flateWriter, err
function flate.NewWriter(w, level) end

--- NewWriterDict is like NewWriter but initializes the new
--- Writer with a preset dictionary. The returned Writer behaves
--- as if the dictionary had been written to it without producing
--- any compressed output. The compressed data written to w
--- can only be decompressed by a Reader initialized with the
--- same dictionary.
---@param w ioWriter
---@param level number
---@param dict byte[]
---@return flateWriter, err
function flate.NewWriterDict(w, level, dict) end

--- NewReader returns a new ReadCloser that can be used
--- to read the uncompressed version of r.
--- If r does not also implement io.ByteReader,
--- the decompressor may read more data than necessary from r.
--- The reader returns io.EOF after the final block in the DEFLATE stream has
--- been encountered. Any trailing data after the final block is ignored.
---
--- The ReadCloser returned by NewReader also implements Resetter.
---@param r ioReader
---@return any
function flate.NewReader(r) end

--- NewReaderDict is like NewReader but initializes the reader
--- with a preset dictionary. The returned Reader behaves as if
--- the uncompressed data stream started with the given dictionary,
--- which has already been read. NewReaderDict is typically used
--- to read data compressed by NewWriterDict.
---
--- The ReadCloser returned by NewReader also implements Resetter.
---@param r ioReader
---@param dict byte[]
---@return any
function flate.NewReaderDict(r, dict) end

--- A ReadError reports an error encountered while reading input.
---
--- Deprecated: No longer returned.
---@class flateReadError
---@field Offset number
---@field Err err
local flateReadError = {}


---@return string
function flateReadError:Error() end

--- A WriteError reports an error encountered while writing output.
---
--- Deprecated: No longer returned.
---@class flateWriteError
---@field Offset number
---@field Err err
local flateWriteError = {}


---@return string
function flateWriteError:Error() end

--- Resetter resets a ReadCloser returned by NewReader or NewReaderDict
--- to switch to a new underlying Reader. This permits reusing a ReadCloser
--- instead of allocating a new one.
---@class flateResetter
local flateResetter = {}

--- The actual read interface needed by NewReader.
--- If the passed in io.Reader does not also have ReadByte,
--- the NewReader will introduce its own buffering.
---@class flateReader
local flateReader = {}

--- A CorruptInputError reports the presence of corrupt input at a given offset.
---@class flateCorruptInputError
local flateCorruptInputError = {}


---@return string
function flateCorruptInputError:Error() end

--- An InternalError reports an error in the flate code itself.
---@class flateInternalError
local flateInternalError = {}


---@return string
function flateInternalError:Error() end

--- A Writer takes data written to it and writes the compressed
--- form of that data to an underlying writer (see NewWriter).
---@class flateWriter
local flateWriter = {}

--- Write writes data to w, which will eventually write the
--- compressed form of data to its underlying writer.
---@param data byte[]
---@return number, err
function flateWriter:Write(data) end

--- Flush flushes any pending data to the underlying writer.
--- It is useful mainly in compressed network protocols, to ensure that
--- a remote reader has enough data to reconstruct a packet.
--- Flush does not return until the data has been written.
--- Calling Flush when there is no pending data still causes the Writer
--- to emit a sync marker of at least 4 bytes.
--- If the underlying writer returns an error, Flush returns that error.
---
--- In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.
---@return err
function flateWriter:Flush() end

--- Close flushes and closes the writer.
---@return err
function flateWriter:Close() end

--- Reset discards the writer's state and makes it equivalent to
--- the result of NewWriter or NewWriterDict called with dst
--- and w's level and dictionary.
---@param dst ioWriter
function flateWriter:Reset(dst) end
