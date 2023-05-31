--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "uninstall [module-name]" },
    run = function(cmd, args)
        if #args < 1 then
            yassert("arguments too little")
        end
        local uninstallParameter = env.params["/ypm/uninstall"]
        local g = uninstallParameter["g"]:Var()
        local module = args[1]

        local wd, err = pwd()
        yassert(err)

        local ypm_path = path.join(env.args[1], "..")
        yassert(cd(ypm_path))

        if #module > 0 then
            local arg_builder = argBuilder:new():add("yock run cmd/uninstall.lua --"):add_str("-wd " .. wd, wd)
                :add_str("-m " .. module, module):add_bool("-g", g)
            sh({
                redirect = true,
                debug = true
            }, arg_builder:build())
        end
    end,
    flags = {
        {
            default = false,
            type = flag_type.bool,
            name = "global",
            shorthand = "g",
            usage = ""
        }
    }
}
