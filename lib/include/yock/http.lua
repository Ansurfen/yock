--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@meta _

httplib = {}

---@param opt table
function httplib.Client(opt) end

---@param pattern string
---@param handle function
function httplib.GET(pattern, handle) end

---@param port integer
function httplib.run(port) end

formdata = {}

---@param v table<string, string[]>
---@return string
function formdata.encode(v) end

---@param v string
---@return urlValues
function formdata.decode(v) end
