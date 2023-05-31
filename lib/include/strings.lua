-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class strings
strings = {}

---{{.Join}}
---@param elems table
---@param sep string
---@return table
function strings.Join(elems, sep)
    return {}
end

---{{.HasPrefix}}
---@param s string
---@param prefix string
---@return boolean
function strings.HasPrefix(s, prefix)
    return false
end

---{{.HasSuffix}}
---@param s string
---@param suffix string
---@return boolean
function strings.HasSuffix(s, suffix)
    return false
end

---@param s string
---@param sep string
---@return string, string, boolean
function strings.Cut(s, sep)
    return "", "", false
end

---@param s string
---@param prefix string
---@return string, boolean
function strings.CutPrefix(s, prefix)
    return "", false
end

---@param s string
---@param suffix string
---@return string, boolean
function strings.CutSuffix(s, suffix)
    return "", false
end

---@param s string
---@param substr string
---@return boolean
function strings.Contains(s, substr)
    return false
end

---@param s string
---@param chars string
---@return boolean
function strings.ContainsAny(s, chars)
    return false
end

---@param s string
---@param r string
---@return boolean
function strings.ContainsRune(s, r)
    return false
end

---@param s string
---@param substr string
---@return number
function strings.Count(s, substr)
    return 0
end

---@param s string
---@param old string
---@param new string
---@param n number
---@return string
function strings.Replace(s, old, new, n)
    return ""
end

---@param s string
---@param old string
---@param new string
---@return string
function strings.ReplaceAll(s, old, new)
    return ""
end

---@param s string
---@return string
function strings.Clone(s)
    return ""
end

---@param a string
---@param b string
---@return number
function strings.Compare(a, b)
    return 0
end

---@param str string
---@param sep string
---@return table
function strings.Split(str, sep)
    return {}
end
