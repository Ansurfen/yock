local versionf = function(v)
    return strings.ReplaceAll(v, ".", "_")
end

function import(module)
    if not strings.Contains(package.path, "yock_modules\\?.lua") then
        local wd, ok = pwd()
        if not ok then
            assert("fail to get work directory")
        end
        package.path = "?.lua" ..
            [[;]] ..
            path.join(wd, "yock_modules", "?", "index.lua") ..
            [[;]] ..
            path.join(wd, "yock_modules", "?.lua") ..
            [[;]] ..
            path.join(env.yock_path, "yock_modules", "?", "index.lua") ..
            [[;]] .. path.join(env.yock_path, "yock_modules", "?.lua")
    end
    local version = ""
    -- module@version
    if strings.Contains(module, "@") then
        local before, after, ok = strings.Cut(module, "@")
        if not ok then
            yassert("invalid module")
        end
        module = before
        version = after
    end
    -- only import file
    if strings.HasPrefix(module, "./") or strings.HasPrefix(module, "../") then
        return require(path.join(debug.getinfo(2, "S").source, "..", module))
    elseif path.abs(module) == module then
        return require(versionf(module))
    end
    local root = path.join(debug.getinfo(2, "S").source, "..")
    local pkg, err = io.open(path.join(root, "modules.json"), "r")
    local pkgFile
    if type(pkg) ~= "nil" then
        pkgFile = json.decode(pkg:read("*a"))
    else
        -- global env
        pkg = io.open(path.join(env.workdir, "ypm", "modules.json"))
        if type(pkg) ~= "nil" then
            pkgFile = json.decode(pkg:read("*a"))
        end
    end
    if #version == 0 then
        version = pkgFile["dependency"][module]
    end
    version = versionf(version)
    return require(path.join(module, version))
end

function cur_dir()
    return path.join(debug.getinfo(2, "S").source, "..")
end
