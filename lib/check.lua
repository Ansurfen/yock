---@diagnostic disable: lowercase-global
function check(...)
    local args = { ... }
    local flag = true
    local idx = 0
    for _, b in ipairs(args) do
        if not flag then
            break
        end
        if (type(b) == "boolean") then
            flag = flag and b
        else
            if type(b) == "function" then
                local ok = b()
                if type(ok) ~= "boolean" then
                    flag = false
                else
                    flag = flag and ok
                end
            end
        end
        if flag then
            idx = idx + 1
        end
    end
    return flag, idx
end

function OS(want_os, want_ver)
    if want_os ~= env.platform.OS then
        return false
    end
    local cc = require("check")
    return cc.CheckVersion(want_ver, env.platform.Ver)
end

function Windows()
    return OS("windows", "-")
end

function Darwin()
    return OS("darwin", "-")
end

function Linux()
    return OS("linux", "-")
end
