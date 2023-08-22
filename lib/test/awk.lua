job("plain", function(ctx)
    local new, err = awk({
        prog = "{ print $1 + $3 }",
        str = { "1 2 3" }
    })
    yassert(err)
    table.dump(strings.Split(new, "\n"))
end)

job("var", function(ctx)
    local new, err = awk({
        prog = "{ print $1, name }",
        str = { "'Hello World'" },
        var = {
            name = "yock"
        }
    })
    yassert(err)
    table.dump(strings.Split(new, "\n"))
end)
