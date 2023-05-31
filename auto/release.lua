local script

if env.platform.OS == "windows" then

else
    local wd, err = pwd()
    yassert(err)
    script = cmdf("export", "path", "=", "$path", ";", wd)
end

sh({
    debug = true
}, script)
