--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global
httplib = {}

function httplib:Client(opt)
    local c = http.Client()
    c.Timeout = (opt["timeout"] or time.second * 10)
    return c
end
