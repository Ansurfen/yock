--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "install [module-name]" },
    run = function(cmd, args)
        local module
        if #args >= 1 and not strings.HasPrefix(args[1], "-") then
            module = args[1]
        end
        local installParameter = env.params["/ypm/install"]
        local g = installParameter["g"]:Var()
        local ypm_path = path.join(env.args[1], "..")
        local wd, err = pwd()
        yassert(err)
        cd(ypm_path)

        local arg_builder = argBuilder:new():add("yock run cmd/install.lua --"):add_str(
            "-wd " .. wd, wd):add_bool("-g", g):add_str("-m " .. module,
            module)

        sh({
            redirect = true,
            debug = true
        }, arg_builder:build())
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
