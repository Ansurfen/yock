--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field

---@type YPM
ypm = {}

function ymodule(todo)
    local modules_path = path.join(env.yock_path, "yock_modules")
    local tmp_path = path.join(env.yock_path, "yock_tmp")
    mkdir(modules_path)
    mkdir(tmp_path)
    local url = todo["url"]
    local suffix
    if env.platform.OS == "windows" then
        suffix = ".zip"
    else
        suffix = ".tar.gz"
    end
    local init = false
    if url == nil and todo["tag"] then
        local ypm_release = pathf("~/ypm/release.json")
        if not is_exist(ypm_release) then
            http({
                debug = true,
                strict = true,
                save = true,
                dir = pathf("~/ypm"),
                filename = function(s)
                    return "release.json"
                end
            }, [[https://raw.githubusercontent.com/Ansurfen/yock-todo/main/release/release.json]])
        end
        local release_json = jsonfile:open(pathf("~/ypm/release.json"))
        url = string.format([[https://github.com/Ansurfen/yock-todo/releases/download/%s/ypm%s]],
            release_json.buf["ypm"]["tag"], suffix)
        init = true
    end
    local version = todo["version"]
    local module = todo["name"]
    local file = fetch.zip(url)
    return function(new_todo)
        if new_todo["version"] ~= nil and #new_todo["version"] ~= 0 then
            version = new_todo["version"]
        else
            new_todo["version"] = version
        end
        if new_todo["modules_path"] ~= nil then
            modules_path = new_todo["modules_path"]
        end
        if new_todo["info"] ~= nil and new_todo["info"] then
            return
        end
        if init then
            if env.platform.OS == "windows" then
            else
                exec({
                        strict = true
                    },
                    string.format([[tar -zxvf %s.tar.gz -C ../yock_modules/ --strip-components=2]],
                        path.join(tmp_path, file)))
            end
        else
            yassert(unzip(path.join(tmp_path, file .. suffix), path.join(modules_path, module)))
        end
        ypm:new_module(module, version)
        return require(path.join(modules_path, module, version, "index"))
    end
end

function ypm:open()
    local ypm_path = path.join(env.yock_path, "ypm")
    local ypm_env = path.join(ypm_path, "env.json")
    local ypm_cache = path.join(ypm_path, "cache.json")
    local ypm_modules = path.join(ypm_path, "modules.json")
    mkdir(ypm_path)
    safe_write(ypm_env, [[{
        "source": [
            "github",
            "gitlab",
            "gitee"
        ]
    }]])
    safe_write(ypm_cache, "{}")
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
    self.cache = jsonfile:open(ypm_cache)
end

function ypm:tidy()
    local files = ls({
        dir = pathf("~/yock_modules"),
        str = false
    })
    if type(files) == "table" then
        for _, file in ipairs(files) do
            local name = file[4]
            print(pathf("~/yock_modules/" .. name .. "/boot"))
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
        res = repo:get(k)
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

function ypm:set_cache(url, file)
    if self.cache ~= nil then
        self.cache:set(url, file)
        self.cache:write()
    end
end

function ypm:get_cache(url)
    if self.cache ~= nil then
        return self.cache:get(url)
    end
    return ""
end

function ypm:clear_cache()
    ypm.cache.buf = {}
    ypm.cache:write()
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

---@see load_module
function load_module(module)
    ypm:open()
    local tmp_path = path.join(env.yock_path, "yock_tmp")
    local modules_path = path.join(env.yock_path, "yock_modules")
    mkdir(tmp_path)
    mkdir(modules_path)

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

    local url = ypm:get(module)

    -- check local modules
    local wd, err = pwd()
    yassert(err)
    local exist = is_exist(path.join(wd, "yock_modules", module))
    if exist then
        return import(module)
    end
    -- check global modules
    exist = is_exist(path.join(modules_path, module))
    local mod
    if not exist then
        local file = fetch.script(url)
        local boot = import(path.join(tmp_path, file))
        mod = boot({
            name = module,
            version = version
        })
    else
        if #version <= 0 then
            version = ypm.modules:get("dependency")[module]
        end
        mod = import(path.join(modules_path, module, version, "index"))
    end
    ypm:close()
    return mod
end
