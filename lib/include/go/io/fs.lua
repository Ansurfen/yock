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

---{{.fsSub}}
---@param fsys fsFS
---@param dir string
---@return fsFS, err
function fs.Sub(fsys, dir)
end

---{{.fsWalkDir}}
---@param fsys fsFS
---@param root string
---@param fn fsWalkDirFunc
---@return err
function fs.WalkDir(fsys, root, fn)
end

---{{.fsValidPath}}
---@param name string
---@return boolean
function fs.ValidPath(name)
end

---{{.fsGlob}}
---@param fsys fsFS
---@param pattern string
---@return string[], err
function fs.Glob(fsys, pattern)
end

---{{.fsReadDir}}
---@param fsys fsFS
---@param name string
---@return any, err
function fs.ReadDir(fsys, name)
end

---{{.fsFileInfoToDirEntry}}
---@param info fsFileInfo
---@return fsDirEntry
function fs.FileInfoToDirEntry(info)
end

---{{.fsReadFile}}
---@param fsys fsFS
---@param name string
---@return byte[], err
function fs.ReadFile(fsys, name)
end

---{{.fsStat}}
---@param fsys fsFS
---@param name string
---@return fsFileInfo, err
function fs.Stat(fsys, name)
end

---@class fsFile
local fsFile = {}

---@class fsFileMode
local fsFileMode = {}

---{{.fsFileModePerm}}
---@return fsFileMode
function fsFileMode:Perm()
end

---{{.fsFileModeType}}
---@return fsFileMode
function fsFileMode:Type()
end

---{{.fsFileModeString}}
---@return string
function fsFileMode:String()
end

---{{.fsFileModeIsDir}}
---@return boolean
function fsFileMode:IsDir()
end

---{{.fsFileModeIsRegular}}
---@return boolean
function fsFileMode:IsRegular()
end

---@class fsPathError
---@field Op string
---@field Path string
---@field Err err
local fsPathError = {}

---{{.fsPathErrorTimeout}}
---@return boolean
function fsPathError:Timeout()
end

---{{.fsPathErrorError}}
---@return string
function fsPathError:Error()
end

---{{.fsPathErrorUnwrap}}
---@return err
function fsPathError:Unwrap()
end

---@class fsStatFS
local fsStatFS = {}

---@class fsDirEntry
local fsDirEntry = {}

---@class fsReadDirFile
local fsReadDirFile = {}

---@class fsFileInfo
local fsFileInfo = {}

---@class fsFS
local fsFS = {}

---@class fsGlobFS
local fsGlobFS = {}

---@class fsReadDirFS
local fsReadDirFS = {}

---@class fsReadFileFS
local fsReadFileFS = {}

---@class fsSubFS
local fsSubFS = {}

---@class fsWalkDirFunc
local fsWalkDirFunc = {}
