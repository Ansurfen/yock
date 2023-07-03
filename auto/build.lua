--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

--[[
require:
    go version 1.20

deploy in develop:
    cd ctl
    ./build.bat/sh
]]

---@diagnostic disable: param-type-mismatch

print("start to build")

local zip_name = "release"
local wd, err = pwd()
yassert(err)
local yock_path = path.join(wd, "../yock")
mkdir(yock_path)

job_option({

})

job("build", function(cenv)
    argsparse(cenv, {
        o = flag_type.str,   -- release name (output)
        os = flag_type.str,
        ver = flag_type.str, -- release version
    })
    local os = env.platform.OS
    os = assign.string(os, cenv.flags["os"])
    optional({
        case(os == "windows", function()
            ---@diagnostic disable-next-line: param-type-mismatch
            _, err = sh({ debug = true, redirect = true }, [[
go env -w GOOS=windows
go build -o ../yock/yock.exe -ldflags "-X 'github.com/ansurfen/yock/util.YockBuild=release'" .]])
        end),
    }, function() -- ? PosixOS: linux, darwin, etc.
        ---@diagnostic disable-next-line: param-type-mismatch
        _, err = sh({ debug = true, redirect = true }, [[
go env -w GOOS=linux
go build -o ../yock/yock -ldflags "-X 'github.com/ansurfen/yock/util.YockBuild=release'" .]])
    end)
    yassert(err)
    local yock_lib_path = path.join(yock_path, "lib")
    cp(path.join(wd, "../lib"), yock_lib_path)
    mkdir(path.join(yock_path, "ypm"), path.join(yock_lib_path, "boot"))
    ---@diagnostic disable-next-line: param-type-mismatch
    cp({ recurse = true, debug = true }, {
        [path.join(wd, "../ypm/ypm.lua")] = path.join(yock_lib_path, "boot"),
        [path.join(wd, "../ypm/include/ypm.lua")] = path.join(yock_lib_path, "include"),
        [path.join(wd, "../yock-todo/ypm/source")] = path.join(yock_path, "ypm"),
        [path.join(wd, "../ypm/boot.tpl")] = path.join(yock_path, "ypm")
    })
    rm({ safe = false },
        path.join(yock_lib_path, "test"),
        path.join(yock_lib_path, "bash"),
        path.join(yock_lib_path, "go"),
        path.join(yock_lib_path, "yock"))
    zip_name = assign.string(zip_name, cenv.flags.o)
    if os == "windows" then
        zip_name = zip_name .. ".zip"
    else
        zip_name = zip_name .. ".tar.gz"
    end
    compress(yock_path, zip_name)
    return true
end)

job("depoly-dev", function(cenv)
    local conf, err = open_conf("secret.ini")
    if err ~= nil then
        write_file("secret.ini", "path = ")
        print("please set path in secret.ini")
        yassert(err)
    end
    local p = conf:GetString("default.path")
    if #p == 0 then
        yassert("path not set")
    end
    cp({ force = true, debug = true, redirect = true },
        string.format([[%s %s]], path.join(wd, "../yock/*"), conf:GetString("default.path")))
    return true
end)

job("clean", function(cenv)
    rm({
        safe = false
    }, yock_path)
    return true
end)

jobs("all", "build", "clean")
jobs("all-dev", "build", "depoly-dev", "clean")
jobs("dist", "build")
