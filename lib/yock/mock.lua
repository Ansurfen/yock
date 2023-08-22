-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field

local mock_server = {}

function mock_server:run(port)
    self.engine:Run(string.format(":%d", port))
end

function mock_server:get(pattern, handle)
    self.engine:GET(pattern, handle)
end

function mock_server:post(pattern, handle)
    self.engine:POST(pattern, handle)
end

mock = {}

---@return mock_server
function mock.new()
    gin.SetMode(gin.ReleaseMode)
    local obj = {
        engine = gin.Default()
    }
    setmetatable(obj, { __index = mock_server })
    return obj
end

---@param method string
---@param url string
---@param body string
---@return httpRequest
function mock.request(method, url, body)
    ---@diagnostic disable-next-line: param-type-mismatch
    return http.NewRequest(method, url, strings.NewReader(body))
end
