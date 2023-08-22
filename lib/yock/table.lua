--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field
---@see table.dump
function table.dump(tbl, level, filteDefault)
    filteDefault = filteDefault or true --default filter keywords（DeleteMe, _class_type）
    level = level or 1
    local indent_str = ""
    for _ = 1, level do
        indent_str = indent_str .. "  "
    end

    print(indent_str .. "{")
    for k, v in pairs(tbl) do
        if filteDefault then
            if k ~= "_class_type" and k ~= "DeleteMe" then
                local item_str = string.format("%s%s = %s", indent_str .. " ", tostring(k), tostring(v))
                print(item_str)
                if type(v) == "table" then
                    table.dump(v, level + 1)
                end
            end
        else
            local item_str = string.format("%s%s = %s", indent_str .. " ", tostring(k), tostring(v))
            print(item_str)
            if type(v) == "table" then
                table.dump(v, level + 1)
            end
        end
    end
    print(indent_str .. "}")
end

table.clone = function(tbl)
    if tbl == nil then
        return nil
    end
    if type(tbl) ~= "table" then
        return tbl
    end
    local new_tab = {}
    local mt = getmetatable(tbl)
    if mt ~= nil then
        setmetatable(new_tab, mt)
    end
    for i, v in pairs(tbl) do
        if type(v) == "table" then
            new_tab[i] = table.clone(v)
        else
            new_tab[i] = v
        end
    end
    return new_tab
end

table.unpack = unpack

strings.Join = table.concat

path.walk = function(root, fn)
    ls(root, function(path, info)
        return fn(path, info, nil)
    end)
end
