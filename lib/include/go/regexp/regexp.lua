-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

regexp = {}

---@param str string
---@return reg
function regexp.MustCompile(str)
    return {}
end

---@param str string
---@return reg, err
function regexp.Compile(str)
end

---@class reg
local reg = {}

---@param s string
---@return userdata
function reg:FindStringSubmatch(s)
end

---@param s string
---@param n number
---@return string[][]
function reg:FindAllStringSubmatch(s, n)
end

---@return string
function reg:String()
end

---@return reg
function reg:Copy()
end

function reg:Longest()
end

---@param src string
---@param repl string
---@return string
function reg:ReplaceAllString(src, repl)
end

---@param src string
---@param repl string
---@return string
function reg:ReplaceAllLiteralString(src, repl)
end

---@param src string
---@param repl fun(s:string):string
---@return string
function reg:ReplaceAllStringFunc(src, repl)
end

---@param src string
---@param repl string
---@return string
function reg:ReplaceAll(src, repl)
end

---@param src string
---@param repl string
---@return string
function reg:ReplaceAllLiteral(src, repl)
end

---@param src string
---@param repl fun(s:string):string
---@return string
function reg:ReplaceAllFunc(src, repl)
end

---@param b string
---@return boolean
function reg:Match(b)
end

---@param r userdata
---@return boolean
function reg:MatchReader(r)
end

---@param s string
---@return boolean
function reg:MatchString(s)
end
