-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

fmt = {}

---@vararg any
---@return integer, err
function fmt.Print(...)
end

---@vararg any
---@return integer, err
function fmt.Println(...)
end

---@param format string
---@param ... any
---@return integer, err
function fmt.Printf(format, ...)
end

---@param w ioWriter
---@vararg any
---@return integer, err
function fmt.Fprint(w, ...)
end

---@param w ioWriter
---@param format string
---@param ... any
---@return integer, err
function fmt.Fprintf(w, format, ...)
end

---@param w ioWriter
---@vararg any
---@return integer, err
function fmt.Fprintln(w, ...)
end
