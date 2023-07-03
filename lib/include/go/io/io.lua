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

---{{.ioCopyN}}
---@param dst ioWriter
---@param src ioReader
---@param n number
---@return number, err
function io.CopyN(dst, src, n)
end

---{{.ioCopy}}
---@param dst ioWriter
---@param src ioReader
---@return number, err
function io.Copy(dst, src)
end

---{{.ioCopyBuffer}}
---@param dst ioWriter
---@param src ioReader
---@param buf byte[]
---@return number, err
function io.CopyBuffer(dst, src, buf)
end

---{{.ioPipe}}
---@return ioPipeReader, ioPipeWriter
function io.Pipe()
end

---{{.ioReadAtLeast}}
---@param r ioReader
---@param buf byte[]
---@param min number
---@return number, err
function io.ReadAtLeast(r, buf, min)
end

---{{.ioWriteString}}
---@param w ioWriter
---@param s string
---@return number, err
function io.WriteString(w, s)
end

---{{.ioNewSectionReader}}
---@param r ioReaderAt
---@param off number
---@param n number
---@return ioSectionReader
function io.NewSectionReader(r, off, n)
end

---{{.ioMultiReader}}
---@vararg ioReader
---@return ioReader
function io.MultiReader(...)
end

---{{.ioLimitReader}}
---@param r ioReader
---@param n number
---@return ioReader
function io.LimitReader(r, n)
end

---{{.ioReadAll}}
---@param r ioReader
---@return byte[], err
function io.ReadAll(r)
end

---{{.ioReadFull}}
---@param r ioReader
---@param buf byte[]
---@return number, err
function io.ReadFull(r, buf)
end

---{{.ioNewOffsetWriter}}
---@param w ioWriterAt
---@param off number
---@return ioOffsetWriter
function io.NewOffsetWriter(w, off)
end

---{{.ioNopCloser}}
---@param r ioReader
---@return ioReadCloser
function io.NopCloser(r)
end

---{{.ioTeeReader}}
---@param r ioReader
---@param w ioWriter
---@return ioReader
function io.TeeReader(r, w)
end

---{{.ioMultiWriter}}
---@vararg ioWriter
---@return ioWriter
function io.MultiWriter(...)
end

---@class ioReader
local ioReader = {}

---@class ioReadSeekCloser
local ioReadSeekCloser = {}

---@class ioWriterAt
local ioWriterAt = {}

---@class ioLimitedReader
---@field R ioReader
---@field N number
local ioLimitedReader = {}

---{{.ioLimitedReaderRead}}
---@param p byte[]
---@return number, err
function ioLimitedReader:Read(p)
end

---@class ioByteWriter
local ioByteWriter = {}

---@class ioRuneScanner
local ioRuneScanner = {}

---@class ioPipeWriter
local ioPipeWriter = {}

---{{.ioPipeWriterWrite}}
---@param data byte[]
---@return number, err
function ioPipeWriter:Write(data)
end

---{{.ioPipeWriterClose}}
---@return err
function ioPipeWriter:Close()
end

---{{.ioPipeWriterCloseWithError}}
---@param err err
---@return err
function ioPipeWriter:CloseWithError(err)
end

---@class ioReadWriteSeeker
local ioReadWriteSeeker = {}

---@class ioReadSeeker
local ioReadSeeker = {}

---@class ioReaderFrom
local ioReaderFrom = {}

---@class ioStringWriter
local ioStringWriter = {}

---@class ioWriterTo
local ioWriterTo = {}

---@class ioWriteSeeker
local ioWriteSeeker = {}

---@class ioWriter
local ioWriter = {}

---@class ioRuneReader
local ioRuneReader = {}

---@class ioReadWriter
local ioReadWriter = {}

---@class ioReaderAt
local ioReaderAt = {}

---@class ioByteReader
local ioByteReader = {}

---@class ioByteScanner
local ioByteScanner = {}

---@class ioSectionReader
local ioSectionReader = {}

---{{.ioSectionReaderRead}}
---@param p byte[]
---@return number, err
function ioSectionReader:Read(p)
end

---{{.ioSectionReaderSeek}}
---@param offset number
---@param whence number
---@return number, err
function ioSectionReader:Seek(offset, whence)
end

---{{.ioSectionReaderReadAt}}
---@param p byte[]
---@param off number
---@return number, err
function ioSectionReader:ReadAt(p, off)
end

---{{.ioSectionReaderSize}}
---@return number
function ioSectionReader:Size()
end

---@class ioSeeker
local ioSeeker = {}

---@class ioCloser
local ioCloser = {}

---@class ioPipeReader
local ioPipeReader = {}

---{{.ioPipeReaderRead}}
---@param data byte[]
---@return number, err
function ioPipeReader:Read(data)
end

---{{.ioPipeReaderClose}}
---@return err
function ioPipeReader:Close()
end

---{{.ioPipeReaderCloseWithError}}
---@param err err
---@return err
function ioPipeReader:CloseWithError(err)
end

---@class ioOffsetWriter
local ioOffsetWriter = {}

---{{.ioOffsetWriterWrite}}
---@param p byte[]
---@return number, err
function ioOffsetWriter:Write(p)
end

---{{.ioOffsetWriterWriteAt}}
---@param p byte[]
---@param off number
---@return number, err
function ioOffsetWriter:WriteAt(p, off)
end

---{{.ioOffsetWriterSeek}}
---@param offset number
---@param whence number
---@return number, err
function ioOffsetWriter:Seek(offset, whence)
end

---@class ioReadWriteCloser
local ioReadWriteCloser = {}

---@class ioWriteCloser
local ioWriteCloser = {}

---@class ioReadCloser
local ioReadCloser = {}
