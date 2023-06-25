-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

http = {}

---@return userdata
function http.Client()
end

---@param route string
---@param handle function
function http.HandleFunc(route, handle)
end

---@param addr string
---@param handle any
function http.ListenAndServe(addr, handle)
end
