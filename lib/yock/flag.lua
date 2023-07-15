--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

-- flag_type identifies the number of parameters to be extracted after matching the flag. 
-- If it is of type string or number, after matching -x, it needs to take a step back to accept the parameter.
-- If it is a bool type, it is not required. 
-- As for the array type, for repeated flags, the received parameters are not overwritten by storage but saved as an array.
flag_type = {
    str = 0,
    num = 0,
    arr = 1,
    bool = 2
}

function argsparse(env, todo)
    local idx = 0
    if env.flags == nil then
        env.flags = {}
    end
    while idx ~= #env.args do
        idx = idx + 1
        local arg = env.args[idx]
        if strings.HasPrefix(arg, "--") then
            arg = string.sub(arg, 3, #arg)
            local job, jobflag, ok = strings.Cut(arg, "-")
            if ok and job == env.job then
                if todo[jobflag] == flag_type.str then
                    idx = idx + 1
                    if idx > #env.args then
                        break
                    end
                    env.flags[jobflag] = env.args[idx]
                elseif todo[jobflag] == flag_type.bool then
                    env.flags[jobflag] = true
                elseif todo[jobflag] == flag_type.arr then
                    if env.flags[jobflag] == nil or type(env.flags[jobflag]) ~= "table" then
                        env.flags[jobflag] = {}
                    end
                    idx = idx + 1
                    if idx > #env.args then
                        break
                    end
                    local value = env.args[idx]
                    if strings.Contains(value, ",") then
                        local s = strings.Split(value, ",")
                        for _, v in ipairs(s) do
                            table.insert(env.flags[arg], v)
                        end
                    else
                        table.insert(env.flags[arg], value)
                    end
                end
            end
        elseif strings.HasPrefix(arg, "-") then
            arg = string.sub(arg, 2, #arg)
            if todo[arg] == flag_type.str then
                idx = idx + 1
                if idx > #env.args then
                    break
                end
                env.flags[arg] = env.args[idx]
            elseif todo[arg] == flag_type.bool then
                env.flags[arg] = true
            elseif todo[arg] == flag_type.arr then
                if env.flags[arg] == nil or type(env.flags[arg]) ~= "table" then
                    env.flags[arg] = {}
                end
                idx = idx + 1
                if idx > #env.args then
                    break
                end
                local value = env.args[idx]
                if strings.Contains(value, ",") then
                    local s = strings.Split(value, ",")
                    for _, v in ipairs(s) do
                        table.insert(env.flags[arg], v)
                    end
                else
                    table.insert(env.flags[arg], value)
                end
            end
        end
    end
end
