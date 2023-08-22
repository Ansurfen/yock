-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class io
---@field SeekStart any
---@field SeekCurrent any
---@field SeekEnd any
---@field ErrShortWrite any
---@field ErrShortBuffer any
---@field EOF any
---@field ErrUnexpectedEOF any
---@field ErrNoProgress any
---@field Discard any
---@field ErrClosedPipe any
io = {}

---@class ioWriter: any

--- ReadFull reads exactly len(buf) bytes from r into buf.
--- It returns the number of bytes copied and an error if fewer bytes were read.
--- The error is EOF only if no bytes were read.
--- If an EOF happens after reading some but not all the bytes,
--- ReadFull returns ErrUnexpectedEOF.
--- On return, n == len(buf) if and only if err == nil.
--- If r returns an error having read at least len(buf) bytes, the error is dropped.
---@param r flateReader
---@param buf byte[]
---@return number, err
function io.ReadFull(r, buf) end

--- CopyN copies n bytes (or until an error) from src to dst.
--- It returns the number of bytes copied and the earliest
--- error encountered while copying.
--- On return, written == n if and only if err == nil.
---
--- If dst implements the ReaderFrom interface,
--- the copy is implemented using it.
---@param dst flateWriter
---@param src flateReader
---@param n number
---@return number, err
function io.CopyN(dst, src, n) end

--- NewOffsetWriter returns an OffsetWriter that writes to w
--- starting at offset off.
---@param w ioWriterAt
---@param off number
---@return ioOffsetWriter
function io.NewOffsetWriter(w, off) end

--- ReadAtLeast reads from r into buf until it has read at least min bytes.
--- It returns the number of bytes copied and an error if fewer bytes were read.
--- The error is EOF only if no bytes were read.
--- If an EOF happens after reading fewer than min bytes,
--- ReadAtLeast returns ErrUnexpectedEOF.
--- If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
--- On return, n >= min if and only if err == nil.
--- If r returns an error having read at least min bytes, the error is dropped.
---@param r flateReader
---@param buf byte[]
---@param min number
---@return number, err
function io.ReadAtLeast(r, buf, min) end

--- Pipe creates a synchronous in-memory pipe.
--- It can be used to connect code expecting an io.Reader
--- with code expecting an io.Writer.
---
--- Reads and Writes on the pipe are matched one to one
--- except when multiple Reads are needed to consume a single Write.
--- That is, each Write to the PipeWriter blocks until it has satisfied
--- one or more Reads from the PipeReader that fully consume
--- the written data.
--- The data is copied directly from the Write to the corresponding
--- Read (or Reads); there is no internal buffering.
---
--- It is safe to call Read and Write in parallel with each other or with Close.
--- Parallel calls to Read and parallel calls to Write are also safe:
--- the individual calls will be gated sequentially.
---@return ioPipeReader, ioPipeWriter
function io.Pipe() end

--- LimitReader returns a Reader that reads from r
--- but stops with EOF after n bytes.
--- The underlying implementation is a *LimitedReader.
---@param r flateReader
---@param n number
---@return flateReader
function io.LimitReader(r, n) end

--- ReadAll reads from r until an error or EOF and returns the data it read.
--- A successful call returns err == nil, not err == EOF. Because ReadAll is
--- defined to read from src until EOF, it does not treat an EOF from Read
--- as an error to be reported.
---@param r flateReader
---@return byte[], err
function io.ReadAll(r) end

--- NopCloser returns a ReadCloser with a no-op Close method wrapping
--- the provided Reader r.
--- If r implements WriterTo, the returned ReadCloser will implement WriterTo
--- by forwarding calls to r.
---@param r flateReader
---@return ioReadCloser
function io.NopCloser(r) end

--- WriteString writes the contents of the string s to w, which accepts a slice of bytes.
--- If w implements StringWriter, its WriteString method is invoked directly.
--- Otherwise, w.Write is called exactly once.
---@param w flateWriter
---@param s string
---@return number, err
function io.WriteString(w, s) end

--- Copy copies from src to dst until either EOF is reached
--- on src or an error occurs. It returns the number of bytes
--- copied and the first error encountered while copying, if any.
---
--- A successful Copy returns err == nil, not err == EOF.
--- Because Copy is defined to read from src until EOF, it does
--- not treat an EOF from Read as an error to be reported.
---
--- If src implements the WriterTo interface,
--- the copy is implemented by calling src.WriteTo(dst).
--- Otherwise, if dst implements the ReaderFrom interface,
--- the copy is implemented by calling dst.ReadFrom(src).
---@param dst flateWriter
---@param src flateReader
---@return number, err
function io.Copy(dst, src) end

--- CopyBuffer is identical to Copy except that it stages through the
--- provided buffer (if one is required) rather than allocating a
--- temporary one. If buf is nil, one is allocated; otherwise if it has
--- zero length, CopyBuffer panics.
---
--- If either src implements WriterTo or dst implements ReaderFrom,
--- buf will not be used to perform the copy.
---@param dst flateWriter
---@param src flateReader
---@param buf byte[]
---@return number, err
function io.CopyBuffer(dst, src, buf) end

--- MultiReader returns a Reader that's the logical concatenation of
--- the provided input readers. They're read sequentially. Once all
--- inputs have returned EOF, Read will return EOF.  If any of the readers
--- return a non-nil, non-EOF error, Read will return that error.
---@vararg flateReader
---@return flateReader
function io.MultiReader(...) end

--- MultiWriter creates a writer that duplicates its writes to all the
--- provided writers, similar to the Unix tee(1) command.
---
--- Each write is written to each listed writer, one at a time.
--- If a listed writer returns an error, that overall write operation
--- stops and returns the error; it does not continue down the list.
---@vararg flateWriter
---@return flateWriter
function io.MultiWriter(...) end

--- NewSectionReader returns a SectionReader that reads from r
--- starting at offset off and stops with EOF after n bytes.
---@param r ioReaderAt
---@param off number
---@param n number
---@return ioSectionReader
function io.NewSectionReader(r, off, n) end

--- TeeReader returns a Reader that writes to w what it reads from r.
--- All reads from r performed through it are matched with
--- corresponding writes to w. There is no internal buffering -
--- the write must complete before the read completes.
--- Any error encountered while writing is reported as a read error.
---@param r flateReader
---@param w flateWriter
---@return flateReader
function io.TeeReader(r, w) end

--- RuneScanner is the interface that adds the UnreadRune method to the
--- basic ReadRune method.
---
--- UnreadRune causes the next call to ReadRune to return the last rune read.
--- If the last operation was not a successful call to ReadRune, UnreadRune may
--- return an error, unread the last rune read (or the rune prior to the
--- last-unread rune), or (in implementations that support the Seeker interface)
--- seek to the start of the rune before the current offset.
---@class ioRuneScanner
local ioRuneScanner = {}

--- ReadWriter is the interface that groups the basic Read and Write methods.
---@class ioReadWriter
local ioReadWriter = {}

--- An OffsetWriter maps writes at offset base to offset base+off in the underlying writer.
---@class ioOffsetWriter
local ioOffsetWriter = {}


---@param p byte[]
---@return number, err
function ioOffsetWriter:Write(p) end


---@param p byte[]
---@param off number
---@return number, err
function ioOffsetWriter:WriteAt(p, off) end


---@param offset number
---@param whence number
---@return number, err
function ioOffsetWriter:Seek(offset, whence) end

--- Seeker is the interface that wraps the basic Seek method.
---
--- Seek sets the offset for the next Read or Write to offset,
--- interpreted according to whence:
--- SeekStart means relative to the start of the file,
--- SeekCurrent means relative to the current offset, and
--- SeekEnd means relative to the end
--- (for example, offset = -2 specifies the penultimate byte of the file).
--- Seek returns the new offset relative to the start of the
--- file or an error, if any.
---
--- Seeking to an offset before the start of the file is an error.
--- Seeking to any positive offset may be allowed, but if the new offset exceeds
--- the size of the underlying object the behavior of subsequent I/O operations
--- is implementation-dependent.
---@class ioSeeker
local ioSeeker = {}

--- ReadWriteSeeker is the interface that groups the basic Read, Write and Seek methods.
---@class ioReadWriteSeeker
local ioReadWriteSeeker = {}

--- WriterAt is the interface that wraps the basic WriteAt method.
---
--- WriteAt writes len(p) bytes from p to the underlying data stream
--- at offset off. It returns the number of bytes written from p (0 <= n <= len(p))
--- and any error encountered that caused the write to stop early.
--- WriteAt must return a non-nil error if it returns n < len(p).
---
--- If WriteAt is writing to a destination with a seek offset,
--- WriteAt should not affect nor be affected by the underlying
--- seek offset.
---
--- Clients of WriteAt can execute parallel WriteAt calls on the same
--- destination if the ranges do not overlap.
---
--- Implementations must not retain p.
---@class ioWriterAt
local ioWriterAt = {}

--- SectionReader implements Read, Seek, and ReadAt on a section
--- of an underlying ReaderAt.
---@class ioSectionReader
local ioSectionReader = {}


---@param p byte[]
---@return number, err
function ioSectionReader:Read(p) end


---@param offset number
---@param whence number
---@return number, err
function ioSectionReader:Seek(offset, whence) end


---@param p byte[]
---@param off number
---@return number, err
function ioSectionReader:ReadAt(p, off) end

--- Size returns the size of the section in bytes.
---@return number
function ioSectionReader:Size() end

--- ReadWriteCloser is the interface that groups the basic Read, Write and Close methods.
---@class ioReadWriteCloser
local ioReadWriteCloser = {}

--- A LimitedReader reads from R but limits the amount of
--- data returned to just N bytes. Each call to Read
--- updates N to reflect the new amount remaining.
--- Read returns EOF when N <= 0 or when the underlying R returns EOF.
---@class ioLimitedReader
---@field R flateReader
---@field N number
local ioLimitedReader = {}


---@param p byte[]
---@return number, err
function ioLimitedReader:Read(p) end

--- ReadCloser is the interface that groups the basic Read and Close methods.
---@class ioReadCloser
local ioReadCloser = {}

--- A PipeWriter is the write half of a pipe.
---@class ioPipeWriter
local ioPipeWriter = {}

--- Close closes the writer; subsequent reads from the
--- read half of the pipe will return no bytes and EOF.
---@return err
function ioPipeWriter:Close() end

--- CloseWithError closes the writer; subsequent reads from the
--- read half of the pipe will return no bytes and the error err,
--- or EOF if err is nil.
---
--- CloseWithError never overwrites the previous error if it exists
--- and always returns nil.
---@param err err
---@return err
function ioPipeWriter:CloseWithError(err) end

--- Write implements the standard Write interface:
--- it writes data to the pipe, blocking until one or more readers
--- have consumed all the data or the read end is closed.
--- If the read end is closed with an error, that err is
--- returned as err; otherwise err is ErrClosedPipe.
---@param data byte[]
---@return number, err
function ioPipeWriter:Write(data) end

--- ByteReader is the interface that wraps the ReadByte method.
---
--- ReadByte reads and returns the next byte from the input or
--- any error encountered. If ReadByte returns an error, no input
--- byte was consumed, and the returned byte value is undefined.
---
--- ReadByte provides an efficient interface for byte-at-time
--- processing. A Reader that does not implement  ByteReader
--- can be wrapped using bufio.NewReader to add this method.
---@class ioByteReader
local ioByteReader = {}

--- Closer is the interface that wraps the basic Close method.
---
--- The behavior of Close after the first call is undefined.
--- Specific implementations may document their own behavior.
---@class ioCloser
local ioCloser = {}

--- ReaderAt is the interface that wraps the basic ReadAt method.
---
--- ReadAt reads len(p) bytes into p starting at offset off in the
--- underlying input source. It returns the number of bytes
--- read (0 <= n <= len(p)) and any error encountered.
---
--- When ReadAt returns n < len(p), it returns a non-nil error
--- explaining why more bytes were not returned. In this respect,
--- ReadAt is stricter than Read.
---
--- Even if ReadAt returns n < len(p), it may use all of p as scratch
--- space during the call. If some data is available but not len(p) bytes,
--- ReadAt blocks until either all the data is available or an error occurs.
--- In this respect ReadAt is different from Read.
---
--- If the n = len(p) bytes returned by ReadAt are at the end of the
--- input source, ReadAt may return either err == EOF or err == nil.
---
--- If ReadAt is reading from an input source with a seek offset,
--- ReadAt should not affect nor be affected by the underlying
--- seek offset.
---
--- Clients of ReadAt can execute parallel ReadAt calls on the
--- same input source.
---
--- Implementations must not retain p.
---@class ioReaderAt
local ioReaderAt = {}

--- StringWriter is the interface that wraps the WriteString method.
---@class ioStringWriter
local ioStringWriter = {}

--- A PipeReader is the read half of a pipe.
---@class ioPipeReader
local ioPipeReader = {}

--- Read implements the standard Read interface:
--- it reads data from the pipe, blocking until a writer
--- arrives or the write end is closed.
--- If the write end is closed with an error, that error is
--- returned as err; otherwise err is EOF.
---@param data byte[]
---@return number, err
function ioPipeReader:Read(data) end

--- Close closes the reader; subsequent writes to the
--- write half of the pipe will return the error ErrClosedPipe.
---@return err
function ioPipeReader:Close() end

--- CloseWithError closes the reader; subsequent writes
--- to the write half of the pipe will return the error err.
---
--- CloseWithError never overwrites the previous error if it exists
--- and always returns nil.
---@param err err
---@return err
function ioPipeReader:CloseWithError(err) end

--- WriteSeeker is the interface that groups the basic Write and Seek methods.
---@class ioWriteSeeker
local ioWriteSeeker = {}

--- ReadSeekCloser is the interface that groups the basic Read, Seek and Close
--- methods.
---@class ioReadSeekCloser
local ioReadSeekCloser = {}

--- ByteWriter is the interface that wraps the WriteByte method.
---@class ioByteWriter
local ioByteWriter = {}

--- Reader is the interface that wraps the basic Read method.
---
--- Read reads up to len(p) bytes into p. It returns the number of bytes
--- read (0 <= n <= len(p)) and any error encountered. Even if Read
--- returns n < len(p), it may use all of p as scratch space during the call.
--- If some data is available but not len(p) bytes, Read conventionally
--- returns what is available instead of waiting for more.
---
--- When Read encounters an error or end-of-file condition after
--- successfully reading n > 0 bytes, it returns the number of
--- bytes read. It may return the (non-nil) error from the same call
--- or return the error (and n == 0) from a subsequent call.
--- An instance of this general case is that a Reader returning
--- a non-zero number of bytes at the end of the input stream may
--- return either err == EOF or err == nil. The next Read should
--- return 0, EOF.
---
--- Callers should always process the n > 0 bytes returned before
--- considering the error err. Doing so correctly handles I/O errors
--- that happen after reading some bytes and also both of the
--- allowed EOF behaviors.
---
--- Implementations of Read are discouraged from returning a
--- zero byte count with a nil error, except when len(p) == 0.
--- Callers should treat a return of 0 and nil as indicating that
--- nothing happened; in particular it does not indicate EOF.
---
--- Implementations must not retain p.
---@class flateReader
local flateReader = {}

--- ReadSeeker is the interface that groups the basic Read and Seek methods.
---@class ioReadSeeker
local ioReadSeeker = {}

--- RuneReader is the interface that wraps the ReadRune method.
---
--- ReadRune reads a single encoded Unicode character
--- and returns the rune and its size in bytes. If no character is
--- available, err will be set.
---@class ioRuneReader
local ioRuneReader = {}

--- ReaderFrom is the interface that wraps the ReadFrom method.
---
--- ReadFrom reads data from r until EOF or error.
--- The return value n is the number of bytes read.
--- Any error except EOF encountered during the read is also returned.
---
--- The Copy function uses ReaderFrom if available.
---@class ioReaderFrom
local ioReaderFrom = {}

--- WriterTo is the interface that wraps the WriteTo method.
---
--- WriteTo writes data to w until there's no more data to write or
--- when an error occurs. The return value n is the number of bytes
--- written. Any error encountered during the write is also returned.
---
--- The Copy function uses WriterTo if available.
---@class ioWriterTo
local ioWriterTo = {}

--- WriteCloser is the interface that groups the basic Write and Close methods.
---@class ioWriteCloser
local ioWriteCloser = {}

--- Writer is the interface that wraps the basic Write method.
---
--- Write writes len(p) bytes from p to the underlying data stream.
--- It returns the number of bytes written from p (0 <= n <= len(p))
--- and any error encountered that caused the write to stop early.
--- Write must return a non-nil error if it returns n < len(p).
--- Write must not modify the slice data, even temporarily.
---
--- Implementations must not retain p.
---@class flateWriter
local flateWriter = {}

--- ByteScanner is the interface that adds the UnreadByte method to the
--- basic ReadByte method.
---
--- UnreadByte causes the next call to ReadByte to return the last byte read.
--- If the last operation was not a successful call to ReadByte, UnreadByte may
--- return an error, unread the last byte read (or the byte prior to the
--- last-unread byte), or (in implementations that support the Seeker interface)
--- seek to one byte before the current offset.
---@class ioByteScanner
local ioByteScanner = {}
