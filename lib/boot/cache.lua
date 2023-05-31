--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field

local yock_cache = path.join(env.yock_tmp, "cache.json")

mkdir(env.yock_tmp)
safe_write(yock_cache, "{}")

ycache = jsonfile:open(yock_cache)

-- put caches key and value in cache.json
---@param k string
---@param v string
function ycache:put(k, v)
    ycache.buf[k] = v
    ycache:write()
end

-- get returns key's value and returns nil when key isn't exist.
---@param k string
---@return string|nil
function ycache:get(k)
    return ycache.buf[k]
end

-- free all cache in cache.json
function ycache:free()
    ycache.buf = nil
    ycache:write()
end
