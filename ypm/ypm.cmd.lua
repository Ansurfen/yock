ctl({
    desc = {
        use = "ypm",
    },
    sub = {
        {
            desc = { use = "rm" },
            run = function()
                local rmParameter = env.params["/ypm/rm"]
                local module = rmParameter["m"]:Var()

                local ypm_path = path.join(env.args[1], "..")
                cd(ypm_path)
                if #module > 0 then
                    exec({
                        redirect = true,
                        debug = true
                    }, "yock run cmd/rm_module.lua -- -m " .. module)
                end
            end,
            flags = {
                {
                    type = flag_type.string_type,
                    default = "",
                    name = "module",
                    shorthand = "m",
                    usage = ""
                }
            }
        }
    },
    flags = {}
})
