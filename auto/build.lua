--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

--[[
require:
    go version 1.20

deploy in develop:
    cd ctl
    ./build.bat/sh dev
]]

print("start to build")

local zip_name = "yock"
local wd, err = pwd()
yassert(err)
local yock_path = pathf(wd, "../yock")
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

    if os == "windows" then
        alias("yock", "../yock/yock.exe")
    else
        alias("yock", "../yock/yock")
    end

    optional({
        case(os == "windows", function()
            ---@diagnostic disable-next-line: param-type-mismatch
            _, err = sh({ debug = true, redirect = true }, [[
go env -w GOOS=windows
go build -o $yock -ldflags "-X 'github.com/ansurfen/yock/util.YockBuild=release'" .]])
        end),
    }, function() -- ? PosixOS: linux, darwin, etc.
        ---@diagnostic disable-next-line: param-type-mismatch
        _, err = sh({ debug = true, redirect = true }, string.format([[
go env -w GOOS=%s
go build -o $yock -ldflags "-X 'github.com/ansurfen/yock/util.YockBuild=release'" .]], os))
    end)
    yassert(err)
    local yock_lib_path = pathf(yock_path, "lib")
    cp(pathf(wd, "../lib"), yock_lib_path)
    cp("install.lua", yock_path)
    mkdir(pathf(yock_path, "ypm"), pathf(yock_lib_path, "boot"))
    cp({ recurse = true, debug = true }, {
        [pathf(wd, "../ypm/ypm.lua")]          = pathf(yock_lib_path, "boot"),
        [pathf(wd, "../ypm/include/ypm.lua")]  = pathf(yock_lib_path, "include"),
        [pathf(wd, "../ypm/boot.tpl")]         = pathf(yock_path, "ypm"),
        [pathf(wd, "../ypm/cmd")]              = pathf(yock_path, "ypm"),
        [pathf(wd, "../ypm/proxy")]            = pathf(yock_path, "ypm"),
        [pathf(wd, "../ypm/ctl.lua")]          = pathf(yock_path, "ypm"),
    })
    rm({ safe = false },
        pathf(yock_lib_path, "test"),
        pathf(yock_lib_path, "bash"),
        pathf(yock_lib_path, "go"),
        pathf(yock_lib_path, "yock"))
    -- sh("$yock run ../auto/bin-tidy.lua")
    -- mv(path.join(wd, "../bin"), path.join(yock_path, "bin"))
    zip_name = assign.string(zip_name, cenv.flags.o)
    if os == "windows" then
        zip_name = zip_name .. ".zip"
    else
        zip_name = zip_name .. ".tar.gz"
    end
    compress(yock_path, pathf("..", zip_name))
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
        string.format([[%s %s]], pathf(wd, "../yock/*"), conf:GetString("default.path")))
    return true
end)

job("clean", function(cenv)
    rm({
        safe = false
    }, yock_path)
    return true
end)

job("remote", function(cenv)
    ssh({
        user = "ubuntu",
        pwd = "root",
        ip = "192.168.127.128",
        network = "tcp",
        redirect = true,
    }, function(s)
        s:Put("../yock.tar.gz", "yock.tar")
        s:Exec("tar -xf yock.tar -C .")
    end)
    -- TODO
    ---@diagnostic disable: undefined-global
    -- sandbox(s, function()
    --     mkdir("/")
    --     tarc("yock.tar", ".")
    -- end)
    return true
end)

jobs("all", "build", "clean", "remote")
jobs("all-dev", "build", "depoly-dev", "clean")
jobs("dist", "build")
