--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "rm" },
    run = function()
        local rmParameter = env.params["/ypm/rm"]
        local module = rmParameter["m"]:Var()

        local ypm_path = path.join(env.args[1], "..")
        cd(ypm_path)
        if #module > 0 then
            sh({
                redirect = true,
                debug = true
            }, "yock run cmd/rm_module.lua -- -m " .. module)
        end
    end,
    flags = {
        {
            type = flag_type.str,
            default = "",
            name = "module",
            shorthand = "m",
            usage = ""
        }
    }
}
