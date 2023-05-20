parse_flags(env, {
    k = flag_type.string_type,
    v = flag_type.string_type
})

ypm:open()
ypm:set_cache(env.flags["k"], env.flags["v"])
ypm:close()
