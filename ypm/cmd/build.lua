--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "build" },
    run = function(cmd, args)
        local meta = import(pathf("$", "boot"))
        local version = meta.version or ""
        local name = meta.name or ""
        local package = name
        if #version == 0 then
            yassert("unknown version")
        end
        local buildPath = pathf("tmp", strings.ReplaceAll(version, ".", "_"))
        if not find(buildPath) then
            mkdir(buildPath)
        end
        local data, err = cat(".yockignore")
        if err ~= nil then
            write(".yockignore", [[include\**
*.tar.gz
*.zip]])
            data, err = cat(".yockignore")
            yassert(err)
        end
        local rules = { ".git\\**" }
        for _, rule in ipairs(strings.Split(data, "\n")) do
            rule = strings.TrimSpace(rule)
            if #rule ~= 0 and not strings.HasPrefix(rule, "#") then
                table.insert(rules, rule)
            end
        end

        if find(".gitignore") then
            data = cat(".gitignore")
            local idx = strings.LastIndex(data, [[# Power by yock, and NOT EDIT!!!]])
            if idx > 0 then
                data = string.sub(data, 1, idx)
            end
            data = data .. "# Power by yock, and NOT EDIT!!!\n"
            for _, rule in ipairs(rules) do
                if rule == [[.git\**]] then
                    goto continue
                end
                if strings.HasSuffix(rule, "**") then
                    data = data .. "\n" .. string.sub(rule, 1, #rule - 3)
                elseif strings.HasSuffix(rule, "*") then
                    data = data .. "\n" .. string.sub(rule, 1, #rule - 2)
                else
                    data = data .. "\n" .. rule
                end
                ::continue::
            end
            write(".gitignore", data)
        end

        local shouldIgnore = function(file)
            for _, rule in ipairs(rules) do
                local match, err = filepath.Match(rule, file)
                if err ~= nil then
                    goto continue
                end
                if match then
                    return true
                end
                if strings.HasSuffix(rule, "**") then
                    if strings.HasPrefix(file, filepath.Dir(rule) .. string.char(filepath.Separator)) then
                        return true
                    end
                end
                ::continue::
            end
            return false
        end

        local fileList = {}
        path.walk(".", function(p, info, err)
            if err ~= nil then
                return false
            end
            if not info:IsDir() and p ~= "." then
                if not shouldIgnore(p) then
                    table.insert(fileList, p)
                end
            end
            return true
        end)

        for _, file in ipairs(fileList) do
            local target = pathf(buildPath, file)
            mkdir(path.dir(target))
            cp(file, target)
        end
        compress(buildPath, pathf(package .. ".zip"))
        compress(buildPath, pathf(package .. ".tar.gz"))
        rm({ safe = false }, "tmp")
    end
}
