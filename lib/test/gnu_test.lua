echo("$GOPATH")
print(echo("$GOPATH not auto print", false))
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
