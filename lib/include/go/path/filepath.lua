-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class filepath
---@field Separator any
---@field ListSeparator any
---@field ErrBadPattern any
---@field SkipDir any
---@field SkipAll any
filepath = {}

---{{.filepathClean}}
---@param path string
---@return string
function filepath.Clean(path)
end

---{{.filepathSplitList}}
---@param path string
---@return string[]
function filepath.SplitList(path)
end

---{{.filepathRel}}
---@param basepath string
---@param targpath string
---@return string, err
function filepath.Rel(basepath, targpath)
end

---{{.filepathEvalSymlinks}}
---@param path string
---@return string, err
function filepath.EvalSymlinks(path)
end

---{{.filepathAbs}}
---@param path string
---@return string, err
function filepath.Abs(path)
end

---{{.filepathWalk}}
---@param root string
---@param fn filepathWalkFunc
---@return err
function filepath.Walk(root, fn)
end

---{{.filepathToSlash}}
---@param path string
---@return string
function filepath.ToSlash(path)
end

---{{.filepathFromSlash}}
---@param path string
---@return string
function filepath.FromSlash(path)
end

---{{.filepathIsLocal}}
---@param path string
---@return boolean
function filepath.IsLocal(path)
end

---{{.filepathSplit}}
---@param path string
---@return string
function filepath.Split(path)
end

---{{.filepathVolumeName}}
---@param path string
---@return string
function filepath.VolumeName(path)
end

---{{.filepathWalkDir}}
---@param root string
---@param fn fsWalkDirFunc
---@return err
function filepath.WalkDir(root, fn)
end

---{{.filepathIsAbs}}
---@param path string
---@return boolean
function filepath.IsAbs(path)
end

---{{.filepathGlob}}
---@param pattern string
---@return string[], err
function filepath.Glob(pattern)
end

---{{.filepathBase}}
---@param path string
---@return string
function filepath.Base(path)
end

---{{.filepathDir}}
---@param path string
---@return string
function filepath.Dir(path)
end

---{{.filepathExt}}
---@param path string
---@return string
function filepath.Ext(path)
end

---{{.filepathHasPrefix}}
---@param p string
---@param prefix string
---@return boolean
function filepath.HasPrefix(p, prefix)
end

---{{.filepathMatch}}
---@param pattern string
---@param name string
---@return boolean, err
function filepath.Match(pattern, name)
end

---{{.filepathJoin}}
---@vararg string
---@return string
function filepath.Join(...)
end

---@class filepathWalkFunc
local filepathWalkFunc = {}
