-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: param-type-mismatch

---@param opt grep_opt
---@return string, err
grep = function(opt)
    local arg_builder = argBuilder:new():add(pathf(env.yock_bin, "grep", wrapexf("rg")))
    if opt["color"] ~= nil then
        arg_builder:add_str("--color=" .. opt["color"], opt["color"])
    end
    arg_builder:add_bool("-S", opt["case"])
    arg_builder:add_str(opt["pattern"], opt["pattern"])

    if opt["file"] ~= nil then
        local files = {}
        for _, file in ipairs(opt["file"]) do
            table.insert(files, file)
        end
        arg_builder:add(table.concat(files, " "))
    end

    local echo = ""
    if opt["str"] ~= nil then
        echo = string.format("echo %s | ", table.concat(opt["str"], " "))
    end

    local res, err = sh({ redirect = false, quiet = true }, echo .. arg_builder:build())
    if #res > 0 then
        return res[1], err
    end
    return "", err
end

---@param opt awk_opt
---@return string, err
awk = function(opt)
    local varset = ""
    if opt["var"] ~= nil then
        local vars = {}
        for k, v in pairs(opt["var"]) do
            table.insert(vars, string.format("-v %s=%s", k, v))
        end
        varset = table.concat(vars, " ")
    end

    local progs
    if type(opt["prog"]) == "string" then
        progs = string.format("'%s'", opt["prog"])
    elseif type(opt["prog"]) == "table" then
        for _, prog in ipairs(opt["prog"]) do
            if progs == nil then
                progs = {}
            end
            table.insert(progs, "-f " .. prog)
        end
        progs = table.concat(progs, " ")
    end

    local fileset = ""
    if opt["file"] ~= nil then
        local files = {}
        for _, file in ipairs(opt["file"]) do
            table.insert(files, file)
        end
        fileset = table.concat(files, " ")
    end

    local echo = ""
    if opt["str"] ~= nil then
        echo = string.format("echo %s | ", table.concat(opt["str"], " "))
    end

    local arg_builder = argBuilder:new():add(path.join(env.yock_bin, "awk", wrapexf("goawk")))
    arg_builder:add(varset, progs, fileset)
    local res, err = sh({ redirect = false, quiet = true }, echo .. arg_builder:build())
    if #res > 0 then
        return res[1], err
    end
    return "", err
end

---@param opt sed_opt
---@return string, err
sed = function(opt)
    local arg_builder = argBuilder:new():add(path.join(env.yock_bin, "sed", wrapexf("sd")))
    arg_builder:add_str("-f " .. (opt["f"] or ""), opt["f"])
    arg_builder:add_str(string.format([['%s']], opt["old"]), opt["old"])
    arg_builder:add_str(string.format([['%s']], opt["new"]), opt["new"])
    if #arg_builder.params == 1 then
        arg_builder:add("-h")
    end

    if opt["file"] ~= nil then
        local files = table.concat(opt["file"], " ")
        arg_builder:add(files)
    end

    local echo = ""
    if opt["str"] ~= nil then
        echo = string.format("echo '%s' | ", table.concat(opt["str"], " "))
        arg_builder:add("-s")
    end

    local res, err = sh({ redirect = false, quiet = true }, echo .. arg_builder:build())
    if #res > 0 then
        return res[1], err
    end
    return "", err
end

--TODO, sed -> strings.ReplaceAll, cut -> strings.Cut, strings.Split,  grep -> HasPrefix