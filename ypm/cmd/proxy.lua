-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

return {
    desc = { use = "proxy" },
    run = function(cmd, args)
        local rows = {}
        path.walk(pathf("#1", "../../proxy"), function(p, info, err)
            yassert(err)
            if path.ext(p) == ".lua" then
                local proxy = import(string.sub(p, 1, #p - 4))
                if type(proxy) == "table" and proxy["name"] ~= nil then
                    table.insert(rows, { proxy["name"], proxy["url"] })
                end
            end
            return true
        end)
        printf({ "Name", "URL" }, rows)
    end,
    sub = {
        {
            desc = { use = "new" },
            run = function(cmd, args)
                if #args == 0 then
                    yassert("arguments too little")
                end
                local newParameter = env.params["/ypm/proxy/new"]
                local author = newParameter["a"]:Var()
                local license = newParameter["l"]:Var()

                if #author == 0 then
                    author = "[AUTHOR]"
                end
                if #license == 0 then
                    license = "[LICENSE]"
                end

                local name = args[1]
                local stream, err = cat(pathf("#1", "../../template/proxy.tpl"))
                yassert(err)
                write(name .. ".lua", strf(stream, {
                    Name = name,
                    Author = author,
                    License = license,
                }))
            end,
            flags = {
                {
                    type = flag_type.str,
                    name = "author",
                    shorthand = "a",
                    default = "",
                    usage = ""
                },
                {
                    type = flag_type.str,
                    name = "license",
                    shorthand = "l",
                    default = "",
                    usage = ""
                }
            }
        },
        {
            desc = { use = "get" },
            run = function(cmd, args)
                if #args == 0 then
                    yassert("arguments too little")
                end
                local url = args[1]
                -- fetch.file()
                print(url)
            end
        },
        {
            desc = { use = "add" },
            run = function(cmd, args)
                if #args == 0 then
                    yassert("arguments too little")
                end

                local name = args[1]
                write(pathf("#1", "../../proxy", filepath.Base(name)), cat(name))
            end,
        },
        {
            desc = { use = "del" },
            run = function(cmd, args)
                if #args == 0 then
                    yassert("arguments too little")
                end
                local name = args[1]
                rm(pathf("#1", "../../proxy", name .. ".lua"))
            end
        }
    }
}
