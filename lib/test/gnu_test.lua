-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: param-type-mismatch

job("echo", function(cenv)
    echo("$GOPATH")
    print(echo("$GOPATH not auto print", false))
    return true
end)

job("default", function(cenv)
    print(whoami())
    print(ls({
        dir = ".",
        str = true
    }))

    local tbl = ls({
        dir = "."
    })
    table.dump(tbl)

    clear()
    -- chmod("main.go", 0777)
    print(pwd())
    cd("..")
    print(pwd())
    print(touch("tmp.txt"))
    print("tmp.txt: ", cat("tmp.txt"))
    rm({
        safe = false
    }, "tmp.txt")
    -- test("pwd", function()

    -- end)
    return true
end)

job("awk", function(cenv)
    awk({
        prog = {
            "../bin/test.awk",
            "../bin/test2.awk"
        },
        file = {
            "awk_test.txt"
        },
        var = {
            name = "ansurfen",
            age = 20
        }
    })
    return true
end)

job("sed", function(cenv)
    local out, err = sed({
        old = "(.*)",
        new = "//$1",
        file = { "t.txt" },
    })
    print(out, err)
    return true
end)

job("grep", function(cenv)
    grep({
        case = true,
        color = "never",
        pattern = "get",
        file = { "awk_test.txt" }
    })
    return true
end)

job("alias", function(cenv)
    alias("CC", "go")
    sh("$CC -v")
    unalias("CC")
    sh("$CC -v")
    return true
end)

job("sudo", function(cenv)
    sudo("go -v")
    return true
end)
