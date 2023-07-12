--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@param opt promptopt
function prompt(opt)
    if env.params == nil then
        env.params = {}
    end
    local get_first_word = function(str)
        local _, index = string.find(str, " ")
        if index == nil then
            return str
        else
            return string.sub(str, 1, index - 1)
        end
    end

    local builder_command = function(cmd, root, path)
        local c
        path = path .. "/"
        if root ~= nil then
            c = root
        else
            c = new_command()
        end
        for name, v in pairs(cmd) do
            if name == "desc" then
                c.Use = assign.string(c.Use, v["use"])
                c.Short = assign.string(c.Short, v["short"])
                c.Long = assign.string(c.Long, v["long"])
                path = path .. get_first_word(c.Use)
            elseif name == "sub" then
                for _, vv in ipairs(v) do
                    local cc = new_command()
                    c:AddCommand(cc)
                    ---@diagnostic disable-next-line: undefined-global
                    builder_command(vv, cc, path)
                end
            elseif name == "flags" then
                for _, flag in ipairs(v) do
                    if flag ~= nil and flag.type ~= nil then
                        if env.params[path] == nil then
                            env.params[path] = {}
                        end
                        if flag.type == flag_type.str then
                            env.params[path][flag.shorthand] = String("")
                            c:PersistentFlags():StringVarP(env.params[path][flag.shorthand]:Ptr(), flag.name,
                                flag.shorthand,
                                flag.default, flag.usage)
                        elseif flag.type == flag_type.bool then
                            env.params[path][flag.shorthand] = Boolean(false)
                            c:PersistentFlags():BoolVarP(env.params[path][flag.shorthand]:Ptr(), flag.name,
                                flag.shorthand,
                                flag.default, flag.usage)
                        elseif flag.type == flag_type.arr then
                            env.params[path][flag.shorthand] = StringArray()
                            c:PersistentFlags():StringArrayVarP(env.params[path][flag.shorthand]:Ptr(), flag.name, flag
                                .shorthand,
                                flag.default,
                                flag.usage)
                        end
                    end
                end
            elseif name == "run" then
                c.Run = v
            end
        end
        return c
    end

    local c = builder_command(opt, nil, "")
    local args = {}
    for i = 2, #env.args do
        args[i - 1] = env.args[i]
    end
    env.set_args(args)

    yassert(c:Execute())
end
