job("plain", function(ctx)
    local res, err = grep({
        pattern = "a",
        str = { "aaaaa", "bcd", "abc" }
    })
    yassert(err)
    print(res)
end)
