-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class fs
---@field ModeDir any
---@field ModeAppend any
---@field ModeExclusive any
---@field ModeTemporary any
---@field ModeSymlink any
---@field ModeDevice any
---@field ModeNamedPipe any
---@field ModeSocket any
---@field ModeSetuid any
---@field ModeSetgid any
---@field ModeCharDevice any
---@field ModeSticky any
---@field ModeIrregular any
---@field ModeType any
---@field ModePerm any
---@field ErrInvalid any
---@field ErrPermission any
---@field ErrExist any
---@field ErrNotExist any
---@field ErrClosed any
---@field SkipDir any
---@field SkipAll any
fs = {}

--- ReadDir reads the named directory
--- and returns a list of directory entries sorted by filename.
---
--- If fs implements ReadDirFS, ReadDir calls fs.ReadDir.
--- Otherwise ReadDir calls fs.Open and uses ReadDir and Close
--- on the returned file.
---@param fsys fsFS
---@param name string
---@return any, err
function fs.ReadDir(fsys, name) end

--- FileInfoToDirEntry returns a DirEntry that returns information from info.
--- If info is nil, FileInfoToDirEntry returns nil.
---@param info osFileInfo
---@return osDirEntry
function fs.FileInfoToDirEntry(info) end

--- ReadFile reads the named file from the file system fs and returns its contents.
--- A successful call returns a nil error, not io.EOF.
--- (Because ReadFile reads the whole file, the expected EOF
--- from the final Read is not treated as an error to be reported.)
---
--- If fs implements ReadFileFS, ReadFile calls fs.ReadFile.
--- Otherwise ReadFile calls fs.Open and uses Read and Close
--- on the returned file.
---@param fsys fsFS
---@param name string
---@return byte[], err
function fs.ReadFile(fsys, name) end

--- Stat returns a FileInfo describing the named file from the file system.
---
--- If fs implements StatFS, Stat calls fs.Stat.
--- Otherwise, Stat opens the file to stat it.
---@param fsys fsFS
---@param name string
---@return osFileInfo, err
function fs.Stat(fsys, name) end

--- Sub returns an FS corresponding to the subtree rooted at fsys's dir.
---
--- If dir is ".", Sub returns fsys unchanged.
--- Otherwise, if fs implements SubFS, Sub returns fsys.Sub(dir).
--- Otherwise, Sub returns a new FS implementation sub that,
--- in effect, implements sub.Open(name) as fsys.Open(path.Join(dir, name)).
--- The implementation also translates calls to ReadDir, ReadFile, and Glob appropriately.
---
--- Note that Sub(os.DirFS("/"), "prefix") is equivalent to os.DirFS("/prefix")
--- and that neither of them guarantees to avoid operating system
--- accesses outside "/prefix", because the implementation of os.DirFS
--- does not check for symbolic links inside "/prefix" that point to
--- other directories. That is, os.DirFS is not a general substitute for a
--- chroot-style security mechanism, and Sub does not change that fact.
---@param fsys fsFS
---@param dir string
---@return fsFS, err
function fs.Sub(fsys, dir) end

--- WalkDir walks the file tree rooted at root, calling fn for each file or
--- directory in the tree, including root.
---
--- All errors that arise visiting files and directories are filtered by fn:
--- see the fs.WalkDirFunc documentation for details.
---
--- The files are walked in lexical order, which makes the output deterministic
--- but requires WalkDir to read an entire directory into memory before proceeding
--- to walk that directory.
---
--- WalkDir does not follow symbolic links found in directories,
--- but if root itself is a symbolic link, its target will be walked.
---@param fsys fsFS
---@param root string
---@param fn fsWalkDirFunc
---@return err
function fs.WalkDir(fsys, root, fn) end

--- ValidPath reports whether the given path name
--- is valid for use in a call to Open.
---
--- Path names passed to open are UTF-8-encoded,
--- unrooted, slash-separated sequences of path elements, like “x/y/z”.
--- Path names must not contain an element that is “.” or “..” or the empty string,
--- except for the special case that the root directory is named “.”.
--- Paths must not start or end with a slash: “/x” and “x/” are invalid.
---
--- Note that paths are slash-separated on all systems, even Windows.
--- Paths containing other characters such as backslash and colon
--- are accepted as valid, but those characters must never be
--- interpreted by an FS implementation as path element separators.
---@param name string
---@return boolean
function fs.ValidPath(name) end

--- Glob returns the names of all files matching pattern or nil
--- if there is no matching file. The syntax of patterns is the same
--- as in path.Match. The pattern may describe hierarchical names such as
--- usr/*/bin/ed.
---
--- Glob ignores file system errors such as I/O errors reading directories.
--- The only possible returned error is path.ErrBadPattern, reporting that
--- the pattern is malformed.
---
--- If fs implements GlobFS, Glob calls fs.Glob.
--- Otherwise, Glob uses ReadDir to traverse the directory tree
--- and look for matches for the pattern.
---@param fsys fsFS
---@param pattern string
---@return string[], err
function fs.Glob(fsys, pattern) end

--- A DirEntry is an entry read from a directory
--- (using the ReadDir function or a ReadDirFile's ReadDir method).
---@class osDirEntry
local osDirEntry = {}

--- PathError records an error and the operation and file path that caused it.
---@class osPathError
---@field Op string
---@field Path string
---@field Err err
local osPathError = {}


---@return string
function osPathError:Error() end


---@return err
function osPathError:Unwrap() end

--- Timeout reports whether this error represents a timeout.
---@return boolean
function osPathError:Timeout() end

--- An FS provides access to a hierarchical file system.
---
--- The FS interface is the minimum implementation required of the file system.
--- A file system may implement additional interfaces,
--- such as ReadFileFS, to provide additional or optimized functionality.
---@class fsFS
local fsFS = {}

--- ReadDirFS is the interface implemented by a file system
--- that provides an optimized implementation of ReadDir.
---@class fsReadDirFS
local fsReadDirFS = {}

--- ReadFileFS is the interface implemented by a file system
--- that provides an optimized implementation of ReadFile.
---@class fsReadFileFS
local fsReadFileFS = {}

--- A StatFS is a file system with a Stat method.
---@class fsStatFS
local fsStatFS = {}

--- A ReadDirFile is a directory file whose entries can be read with the ReadDir method.
--- Every directory file should implement this interface.
--- (It is permissible for any file to implement this interface,
--- but if so ReadDir should return an error for non-directories.)
---@class fsReadDirFile
local fsReadDirFile = {}

--- A FileInfo describes a file and is returned by Stat.
---@class osFileInfo
local osFileInfo = {}

--- A FileMode represents a file's mode and permission bits.
--- The bits have the same definition on all systems, so that
--- information about files can be moved from one system
--- to another portably. Not all bits apply to all systems.
--- The only required bit is ModeDir for directories.
---@class osFileMode
local osFileMode = {}


---@return string
function osFileMode:String() end

--- IsDir reports whether m describes a directory.
--- That is, it tests for the ModeDir bit being set in m.
---@return boolean
function osFileMode:IsDir() end

--- IsRegular reports whether m describes a regular file.
--- That is, it tests that no mode type bits are set.
---@return boolean
function osFileMode:IsRegular() end

--- Perm returns the Unix permission bits in m (m & ModePerm).
---@return osFileMode
function osFileMode:Perm() end

--- Type returns type bits in m (m & ModeType).
---@return osFileMode
function osFileMode:Type() end

--- A File provides access to a single file.
--- The File interface is the minimum implementation required of the file.
--- Directory files should also implement ReadDirFile.
--- A file may implement io.ReaderAt or io.Seeker as optimizations.
---@class osFile
local osFile = {}

--- A GlobFS is a file system with a Glob method.
---@class fsGlobFS
local fsGlobFS = {}

--- A SubFS is a file system with a Sub method.
---@class fsSubFS
local fsSubFS = {}

--- WalkDirFunc is the type of the function called by WalkDir to visit
--- each file or directory.
---
--- The path argument contains the argument to WalkDir as a prefix.
--- That is, if WalkDir is called with root argument "dir" and finds a file
--- named "a" in that directory, the walk function will be called with
--- argument "dir/a".
---
--- The d argument is the fs.DirEntry for the named path.
---
--- The error result returned by the function controls how WalkDir
--- continues. If the function returns the special value SkipDir, WalkDir
--- skips the current directory (path if d.IsDir() is true, otherwise
--- path's parent directory). If the function returns the special value
--- SkipAll, WalkDir skips all remaining files and directories. Otherwise,
--- if the function returns a non-nil error, WalkDir stops entirely and
--- returns that error.
---
--- The err argument reports an error related to path, signaling that
--- WalkDir will not walk into that directory. The function can decide how
--- to handle that error; as described earlier, returning the error will
--- cause WalkDir to stop walking the entire tree.
---
--- WalkDir calls the function with a non-nil err argument in two cases.
---
--- First, if the initial fs.Stat on the root directory fails, WalkDir
--- calls the function with path set to root, d set to nil, and err set to
--- the error from fs.Stat.
---
--- Second, if a directory's ReadDir method fails, WalkDir calls the
--- function with path set to the directory's path, d set to an
--- fs.DirEntry describing the directory, and err set to the error from
--- ReadDir. In this second case, the function is called twice with the
--- path of the directory: the first call is before the directory read is
--- attempted and has err set to nil, giving the function a chance to
--- return SkipDir or SkipAll and avoid the ReadDir entirely. The second call
--- is after a failed ReadDir and reports the error from ReadDir.
--- (If ReadDir succeeds, there is no second call.)
---
--- The differences between WalkDirFunc compared to filepath.WalkFunc are:
---
---   - The second argument has type fs.DirEntry instead of fs.FileInfo.
---   - The function is called before reading a directory, to allow SkipDir
---     or SkipAll to bypass the directory read entirely or skip all remaining
---     files and directories respectively.
---   - If a directory read fails, the function is called a second time
---     for that directory to report the error.
---@class fsWalkDirFunc
local fsWalkDirFunc = {}
