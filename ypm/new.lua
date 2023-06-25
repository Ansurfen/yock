-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: undefined-global
return {
    desc = {
        use = "new [module-name]",
        short = "Initialize new module in current directory",
        long = [[Init initializes and writes a new index.lua file in the current directory, in
		effect creating a new module rooted at the current directory. The index.lua file
		must not already exist.]]
    },
    run = function(cmd, args)
        if #args == 0 or args[1][1] == '-' then
            yassert("arguments too little")
        end
        local module = args[1]
        local newParameter = env.params["/ypm/new"]
        local lang = newParameter["l"]:Var()
        local ver = newParameter["v"]:Var()
        local create = newParameter["c"]:Var()

        local out, err = read_file(pathf("~/ypm/boot.tpl"))
        yassert(err)

        local t = tmpl()
        out, err = t:OnceParse(out, {
            version = ver,
            name = module
        })
        if create then
            mkdir(module)
            yassert(cd(module))
        end
        yassert(write_file("boot.lua", out))
        local include_path = pathf("~/lib/include")
        local files, err = ls({
            dir = include_path
        })
        yassert(err)
        if #lang == 0 then
            lang = env.conf.Lang
        end
        mkdir("include")
        ---@diagnostic disable-next-line: param-type-mismatch
        for _, file in ipairs(files) do
            local filename = file[4]

            if path.ext(filename) == ".lua" then
                out, err = read_file(path.join(include_path, filename))
                yassert(err)

                local doc = path.join(include_path, "lang", lang, path.filename(filename) .. ".json")
                if is_exist(doc) then
                    local doc_str, err = read_file(doc)
                    yassert(err)
                    local doc_tbl = json.decode(doc_str)
                    yassert(err)
                    out, err = t:OnceParse(out, doc_tbl)
                end
                write_file("include/" .. filename, out)
            end
        end

        local go_include_path = path.join(include_path, "go")
        path.walk(go_include_path, function(p, info, err)
            if path.ext(p) == ".lua" then
                out, err = read_file(p)
                yassert(err)
                local file = string.sub(p, #go_include_path + 2, #p)
                mkdir(path.join("include/go", path.dir(file)))
                write_file(path.join("include/go", file), out)
            end
            return true
        end)

        safe_write(".gitignore", "/yock_modules\n/include")
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
            default = "1.0",
            usage = ""
        },
        {
            type = flag_type.bool,
            name = "create",
            shorthand = "c",
            default = false,
            usage = ""
        }
    }
}
