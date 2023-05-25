parse_flags(env, {
    m = flag_type.string_type,
    g = flag_type.bool_type,
    wd = flag_type.string_type
})

ypm:open()
local module = env.flags["m"]
local g = env.flags["g"]
local wd = env.flags["wd"]

if g then
    ypm:rm_module(module)
else
    local fp = jsonfile:open(path.join(wd, "modules.json"))
    fp.buf["dependency"][module] = nil
    fp:write()
    fp:close()
    rm({
        safe = false
    }, path.join(wd, "yock_modules", module))
end

ypm:close()
