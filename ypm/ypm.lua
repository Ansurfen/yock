---@diagnostic disable: duplicate-set-field

---@type YPM
ypm = {}

function ypm:open()
    local ypm_path = path.join(env.workdir, "ypm")
    local ypm_env = path.join(ypm_path, "env.json")
    local ypm_cache = path.join(ypm_path, "cache.json")
    local ypm_modules = path.join(ypm_path, "modules.json")
    safe_write(ypm_env, [[{
        "source": "github"
    }]])
    safe_write(ypm_cache, "{}")
    safe_write(ypm_modules, [[{
        "dependency": {}
    }]])
    self.env = jsonfile:open(ypm_env)
    local ypm_source = path.join(ypm_path, self.env.buf["source"] .. ".json")
    safe_write(ypm_source, "{}")
    self.source = jsonfile:open(ypm_source)
    self.modules = jsonfile:open(ypm_modules)
    self.cache = jsonfile:open(ypm_cache)
end

function ypm:read()
    return self.source.buf:read("*a")
end

function ypm:write()
    self.source:write()
end

function ypm:set(k, v)
    self.source:set(k, v)
end

function ypm:get(k)
    return self.source:get(k)
end

function ypm:close()
    self.source:close()
    self.env:close()
    self.modules:close()
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
    ypm.modules:write()
end

---@see load_module
function load_module(module)
    ypm:open()
    local tmp_path = path.join(env.workdir, "..", "yock_tmp")
    mkdir(tmp_path)

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
    local modules_path = path.join(env.workdir, "..", "lua_modules")
    local exist = is_exist(path.join(modules_path, module))
    local mod
    if not exist then
        local file = fetch.script(url)
        local boot = import(path.join(tmp_path, file))
        mod = boot(module)
    else
        version = ypm.modules:get("dependency")[module]
        mod = import(path.join(modules_path, module, version, "index"))
    end
    ypm:close()
    return mod
end
