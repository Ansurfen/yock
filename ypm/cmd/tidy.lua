--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "tidy" },
    run = function(cmd, args)
        local proxies, err = find({
            dir = false,
            pattern = "\\.lua"
        }, pathf("#1", "../../proxy"))
        if err ~= nil or #proxies == 0 then
            cp(cat(pathf("#1", "../../template/defaultSource.tpl")), pathf("#1", "../../proxy"))
        end
        cp(pathf("~/lib/include"), pathf("$"))
        if find(pathf("$/modules.json")) then
            local modules = json.open(pathf("$/modules.json"))
            path.walk(env.yock_modules, function(p, info, err)
                yassert(err)
                if info:IsDir() then
                    if strings.HasSuffix(p, "docs") or strings.HasSuffix(p, "doc") then
                        local got = filepath.Base(filepath.Dir(p))
                        local name = filepath.Base(filepath.Dir(filepath.Dir(p)))
                        local want = modules:get("depend." .. name)
                        if want == nil or strings.ReplaceAll(want, ".", "_") ~= got then
                            return true
                        end
                        mkdir(pathf("$/include", name))
                        cp(p .. "/*", pathf("$/include", name))
                    end
                end
                return true
            end)
        end
    end
}
