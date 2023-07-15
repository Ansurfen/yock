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
---@return table<string>
function strings.Split(str, sep) end

---@param s string
---@return userdata
function strings.NewReader(s)
end

---@param s string
---@return string
function strings.TrimSpace(s)
end

---@param s string
---@param substr string
---@return integer
function strings.LastIndex(s, substr)
end

---@param s string
---@param c integer
---@return integer
function strings.IndexByte(s, c)
end

---@param s string
---@param r integer
---@return integer
function strings.IndexRune(s, r)
end

---@param s string
---@param chars string
---@return integer
function strings.IndexAny(s, chars)
end

---@param s string
---@param chars string
---@return integer
function strings.LastIndexAny(s, chars)
end

---@param s string
---@param c integer
---@return integer
function strings.LastIndexByte(s, c)
end

---@param s string
---@param sep string
---@param n integer
---@return string[]
function strings.SplitN(s, sep, n)
end

---@param s string
---@param sep string
---@param n integer
---@return string[]
function strings.SplitAfterN(s, sep, n)
end

---@param s string
---@param sep string
---@return string[]
function strings.SplitAfter(s, sep)
end

---@param s string
---@return string[]
function strings.Fields(s)
end

---@param s string
---@param count integer
---@return string
function strings.Repeat(s, count)
end

---@param s string
---@return string
function strings.ToUpper(s)
end

---@param s string
---@return string
function strings.ToLower(s)
end

---@param s string
---@return string
function strings.ToTitle(s)
end

---@param s string
---@param f fun(r: integer): boolean
---@return string[]
function strings.FieldsFunc(s, f)
end

---@param mapping fun(r: integer): integer
---@param s string
---@return string
function strings.Map(mapping, s)
end

---@param s string
---@param f fun(r: integer): boolean
---@return string
function strings.TrimLeftFunc(s, f)
end

---@param s string
---@param f fun(r: integer): boolean
---@return string
function strings.TrimRightFunc(s, f)
end

---@param s string
---@param f fun(r: integer): boolean
---@return string
function strings.TrimFunc(s, f)
end

---@param s string
---@param f fun(r: integer): boolean
---@return integer
function strings.IndexFunc(s, f)
end

---@param s string
---@param f fun(r: integer): boolean
---@return integer
function strings.LastIndexFunc(s, f)
end

---@param s string
---@param prefix string
---@return string
function strings.TrimPrefix(s, prefix)
end

---@param s string
---@param suffix string
---@return string
function strings.TrimSuffix(s, suffix)
end
