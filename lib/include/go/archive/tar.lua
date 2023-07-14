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

--- NewReader creates a new Reader reading from r.
---@param r ioReader
---@return tarReader
function tar.NewReader(r)
end

--- NewWriter creates a new Writer writing to w.
---@param w ioWriter
---@return tarWriter
function tar.NewWriter(w)
end

--- FileInfoHeader creates a partially-populated Header from fi.
--- If fi describes a symlink, FileInfoHeader records link as the link target.
--- If fi describes a directory, a slash is appended to the name.
---
--- Since fs.FileInfo's Name method only returns the base name of
--- the file it describes, it may be necessary to modify Header.Name
--- to provide the full path name of the file.
---@param fi fsFileInfo
---@param link string
---@return tarHeader, err
function tar.FileInfoHeader(fi, link)
end

--- A Header represents a single header in a tar archive.
--- Some fields may not be populated.
---
--- For forward compatibility, users that retrieve a Header from Reader.Next,
--- mutate it in some ways, and then pass it back to Writer.WriteHeader
--- should do so by creating a new Header and copying the fields
--- that they are interested in preserving.
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

--- FileInfo returns an fs.FileInfo for the Header.
---@return fsFileInfo
function tarHeader:FileInfo()
end

--- Format represents the tar archive format.
---
--- The original tar format was introduced in Unix V7.
--- Since then, there have been multiple competing formats attempting to
--- standardize or extend the V7 format to overcome its limitations.
--- The most common formats are the USTAR, PAX, and GNU formats,
--- each with their own advantages and limitations.
---
--- The following table captures the capabilities of each format:
---
---	                  |  USTAR |       PAX |       GNU
---	------------------+--------+-----------+----------
---	Name              |   256B | unlimited | unlimited
---	Linkname          |   100B | unlimited | unlimited
---	Size              | uint33 | unlimited |    uint89
---	Mode              | uint21 |    uint21 |    uint57
---	Uid/Gid           | uint21 | unlimited |    uint57
---	Uname/Gname       |    32B | unlimited |       32B
---	ModTime           | uint33 | unlimited |     int89
---	AccessTime        |    n/a | unlimited |     int89
---	ChangeTime        |    n/a | unlimited |     int89
---	Devmajor/Devminor | uint21 |    uint21 |    uint57
---	------------------+--------+-----------+----------
---	string encoding   |  ASCII |     UTF-8 |    binary
---	sub-second times  |     no |       yes |        no
---	sparse files      |     no |       yes |       yes
---
--- The table's upper portion shows the Header fields, where each format reports
--- the maximum number of bytes allowed for each string field and
--- the integer type used to store each numeric field
--- (where timestamps are stored as the number of seconds since the Unix epoch).
---
--- The table's lower portion shows specialized features of each format,
--- such as supported string encodings, support for sub-second timestamps,
--- or support for sparse files.
---
--- The Writer currently provides no support for sparse files.
---@class tarFormat
local tarFormat = {}


---@return string
function tarFormat:String()
end

--- Reader provides sequential access to the contents of a tar archive.
--- Reader.Next advances to the next file in the archive (including the first),
--- and then Reader can be treated as an io.Reader to access the file's data.
---@class tarReader
local tarReader = {}

--- Next advances to the next entry in the tar archive.
--- The Header.Size determines how many bytes can be read for the next file.
--- Any remaining data in the current file is automatically discarded.
--- At the end of the archive, Next returns the error io.EOF.
---
--- If Next encounters a non-local name (as defined by [filepath.IsLocal])
--- and the GODEBUG environment variable contains `tarinsecurepath=0`,
--- Next returns the header with an ErrInsecurePath error.
--- A future version of Go may introduce this behavior by default.
--- Programs that want to accept non-local names can ignore
--- the ErrInsecurePath error and use the returned header.
---@return tarHeader, err
function tarReader:Next()
end

--- Read reads from the current file in the tar archive.
--- It returns (0, io.EOF) when it reaches the end of that file,
--- until Next is called to advance to the next file.
---
--- If the current file is sparse, then the regions marked as a hole
--- are read back as NUL-bytes.
---
--- Calling Read on special types like TypeLink, TypeSymlink, TypeChar,
--- TypeBlock, TypeDir, and TypeFifo returns (0, io.EOF) regardless of what
--- the Header.Size claims.
---@param b byte[]
---@return number, err
function tarReader:Read(b)
end

--- Writer provides sequential writing of a tar archive.
--- Write.WriteHeader begins a new file with the provided Header,
--- and then Writer can be treated as an io.Writer to supply that file's data.
---@class tarWriter
local tarWriter = {}

--- Close closes the tar archive by flushing the padding, and writing the footer.
--- If the current file (from a prior call to WriteHeader) is not fully written,
--- then this returns an error.
---@return err
function tarWriter:Close()
end

--- Write writes to the current file in the tar archive.
--- Write returns the error ErrWriteTooLong if more than
--- Header.Size bytes are written after WriteHeader.
---
--- Calling Write on special types like TypeLink, TypeSymlink, TypeChar,
--- TypeBlock, TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless
--- of what the Header.Size claims.
---@param b byte[]
---@return number, err
function tarWriter:Write(b)
end

--- Flush finishes writing the current file's block padding.
--- The current file must be fully written before Flush can be called.
---
--- This is unnecessary as the next call to WriteHeader or Close
--- will implicitly flush out the file's padding.
---@return err
function tarWriter:Flush()
end

--- WriteHeader writes hdr and prepares to accept the file's contents.
--- The Header.Size determines how many bytes can be written for the next file.
--- If the current file is not fully written, then this returns an error.
--- This implicitly flushes any padding necessary before writing the header.
---@param hdr tarHeader
---@return err
function tarWriter:WriteHeader(hdr)
end
