--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

argsparse(env, {
    -- pwd
    wd = flag_type.str,
    -- local or global
    g = flag_type.bool,
    m = flag_type.str
})

ypm:open()
local module = env.flags["m"]
local g = env.flags["g"]
local wd = env.flags["wd"]

local yock_json = path.join(wd, "modules.json")
local yock_modules = path.join(wd, "yock_modules")

if g then
    yock_modules = env.yock_modules
end

if not is_exist(yock_json) then
    safe_write(yock_json, [[{"dependency":{}}]])
end

local fp

if g then
    fp = ypm.modules
else
    fp = jsonfile:open(yock_json)
    mkdir(yock_modules)
end

if #module > 0 then
    fp.buf.dependency[module] = ""
    fp:write()
end

for name, _ in pairs(fp.buf["dependency"]) do
    local remote_path = ypm:get(name)
    if remote_path == nil then
        yassert("invalid path")
    end
    local file = fetch.script(ypm:get(name))
    local mod = import(path.join(env.yock_tmp, file))
    if mod == nil or mod.load == nil then
        goto continue
    end
    -- mod.load({
    --     name = module,
    --     version = mod.version
    -- })
    mkdir(path.join(yock_modules, name))
    cp(path.join(env.yock_modules, name, "boot.lua"), path.join(yock_modules, name, "boot.lua"))
    local ver = strings.ReplaceAll(mod.version, ".", "_")
    cp(path.join(env.yock_modules, name, ver), path.join(yock_modules, name, ver))
    if #module > 0 and #fp.buf.dependency[module] == 0 then
        fp.buf.dependency[module] = mod.version
        fp:write()
    end
    ::continue::
end

ypm:close()
