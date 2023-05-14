function import(module)
    if not strings.Contains(package.path, "lua_modules\\?.lua") then
        local wd, ok = pwd()
        if not ok then
            assert("fail to get work directory")
        end
        package.path = "?.lua" ..
            [[;]] ..
            path.join(wd, "lua_modules", "?", "index.lua") ..
            [[;]] ..
            path.join(wd, "lua_modules", "?.lua") ..
            [[;]] ..
            path.join(env.workdir, "..", "lua_modules", "?", "index.lua") ..
            [[;]] .. path.join(env.workdir, "..", "lua_modules", "?.lua")
    end
    if strings.HasPrefix(module, "./") or strings.HasPrefix(module, "../") then
        return require(path.join(debug.getinfo(2, "S").source, "..", module))
    end
    return require(path.join(module))
end
