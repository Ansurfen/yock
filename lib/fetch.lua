--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field

fetch = {}

---@param url string
---@param file_type string
---@return string, err
function fetch.file(url, file_type)
    if file_type == nil then
        yassert("invalid file type")
    end
    local cache = cachetable:get("public", "")
    if cache ~= nil then
        cache = cachetable:create("public", 2, 0 * time.Second, "", 0)
    end
    local file = cache:get(url)
    if not (type(file) == "string" and #file > 0) then
        file = random.str(32) .. file_type
        local _, err = curl({
            debug = true,
            save = true,
            strict = true,
            dir = cache.dir,
            filename = function(s)
                return file
            end
        }, url)
        if err ~= nil then
            return "", err
        end
        cache:put(url, file)
        cache:save()
    end
    return pathf(cache.dir, file), nil
end

function fetch.zip(url)
    local suffix
    if env.platform.OS == "windows" then
        suffix = ".zip"
    else
        suffix = ".tar.gz"
    end
    return fetch.file(url, suffix)
end

---@param url string
---@return string
function fetch.script(url)
    return fetch.file(url, ".lua")
end
