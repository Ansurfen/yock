-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class path
---@field Separator integer
path = {}

---@param filepath string
---@return string
function path.filename(filepath) end

--- Join joins any number of path elements into a single path,
--- separating them with slashes. Empty elements are ignored.
--- The result is Cleaned. However, if the argument list is
--- empty or all its elements are empty, Join returns
--- an empty string.
---@vararg string
---@return string
function path.join(...) end

--- Dir returns all but the last element of path, typically the path's directory.
--- After dropping the final element using Split, the path is Cleaned and trailing
--- slashes are removed.
--- If the path is empty, Dir returns ".".
--- If the path consists entirely of slashes followed by non-slash bytes, Dir
--- returns a single slash. In any other case, the returned path does not end in a
--- slash.
---@param path string
---@return string
function path.dir(path) end

--- Base returns the last element of path.
--- Trailing slashes are removed before extracting the last element.
--- If the path is empty, Base returns ".".
--- If the path consists entirely of slashes, Base returns "/".
---@param path string
---@return string
function path.base(path) end

--- Clean returns the shortest path name equivalent to path
--- by purely lexical processing. It applies the following rules
--- iteratively until no further processing can be done:
---
---  1. Replace multiple slashes with a single slash.
---  2. Eliminate each . path name element (the current directory).
---  3. Eliminate each inner .. path name element (the parent directory)
---     along with the non-.. element that precedes it.
---  4. Eliminate .. elements that begin a rooted path:
---     that is, replace "/.." by "/" at the beginning of a path.
---
--- The returned path ends in a slash only if it is the root "/".
---
--- If the result of this process is an empty string, Clean
--- returns the string ".".
---
--- See also Rob Pike, “Lexical File Names in Plan 9 or
--- Getting Dot-Dot Right,”
--- https://9p.io/sys/doc/lexnames.html
---@param path string
---@return string
function path.clean(path) end

--- Ext returns the file name extension used by path.
--- The extension is the suffix beginning at the final dot
--- in the final slash-separated element of path;
--- it is empty if there is no dot.
---@param path string
---@return string
function path.ext(path) end

---@param path string
---@return string, string
function path.abs(path) end

---@param root string
---@param fn fun(path: string, info: fileinfo, err:err): boolean
---@return err
function path.walk(root, fn) end

---@class fileinfo
---@field Name fun(): string
---@field Size fun(): number
---@field Mode fun(): userdata
---@field ModTime fun(): userdata
---@field IsDir fun(): boolean
---@field Sys fun(): userdata
local fileinfo = {}
