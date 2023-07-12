-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: undefined-global

return {
    desc = {
        use = "init [module-name]",
        short = "Initialize new module in current directory",
        long = [[Init initializes and writes a new index.lua file in the current directory, in
		effect creating a new module rooted at the current directory. The index.lua file
		must not already exist.]]
    },
    run = function(cmd, args)
        if #args == 0 then
            yassert("arguments too little")
        end
        local name = args[1]
        local dir = false
        local ver = "1.0.0"
        local initParameter = env.params["/ypm/init"]
        if type(initParameter) == "table" then
            dir = initParameter["d"]:Var()
            ver = initParameter["v"]:Var()
        end
        if dir then
            if find(name) then
                yassert("directory is exist")
            end
            mkdir(name)
            cd(name)
        end
        if find("boot.lua") then
            yassert("boot file exist")
        end
        write("boot.lua", strf(cat(pathf("#1", "../../template/boot.tpl")), {
            name = name,
            version = ver
        }))
        ver = strings.ReplaceAll(ver, ".", "_")
        mkdir(ver, "include")
        write(pathf(ver, "index.lua"), cat(pathf("#1", "../../template/index.tpl")))
        cp(pathf("~/lib/include"), pathf("$"))
        rm({ safe = false }, pathf("$/include/lang"), pathf("$/include/misc"))
        write("README.md", string.format("# %s", name))
        write(".gitignore", cat(pathf("#1", "../../template/gitignore.tpl")))
    end,
    flags = {
        {
            type = flag_type.str,
            name = "lang",
            default = "en_us",
            shorthand = "l",
            usage = ""
        },
        {
            type = flag_type.str,
            name = "ver",
            shorthand = "v",
            default = "1.0.0",
            usage = ""
        },
        {
            type = flag_type.bool,
            name = "dir",
            shorthand = "d",
            default = false,
            usage = ""
        }
    }
}
