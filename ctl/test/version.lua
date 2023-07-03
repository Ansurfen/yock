---@diagnostic disable: param-type-mismatch
alias("yock", "yock")

job("default", function(cenv)
    sh([[$yock version]])
    return true
end)

job("logo", function(cenv)
    sh([[$yock version -l]])
    return true
end)
