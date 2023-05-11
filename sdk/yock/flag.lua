flag_type = {
    string_type = 0,
    number_type = 1,
    array_type = 2,
    bool_type = 3
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
                -- ! ip
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
                    table.insert(env.flags[jobflag], env.args[idx])
                end
            end
        end
    end
end
