parse_flags(env, {
    m = flag_type.string_type,
    v = flag_type.string_type
})

ypm:open()
ypm:new_module(env.flags["m"], env.flags["v"])
ypm:close()
