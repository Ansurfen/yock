--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field
---@diagnostic disable: duplicate-set-field

---@class ycache
---@field dir string
---@field jf jsonobject
ycache = {}

---@param index string
---@param dir string
---@return ycache
function ycache:new(index, dir)
    if not find(dir) then
        mkdir(dir)
    end
    local obj = {
        dir = dir,
        jf = json.create(index)
    }
    setmetatable(obj, { __index = self })
    return obj
end

-- get returns key's value and returns nil when key isn't exist.
---@param k string
---@return any
function ycache:get(k)
    return self.jf:rawget(k)
end

---@param k string
---@param v any
function ycache:put(k, v)
    self.jf:rawset(k, v)
end

function ycache:free()
    self.jf.buf = {}
    rm({ safe = false }, self.dir)
    self:save()
end

function ycache:save()
    self.jf:save(true)
end

cachetable = {
    index = json.create(pathf(env.yock_tmp, "index.json")),
    caches = {}
}

---@param name string
---@param level integer
---@param expire timeTime
---@param lock string
---@param attr integer
---@return ycache
function cachetable:create(name, level, expire, lock, attr)
    if self.caches[name] ~= nil then
        return self.caches[name]
    end
    local cacheIndex = self.index:get(name)
    if cacheIndex == nil then
        local now = time.Now():Unix()
        cacheIndex = {
            level = level,
            lock = lock or "",
            updateAt = now,
            expire = expire or 0,
            attr = attr or 0,
            dir = random.str(8),
        }
        self.index:set(name, cacheIndex)
        self.index:save(true)
        local cache = ycache:new(pathf(env.yock_tmp, name .. ".json"), pathf(env.yock_tmp, cacheIndex.dir))
        self.caches[name] = cache
        return cache
    else
        if cacheIndex.expire ~= 0 and
            time.Now():Compare(time.Unix(cacheIndex.updateAt, 0):Add(cacheIndex.expire)) == 1 then
            rm({ safe = false }, pathf(env.yock_tmp, cacheIndex.dir))
            cacheIndex.updateAt = time.Now():Unix()
            self.index:set(name, cacheIndex)
            self.index:save(true)
        end
        if #cacheIndex.lock ~= 0 and cacheIndex.lock ~= lock then
            yassert("lack token")
        end
        local cache = ycache:new(pathf(env.yock_tmp, name .. ".json"), pathf(env.yock_tmp, cacheIndex.dir))
        self.caches[name] = cache
        return cache
    end
end

---@param name string
---@param lock string
---@return ycache|nil
function cachetable:get(name, lock)
    if self.caches[name] ~= nil then
        return self.caches[name]
    end
    local cacheIndex = self.index:get(name)
    if cacheIndex ~= nil then
        if cacheIndex.expire ~= 0 and
            time.Now():Compare(time.Unix(cacheIndex.updateAt, 0):Add(cacheIndex.expire)) == 1 then
            rm({ safe = false }, pathf(env.yock_tmp, cacheIndex.dir))
            cacheIndex.updateAt = time.Now():Unix()
            self.index:set(name, cacheIndex)
            self.index:save(true)
        end
        if #cacheIndex.lock ~= 0 and cacheIndex.lock ~= lock then
            yassert("lack token")
        end
        local cache = ycache:new(pathf(env.yock_tmp, name .. ".json"), pathf(env.yock_tmp, cacheIndex.dir))
        self.caches[name] = cache
        return cache
    end
    return nil
end

---@param level integer
---@param expire? boolean
function cachetable:free(level, expire)
    for name, cacheIndex in pairs(self.index.buf) do
        if type(expire) == "boolean" and expire then
            if cacheIndex.expire ~= 0 and
                time.Now():Compare(time.Unix(cacheIndex.updateAt, 0):Add(cacheIndex.expire)) == 1 then
                rm({ safe = false }, pathf(env.yock_tmp, cacheIndex.dir))
                cacheIndex.updateAt = time.Now():Unix()
                self.index:set(name, cacheIndex)
                self.index:save(true)
            end
        else
            if cacheIndex.level <= level then
                local cache
                if self.caches[name] ~= nil then
                    cache = self.caches[name]
                else
                    cache = ycache:new(pathf(env.yock_tmp, name .. ".json"),
                        pathf(env.yock_tmp, cacheIndex.dir))
                    self.caches[name] = cache
                end
                cache:free()
            end
        end
    end
end

cachetable:create("public", 2, 7 * 24 * time.Hour, "", 0)
cachetable:create("protected", 2, 0 * time.Second, "", 0)
cachetable:create("private", 2, 0 * time.Second, "yock", 0)
