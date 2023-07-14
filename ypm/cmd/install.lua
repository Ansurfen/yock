-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

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
        local parse = import("../util/parse")

        parse(mod, function(url)
            local file = fetch.file(url, ".lua")

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
        end)
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
