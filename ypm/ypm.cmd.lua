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
        },
        {
            desc = { use = "install [module-name]" },
            run = function(cmd, args)
                local module
                if #args >= 1 and not strings.HasPrefix(args[1], "-") then
                    module = args[1]
                end
                local installParameter = env.params["/ypm/install [module-name]"]
                local g = installParameter["g"]:Var()
                local ypm_path = path.join(env.args[1], "..")
                local wd, err = pwd()
                yassert(err)
                cd(ypm_path)

                local arg_builder = argBuilder:new():add("yock run cmd/install.lua --"):add_str(
                    "-wd " .. wd, wd):add_bool("-g", g):add_str("-m " .. module,
                    module)

                exec({
                    redirect = true,
                    debug = true
                }, arg_builder:build())
            end,
            flags = {
                {
                    default = false,
                    type = flag_type.bool_type,
                    name = "global",
                    shorthand = "g",
                    usage = ""
                }
            }
        },
        {
            desc = { use = "uninstall [module-name]" },
            run = function(cmd, args)
                if #args < 1 then
                    yassert("arguments too little")
                end
                local uninstallParameter = env.params["/ypm/uninstall [module-name]"]
                local g = uninstallParameter["g"]:Var()
                local module = args[1]

                local wd, err = pwd()
                yassert(err)

                local ypm_path = path.join(env.args[1], "..")
                cd(ypm_path)


                if #module > 0 then
                    local arg_builder = argBuilder:new():add("yock run cmd/uninstall.lua --"):add_str("-wd " .. wd, wd)
                        :add_str("-m " .. module, module):add_bool("-g", g)
                    exec({
                        redirect = true,
                        debug = true
                    }, arg_builder:build())
                end
            end,
            flags = {
                {
                    default = false,
                    type = flag_type.bool_type,
                    name = "global",
                    shorthand = "g",
                    usage = ""
                }
            }
        }
    },
    flags = {}
})
