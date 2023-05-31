--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: deprecated
function case(...)
    local args = { ... }
    if #args >= 2 then
        return function()
            local flag, idx = check(unpack(args, 1, #args - 1))
            return flag, idx, args[#args]
        end
    end
    return function()
        return false, -1, nil
    end
end

function optional(cases, bad_case)
    local max = -1
    local fn
    for _, case in ipairs(cases) do
        if type(case) == "function" then
            local flag, idx, f = case()
            if flag and idx > max then
                max = idx
                fn = f
            end
        end
    end
    if max ~= -1 then
        fn()
    else
        bad_case()
    end
end
