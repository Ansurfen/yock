--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "tidy", short = "Complement definitions" },
    run = function(cmd, args)
        local proxies, err = find({
            dir = false,
            pattern = "\\.lua"
        }, pathf("#1", "../../proxy"))
        if err ~= nil or #proxies == 0 then
            cp(cat(pathf("#1", "../../template/defaultSource.tpl")), pathf("#1", "../../proxy"))
        end

        local tidyParameter = env.params["/ypm/tidy"]
        local vsc = tidyParameter["c"]:Var()
        if type(vsc) == "boolean" and vsc then
            if not find(".vscode") then
                mkdir(".vscode")
            end

            local new_libs = { pathf(env.yock_path, "lib", "include") }

            if find(pathf("$/modules.json")) then
                local modules = json.open(pathf("$/modules.json"))
                local deps = modules:rawget("depend")
                for name, dep in pairs(deps) do
                    local path = pathf(env.yock_modules, name, strings.ReplaceAll(dep.version, ".", "_"), "doc")
                    if find(path) then
                        table.insert(new_libs, path)
                    end
                end
            end

            local settings = json.create(".vscode/settings.json", "{}")
            local libs = settings:rawget("Lua.workspace.library")
            if libs == nil then
                libs = {}
            end
            for _, lib in ipairs(new_libs) do
                local found = false
                for _, value in ipairs(libs) do
                    if value == lib then
                        found = true
                        break
                    end
                end
                if not found then
                    table.insert(libs, lib)
                end
            end

            settings:rawset("Lua.workspace.library", libs)
            settings:save(true)
        else
            rm({ safe = false }, pathf("$/include"))
            cp(pathf("~/lib/include"), pathf("$"))
            if find(pathf("$/modules.json")) then
                local modules = json.open(pathf("$/modules.json"))
                ls(env.yock_modules, function(p, info)
                    if info:IsDir() then
                        if strings.HasSuffix(p, "docs") or strings.HasSuffix(p, "doc") then
                            local got = filepath.Base(filepath.Dir(p))
                            local name = filepath.Base(filepath.Dir(filepath.Dir(p)))
                            local want = modules:get("depend." .. name)
                            if want == nil or strings.ReplaceAll(want.version, ".", "_") ~= got then
                                return
                            end
                            mkdir(pathf("$/include", name))
                            cp(p .. "/*", pathf("$/include", name))
                        end
                    end
                end)
            end
        end
    end,
    flags = {
        {
            default = false,
            type = flag_type.bool,
            shorthand = "c",
            name = "vscode",
            usage = ""
        }
    }
}
