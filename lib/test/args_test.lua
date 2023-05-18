table.dump(env.args)

print(cmdf("yock", "new", "test.lua", "-f", "-t", "null"))

print(env.platform.OS, env.platform.Arch, env.platform.Ver)
job("c", function(cenv)
    print("c")
    table.dump(cenv)
    cenv.c = 6
    table.dump(cenv)
    a = 10
    return true
end)

job("b", function(cenv)
    time.sleep(1 * time.second)
    print("b")
    table.dump(cenv)
    cenv.b = "aaa"
    table.dump(cenv)
    return true
end)

job("d", function(cenv)
    time.sleep(3 * time.second)
    print("d")
    table.dump(cenv)
    return true
end)
