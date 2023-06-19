local grep = function(pattern, file)
    local out, err = sh(cmdf("rg", "-N", [["]] .. pattern .. [["]], file))
    if err ~= nil then
        return ""
    end
    return out
end

print(grep("(\\{|j)", "yock.json"))
