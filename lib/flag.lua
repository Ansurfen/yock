flag_type = {
    string_type = 0,
    number_type = 0,
    array_type = 1,
    bool_type = 2
}

function parse_flags(env, todo)
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
                if todo[jobflag] == flag_type.string_type then
                    idx = idx + 1
                    if idx > #env.args then
                        break
                    end
                    env.flags[jobflag] = env.args[idx]
                elseif todo[jobflag] == flag_type.bool_type then
                    env.flags[jobflag] = true
                elseif todo[jobflag] == flag_type.array_type then
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
            if todo[arg] == flag_type.string_type then
                idx = idx + 1
                if idx > #env.args then
                    break
                end
                env.flags[arg] = env.args[idx]
            elseif todo[arg] == flag_type.bool_type then
                env.flags[arg] = true
            elseif todo[arg] == flag_type.array_type then
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
