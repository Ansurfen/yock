---@diagnostic disable: duplicate-set-field

---@type YPM
ypm = {}

function ypm:open()
    local ypm_path = path.join(env.workdir, "ypm")
    local ypm_env = path.join(ypm_path, "env.json")
    safe_write(ypm_env, [[{
        "source": "github"
    }]])
    local env_fp = io.open(ypm_env, "r+")
    if type(env_fp) ~= "nil" then
        self.env = json.decode(env_fp:read("*a"))
    end
    local file = path.join(ypm_path, self.env["source"] .. ".json")
    safe_write(file, "{}")
    local fp = io.open(file, "r+")
    self.fp = fp
    if type(fp) ~= "nil" then
        self.buf = json.decode(fp:read("*a"))
    end
end

function ypm:read()
    return self.fp:read("*a")
end

function ypm:write()
    self.fp:seek("set")
    self.fp:write(json.encode(self.buf))
end

function ypm:set(k, v)
    self.buf[k] = v
end

function ypm:get(k)
    return self.buf[k]
end

function ypm:close()
    self.fp:close()
end

function ypm:set_cache(url, file)
    local ypm_path = path.join(env.workdir, "ypm")
    local ypm_cache = path.join(ypm_path, "cache.json")
    safe_write(ypm_cache, "{}")
    if self.cache ~= nil then
        self.cache:seek("set")
        self.cache_buf[url] = file
        self.cache:write(json.encode(self.cache_buf))
    end
end

function ypm:get_cache(url)
    local ypm_path = path.join(env.workdir, "ypm")
    local ypm_cache = path.join(ypm_path, "cache.json")
    safe_write(ypm_cache, "{}")
    local cache_buf = io.open(ypm_cache, "r+")
    self.cache = cache_buf
    if type(cache_buf) ~= "nil" then
        self.cache_buf = json.decode(cache_buf:read("*a"))
        return self.cache_buf[url]
    end
    return ""
end

---@see load_module
function load_module(module)
    ypm:open()
    local tmp_path = path.join(env.workdir, "..", "yock_tmp")
    mkdir(tmp_path)
    local url = ypm:get(module)
    -- 缓存删了,但是本地还是存在就不需要去爬取
    local file = fetch.script(url)
    local boot = import(path.join(tmp_path, file))
    local mod = boot()
    ypm:close()
    return mod
end
