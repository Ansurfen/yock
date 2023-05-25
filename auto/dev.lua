--[[
require:
    go version 1.20

deploy in develop:
    cd ctl
    go run . run ..\auto\dev.lua
]]
print('start to deploy')
-- TODO: using ".yock-dev" to replace in ".yock" in develop enviroment
-- env.workdir = path.join(env.workdir, "..", ".yock-dev")
mkdir(env.workdir)
local sdk_path = path.join(env.workdir, "sdk")
cp("../sdk/yock", path.join(sdk_path))
cp("../parser/decompose.tpl", path.join(sdk_path, "yock"))
mkdir(path.join(sdk_path, "yock", "deps"))
cp("../sdk/deps/stdlib.json", path.join(sdk_path, "yock", "deps", "stdlib.json"))
cp("./test/include.yaml", path.join(env.workdir, "include.yaml"))
print("deploy finish")
