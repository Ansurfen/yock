-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class ioutil
---@field Discard any
ioutil = {}

---{{.ioutilTempFile}}
---@param dir string
---@param pattern string
---@return any, err
function ioutil.TempFile(dir, pattern)
end

---{{.ioutilTempDir}}
---@param dir string
---@param pattern string
---@return string, err
function ioutil.TempDir(dir, pattern)
end

---{{.ioutilReadAll}}
---@param r ioReader
---@return byte[], err
function ioutil.ReadAll(r)
end

---{{.ioutilReadFile}}
---@param filename string
---@return byte[], err
function ioutil.ReadFile(filename)
end

---{{.ioutilWriteFile}}
---@param filename string
---@param data byte[]
---@param perm fsFileMode
---@return err
function ioutil.WriteFile(filename, data, perm)
end

---{{.ioutilReadDir}}
---@param dirname string
---@return any, err
function ioutil.ReadDir(dirname)
end

---{{.ioutilNopCloser}}
---@param r ioReader
---@return any
function ioutil.NopCloser(r)
end
