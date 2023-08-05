-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

local find_modules = function(path, modules)
    if find(path) then
        local module = json.open(path)
        for name, value in pairs(module:rawget("depend")) do
            local version = strings.ReplaceAll(value.version, ".", "_")
            local key = name
            if modules[key] == nil then
                modules[key] = {}
            end
            local file = pathf(env.yock_modules, name, version, "modules.json")
            modules[key][version] = {
                file = file,
                resolved = value.resolved or {}
            }
            ---@diagnostic disable-next-line: undefined-global
            find_modules(file, modules)
        end
    end
end

return {
    desc = { use = "install", short = "Install module to be specified" },
    run = function(cmd, args)
        local installParameter = env.params["/ypm/install"]
        local p = ""
        local w = true
        if type(installParameter) == "table" then
            p = installParameter["p"]:Var()
            if installParameter["w"]:Var() then
                w = false
            end
        end
        if #args == 0 then
            local modulesPath = pathf("$", "modules.json")
            local modules = {}
            find_modules(modulesPath, modules)
            for name, mod in pairs(modules) do
                for ver, meta in pairs(mod) do
                    ycho.info(string.format("install %s@%s", name, ver))
                    if #meta.resolved > 0 then
                        sh(string.format("ypm install %s -w", meta.resolved[1]))
                    end
                end
            end
        else
            local mod = args[1]
            if strings.Contains(mod, "/") and strings.Contains(mod, "@") then
                local before, after, ok = strings.Cut(mod, "@")
                if ok then
                    local idx = strings.IndexAny(before, "/")
                    if idx > 0 then
                        local policy = string.sub(before, 1, idx)
                        local _switch = json.open(pathf("#1", "../install.json")).buf
                        local name = string.sub(before, strings.LastIndex(before, "/") + 2, #before)
                        local url = _switch[policy]
                        if url ~= nil then
                            url = strf(url, {
                                Repo = string.sub(before, idx + 2, #before),
                                TagPack = wrapzip(after),
                                Tag = after,
                                ReleasePack = wrapzip(name),
                            })
                            local file = fetch.file(url, env.platform:Zip())
                            local dir = pathf(path.dir(file), path.filename(path.filename(file)))
                            if #file == 0 then
                                yassert("fail to fetch package")
                            end
                            uncompress(file, dir)
                            local boot
                            path.walk(dir, function(p, info, err)
                                yassert(err)
                                if info:IsDir() then
                                    return true
                                end
                                if strings.HasSuffix(p, "boot.lua") then
                                    boot = p
                                    return false
                                end
                                return true
                            end)
                            local newDir = strings.ReplaceAll(filepath.Dir(boot), ".", "_")
                            local bootDir = filepath.Dir(boot)
                            if strings.Contains(bootDir, ".") and not find(newDir) then
                                mv(bootDir, newDir)
                                boot = pathf(newDir, "boot")
                            end
                            local meta = import(boot)
                            local version = strings.ReplaceAll(after, ".", "_")
                            if not find(pathf(env.yock_modules, meta.name, version)) then
                                mkdir(pathf(env.yock_modules, meta.name, version))
                            end

                            local files = ioutil.ReadDir(pathf(env.yock_modules, meta.name, version))
                            if #files == 0 then
                                cp(pathf(newDir, "*"), pathf(env.yock_modules, meta.name, version))
                            end

                            local boot_lua = import(pathf(env.yock_modules, meta.name, version, "boot"))
                            if type(boot_lua.load) == "function" then
                                boot_lua.load({
                                    ver = version
                                })
                            end

                            if w then
                                local modules_json = json.create(pathf("$/modules.json"), [[{"depend":{}}]])
                                modules_json:set(string.format("depend.%s.version", meta.name), after)
                                local resolved = modules_json:get(strf("depend.%s.resolved", meta.name)) or {}
                                local found = false
                                for _, resolve in ipairs(resolved) do
                                    if resolve == mod then
                                        found = true
                                        break
                                    end
                                end
                                if not found then
                                    table.insert(resolved, mod)
                                end
                                modules_json:set(strf("depend.%s.resolved", meta.name), resolved)
                                modules_json:save(true)
                            end
                        end
                    end
                end
            else
                local parse = import("../util/parse")

                parse(mod, function(url)
                    local file = fetch.file(url, ".lua")
                    if #file == 0 then
                        return
                    end
                    ---@type module
                    local module = import(file)

                    if type(p) == "string" and #p > 0 then
                        local target = module.name
                        if env.platform.OS == "windows" then
                            target = target .. ".zip"
                        else
                            target = target .. ".tar.gz"
                        end
                        p = strf(p, {
                            ver = module.version,
                            target = target
                        })
                    end
                    if #p == 0 then
                        ---@diagnostic disable-next-line: cast-local-type
                        p = nil
                    end

                    module.load({
                        ver = module.version,
                        url = p
                    })

                    local modules_json = json.create(pathf("~/ypm/modules.json"), [[{"depend":{}}]])
                    modules_json:set(string.format("depend.%s", mod), module.version)
                    modules_json:save(true)
                    if w then
                        modules_json = json.create(pathf("$/modules.json"), [[{"depend":{}}]])
                        modules_json:set(string.format("depend.%s", mod), module.version)
                        modules_json:save(true)
                    end
                end)
            end
        end
    end,
    flags = {
        {
            type = flag_type.str,
            name = "proxy",
            shorthand = "p",
            default = "",
            usage = ""
        },
        {
            type = flag_type.bool,
            name = "write",
            shorthand = "w",
            default = false,
            usage = ""
        }
    }
}
