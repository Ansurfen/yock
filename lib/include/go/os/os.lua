-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class os
---@field O_RDONLY any
---@field O_WRONLY any
---@field O_RDWR any
---@field O_APPEND any
---@field O_CREATE any
---@field O_EXCL any
---@field O_SYNC any
---@field O_TRUNC any
---@field SEEK_SET any
---@field SEEK_CUR any
---@field SEEK_END any
---@field DevNull any
---@field DevNull any
---@field DevNull any
---@field PathSeparator any
---@field PathListSeparator any
---@field PathSeparator any
---@field PathListSeparator any
---@field PathSeparator any
---@field PathListSeparator any
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
---@field ErrNoDeadline any
---@field ErrDeadlineExceeded any
---@field ErrProcessDone any
---@field Interrupt any
---@field Kill any
---@field Interrupt any
---@field Kill any
---@field Stdin any
---@field Stdout any
---@field Stderr any
---@field Args any
os = {}

--- OpenFile is the generalized open call; most users will use Open
--- or Create instead. It opens the named file with specified flag
--- (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag
--- is passed, it is created with mode perm (before umask). If successful,
--- methods on the returned File can be used for I/O.
--- If there is an error, it will be of type *PathError.
---@param name string
---@param flag number
---@param perm osFileMode
---@return osFile, err
function os.OpenFile(name, flag, perm) end

--- Link creates newname as a hard link to the oldname file.
--- If there is an error, it will be of type *LinkError.
---@param oldname string
---@param newname string
---@return err
function os.Link(oldname, newname) end

--- Exit causes the current program to exit with the given status code.
--- Conventionally, code zero indicates success, non-zero an error.
--- The program terminates immediately; deferred functions are not run.
---
--- For portability, the status code should be in the range [0, 125].
---@param code number
function os.Exit(code) end

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

--- Truncate changes the size of the named file.
--- If the file is a symbolic link, it changes the size of the link's target.
---@param name string
---@param size number
---@return err
function os.Truncate(name, size) end

--- Chown changes the numeric uid and gid of the named file.
--- If the file is a symbolic link, it changes the uid and gid of the link's target.
--- A uid or gid of -1 means to not change that value.
--- If there is an error, it will be of type *PathError.
---
--- On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or
--- EPLAN9 error, wrapped in *PathError.
---@param name string
---@param uid number
---@param gid number
---@return err
function os.Chown(name, uid, gid) end

--- Hostname returns the host name reported by the kernel.
---@return string, err
function os.Hostname() end

--- Environ returns a copy of strings representing the environment,
--- in the form "key=value".
---@return string[]
function os.Environ() end

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

--- Remove removes the named file or directory.
--- If there is an error, it will be of type *PathError.
---@param name string
---@return err
function os.Remove(name) end

--- Readlink returns the destination of the named symbolic link.
--- If there is an error, it will be of type *PathError.
---@param name string
---@return string, err
function os.Readlink(name) end

--- RemoveAll removes path and any children it contains.
--- It removes everything it can but returns the first error
--- it encounters. If the path does not exist, RemoveAll
--- returns nil (no error).
--- If there is an error, it will be of type *PathError.
---@param path string
---@return err
function os.RemoveAll(path) end

--- Clearenv deletes all environment variables.
function os.Clearenv() end

--- Open opens the named file for reading. If successful, methods on
--- the returned file can be used for reading; the associated file
--- descriptor has mode O_RDONLY.
--- If there is an error, it will be of type *PathError.
---@param name string
---@return osFile, err
function os.Open(name) end

--- Symlink creates newname as a symbolic link to oldname.
--- On Windows, a symlink to a non-existent oldname creates a file symlink;
--- if oldname is later created as a directory the symlink will not work.
--- If there is an error, it will be of type *LinkError.
---@param oldname string
---@param newname string
---@return err
function os.Symlink(oldname, newname) end

--- ExpandEnv replaces ${var} or $var in the string according to the values
--- of the current environment variables. References to undefined
--- variables are replaced by the empty string.
---@param s string
---@return string
function os.ExpandEnv(s) end

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

--- Chmod changes the mode of the named file to mode.
--- If the file is a symbolic link, it changes the mode of the link's target.
--- If there is an error, it will be of type *PathError.
---
--- A different subset of the mode bits are used, depending on the
--- operating system.
---
--- On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and
--- ModeSticky are used.
---
--- On Windows, only the 0200 bit (owner writable) of mode is used; it
--- controls whether the file's read-only attribute is set or cleared.
--- The other bits are currently unused. For compatibility with Go 1.12
--- and earlier, use a non-zero mode. Use mode 0400 for a read-only
--- file and 0600 for a readable+writable file.
---
--- On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive,
--- and ModeTemporary are used.
---@param name string
---@param mode osFileMode
---@return err
function os.Chmod(name, mode) end

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

--- Mkdir creates a new directory with the specified name and permission
--- bits (before umask).
--- If there is an error, it will be of type *PathError.
---@param name string
---@param perm osFileMode
---@return err
function os.Mkdir(name, perm) end

--- Getpid returns the process id of the caller.
---@return number
function os.Getpid() end

--- Create creates or truncates the named file. If the file already exists,
--- it is truncated. If the file does not exist, it is created with mode 0666
--- (before umask). If successful, methods on the returned File can
--- be used for I/O; the associated file descriptor has mode O_RDWR.
--- If there is an error, it will be of type *PathError.
---@param name string
---@return osFile, err
function os.Create(name) end

--- MkdirTemp creates a new temporary directory in the directory dir
--- and returns the pathname of the new directory.
--- The new directory's name is generated by adding a random string to the end of pattern.
--- If pattern includes a "*", the random string replaces the last "*" instead.
--- If dir is the empty string, MkdirTemp uses the default directory for temporary files, as returned by TempDir.
--- Multiple programs or goroutines calling MkdirTemp simultaneously will not choose the same directory.
--- It is the caller's responsibility to remove the directory when it is no longer needed.
---@param dir string
---@param pattern string
---@return string, err
function os.MkdirTemp(dir, pattern) end

--- Setenv sets the value of the environment variable named by the key.
--- It returns an error, if any.
---@param key string
---@param value string
---@return err
function os.Setenv(key, value) end

--- Getwd returns a rooted path name corresponding to the
--- current directory. If the current directory can be
--- reached via multiple paths (due to symbolic links),
--- Getwd may return any one of them.
---@return string, err
function os.Getwd() end

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

--- CreateTemp creates a new temporary file in the directory dir,
--- opens the file for reading and writing, and returns the resulting file.
--- The filename is generated by taking pattern and adding a random string to the end.
--- If pattern includes a "*", the random string replaces the last "*".
--- If dir is the empty string, CreateTemp uses the default directory for temporary files, as returned by TempDir.
--- Multiple programs or goroutines calling CreateTemp simultaneously will not choose the same file.
--- The caller can use the file's Name method to find the pathname of the file.
--- It is the caller's responsibility to remove the file when it is no longer needed.
---@param dir string
---@param pattern string
---@return osFile, err
function os.CreateTemp(dir, pattern) end

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

--- NewFile returns a new File with the given file descriptor and
--- name. The returned value will be nil if fd is not a valid file
--- descriptor.
---@param fd any
---@param name string
---@return osFile
function os.NewFile(fd, name) end

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

--- Chdir changes the current working directory to the named directory.
--- If there is an error, it will be of type *PathError.
---@param dir string
---@return err
function os.Chdir(dir) end

--- TempDir returns the default directory to use for temporary files.
---
--- On Unix systems, it returns $TMPDIR if non-empty, else /tmp.
--- On Windows, it uses GetTempPath, returning the first non-empty
--- value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory.
--- On Plan 9, it returns /tmp.
---
--- The directory is neither guaranteed to exist nor have accessible
--- permissions.
---@return string
function os.TempDir() end

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

--- IsExist returns a boolean indicating whether the error is known to report
--- that a file or directory already exists. It is satisfied by ErrExist as
--- well as some syscall errors.
---
--- This function predates errors.Is. It only supports errors returned by
--- the os package. New code should use errors.Is(err, fs.ErrExist).
---@param err err
---@return boolean
function os.IsExist(err) end

--- LookupEnv retrieves the value of the environment variable named
--- by the key. If the variable is present in the environment the
--- value (which may be empty) is returned and the boolean is true.
--- Otherwise the returned value will be empty and the boolean will
--- be false.
---@param key string
---@return string, boolean
function os.LookupEnv(key) end

--- Unsetenv unsets a single environment variable.
---@param key string
---@return err
function os.Unsetenv(key) end

--- ReadFile reads the named file and returns the contents.
--- A successful call returns err == nil, not err == EOF.
--- Because ReadFile reads the whole file, it does not treat an EOF from Read
--- as an error to be reported.
---@param name string
---@return byte[], err
function os.ReadFile(name) end

--- Getenv retrieves the value of the environment variable named by the key.
--- It returns the value, which will be empty if the variable is not present.
--- To distinguish between an empty value and an unset value, use LookupEnv.
---@param key string
---@return string
function os.Getenv(key) end

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

--- Rename renames (moves) oldpath to newpath.
--- If newpath already exists and is not a directory, Rename replaces it.
--- OS-specific restrictions may apply when oldpath and newpath are in different directories.
--- Even within the same directory, on non-Unix platforms Rename is not an atomic operation.
--- If there is an error, it will be of type *LinkError.
---@param oldpath string
---@param newpath string
---@return err
function os.Rename(oldpath, newpath) end

--- MkdirAll creates a directory named path,
--- along with any necessary parents, and returns nil,
--- or else returns an error.
--- The permission bits perm (before umask) are used for all
--- directories that MkdirAll creates.
--- If path is already a directory, MkdirAll does nothing
--- and returns nil.
---@param path string
---@param perm osFileMode
---@return err
function os.MkdirAll(path, perm) end

--- UserHomeDir returns the current user's home directory.
---
--- On Unix, including macOS, it returns the $HOME environment variable.
--- On Windows, it returns %USERPROFILE%.
--- On Plan 9, it returns the $home environment variable.
---@return string, err
function os.UserHomeDir() end

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


---@class any
local any = {}

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


---@class any
local any = {}

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
