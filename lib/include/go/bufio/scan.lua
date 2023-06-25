-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

bufio = {}

---@return bufioScanner
function bufio.NewScanner()
end

---@param data string
---@param atEOF boolean
---@return number,string,err
function bufio.ScanLines(data, atEOF)
end

---@class bufioScanner
local bufioScanner = {}

---@param split function
function bufioScanner:Split(split)
end

---@return boolean
function bufioScanner:Scan()
end

---@return string
function bufioScanner:Text()
end

---@return string
function bufioScanner:Bytes()
end

---@return err
function bufioScanner:Err()
end

---@param buf string
---@param max integer
function bufioScanner:Buffer(buf, max)
end
