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
    local code_path = path.join(debug.getinfo(2, "S").source, "..")

    -- import file by relative path
    if (strings.Contains(target, "/") or strings.Contains(target, "\\"))
        and string.sub(target, 1, 1) == "." then
        return require(versionf(path.join(code_path, target)))
        -- import file by absolute path
    elseif path.abs(target) == target then
        return require(versionf(target))
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
        version = ver
    end
    -- check local modules
    local modules_json = jsonfile:open(path.join(code_path, "modules.json"), true)
    -- check global modules
    if type(modules_json.fp) == "nil" then
        modules_json = jsonfile:open(path.join(code_path, "modules.json"), true)
    end

    if #version == 0 and modules_json.buf ~= nil then
        version = modules_json.buf["dependency"][module]
    end

    return require(path.join(module, versionf(version)))
end

function cur_dir()
    return path.join(debug.getinfo(2, "S").source, "..")
end

-- load_module layer

---@type YPM
ypm = {}

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

function github_completion(url, version)
    fetch.file(string.format(url, version))
end

function github_loader(url)
    return function(opt)
        github_completion(url, opt["version"])
        return import(path.join(env.yock_modules, opt["name"], opt["version"], "index"))
    end
end

function ypm:open()
    local ypm_path = path.join(env.yock_path, "ypm")
    local ypm_env = path.join(ypm_path, "env.json")
    local ypm_modules = path.join(ypm_path, "modules.json")
    mkdir(ypm_path)
    safe_write(ypm_env, [[{
        "source": [
            "github",
            "gitlab",
            "gitee"
        ]
    }]])
    safe_write(ypm_modules, [[{
        "dependency": {}
    }]])
    self.env = jsonfile:open(ypm_env)
    self.sources = {}
    for _, source in ipairs(self.env.buf["source"]) do
        local repo_json = path.join(ypm_path, source .. ".json")
        safe_write(repo_json, "{}")
        local repo = jsonfile:open(repo_json)
        table.insert(self.sources, repo)
    end
    self.modules = jsonfile:open(ypm_modules)
end

function ypm:tidy()
    local files = ls({
        dir = pathf("~/yock_modules"),
        str = false
    })
    if type(files) == "table" then
        for _, file in ipairs(files) do
            local name = file[4]
            local boot = import(pathf("~/yock_modules/" .. name .. "/boot"))
            local todo = { info = true }
            boot(todo)
            ypm:new_module(name, todo["version"])
        end
    end
end

function ypm:get(k)
    local res
    for _, repo in ipairs(self.sources) do
        res = repo.buf[k]
        if res ~= nil and type(res) == "string" and #res > 0 then
            return res
        end
    end
    return res
end

function ypm:close()
    self.env:close()
    self.modules:close()
    for _, source in ipairs(self.sources) do
        source:close()
    end
end

function ypm:new_module(module, version)
    ypm.modules.buf["dependency"][module] = version
    ypm.modules:write()
end

function ypm:rm_module(module)
    ypm.modules.buf["dependency"][module] = nil
    local modules_path = path.join(env.yock_path, "yock_modules")
    rm({
        safe = false,
        debug = true
    }, path.join(modules_path, module))
    ypm.modules:write()
end

---@param target string
---@return any
function load_module(target)
    mkdir(env.yock_tmp)
    mkdir(env.yock_modules)

    local module = ""
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
    local wd, err = pwd()
    yassert(err)

    if is_exist(path.join(wd, "yock_modules", module)) then
        if #version == 0 then
            ---@diagnostic disable-next-line: redundant-return-value
            return import(module)
        else
            ---@diagnostic disable-next-line: redundant-return-value
            return import(path.join(wd, "yock_modules", module, version, "index"))
        end
    end

    -- check global modules
    if is_exist(path.join(env.yock_modules, module)) then
        if #version == 0 then
            version = ypm.modules.buf["dependency"][module]
        end
        if version == nil or #version == 0 then
            ---@diagnostic disable-next-line: redundant-return-value
            return import(module)
        end
        ---@diagnostic disable-next-line: redundant-return-value
        return import(path.join(env.yock_modules, module, version, "index"))
    else
        local remote_path = ypm:get(module)
        local file = fetch.script(remote_path)
        local mod = import(path.join(env.yock_tmp, file))
        if mod.load ~= nil then
            return mod.load({
                name = module,
                version = version,
            })
        end
        yassert("invalid mod loader")
    end
end

ypm:open()
