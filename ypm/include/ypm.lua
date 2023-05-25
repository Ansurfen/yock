---@meta _

---@class YPM
ypm = {}

function ypm:open()

end

---@return string
function ypm:read()
    return ""
end

function ypm:write()

end

---@param opt table
---@return function
function ymodule(opt)
    return function()
    end
end
