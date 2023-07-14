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

--- NewReader creates a new Reader reading the given reader.
--- If r does not also implement io.ByteReader,
--- the decompressor may read more data than necessary from r.
---
--- It is the caller's responsibility to call Close on the Reader when done.
---
--- The Reader.Header fields will be valid in the Reader returned.
---@param r ioReader
---@return flateReader, err
function gzip.NewReader(r) end

--- NewWriter returns a new Writer.
--- Writes to the returned writer are compressed and written to w.
---
--- It is the caller's responsibility to call Close on the Writer when done.
--- Writes may be buffered and not flushed until Close.
---
--- Callers that wish to set the fields in Writer.Header must do so before
--- the first call to Write, Flush, or Close.
---@param w ioWriter
---@return flateWriter
function gzip.NewWriter(w) end

--- NewWriterLevel is like NewWriter but specifies the compression level instead
--- of assuming DefaultCompression.
---
--- The compression level can be DefaultCompression, NoCompression, HuffmanOnly
--- or any integer value between BestSpeed and BestCompression inclusive.
--- The error returned will be nil if the level is valid.
---@param w ioWriter
---@param level number
---@return flateWriter, err
function gzip.NewWriterLevel(w, level) end

--- A Reader is an io.Reader that can be read to retrieve
--- uncompressed data from a gzip-format compressed file.
---
--- In general, a gzip file can be a concatenation of gzip files,
--- each with its own header. Reads from the Reader
--- return the concatenation of the uncompressed data of each.
--- Only the first header is recorded in the Reader fields.
---
--- Gzip files store a length and checksum of the uncompressed data.
--- The Reader will return an ErrChecksum when Read
--- reaches the end of the uncompressed data if it does not
--- have the expected length or checksum. Clients should treat data
--- returned by Read as tentative until they receive the io.EOF
--- marking the end of the data.
---@class flateReader
local flateReader = {}

--- Read implements io.Reader, reading uncompressed bytes from its underlying Reader.
---@param p byte[]
---@return number, err
function flateReader:Read(p) end

--- Close closes the Reader. It does not close the underlying io.Reader.
--- In order for the GZIP checksum to be verified, the reader must be
--- fully consumed until the io.EOF.
---@return err
function flateReader:Close() end

--- Reset discards the Reader z's state and makes it equivalent to the
--- result of its original state from NewReader, but reading from r instead.
--- This permits reusing a Reader rather than allocating a new one.
---@param r ioReader
---@return err
function flateReader:Reset(r) end

--- Multistream controls whether the reader supports multistream files.
---
--- If enabled (the default), the Reader expects the input to be a sequence
--- of individually gzipped data streams, each with its own header and
--- trailer, ending at EOF. The effect is that the concatenation of a sequence
--- of gzipped files is treated as equivalent to the gzip of the concatenation
--- of the sequence. This is standard behavior for gzip readers.
---
--- Calling Multistream(false) disables this behavior; disabling the behavior
--- can be useful when reading file formats that distinguish individual gzip
--- data streams or mix gzip data streams with other data streams.
--- In this mode, when the Reader reaches the end of the data stream,
--- Read returns io.EOF. The underlying reader must implement io.ByteReader
--- in order to be left positioned just after the gzip stream.
--- To start the next stream, call z.Reset(r) followed by z.Multistream(false).
--- If there is no next stream, z.Reset(r) will return io.EOF.
---@param ok boolean
function flateReader:Multistream(ok) end

--- A Writer is an io.WriteCloser.
--- Writes to a Writer are compressed and written to w.
---@class flateWriter
local flateWriter = {}

--- Close closes the Writer by flushing any unwritten data to the underlying
--- io.Writer and writing the GZIP footer.
--- It does not close the underlying io.Writer.
---@return err
function flateWriter:Close() end

--- Reset discards the Writer z's state and makes it equivalent to the
--- result of its original state from NewWriter or NewWriterLevel, but
--- writing to w instead. This permits reusing a Writer rather than
--- allocating a new one.
---@param w ioWriter
function flateWriter:Reset(w) end

--- Write writes a compressed form of p to the underlying io.Writer. The
--- compressed bytes are not necessarily flushed until the Writer is closed.
---@param p byte[]
---@return number, err
function flateWriter:Write(p) end

--- Flush flushes any pending compressed data to the underlying writer.
---
--- It is useful mainly in compressed network protocols, to ensure that
--- a remote reader has enough data to reconstruct a packet. Flush does
--- not return until the data has been written. If the underlying
--- writer returns an error, Flush returns that error.
---
--- In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.
---@return err
function flateWriter:Flush() end

--- The gzip file stores a header giving metadata about the compressed file.
--- That header is exposed as the fields of the Writer and Reader structs.
---
--- Strings must be UTF-8 encoded and may only contain Unicode code points
--- U+0001 through U+00FF, due to limitations of the GZIP file format.
---@class gzipHeader
---@field Comment string
---@field Extra any
---@field ModTime any
---@field Name string
---@field OS byte
local gzipHeader = {}
