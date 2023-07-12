--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "uninstall" },
    run = function(cmd, args)
        if #args == 0 then
            yassert("arguments too little")
        end
        local uninstallParameter = env.params["/ypm/uninstall"]
        local g = false
        if type(uninstallParameter) == "table" then
            g = uninstallParameter["g"]:Var()
        end
        local name = args[1]
        local modules_path
        if g then
            modules_path = pathf("~/ypm/modules.json")
        else
            modules_path = "./modules.json"
        end
        local mod_path = pathf("~/yock_modules", name, "boot.lua")
        if find(mod_path) then
            local mod = import(mod_path)
            if type(mod) == "table" and type(mod.unload) == "function" then
                mod.unload()
            end
        end
        local jf = json.open(modules_path)
        jf:set(string.format("denpend.%s", name), nil)
        jf:save(true)
    end,
    flags = {
        {
            type = flag_type.bool,
            name = "global",
            shorthand = "g",
            default = false,
            usage = ""
        }
    }
}
