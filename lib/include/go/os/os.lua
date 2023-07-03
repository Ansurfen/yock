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
---@field Stdin any
---@field Stdout any
---@field Stderr any
---@field Args any
os = {}

---{{.osReadDir}}
---@param name string
---@return any, err
function os.ReadDir(name)
end

---{{.osIsExist}}
---@param err err
---@return boolean
function os.IsExist(err)
end

---{{.osStat}}
---@param name string
---@return osFileInfo, err
function os.Stat(name)
end

---{{.osNewSyscallError}}
---@param syscall string
---@param err err
---@return err
function os.NewSyscallError(syscall, err)
end

---{{.osDirFS}}
---@param dir string
---@return any
function os.DirFS(dir)
end

---{{.osTruncate}}
---@param name string
---@param size number
---@return err
function os.Truncate(name, size)
end

---{{.osReadlink}}
---@param name string
---@return string, err
function os.Readlink(name)
end

---{{.osGetwd}}
---@return string, err
function os.Getwd()
end

---{{.osMkdirAll}}
---@param path string
---@param perm osFileMode
---@return err
function os.MkdirAll(path, perm)
end

---{{.osGetgroups}}
---@return any, err
function os.Getgroups()
end

---{{.osGeteuid}}
---@return number
function os.Geteuid()
end

---{{.osExecutable}}
---@return string, err
function os.Executable()
end

---{{.osUserConfigDir}}
---@return string, err
function os.UserConfigDir()
end

---{{.osGetppid}}
---@return number
function os.Getppid()
end

---{{.osGetpagesize}}
---@return number
function os.Getpagesize()
end

---{{.osUnsetenv}}
---@param key string
---@return err
function os.Unsetenv(key)
end

---{{.osClearenv}}
function os.Clearenv()
end

---{{.osMkdir}}
---@param name string
---@param perm osFileMode
---@return err
function os.Mkdir(name, perm)
end

---{{.osChtimes}}
---@param name string
---@param atime timeTime
---@param mtime timeTime
---@return err
function os.Chtimes(name, atime, mtime)
end

---{{.osSameFile}}
---@param fi1 osFileInfo
---@param fi2 osFileInfo
---@return boolean
function os.SameFile(fi1, fi2)
end

---{{.osEnviron}}
---@return string[]
function os.Environ()
end

---{{.osGetpid}}
---@return number
function os.Getpid()
end

---{{.osWriteFile}}
---@param name string
---@param data byte[]
---@param perm osFileMode
---@return err
function os.WriteFile(name, data, perm)
end

---{{.osUserHomeDir}}
---@return string, err
function os.UserHomeDir()
end

---{{.osExpandEnv}}
---@param s string
---@return string
function os.ExpandEnv(s)
end

---{{.osIsPermission}}
---@param err err
---@return boolean
function os.IsPermission(err)
end

---{{.osGetgid}}
---@return number
function os.Getgid()
end

---{{.osLstat}}
---@param name string
---@return osFileInfo, err
function os.Lstat(name)
end

---{{.osTempDir}}
---@return string
function os.TempDir()
end

---{{.osLchown}}
---@param name string
---@param uid number
---@param gid number
---@return err
function os.Lchown(name, uid, gid)
end

---{{.osIsPathSeparator}}
---@param c any
---@return boolean
function os.IsPathSeparator(c)
end

---{{.osRemove}}
---@param name string
---@return err
function os.Remove(name)
end

---{{.osNewFile}}
---@param fd any
---@param name string
---@return osFile
function os.NewFile(fd, name)
end

---{{.osExpand}}
---@param s string
---@param mapping function
---@return string
function os.Expand(s, mapping)
end

---{{.osRename}}
---@param oldpath string
---@param newpath string
---@return err
function os.Rename(oldpath, newpath)
end

---{{.osRemoveAll}}
---@param path string
---@return err
function os.RemoveAll(path)
end

---{{.osChmod}}
---@param name string
---@param mode osFileMode
---@return err
function os.Chmod(name, mode)
end

---{{.osUserCacheDir}}
---@return string, err
function os.UserCacheDir()
end

---{{.osSymlink}}
---@param oldname string
---@param newname string
---@return err
function os.Symlink(oldname, newname)
end

---{{.osGetuid}}
---@return number
function os.Getuid()
end

---{{.osOpen}}
---@param name string
---@return osFile, err
function os.Open(name)
end

---{{.osChdir}}
---@param dir string
---@return err
function os.Chdir(dir)
end

---{{.osExit}}
---@param code number
function os.Exit(code)
end

---{{.osMkdirTemp}}
---@param dir string
---@param pattern string
---@return string, err
function os.MkdirTemp(dir, pattern)
end

---{{.osLookupEnv}}
---@param key string
---@return string, boolean
function os.LookupEnv(key)
end

---{{.osPipe}}
---@return osFile, osFile, err
function os.Pipe()
end

---{{.osStartProcess}}
---@param name string
---@param argv string[]
---@param attr osProcAttr
---@return osProcess, err
function os.StartProcess(name, argv, attr)
end

---{{.osIsTimeout}}
---@param err err
---@return boolean
function os.IsTimeout(err)
end

---{{.osFindProcess}}
---@param pid number
---@return osProcess, err
function os.FindProcess(pid)
end

---{{.osIsNotExist}}
---@param err err
---@return boolean
function os.IsNotExist(err)
end

---{{.osCreate}}
---@param name string
---@return osFile, err
function os.Create(name)
end

---{{.osChown}}
---@param name string
---@param uid number
---@param gid number
---@return err
function os.Chown(name, uid, gid)
end

---{{.osLink}}
---@param oldname string
---@param newname string
---@return err
function os.Link(oldname, newname)
end

---{{.osSetenv}}
---@param key string
---@param value string
---@return err
function os.Setenv(key, value)
end

---{{.osGetenv}}
---@param key string
---@return string
function os.Getenv(key)
end

---{{.osGetegid}}
---@return number
function os.Getegid()
end

---{{.osHostname}}
---@return string, err
function os.Hostname()
end

---{{.osCreateTemp}}
---@param dir string
---@param pattern string
---@return osFile, err
function os.CreateTemp(dir, pattern)
end

---{{.osOpenFile}}
---@param name string
---@param flag number
---@param perm osFileMode
---@return osFile, err
function os.OpenFile(name, flag, perm)
end

---{{.osReadFile}}
---@param name string
---@return byte[], err
function os.ReadFile(name)
end

---@class any
local any = {}

---@class any
local any = {}

---@class osLinkError
---@field Op string
---@field Old string
---@field New string
---@field Err err
local osLinkError = {}

---{{.osLinkErrorError}}
---@return string
function osLinkError:Error()
end

---{{.osLinkErrorUnwrap}}
---@return err
function osLinkError:Unwrap()
end

---@class osFileInfo
local osFileInfo = {}

---@class osFile
local osFile = {}

---@class osDirEntry
local osDirEntry = {}

---@class osPathError
local osPathError = {}

---@class osSyscallError
---@field Syscall string
---@field Err err
local osSyscallError = {}

---{{.osSyscallErrorError}}
---@return string
function osSyscallError:Error()
end

---{{.osSyscallErrorUnwrap}}
---@return err
function osSyscallError:Unwrap()
end

---{{.osSyscallErrorTimeout}}
---@return boolean
function osSyscallError:Timeout()
end

---@class osProcAttr
---@field Dir string
---@field Env any
---@field Files any
---@field Sys any
local osProcAttr = {}

---@class osSignal
local osSignal = {}

---@class osFileMode
local osFileMode = {}
