parse_flags(env, {
    m = flag_type.array_type
})

ypm:open()
ypm:rm_module(env.flags["m"])
ypm:close()
