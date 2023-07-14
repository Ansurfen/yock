-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class bufio
---@field MaxScanTokenSize any
---@field ErrInvalidUnreadByte any
---@field ErrInvalidUnreadRune any
---@field ErrBufferFull any
---@field ErrNegativeCount any
---@field ErrTooLong any
---@field ErrNegativeAdvance any
---@field ErrAdvanceTooFar any
---@field ErrBadReadCount any
---@field ErrFinalToken any
bufio = {}

--- NewReader returns a new Reader whose buffer has the default size.
---@param rd ioReader
---@return bufioReader
function bufio.NewReader(rd)
end

--- NewWriter returns a new Writer whose buffer has the default size.
--- If the argument io.Writer is already a Writer with large enough buffer size,
--- it returns the underlying Writer.
---@param w ioWriter
---@return bufioWriter
function bufio.NewWriter(w)
end

--- NewReadWriter allocates a new ReadWriter that dispatches to r and w.
---@param r bufioReader
---@param w bufioWriter
---@return bufioReadWriter
function bufio.NewReadWriter(r, w)
end

--- ScanWords is a split function for a Scanner that returns each
--- space-separated word of text, with surrounding spaces deleted. It will
--- never return an empty string. The definition of space is set by
--- unicode.IsSpace.
---@param data byte[]
---@param atEOF boolean
---@return number, byte[], err
function bufio.ScanWords(data, atEOF)
end

--- NewScanner returns a new Scanner to read from r.
--- The split function defaults to ScanLines.
---@param r ioReader
---@return bufioScanner
function bufio.NewScanner(r)
end

--- NewWriterSize returns a new Writer whose buffer has at least the specified
--- size. If the argument io.Writer is already a Writer with large enough
--- size, it returns the underlying Writer.
---@param w ioWriter
---@param size number
---@return bufioWriter
function bufio.NewWriterSize(w, size)
end

--- NewReaderSize returns a new Reader whose buffer has at least the specified
--- size. If the argument io.Reader is already a Reader with large enough
--- size, it returns the underlying Reader.
---@param rd ioReader
---@param size number
---@return bufioReader
function bufio.NewReaderSize(rd, size)
end

--- ScanRunes is a split function for a Scanner that returns each
--- UTF-8-encoded rune as a token. The sequence of runes returned is
--- equivalent to that from a range loop over the input as a string, which
--- means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
--- Because of the Scan interface, this makes it impossible for the client to
--- distinguish correctly encoded replacement runes from encoding errors.
---@param data byte[]
---@param atEOF boolean
---@return number, byte[], err
function bufio.ScanRunes(data, atEOF)
end

--- ScanLines is a split function for a Scanner that returns each line of
--- text, stripped of any trailing end-of-line marker. The returned line may
--- be empty. The end-of-line marker is one optional carriage return followed
--- by one mandatory newline. In regular expression notation, it is `\r?\n`.
--- The last non-empty line of input will be returned even if it has no
--- newline.
---@param data byte[]
---@param atEOF boolean
---@return number, byte[], err
function bufio.ScanLines(data, atEOF)
end

--- ScanBytes is a split function for a Scanner that returns each byte as a token.
---@param data byte[]
---@param atEOF boolean
---@return number, byte[], err
function bufio.ScanBytes(data, atEOF)
end

--- Reader implements buffering for an io.Reader object.
---@class bufioReader
local bufioReader = {}

--- Size returns the size of the underlying buffer in bytes.
---@return number
function bufioReader:Size()
end

--- Peek returns the next n bytes without advancing the reader. The bytes stop
--- being valid at the next read call. If Peek returns fewer than n bytes, it
--- also returns an error explaining why the read is short. The error is
--- ErrBufferFull if n is larger than b's buffer size.
---
--- Calling Peek prevents a UnreadByte or UnreadRune call from succeeding
--- until the next read operation.
---@param n number
---@return byte[], err
function bufioReader:Peek(n)
end

--- Discard skips the next n bytes, returning the number of bytes discarded.
---
--- If Discard skips fewer than n bytes, it also returns an error.
--- If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without
--- reading from the underlying io.Reader.
---@param n number
---@return number, err
function bufioReader:Discard(n)
end

--- ReadRune reads a single UTF-8 encoded Unicode character and returns the
--- rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
--- and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
---@return any, number, err
function bufioReader:ReadRune()
end

--- ReadString reads until the first occurrence of delim in the input,
--- returning a string containing the data up to and including the delimiter.
--- If ReadString encounters an error before finding a delimiter,
--- it returns the data read before the error and the error itself (often io.EOF).
--- ReadString returns err != nil if and only if the returned data does not end in
--- delim.
--- For simple uses, a Scanner may be more convenient.
---@param delim byte
---@return string, err
function bufioReader:ReadString(delim)
end

--- UnreadRune unreads the last rune. If the most recent method called on
--- the Reader was not a ReadRune, UnreadRune returns an error. (In this
--- regard it is stricter than UnreadByte, which will unread the last byte
--- from any read operation.)
---@return err
function bufioReader:UnreadRune()
end

--- Buffered returns the number of bytes that can be read from the current buffer.
---@return number
function bufioReader:Buffered()
end

--- ReadSlice reads until the first occurrence of delim in the input,
--- returning a slice pointing at the bytes in the buffer.
--- The bytes stop being valid at the next read.
--- If ReadSlice encounters an error before finding a delimiter,
--- it returns all the data in the buffer and the error itself (often io.EOF).
--- ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.
--- Because the data returned from ReadSlice will be overwritten
--- by the next I/O operation, most clients should use
--- ReadBytes or ReadString instead.
--- ReadSlice returns err != nil if and only if line does not end in delim.
---@param delim byte
---@return byte[], err
function bufioReader:ReadSlice(delim)
end

--- Reset discards any buffered data, resets all state, and switches
--- the buffered reader to read from r.
--- Calling Reset on the zero value of Reader initializes the internal buffer
--- to the default size.
---@param r ioReader
function bufioReader:Reset(r)
end

--- Read reads data into p.
--- It returns the number of bytes read into p.
--- The bytes are taken from at most one Read on the underlying Reader,
--- hence n may be less than len(p).
--- To read exactly len(p) bytes, use io.ReadFull(b, p).
--- If the underlying Reader can return a non-zero count with io.EOF,
--- then this Read method can do so as well; see the [io.Reader] docs.
---@param p byte[]
---@return number, err
function bufioReader:Read(p)
end

--- ReadByte reads and returns a single byte.
--- If no byte is available, returns an error.
---@return byte, err
function bufioReader:ReadByte()
end

--- UnreadByte unreads the last byte. Only the most recently read byte can be unread.
---
--- UnreadByte returns an error if the most recent method called on the
--- Reader was not a read operation. Notably, Peek, Discard, and WriteTo are not
--- considered read operations.
---@return err
function bufioReader:UnreadByte()
end

--- ReadBytes reads until the first occurrence of delim in the input,
--- returning a slice containing the data up to and including the delimiter.
--- If ReadBytes encounters an error before finding a delimiter,
--- it returns the data read before the error and the error itself (often io.EOF).
--- ReadBytes returns err != nil if and only if the returned data does not end in
--- delim.
--- For simple uses, a Scanner may be more convenient.
---@param delim byte
---@return byte[], err
function bufioReader:ReadBytes(delim)
end

--- ReadLine is a low-level line-reading primitive. Most callers should use
--- ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
---
--- ReadLine tries to return a single line, not including the end-of-line bytes.
--- If the line was too long for the buffer then isPrefix is set and the
--- beginning of the line is returned. The rest of the line will be returned
--- from future calls. isPrefix will be false when returning the last fragment
--- of the line. The returned buffer is only valid until the next call to
--- ReadLine. ReadLine either returns a non-nil line or it returns an error,
--- never both.
---
--- The text returned from ReadLine does not include the line end ("\r\n" or "\n").
--- No indication or error is given if the input ends without a final line end.
--- Calling UnreadByte after ReadLine will always unread the last byte read
--- (possibly a character belonging to the line end) even if that byte is not
--- part of the line returned by ReadLine.
---@return byte[], boolean, err
function bufioReader:ReadLine()
end

--- WriteTo implements io.WriterTo.
--- This may make multiple calls to the Read method of the underlying Reader.
--- If the underlying reader supports the WriteTo method,
--- this calls the underlying WriteTo without buffering.
---@param w ioWriter
---@return number, err
function bufioReader:WriteTo(w)
end

--- Writer implements buffering for an io.Writer object.
--- If an error occurs writing to a Writer, no more data will be
--- accepted and all subsequent writes, and Flush, will return the error.
--- After all data has been written, the client should call the
--- Flush method to guarantee all data has been forwarded to
--- the underlying io.Writer.
---@class bufioWriter
local bufioWriter = {}

--- Flush writes any buffered data to the underlying io.Writer.
---@return err
function bufioWriter:Flush()
end

--- Available returns how many bytes are unused in the buffer.
---@return number
function bufioWriter:Available()
end

--- Buffered returns the number of bytes that have been written into the current buffer.
---@return number
function bufioWriter:Buffered()
end

--- Write writes the contents of p into the buffer.
--- It returns the number of bytes written.
--- If nn < len(p), it also returns an error explaining
--- why the write is short.
---@param p byte[]
---@return number, err
function bufioWriter:Write(p)
end

--- WriteRune writes a single Unicode code point, returning
--- the number of bytes written and any error.
---@param r any
---@return number, err
function bufioWriter:WriteRune(r)
end

--- WriteString writes a string.
--- It returns the number of bytes written.
--- If the count is less than len(s), it also returns an error explaining
--- why the write is short.
---@param s string
---@return number, err
function bufioWriter:WriteString(s)
end

--- ReadFrom implements io.ReaderFrom. If the underlying writer
--- supports the ReadFrom method, this calls the underlying ReadFrom.
--- If there is buffered data and an underlying ReadFrom, this fills
--- the buffer and writes it before calling ReadFrom.
---@param r ioReader
---@return number, err
function bufioWriter:ReadFrom(r)
end

--- Size returns the size of the underlying buffer in bytes.
---@return number
function bufioWriter:Size()
end

--- Reset discards any unflushed buffered data, clears any error, and
--- resets b to write its output to w.
--- Calling Reset on the zero value of Writer initializes the internal buffer
--- to the default size.
---@param w ioWriter
function bufioWriter:Reset(w)
end

--- AvailableBuffer returns an empty buffer with b.Available() capacity.
--- This buffer is intended to be appended to and
--- passed to an immediately succeeding Write call.
--- The buffer is only valid until the next write operation on b.
---@return byte[]
function bufioWriter:AvailableBuffer()
end

--- WriteByte writes a single byte.
---@param c byte
---@return err
function bufioWriter:WriteByte(c)
end

--- ReadWriter stores pointers to a Reader and a Writer.
--- It implements io.ReadWriter.
---@class bufioReadWriter
local bufioReadWriter = {}

--- Scanner provides a convenient interface for reading data such as
--- a file of newline-delimited lines of text. Successive calls to
--- the Scan method will step through the 'tokens' of a file, skipping
--- the bytes between the tokens. The specification of a token is
--- defined by a split function of type SplitFunc; the default split
--- function breaks the input into lines with line termination stripped. Split
--- functions are defined in this package for scanning a file into
--- lines, bytes, UTF-8-encoded runes, and space-delimited words. The
--- client may instead provide a custom split function.
---
--- Scanning stops unrecoverably at EOF, the first I/O error, or a token too
--- large to fit in the buffer. When a scan stops, the reader may have
--- advanced arbitrarily far past the last token. Programs that need more
--- control over error handling or large tokens, or must run sequential scans
--- on a reader, should use bufio.Reader instead.
---@class bufioScanner
local bufioScanner = {}

--- Buffer sets the initial buffer to use when scanning and the maximum
--- size of buffer that may be allocated during scanning. The maximum
--- token size is the larger of max and cap(buf). If max <= cap(buf),
--- Scan will use this buffer only and do no allocation.
---
--- By default, Scan uses an internal buffer and sets the
--- maximum token size to MaxScanTokenSize.
---
--- Buffer panics if it is called after scanning has started.
---@param buf byte[]
---@param max number
function bufioScanner:Buffer(buf, max)
end

--- Split sets the split function for the Scanner.
--- The default split function is ScanLines.
---
--- Split panics if it is called after scanning has started.
---@param split bufioSplitFunc
function bufioScanner:Split(split)
end

--- Err returns the first non-EOF error that was encountered by the Scanner.
---@return err
function bufioScanner:Err()
end

--- Bytes returns the most recent token generated by a call to Scan.
--- The underlying array may point to data that will be overwritten
--- by a subsequent call to Scan. It does no allocation.
---@return byte[]
function bufioScanner:Bytes()
end

--- Text returns the most recent token generated by a call to Scan
--- as a newly allocated string holding its bytes.
---@return string
function bufioScanner:Text()
end

--- Scan advances the Scanner to the next token, which will then be
--- available through the Bytes or Text method. It returns false when the
--- scan stops, either by reaching the end of the input or an error.
--- After Scan returns false, the Err method will return any error that
--- occurred during scanning, except that if it was io.EOF, Err
--- will return nil.
--- Scan panics if the split function returns too many empty
--- tokens without advancing the input. This is a common error mode for
--- scanners.
---@return boolean
function bufioScanner:Scan()
end

--- SplitFunc is the signature of the split function used to tokenize the
--- input. The arguments are an initial substring of the remaining unprocessed
--- data and a flag, atEOF, that reports whether the Reader has no more data
--- to give. The return values are the number of bytes to advance the input
--- and the next token to return to the user, if any, plus an error, if any.
---
--- Scanning stops if the function returns an error, in which case some of
--- the input may be discarded. If that error is ErrFinalToken, scanning
--- stops with no error.
---
--- Otherwise, the Scanner advances the input. If the token is not nil,
--- the Scanner returns it to the user. If the token is nil, the
--- Scanner reads more data and continues scanning; if there is no more
--- data--if atEOF was true--the Scanner returns. If the data does not
--- yet hold a complete token, for instance if it has no newline while
--- scanning lines, a SplitFunc can return (0, nil, nil) to signal the
--- Scanner to read more data into the slice and try again with a
--- longer slice starting at the same point in the input.
---
--- The function is never called with an empty data slice unless atEOF
--- is true. If atEOF is true, however, data may be non-empty and,
--- as always, holds unprocessed text.
---@class bufioSplitFunc
local bufioSplitFunc = {}
