job("plain", function(ctx)
    local new, err = sed({
        old = "((([])))",
        new = " ",
        str = { "lots((([]))) of special chars" }
    })
    yassert(err)
    print(new)
end)
