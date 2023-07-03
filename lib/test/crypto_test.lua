-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

print(crypto.md5("Hello World"))
print(crypto.sha256("Hello World"))
local key = "Yock Key "
local hash = crypto.encode_aes(key, "Hello World!")
print(hash)
print(crypto.decode_aes(key, hash))
