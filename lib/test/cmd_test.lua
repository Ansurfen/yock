---@diagnostic disable: param-type-mismatch
table.dump(sh([[
    echo Hello
    echo World
]]))

table.dump(sh({
    redirect = false,
    debug = true
}, "go version", "go version"))

table.dump(sh({
    redirect = false,
    debug = true
}, [[
    go version
    go version
]], "go version"))