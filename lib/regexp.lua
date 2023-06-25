--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global
---@diagnostic disable: duplicate-set-field

reglib = {}

function reglib:new(patterns)
    local obj = {
        regs = {}
    }
    for name, pattern in pairs(patterns) do
        obj.regs[name] = regexp.MustCompile(pattern)
    end
    setmetatable(obj, { __index = self })
    return obj
end

function reglib:find_str(p, s)
    local str = self.regs[p]:FindStringSubmatch(s)
    if str == nil then
        return nil
    end
    if #str == 1 then
        return str[1]
    end
    return str[2]
end

function reglib:match_str(p, s)
    return self.regs[p]:MatchString(s)
end
