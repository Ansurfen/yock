-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

local data = formdata.encode({
    username = { "ansurfen" },
    password = { "root" }
})
print(data)
print(formdata.decode(data):Get("password"))
print(formdata.decode("pwd=a"):Get("password"))
print(formdata.decode(""):Get("password"))
