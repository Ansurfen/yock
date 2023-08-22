-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field

---@meta _

---@class integer: number

---@class mock
---@field new fun(): mock_server # returns mock server object
---@field request fun(method: string, url: string, body: string): httpRequest # returns a new request
mock = {}

---mock_context is the most important part of mock server.
---It allows us to pass variables between middleware, manage the flow,
---validate the JSON of a request and render a JSON response for example.
---@class mock_context
---@field Writer httpResponseWriter # an abstract to response will be sent to client
---@field Request httpRequest # saves request information from client
---@field Param fun(self: mock_context, key: string): string # returns the value of the URL param.
---@field Query fun(self: mock_context, key: string): string # returns the keyed url query value if it exists, otherwise it returns an empty string
---@field PostForm fun(self: mock_context, key: string): string # returns the specified key from a POST urlencoded form or multipart form
---@field Bind fun(self: mock_context, obj: table): err # checks the Method and Content-Type to select a binding engine automatically,
---@field String fun(self: mock_context, code: integer, format: string, ...) # writes the given string into the response body.
---@field JSON fun(self: mock_context, code: integer, obj: table) # serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
---@field File fun(self: mock_context, code: integer, dir: string) # writes the specified file into the body stream in an efficient way.

---mock_handle is the callback function of mock_server, will be
---called when pattern matches with request's route. when called,
---mock server make and wrap request into mock_context for handling
---business with convenient.
---@alias mock_handle fun(ctx: mock_context)

---@class mock_server: ginEngine
---@field engine ginEngine
---@field run fun(self: mock_server, port: integer): err # runs service on specified port
---@field get fun(self: mock_server, pattern: string, handle: mock_handle) # binds pattern and handle, and will be called when the http GET request arrived by client
---@field post fun(self: mock_server, pattern: string, handle: mock_handle) # binds pattern and handle, and will be called when the http POST request arrived by client

---proxy sends request and writes to writer
---```lua
---#assume you use mock and try to proxy request handle.
---#`NOTE`: it'll copy and return automatically when proxy finished.
---local s = mock.new()
---s:get("/", function(ctx)
---     proxy(ctx.Writer, mock.new("GET", "http://localhost:8080", ""))
---end)
---```
---@param writer httpResponseWriter
---@param request httpRequest
function proxy(writer, request) end
