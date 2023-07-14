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

--- NewReader returns a new Reader reading from r, which is assumed to
--- have the given size in bytes.
---
--- If any file inside the archive uses a non-local name
--- (as defined by [filepath.IsLocal]) or a name containing backslashes
--- and the GODEBUG environment variable contains `zipinsecurepath=0`,
--- NewReader returns the reader with an ErrInsecurePath error.
--- A future version of Go may introduce this behavior by default.
--- Programs that want to accept non-local names can ignore
--- the ErrInsecurePath error and use the returned reader.
---@param r ioReaderAt
---@param size number
---@return tarReader, err
function zip.NewReader(r, size)
end

--- RegisterDecompressor allows custom decompressors for a specified method ID.
--- The common methods Store and Deflate are built in.
---@param method any
---@param dcomp zipDecompressor
function zip.RegisterDecompressor(method, dcomp)
end

--- RegisterCompressor registers custom compressors for a specified method ID.
--- The common methods Store and Deflate are built in.
---@param method any
---@param comp zipCompressor
function zip.RegisterCompressor(method, comp)
end

--- FileInfoHeader creates a partially-populated FileHeader from an
--- fs.FileInfo.
--- Because fs.FileInfo's Name method returns only the base name of
--- the file it describes, it may be necessary to modify the Name field
--- of the returned header to provide the full path name of the file.
--- If compression is desired, callers should set the FileHeader.Method
--- field; it is unset by default.
---@param fi fsFileInfo
---@return zipFileHeader, err
function zip.FileInfoHeader(fi)
end

--- NewWriter returns a new Writer writing a zip file to w.
---@param w ioWriter
---@return tarWriter
function zip.NewWriter(w)
end

--- OpenReader will open the Zip file specified by name and return a ReadCloser.
---@param name string
---@return zipReadCloser, err
function zip.OpenReader(name)
end

--- A File is a single file in a ZIP archive.
--- The file information is in the embedded FileHeader.
--- The file content can be accessed by calling Open.
---@class zipFile
local zipFile = {}

--- DataOffset returns the offset of the file's possibly-compressed
--- data, relative to the beginning of the zip file.
---
--- Most callers should instead use Open, which transparently
--- decompresses data and verifies checksums.
---@return number, err
function zipFile:DataOffset()
end

--- Open returns a ReadCloser that provides access to the File's contents.
--- Multiple files may be read concurrently.
---@return any, err
function zipFile:Open()
end

--- OpenRaw returns a Reader that provides access to the File's contents without
--- decompression.
---@return ioReader, err
function zipFile:OpenRaw()
end

--- A Compressor returns a new compressing writer, writing to w.
--- The WriteCloser's Close method must be used to flush pending data to w.
--- The Compressor itself must be safe to invoke from multiple goroutines
--- simultaneously, but each returned writer will be used only by
--- one goroutine at a time.
---@class zipCompressor
local zipCompressor = {}

--- A Decompressor returns a new decompressing reader, reading from r.
--- The ReadCloser's Close method must be used to release associated resources.
--- The Decompressor itself must be safe to invoke from multiple goroutines
--- simultaneously, but each returned reader will be used only by
--- one goroutine at a time.
---@class zipDecompressor
local zipDecompressor = {}

--- FileHeader describes a file within a ZIP file.
--- See the [ZIP specification] for details.
---
--- [ZIP specification]: https://www.pkware.com/appnote
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

--- ModTime returns the modification time in UTC using the legacy
--- ModifiedDate and ModifiedTime fields.
---
--- Deprecated: Use Modified instead.
---@return any
function zipFileHeader:ModTime()
end

--- SetModTime sets the Modified, ModifiedTime, and ModifiedDate fields
--- to the given time in UTC.
---
--- Deprecated: Use Modified instead.
---@param t timeTime
function zipFileHeader:SetModTime(t)
end

--- Mode returns the permission and mode bits for the FileHeader.
---@return fsFileMode
function zipFileHeader:Mode()
end

--- SetMode changes the permission and mode bits for the FileHeader.
---@param mode fsFileMode
function zipFileHeader:SetMode(mode)
end

--- FileInfo returns an fs.FileInfo for the FileHeader.
---@return fsFileInfo
function zipFileHeader:FileInfo()
end

--- Writer implements a zip file writer.
---@class tarWriter
local tarWriter = {}

--- Flush flushes any buffered data to the underlying writer.
--- Calling Flush is not normally necessary; calling Close is sufficient.
---@return err
function tarWriter:Flush()
end

--- Close finishes writing the zip file by writing the central directory.
--- It does not close the underlying writer.
---@return err
function tarWriter:Close()
end

--- Create adds a file to the zip file using the provided name.
--- It returns a Writer to which the file contents should be written.
--- The file contents will be compressed using the Deflate method.
--- The name must be a relative path: it must not start with a drive
--- letter (e.g. C:) or leading slash, and only forward slashes are
--- allowed. To create a directory instead of a file, add a trailing
--- slash to the name.
--- The file's contents must be written to the io.Writer before the next
--- call to Create, CreateHeader, or Close.
---@param name string
---@return ioWriter, err
function tarWriter:Create(name)
end

--- CreateHeader adds a file to the zip archive using the provided FileHeader
--- for the file metadata. Writer takes ownership of fh and may mutate
--- its fields. The caller must not modify fh after calling CreateHeader.
---
--- This returns a Writer to which the file contents should be written.
--- The file's contents must be written to the io.Writer before the next
--- call to Create, CreateHeader, CreateRaw, or Close.
---@param fh zipFileHeader
---@return ioWriter, err
function tarWriter:CreateHeader(fh)
end

--- Copy copies the file f (obtained from a Reader) into w. It copies the raw
--- form directly bypassing decompression, compression, and validation.
---@param f zipFile
---@return err
function tarWriter:Copy(f)
end

--- RegisterCompressor registers or overrides a custom compressor for a specific
--- method ID. If a compressor for a given method is not found, Writer will
--- default to looking up the compressor at the package level.
---@param method any
---@param comp zipCompressor
function tarWriter:RegisterCompressor(method, comp)
end

--- SetOffset sets the offset of the beginning of the zip data within the
--- underlying writer. It should be used when the zip data is appended to an
--- existing file, such as a binary executable.
--- It must be called before any data is written.
---@param n number
function tarWriter:SetOffset(n)
end

--- SetComment sets the end-of-central-directory comment field.
--- It can only be called before Close.
---@param comment string
---@return err
function tarWriter:SetComment(comment)
end

--- CreateRaw adds a file to the zip archive using the provided FileHeader and
--- returns a Writer to which the file contents should be written. The file's
--- contents must be written to the io.Writer before the next call to Create,
--- CreateHeader, CreateRaw, or Close.
---
--- In contrast to CreateHeader, the bytes passed to Writer are not compressed.
---@param fh zipFileHeader
---@return ioWriter, err
function tarWriter:CreateRaw(fh)
end

--- A Reader serves content from a ZIP archive.
---@class tarReader
---@field File any
---@field Comment string
local tarReader = {}

--- Open opens the named file in the ZIP archive,
--- using the semantics of fs.FS.Open:
--- paths are always slash separated, with no
--- leading / or ../ elements.
---@param name string
---@return any, err
function tarReader:Open(name)
end

--- RegisterDecompressor registers or overrides a custom decompressor for a
--- specific method ID. If a decompressor for a given method is not found,
--- Reader will default to looking up the decompressor at the package level.
---@param method any
---@param dcomp zipDecompressor
function tarReader:RegisterDecompressor(method, dcomp)
end

--- A ReadCloser is a Reader that must be closed when no longer needed.
---@class zipReadCloser
local zipReadCloser = {}

--- Close closes the Zip file, rendering it unusable for I/O.
---@return err
function zipReadCloser:Close()
end
