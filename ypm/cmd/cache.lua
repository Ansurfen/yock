--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: cast-local-type
---@diagnostic disable: undefined-field

return {
    desc = { use = "cache", short = "Manage local ypm' cache" },
    sub = {
        {
            desc = { use = "free" },
            run = function(cmd, args)
                local freeParameter = env.params["/ypm/cache/free"]
                local force = false
                local level = 2
                if type(freeParameter) == "table" then
                    force = freeParameter["f"]:Var()
                    ---@diagnostic disable-next-line: param-type-mismatch
                    level, err = strconv.Atoi(freeParameter["l"]:Var())
                    if err ~= nil then
                        level = 2
                    end
                end
                if force then
                    cachetable:free(level)
                else
                    cachetable:free(level, true)
                end
            end,
            flags = {
                {
                    type = flag_type.str,
                    name = "level",
                    shorthand = "l",
                    default = "",
                    usage = ""
                },
                {
                    type = flag_type.bool,
                    name = "force",
                    shorthand = "f",
                    default = false,
                    usage = ""
                }
            }
        },
        {
            desc = { use = "ls" },
            run = function(cmd, args)
                if #args == 0 then
                    local rows = {}
                    for name, cacheIndex in pairs(cachetable.index.buf) do
                        table.insert(rows, {
                            name,
                            strconv.Itoa(cacheIndex.level),
                            cacheIndex.dir,
                            strconv.Itoa(cacheIndex.expire / time.Second) .. "s",
                            time.Unix(cacheIndex.updateAt, 0):Format("2006-01-02 15:04:05"),
                        })
                    end
                    printf({ "Name", "Level", "Dir", "Expire", "UpdateAt" }, rows)
                else
                    local name = args[1]
                    local cache = cachetable:get(name, "")
                    if cache ~= nil then
                        for key, value in pairs(cache.jf.buf) do
                            print(key, value)
                        end
                    end
                end
            end,
            flags = {
                {
                    type = flag_type.str,
                    name = "lock",
                    shorthand = "l",
                    default = "",
                    usage = ""
                },
            }
        },
        {
            desc = { use = "rm" },
            run = function(cmd, args)
                if #args < 2 then
                    yassert("arguments too little")
                end
                local name = args[1]
                local k = args[2]
                local cache = cachetable:get(name, "")
                if cache ~= nil then
                    for key, value in pairs(cache.jf.buf) do
                        if strings.HasPrefix(key, k) then
                            rm(pathf(cache.dir, value))
                            cache.jf:rawset(key, nil)
                        end
                    end
                    cache:save()
                end
            end
        }
    }
}
