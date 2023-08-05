--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global
---@diagnostic disable: duplicate-set-field

env.yock_modules = path.join(env.yock_path, "yock_modules")
env.pwd = pwd()

local ypm_path = pathf("~/ypm")
if not find(ypm_path) then
    mkdir(ypm_path)
end

config = json.create(pathf("~/ypm/config.json"), [[{"defaultSource": "github"}]])
modules = json.create(pathf("~/ypm/modules.json"), [[{"depend": {}}]])

config:save(true)
modules:save(true)

if not find(env.yock_tmp) then
    mkdir(env.yock_tmp)
end

if not find(env.yock_modules) then
    mkdir(env.yock_modules)
end

---@return unknown
---@return unknown loaderdata
function load_file(path)
    local code_path = filepath.Abs(pathf("#3", ".."))
    code_path = pathf(code_path, path)
    if strings.HasSuffix(code_path, ".lua") then
        code_path = string.sub(code_path, 1, #code_path - 4)
    end
    return require(code_path)
end

---@return string, string|nil
local resolve_name = function(name, sep)
    local before, after, ok = strings.Cut(name, sep)
    if ok then
        return before, after
    end
    return name, nil
end

---@return string name, string|nil version, string sub
local resolve_target = function(target)
    local name, version, sub
    if strings.Contains(target, "@") then
        local before, after, ok = strings.Cut(target, "@")
        if ok then
            version = after
            name = before
        end
    else
        name = target
    end
    if strings.Contains(name, "/") then
        name, sub = resolve_name(name, "/")
    elseif strings.Contains(name, "\\") then
        name, sub = resolve_name(name, "\\")
    end
    if sub == nil then
        sub = ""
    end
    return name, version, sub
end

local find_module = function(path, name)
    if find(path) then
        local jf = json.open(path)
        return jf:get(string.format("depend.%s.version", name))
    end
    return nil
end

---@param target string
---@return any, boolean
function load_module(target)
    local name, version, sub = resolve_target(target)
    if name == nil then
        yassert("invalid name")
    end
    local absPath = filepath.Abs(pathf("#3"))
    if strings.Contains(absPath, env.yock_modules) then
        if version == nil then
            local tmp = string.sub(absPath, #env.yock_modules + 2, #absPath)
            local idx = strings.IndexAny(tmp, string.char(filepath.Separator))
            if idx > 0 then
                local subname = string.sub(tmp, 1, idx)
                tmp = string.sub(tmp, #subname + 2, #tmp)
                idx = strings.IndexAny(tmp, string.char(filepath.Separator))
                if idx > 0 then
                    absPath = pathf(env.yock_modules, subname, string.sub(tmp, 1, idx), "modules.json")
                end
            end
        end
    else
        absPath = pathf(env.pwd, "modules.json")
    end
    local found = true
    if version == nil then
        version = find_module(absPath, name)
        if version == nil then
            if find(pathf(env.yock_modules, name, "boot.lua")) then
                local meta = require(pathf(env.yock_modules, name, "boot"))
                if meta ~= nil and meta.version ~= nil then
                    version = meta.version
                else
                    found = false
                end
            else
                found = false
            end
        end
    end
    if version ~= nil then
        -- Due to the feature of require, all "." will be replaced with "/".
        -- In order to solve the problem that
        -- the version directory name cannot be taken with "." feature, utilizing "_" replace.
        version = strings.ReplaceAll(version, ".", "_")
    else
        version = ""
    end
    if found then
        target = pathf(env.yock_modules, name, version, sub)
        if find(target .. ".lua") then
            return require(target), found
        else
            return require(pathf(target, "index"))
        end
    end
    return require(pathf(name, version, sub)), found
end

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
    if strings.HasPrefix(target, ".") then
        return load_file(target)
    end
    if not strings.HasSuffix(target, ".lua") and find(target .. ".lua") then
        return require(target)
    end
    if find(target) then
        return require(string.sub(target, 1, #target - 4))
    end
    return load_module(target)
end

service = json.create(pathf("~/ypm/service.json"))

---@param name string
---@param cmd fun(port: integer): string
---@return integer
function register_service(name, cmd)
    if service:rawget(name) == nil then
        local port = random.port()
        local c = cmd(port)
        yassert(nohup(c))
        service:rawset(name, {
            port = port,
            cmd = c
        })
        service:save(true)
        return port
    end
    return service:rawget(name).port
end

---@param name string
function unregister_service(name)
    if service:rawget(name) ~= nil then
        local meta = service:rawget(name)
        local infos = lsof(meta.port)
        if #infos > 0 then
            for _, info in ipairs(infos) do
                kill(info.pid)
            end
        end
    end
    service:rawset(name, nil)
    service:save(true)
end

---@param target string
function init(target)
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
    else
        module = target
        local tmp = json.open(pathf("$/modules.json"))
        version = tmp:getstr(strf("depend.%s.version", module))
    end
    version = strings.ReplaceAll(version, ".", "_")
    dofile(pathf(env.yock_modules, module, version, "init.lua"))
end
