-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

local defaultSource = "https://github.com/ansurfen/yock-todo"
local defaultProxyIdent = "github"

---@return string|nil
local parseProxy = function(p, m)
    for _, v in pairs(p["filter"]) do
        if v == m then
            return nil
        end
    end
    local rurl = p["redirect"][m]
    if rurl ~= nil then
        return rurl
    end
    return strf(p["url"], {
        ver = m
    })
end

return {
    desc = { use = "install" },
    run = function(cmd, args)
        if #args == 0 then
            yassert("arguments too little")
        end
        local installParameter = env.params["/ypm/install"]
        local g = false
        local p = ""
        if type(installParameter) == "table" then
            g = installParameter["g"]:Var()
            p = installParameter["p"]:Var()
        end

        local mod = args[1]
        local proxies, err = find({
            dir = false,
            pattern = "\\.lua"
        }, pathf("#1", "../../proxy"))

        local defaultProxy
        local candidates = {}
        if err ~= nil or #proxies == 0 then
            print("prepare to fetch default source...")
            print(defaultSource)
        else
            print("select startegies from proxies")
            if type(proxies) == "table" then
                for _, proxy in ipairs(proxies) do
                    local filename = path.filename(proxy)
                    if filename == defaultProxyIdent then
                        defaultProxy = import(proxy)
                    else
                        table.insert(candidates, import(proxy))
                    end
                end
            end
        end

        local res
        if defaultProxy ~= nil then
            res = parseProxy(defaultProxy, mod)
        end
        for _, proxy in ipairs(candidates) do
            if res ~= nil then
                break
            end
            res = parseProxy(proxy, mod)
        end

        if res ~= nil then
            local file = fetch.file(res, ".lua")

            ---@type module
            local module = import(pathf(env.yock_tmp, file))

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

            local modules_path
            if g then
                modules_path = pathf("~/ypm/modules.json")
            else
                modules_path = pathf("$/modules.json")
            end
            mkdir("yock_modules")
            local modules_json = json.create(modules_path, [[{"denpend":{}}]])
            modules_json:set(string.format("denpend.%s", mod), module.version)
            modules_json:save(true)
        end
    end,
    flags = {
        {
            type = flag_type.bool,
            name = "global",
            shorthand = "g",
            default = false,
            usage = ""
        },
        {
            type = flag_type.str,
            name = "proxy",
            shorthand = "p",
            default = "",
            usage = ""
        }
    }
}
