print("checkpoint one")

job("build1", function(cenv)
    print("build1")
    return true
end)

print("checkpoint two")

job("build2", function(cenv)
    print("build2")
    return true
end)

job("deploy", function(cenv)
    print("deploy")
    return true
end)

job("clean", function(cenv)
    print("clean")
    return true
end)

jobs("all", "build1", "build2", "clean", "deploy")
jobs("pony", "clean", "deploy")
-- yok compile -c all -c build2
