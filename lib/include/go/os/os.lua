-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class os
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
---@field Stdin any
---@field Stdout any
---@field Stderr any
---@field Args any
os = {}

--- Link creates newname as a hard link to the oldname file.
--- If there is an error, it will be of type *LinkError.
---@param oldname string
---@param newname string
---@return err
function os.Link(oldname, newname) end

--- ReadDir reads the named directory,
--- returning all its directory entries sorted by filename.
--- If an error occurs reading the directory,
--- ReadDir returns the entries it was able to read before the error,
--- along with the error.
---@param name string
---@return any, err
function os.ReadDir(name) end

--- Expand replaces ${var} or $var in the string based on the mapping function.
--- For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).
---@param s string
---@param mapping function
---@return string
function os.Expand(s, mapping) end

--- UserCacheDir returns the default root directory to use for user-specific
--- cached data. Users should create their own application-specific subdirectory
--- within this one and use that.
---
--- On Unix systems, it returns $XDG_CACHE_HOME as specified by
--- https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
--- non-empty, else $HOME/.cache.
--- On Darwin, it returns $HOME/Library/Caches.
--- On Windows, it returns %LocalAppData%.
--- On Plan 9, it returns $home/lib/cache.
---
--- If the location cannot be determined (for example, $HOME is not defined),
--- then it will return an error.
---@return string, err
function os.UserCacheDir() end

--- Readlink returns the destination of the named symbolic link.
--- If there is an error, it will be of type *PathError.
---@param name string
---@return string, err
function os.Readlink(name) end

--- Symlink creates newname as a symbolic link to oldname.
--- On Windows, a symlink to a non-existent oldname creates a file symlink;
--- if oldname is later created as a directory the symlink will not work.
--- If there is an error, it will be of type *LinkError.
---@param oldname string
---@param newname string
---@return err
function os.Symlink(oldname, newname) end

--- Getppid returns the process id of the caller's parent.
---@return number
function os.Getppid() end

--- Executable returns the path name for the executable that started
--- the current process. There is no guarantee that the path is still
--- pointing to the correct executable. If a symlink was used to start
--- the process, depending on the operating system, the result might
--- be the symlink or the path it pointed to. If a stable result is
--- needed, path/filepath.EvalSymlinks might help.
---
--- Executable returns an absolute path unless an error occurred.
---
--- The main use case is finding resources located relative to an
--- executable.
---@return string, err
function os.Executable() end

--- WriteFile writes data to the named file, creating it if necessary.
--- If the file does not exist, WriteFile creates it with permissions perm (before umask);
--- otherwise WriteFile truncates it before writing, without changing permissions.
--- Since Writefile requires multiple system calls to complete, a failure mid-operation
--- can leave the file in a partially written state.
---@param name string
---@param data byte[]
---@param perm osFileMode
---@return err
function os.WriteFile(name, data, perm) end

--- IsNotExist returns a boolean indicating whether the error is known to
--- report that a file or directory does not exist. It is satisfied by
--- ErrNotExist as well as some syscall errors.
---
--- This function predates errors.Is. It only supports errors returned by
--- the os package. New code should use errors.Is(err, fs.ErrNotExist).
---@param err err
---@return boolean
function os.IsNotExist(err) end

--- IsPathSeparator reports whether c is a directory separator character.
---@param c any
---@return boolean
function os.IsPathSeparator(c) end

--- Getgroups returns a list of the numeric ids of groups that the caller belongs to.
---
--- On Windows, it returns syscall.EWINDOWS. See the os/user package
--- for a possible alternative.
---@return any, err
function os.Getgroups() end

--- SameFile reports whether fi1 and fi2 describe the same file.
--- For example, on Unix this means that the device and inode fields
--- of the two underlying structures are identical; on other systems
--- the decision may be based on the path names.
--- SameFile only applies to results returned by this package's Stat.
--- It returns false in other cases.
---@param fi1 osFileInfo
---@param fi2 osFileInfo
---@return boolean
function os.SameFile(fi1, fi2) end

--- Getpid returns the process id of the caller.
---@return number
function os.Getpid() end

--- Getegid returns the numeric effective group id of the caller.
---
--- On Windows, it returns -1.
---@return number
function os.Getegid() end

--- Lstat returns a FileInfo describing the named file.
--- If the file is a symbolic link, the returned FileInfo
--- describes the symbolic link. Lstat makes no attempt to follow the link.
--- If there is an error, it will be of type *PathError.
---@param name string
---@return osFileInfo, err
function os.Lstat(name) end

--- FindProcess looks for a running process by its pid.
---
--- The Process it returns can be used to obtain information
--- about the underlying operating system process.
---
--- On Unix systems, FindProcess always succeeds and returns a Process
--- for the given pid, regardless of whether the process exists.
---@param pid number
---@return osProcess, err
function os.FindProcess(pid) end

--- StartProcess starts a new process with the program, arguments and attributes
--- specified by name, argv and attr. The argv slice will become os.Args in the
--- new process, so it normally starts with the program name.
---
--- If the calling goroutine has locked the operating system thread
--- with runtime.LockOSThread and modified any inheritable OS-level
--- thread state (for example, Linux or Plan 9 name spaces), the new
--- process will inherit the caller's thread state.
---
--- StartProcess is a low-level interface. The os/exec package provides
--- higher-level interfaces.
---
--- If there is an error, it will be of type *PathError.
---@param name string
---@param argv string[]
---@param attr osProcAttr
---@return osProcess, err
function os.StartProcess(name, argv, attr) end

--- NewSyscallError returns, as an error, a new SyscallError
--- with the given system call name and error details.
--- As a convenience, if err is nil, NewSyscallError returns nil.
---@param syscall string
---@param err err
---@return err
function os.NewSyscallError(syscall, err) end

--- IsPermission returns a boolean indicating whether the error is known to
--- report that permission is denied. It is satisfied by ErrPermission as well
--- as some syscall errors.
---
--- This function predates errors.Is. It only supports errors returned by
--- the os package. New code should use errors.Is(err, fs.ErrPermission).
---@param err err
---@return boolean
function os.IsPermission(err) end

--- Pipe returns a connected pair of Files; reads from r return bytes written to w.
--- It returns the files and an error, if any.
---@return osFile, osFile, err
function os.Pipe() end

--- Lchown changes the numeric uid and gid of the named file.
--- If the file is a symbolic link, it changes the uid and gid of the link itself.
--- If there is an error, it will be of type *PathError.
---
--- On Windows, it always returns the syscall.EWINDOWS error, wrapped
--- in *PathError.
---@param name string
---@param uid number
---@param gid number
---@return err
function os.Lchown(name, uid, gid) end

--- Getgid returns the numeric group id of the caller.
---
--- On Windows, it returns -1.
---@return number
function os.Getgid() end

--- Stat returns a FileInfo describing the named file.
--- If there is an error, it will be of type *PathError.
---@param name string
---@return osFileInfo, err
function os.Stat(name) end

--- IsTimeout returns a boolean indicating whether the error is known
--- to report that a timeout occurred.
---
--- This function predates errors.Is, and the notion of whether an
--- error indicates a timeout can be ambiguous. For example, the Unix
--- error EWOULDBLOCK sometimes indicates a timeout and sometimes does not.
--- New code should use errors.Is with a value appropriate to the call
--- returning the error, such as os.ErrDeadlineExceeded.
---@param err err
---@return boolean
function os.IsTimeout(err) end

--- Getpagesize returns the underlying system's memory page size.
---@return number
function os.Getpagesize() end

--- Geteuid returns the numeric effective user id of the caller.
---
--- On Windows, it returns -1.
---@return number
function os.Geteuid() end

--- Getuid returns the numeric user id of the caller.
---
--- On Windows, it returns -1.
---@return number
function os.Getuid() end

--- Chtimes changes the access and modification times of the named
--- file, similar to the Unix utime() or utimes() functions.
---
--- The underlying filesystem may truncate or round the values to a
--- less precise time unit.
--- If there is an error, it will be of type *PathError.
---@param name string
---@param atime timeTime
---@param mtime timeTime
---@return err
function os.Chtimes(name, atime, mtime) end

--- DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory dir.
---
--- Note that DirFS("/prefix") only guarantees that the Open calls it makes to the
--- operating system will begin with "/prefix": DirFS("/prefix").Open("file") is the
--- same as os.Open("/prefix/file"). So if /prefix/file is a symbolic link pointing outside
--- the /prefix tree, then using DirFS does not stop the access any more than using
--- os.Open does. Additionally, the root of the fs.FS returned for a relative path,
--- DirFS("prefix"), will be affected by later calls to Chdir. DirFS is therefore not
--- a general substitute for a chroot-style security mechanism when the directory tree
--- contains arbitrary content.
---
--- The directory dir must not be "".
---
--- The result implements fs.StatFS.
---@param dir string
---@return any
function os.DirFS(dir) end

--- ReadFile reads the named file and returns the contents.
--- A successful call returns err == nil, not err == EOF.
--- Because ReadFile reads the whole file, it does not treat an EOF from Read
--- as an error to be reported.
---@param name string
---@return byte[], err
function os.ReadFile(name) end

--- UserConfigDir returns the default root directory to use for user-specific
--- configuration data. Users should create their own application-specific
--- subdirectory within this one and use that.
---
--- On Unix systems, it returns $XDG_CONFIG_HOME as specified by
--- https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
--- non-empty, else $HOME/.config.
--- On Darwin, it returns $HOME/Library/Application Support.
--- On Windows, it returns %AppData%.
--- On Plan 9, it returns $home/lib.
---
--- If the location cannot be determined (for example, $HOME is not defined),
--- then it will return an error.
---@return string, err
function os.UserConfigDir() end

--- SyscallError records an error from a specific system call.
---@class osSyscallError
---@field Syscall string
---@field Err err
local osSyscallError = {}


---@return string
function osSyscallError:Error() end


---@return err
function osSyscallError:Unwrap() end

--- Timeout reports whether this error represents a timeout.
---@return boolean
function osSyscallError:Timeout() end

--- A FileInfo describes a file and is returned by Stat and Lstat.
---@class osFileInfo
local osFileInfo = {}

--- A FileMode represents a file's mode and permission bits.
--- The bits have the same definition on all systems, so that
--- information about files can be moved from one system
--- to another portably. Not all bits apply to all systems.
--- The only required bit is ModeDir for directories.
---@class osFileMode
local osFileMode = {}

--- A DirEntry is an entry read from a directory
--- (using the ReadDir function or a File's ReadDir method).
---@class osDirEntry
local osDirEntry = {}

--- PathError records an error and the operation and file path that caused it.
---@class osPathError
local osPathError = {}

--- ProcAttr holds the attributes that will be applied to a new process
--- started by StartProcess.
---@class osProcAttr
---@field Dir string
---@field Env any
---@field Files any
---@field Sys any
local osProcAttr = {}

--- A Signal represents an operating system signal.
--- The usual underlying implementation is operating system-dependent:
--- on Unix it is syscall.Signal.
---@class osSignal
local osSignal = {}

--- LinkError records an error during a link or symlink or rename
--- system call and the paths that caused it.
---@class osLinkError
---@field Op string
---@field Old string
---@field New string
---@field Err err
local osLinkError = {}


---@return string
function osLinkError:Error() end


---@return err
function osLinkError:Unwrap() end

--- File represents an open file descriptor.
---@class osFile
local osFile = {}
