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

job_option({
    debug = true,
    strict = true
})

local zip_name = "yock"
local wd, err = pwd()
yassert(err)
local yock_path = pathf(wd, "../yock")
mkdir(yock_path)

job("build", function(ctx)
    argsparse(ctx, {
        o = flag_type.str,   -- release name (output)
        os = flag_type.str,
        ver = flag_type.str, -- release version
    })
    local os = env.platform.OS
    os = assign.string(os, ctx.flags["os"])

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
    mkdir(pathf(yock_path, "ypm"),
        pathf(yock_lib_path, "boot"),
        pathf(yock_lib_path, "yock"),
        pathf(yock_lib_path, "include"),
        pathf(yock_path, "bin"),
        pathf(yock_path, "tmp"),
        pathf(yock_lib_path, "include/ypm"))
    cp({ recurse = true }, {
        ["install.lua"]                       = yock_path,
        ["uninstall.lua"]                     = yock_path,
        [pathf(wd, "../lib/yock")]            = yock_lib_path,
        [pathf(wd, "../lib/include")]         = yock_lib_path,
        [pathf(wd, "../lib/boot/*")]          = pathf(yock_lib_path, "boot"),
        [pathf(wd, "../ypm/ypm.lua")]         = pathf(yock_lib_path, "boot"),
        [pathf(wd, "../ypm/include/ypm.lua")] = pathf(yock_lib_path, "include/ypm"),
        [pathf(wd, "../ypm/template")]        = pathf(yock_path, "ypm"),
        [pathf(wd, "../ypm/cmd")]             = pathf(yock_path, "ypm"),
        [pathf(wd, "../ypm/proxy")]           = pathf(yock_path, "ypm"),
        [pathf(wd, "../ypm/ctl.lua")]         = pathf(yock_path, "ypm"),
        [pathf(wd, "../ypm/util")]            = pathf(yock_path, "ypm"),
        [pathf(wd, "../auto/sudo.bat")]       = pathf(yock_path, "bin")
    })
    rm({ safe = false },
        pathf(yock_lib_path, "test"),
        pathf(yock_lib_path, "bash"),
        pathf(yock_lib_path, "go"))
    -- sh("$yock run ../auto/bin-tidy.lua")
    -- mv(path.join(wd, "../bin"), path.join(yock_path, "bin"))

    zip_name = assign.string(zip_name, ctx.flags.o)
    if os == "windows" then
        zip_name = zip_name .. ".zip"
    else
        zip_name = zip_name .. ".tar.gz"
    end
    compress(yock_path, pathf("..", zip_name))
end)

job("depoly-dev", function(ctx)
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
end)

job("clean", function(ctx)
    rm({
        safe = false
    }, yock_path)
end)

job("remote", function(ctx)
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
end)

jobs("all", "build", "remote", "clean")
jobs("all-dev", "build", "depoly-dev", "clean")
jobs("dist", "build")
