-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class zip
---@field Store any
---@field Deflate any
---@field ErrFormat any
---@field ErrAlgorithm any
---@field ErrChecksum any
---@field ErrInsecurePath any
zip = {}

---{{.zipNewReader}}
---@param r ioReaderAt
---@param size number
---@return tarReader, err
function zip.NewReader(r, size)
end

---{{.zipRegisterDecompressor}}
---@param method any
---@param dcomp zipDecompressor
function zip.RegisterDecompressor(method, dcomp)
end

---{{.zipRegisterCompressor}}
---@param method any
---@param comp zipCompressor
function zip.RegisterCompressor(method, comp)
end

---{{.zipFileInfoHeader}}
---@param fi fsFileInfo
---@return zipFileHeader, err
function zip.FileInfoHeader(fi)
end

---{{.zipNewWriter}}
---@param w ioWriter
---@return tarWriter
function zip.NewWriter(w)
end

---{{.zipOpenReader}}
---@param name string
---@return zipReadCloser, err
function zip.OpenReader(name)
end

---@class zipCompressor
local zipCompressor = {}

---@class zipDecompressor
local zipDecompressor = {}

---@class zipFileHeader
---@field Name string
---@field Comment string
---@field NonUTF8 boolean
---@field CreatorVersion any
---@field ReaderVersion any
---@field Flags any
---@field Method any
---@field Modified any
---@field ModifiedTime any
---@field ModifiedDate any
---@field CRC32 any
---@field CompressedSize any
---@field UncompressedSize any
---@field CompressedSize64 number
---@field UncompressedSize64 number
---@field Extra any
---@field ExternalAttrs any
local zipFileHeader = {}

---{{.zipFileHeaderFileInfo}}
---@return fsFileInfo
function zipFileHeader:FileInfo()
end

---{{.zipFileHeaderModTime}}
---@return any
function zipFileHeader:ModTime()
end

---{{.zipFileHeaderSetModTime}}
---@param t timeTime
function zipFileHeader:SetModTime(t)
end

---{{.zipFileHeaderMode}}
---@return any
function zipFileHeader:Mode()
end

---{{.zipFileHeaderSetMode}}
---@param mode fsFileMode
function zipFileHeader:SetMode(mode)
end

---@class tarWriter
local tarWriter = {}

---{{.tarWriterClose}}
---@return err
function tarWriter:Close()
end

---{{.tarWriterCreateHeader}}
---@param fh zipFileHeader
---@return ioWriter, err
function tarWriter:CreateHeader(fh)
end

---{{.tarWriterCreateRaw}}
---@param fh zipFileHeader
---@return ioWriter, err
function tarWriter:CreateRaw(fh)
end

---{{.tarWriterRegisterCompressor}}
---@param method any
---@param comp zipCompressor
function tarWriter:RegisterCompressor(method, comp)
end

---{{.tarWriterSetOffset}}
---@param n number
function tarWriter:SetOffset(n)
end

---{{.tarWriterSetComment}}
---@param comment string
---@return err
function tarWriter:SetComment(comment)
end

---{{.tarWriterCreate}}
---@param name string
---@return ioWriter, err
function tarWriter:Create(name)
end

---{{.tarWriterCopy}}
---@param f zipFile
---@return err
function tarWriter:Copy(f)
end

---{{.tarWriterFlush}}
---@return err
function tarWriter:Flush()
end

---@class zipReadCloser
local zipReadCloser = {}

---{{.zipReadCloserClose}}
---@return err
function zipReadCloser:Close()
end

---@class zipFile
local zipFile = {}

---{{.zipFileDataOffset}}
---@return number, err
function zipFile:DataOffset()
end

---{{.zipFileOpen}}
---@return any, err
function zipFile:Open()
end

---{{.zipFileOpenRaw}}
---@return ioReader, err
function zipFile:OpenRaw()
end

---@class tarReader
---@field File any
---@field Comment string
local tarReader = {}

---{{.tarReaderOpen}}
---@param name string
---@return any, err
function tarReader:Open(name)
end

---{{.tarReaderRegisterDecompressor}}
---@param method any
---@param dcomp zipDecompressor
function tarReader:RegisterDecompressor(method, dcomp)
end
