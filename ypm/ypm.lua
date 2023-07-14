--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global
---@diagnostic disable: duplicate-set-field

-- Due to the feature of require, all "." will be replaced with "/".
-- In order to solve the problem that
-- the version directory name cannot be taken with "." feature, utilizing "_" replace.
local versionf = function(v)
    return strings.ReplaceAll(v, ".", "_")
end

env.yock_modules = path.join(env.yock_path, "yock_modules")

-- import layer
--
-- ypm injects yock_modules into package.path to guarantee
-- that global yock_modules can be found.
--
-- At the same time, the import layer also ensures that
-- the module can be introduced in the form of "module@version" or "module",
-- and the feasibility of local yock_modules.
---@param target string
---@return unknown
---@return unknown loaderdata
function import(target)
    -- injects yock_modules when yock_modules isn't exist
    if not strings.Contains(package.path, "yock_modules") then
        local wd, err = pwd()
        yassert(err)
        package.path = package.path .. string.format([[;%s;%s;%s;%s]],
            -- local yock_modules
            path.join(wd, "yock_modules", "?.lua"),
            path.join(wd, "yock_modules", "?", "index.lua"),
            -- global yock_modules
            path.join(env.yock_modules, "?.lua"),
            path.join(env.yock_modules, "?", "index.lua"))
    end

    -- Preprocessing to adapt to the require layer
    if strings.HasSuffix(target, ".lua") then
        target = string.sub(target, 1, #target - 4)
    end

    -- The actual path to the script file, not the path where the script runs.
    local code_path = pathf("#2", "..")

    -- import file by relative path
    -- if (strings.Contains(target, "/") or strings.Contains(target, "\\"))
    --     and string.sub(target, 1, 1) == "." then
    --     -- todo: pathf("!", code_path, target)
    --     return require(versionf(path.abs(path.join(code_path, target))))
    --     -- import file by absolute path
    -- elseif path.abs(target) == target then
    --     return require(versionf(target))
    -- end
    if find(pathf(code_path, target .. ".lua")) then
        return require(pathf(code_path, target))
    end

    local module = target
    local version = ""

    -- module@version
    if strings.Contains(target, "@") then
        local mod, ver, ok = strings.Cut(target, "@")
        if not ok then
            yassert("invalid module")
        end
        module = mod
        version = versionf(ver)
    end
    -- check local modules
    local modules_json = jsonfile:open(path.join(code_path, "modules.json"), true)
    -- check global modules
    if type(modules_json.fp) == "nil" then
        modules_json = jsonfile:open(path.join(code_path, "modules.json"), true)
    end

    if #version == 0 and modules_json.buf ~= nil then
        version = modules_json.buf["depend"][module]
    end
    local idx = strings.IndexAny(module, "/")
    if idx ~= -1 then
        module = path.join(string.sub(module, 1, idx), version,
            string.sub(module, idx + 1, #module))
    else
        module = pathf(module, version)
    end
    if string.sub(module, 1, 1) == string.char(path.Separator) then
        module = string.sub(module, 2, #module)
    end
    if find(module) then
        return require(module)
    end
    return require(target)
end

function cur_dir()
    return path.join(debug.getinfo(2, "S").source, "..")
end

-- load_module layer

function yock_todo_completion()
    local file = fetch.file([[https://raw.githubusercontent.com/Ansurfen/yock-todo/main/release/release.json]],
        ".json")
    local release_json = jsonfile:open(pathf("~/tmp/") .. file)
    local suffix
    if env.platform.OS == "windows" then
        suffix = ".zip"
    else
        suffix = ".tar.gz"
    end
    file = fetch.file(string.format([[https://github.com/Ansurfen/yock-todo/releases/download/%s/ypm%s]],
        release_json.buf["ypm"]["tag"], suffix), suffix)
    mkdir(pathf("~/yock_modules"))
    if env.platform.OS == "windows" then
        unzip(pathf("~/tmp/") .. file, pathf("~/yock_modules"))
        mv(pathf("~/ypm/modules/*"), pathf("~/yock_modules"))
        rm({ safe = false }, pathf("~/ypm/modules"))
    else
        ---@diagnostic disable-next-line: param-type-mismatch
        sh(string.format([[
        tar -xvf %s -C %s --strip-components=2
    ]], pathf("~/tmp/") .. file, pathf("~/yock_modules")))
    end
    return true
end

function yock_todo_loader(opt)
    yock_todo_completion()
    return import(path.join(env.yock_modules, opt["name"], opt["version"], "index"))
end

mkdir(pathf("~/ypm"))
config = json.create(pathf("~/ypm/config.json"), [[{"defaultSource": "github"}]])
modules = json.create(pathf("~/ypm/modules.json"), [[{"depend": {}}]])

config:save(true)
modules:save(true)

---@param target string
---@return any
function load_module(target)
    mkdir(env.yock_tmp, env.yock_modules)

    local module = target
    local version = ""
    -- module@version
    if strings.Contains(target, "@") then
        local mod, ver, ok = strings.Cut(target, "@")
        if not ok then
            yassert("invalid module")
        end
        module = mod
        version = ver
    end

    -- check local modules
    if find(pathf("$/yock_modules", module)) then
        if #version == 0 then
            ---@diagnostic disable-next-line: redundant-return-value
            return import(target)
        else
            ---@diagnostic disable-next-line: redundant-return-value
            return import(pathf("~/yock_modules", module, version, "index"))
        end
    end

    -- check global modules
    if find(pathf("~/yock_modules", module)) then
        if #version == 0 then
            version = modules:get(string.format("depend.%s", module))
        end
        if version == nil or #version == 0 then
            ---@diagnostic disable-next-line: redundant-return-value
            return import(module)
        end
        ---@diagnostic disable-next-line: redundant-return-value
        return import(path.join(env.yock_modules, module, version, "index"))
    else
        local parse = import(pathf("~/ypm/util/parse"))
        local lib, ok

        parse(module, function(url)
            local file, err = fetch.file(url, ".lua")
            if err == nil then
                ---@type module
                local mod = import(file)
                if mod.load ~= nil then
                    modules:set(string.format("depend.%s", module), mod.version)
                    modules:save(true)
                    lib = mod.load({
                        name = module,
                        version = version,
                    })
                    ok = true
                    return
                end
            end
            return err
        end)

        if lib ~= nil then
            return lib
        end

        if ok then
            return
        end

        yassert("invalid mod loader")
    end
end
