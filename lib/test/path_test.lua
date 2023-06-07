print(path.filename("./abc.lua"))
print(path.exist("./test/undefined"))
print(path.exist("./test"))
path.walk("../lib", function(path, info, err)
    print(path)
    return true
end)
