--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field
---@diagnostic disable: lowercase-global
httplib = {}

---@param opt table
function httplib.Client(opt)
    local c = http.Client()
    c.Timeout = (opt["timeout"] or time.Second * 10)
    return c
end

---@param pattern string
---@param handle function
function httplib.GET(pattern, handle)
    http.HandleFunc(pattern, function(w, req)
        fmt.Fprintf(w, handle(req))
    end)
end

---@param port integer
function httplib.run(port)
    http.ListenAndServe(":" .. strconv.Itoa(port), nil)
end

formdata = {}

---@param v table<string, string[]>
---@return string
function formdata.encode(v)
    return url.Values(v):Encode()
end

---@param v string
---@return urlValues
function formdata.decode(v)
    local formData = url.ParseQuery(v)
    return formData or url.Values({})
end
