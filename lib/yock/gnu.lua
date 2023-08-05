-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: param-type-mismatch

---@param opt grep_opt
---@return string, err
grep = function(opt)
    local arg_builder = argBuilder:new():add(path.join(env.yock_bin, "grep", wrapexf("rg")))
    if opt["color"] ~= nil then
        arg_builder:add_str("--color=" .. opt["color"], opt["color"])
    end
    arg_builder:add_bool("-S", opt["case"])
    arg_builder:add_str(opt["pattern"], opt["pattern"])
    local files = {}
    for _, file in ipairs(opt["file"]) do
        table.insert(files, file)
    end
    arg_builder:add(strings.Join(files, " "))
    local res, err = sh({ redirect = false }, arg_builder:build())
    if #res > 0 then
        return res[1], err
    end
    return "", err
end

---@param opt awk_opt
---@return string, err
awk = function(opt)
    local vars = {}
    for k, v in pairs(opt["var"]) do
        table.insert(vars, "-v " .. k .. "=" .. v)
    end
    local progs
    if type(opt["prog"]) == "string" then
        progs = opt["prog"]
    elseif type(opt["prog"]) == "table" then
        for _, prog in ipairs(opt["prog"]) do
            if progs == nil then
                progs = {}
            end
            table.insert(progs, "-f " .. prog)
        end
        progs = strings.Join(progs, " ")
    end
    local files = {}
    for _, file in ipairs(opt["file"]) do
        table.insert(files, file)
    end
    local arg_builder = argBuilder:new():add(path.join(env.yock_bin, "awk", wrapexf("goawk")))
    arg_builder:add(strings.Join(vars, " ")):add(progs):add(strings.Join(files, " "))
    local res, err = sh({ redirect = false }, arg_builder:build())
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
    local files = strings.Join(opt["file"], " ")
    arg_builder:add(files)
    local res, err = sh({ redirect = false }, arg_builder:build())
    if #res > 0 then
        return res[1], err
    end
    return "", err
end
