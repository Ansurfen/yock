-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class tar
---@field TypeReg any
---@field TypeRegA any
---@field TypeLink any
---@field TypeSymlink any
---@field TypeChar any
---@field TypeBlock any
---@field TypeDir any
---@field TypeFifo any
---@field TypeCont any
---@field TypeXHeader any
---@field TypeXGlobalHeader any
---@field TypeGNUSparse any
---@field TypeGNULongName any
---@field TypeGNULongLink any
---@field FormatUnknown any
---@field FormatUSTAR any
---@field FormatPAX any
---@field FormatGNU any
---@field ErrHeader any
---@field ErrWriteTooLong any
---@field ErrFieldTooLong any
---@field ErrWriteAfterClose any
---@field ErrInsecurePath any
tar = {}

---{{.tarFileInfoHeader}}
---@param fi fsFileInfo
---@param link string
---@return tarHeader, err
function tar.FileInfoHeader(fi, link)
end

---{{.tarNewReader}}
---@param r ioReader
---@return tarReader
function tar.NewReader(r)
end

---{{.tarNewWriter}}
---@param w ioWriter
---@return tarWriter
function tar.NewWriter(w)
end

---@class tarHeader
---@field Typeflag byte
---@field Name string
---@field Linkname string
---@field Size number
---@field Mode number
---@field Uid number
---@field Gid number
---@field Uname string
---@field Gname string
---@field ModTime any
---@field AccessTime any
---@field ChangeTime any
---@field Devmajor number
---@field Devminor number
---@field Xattrs any
---@field PAXRecords any
---@field Format tarFormat
local tarHeader = {}

---{{.tarHeaderFileInfo}}
---@return fsFileInfo
function tarHeader:FileInfo()
end

---@class tarFormat
local tarFormat = {}

---{{.tarFormatString}}
---@return string
function tarFormat:String()
end

---@class tarReader
local tarReader = {}

---{{.tarReaderRead}}
---@param b byte[]
---@return number, err
function tarReader:Read(b)
end

---{{.tarReaderNext}}
---@return tarHeader, err
function tarReader:Next()
end

---@class tarWriter
local tarWriter = {}

---{{.tarWriterFlush}}
---@return err
function tarWriter:Flush()
end

---{{.tarWriterWriteHeader}}
---@param hdr tarHeader
---@return err
function tarWriter:WriteHeader(hdr)
end

---{{.tarWriterWrite}}
---@param b byte[]
---@return number, err
function tarWriter:Write(b)
end

---{{.tarWriterClose}}
---@return err
function tarWriter:Close()
end
