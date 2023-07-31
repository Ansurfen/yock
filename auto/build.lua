--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: missing-fields

--[[
require:
    go version 1.20

deploy in develop:
    cd ctl
    ./build.bat/sh dev
]]

option({
    ycho = {
        stdout = true
    },
    yockw = {
        self_boot = false
    },
    strict = false,
    sync = false
})

local zip_name = "yock"
local yock_path = pathf("$/../yock")
if not find(yock_path) then
    mkdir(yock_path)
end

job("build", function(ctx)
    argsparse(ctx, {
        o = flag_type.str, -- release name (output)
        os = flag_type.str,
        v = flag_type.str, -- release version
        r = flag_type.str,
    })
    local os = env.platform.OS
    os = assign.string(os, ctx.flags["os"])
    if ctx.flags["v"] == "0" then
        ctx.throw([[no version is specified
try add v2.0.0 parameter to mark version
]])
    end
    alias("ver", ctx.flags["v"])

    local build_cmd = [[
go env -w GOOS=$os
go build -o $yockctl -ldflags "-X 'github.com/ansurfen/yock/util.YockBuild=release' -X 'github.com/ansurfen/yock/util.YockVersion=$ver'" .
go build -o $yockd ../daemon/yockd.go
go build -o $yockw ../watch/main.go]]

    if os == "all" then
        rm({ safe = false }, "../tmp")
        for _, o in ipairs({ "linux", "windows", "darwin" }) do
            ctx.set_os(o)
            alias("os", o)
            alias("yockctl", string.format("../tmp/%s/yock/yock" .. ctx.platform:Exf(), o))
            alias("yockd", string.format("../tmp/%s/yock/yockd" .. ctx.platform:Exf(), o))
            alias("yockw", string.format("../tmp/%s/yock/yockw" .. ctx.platform:Exf(), o))

            local _, err = sh({ redirect = true }, build_cmd)

            yassert(err)
        end
    else
        ctx.set_os(os)
        alias("os", os)
        alias("yockctl", "../yock/yock" .. ctx.platform:Exf())
        alias("yockd", "../yock/yockd" .. ctx.platform:Exf())
        alias("yockw", "../yock/yockw" .. ctx.platform:Exf())
        local _, err = sh({ redirect = true }, build_cmd)

        yassert(err)
    end

    local yock_lib_path = pathf(yock_path, "lib")
    mkdir(pathf(yock_path, "ypm"),
        pathf(yock_lib_path, "boot"),
        pathf(yock_lib_path, "yock"),
        pathf(yock_lib_path, "sdk"),
        pathf(yock_lib_path, "include"),
        pathf(yock_path, "bin"),
        pathf(yock_path, "tmp"),
        pathf(yock_lib_path, "include/ypm"))
    cp({ recurse = true, force = true }, {
        ["install.lua"]                     = yock_path,
        ["uninstall.lua"]                   = yock_path,
        [pathf("$/../lib/yock")]            = yock_lib_path,
        [pathf("$/../lib/include")]         = yock_lib_path,
        [pathf("$/../lib/boot/*")]          = pathf(yock_lib_path, "boot"),
        [pathf("$/../ypm/ypm.lua")]         = pathf(yock_lib_path, "boot"),
        [pathf("$/../ypm/include/ypm.lua")] = pathf(yock_lib_path, "include/ypm"),
        [pathf("$/../ypm/*")]               = pathf(yock_path, "ypm"),
        [pathf("$/../auto/sudo.bat")]       = pathf(yock_path, "bin"),
        [pathf("$/../interface/python")]    = pathf(yock_path, "sdk/python")
    })
    rm({ safe = false },
        pathf(yock_lib_path, "test"),
        pathf(yock_lib_path, "bash"),
        pathf(yock_lib_path, "go"),
        pathf(yock_path, "ypm/ypm.lua"))

    -- sh("$yock run ../auto/bin-tidy.lua")
    -- mv(path.join(wd, "../bin"), path.join(yock_path, "bin"))
    if os == "all" then
        for _, o in ipairs({ "linux", "windows", "darwin" }) do
            ctx.set_os(o)
            local target = string.format(pathf("$/../tmp/%s/yock"), o)
            cp(yock_path .. "/*", target)
            compress(target, pathf("..", string.format("%s-%s%s", o, ctx.flags["v"], ctx.platform:Zip())))
        end
    else
        zip_name = assign.string(zip_name, ctx.flags.o)
        compress(yock_path, pathf("..", zip_name .. ctx.platform:Zip()))
    end
    ctx.exit(2)
end)

job("depoly-dev", function(ctx)
    local secret = conf.create("secret.ini", "path = ")
    local p = secret:read("default.path")
    if #p == 0 then
        yassert("path not set")
    end
    for pid, _ in pairs(ps("yockd")) do
        kill(pid)
    end
    for pid, _ in pairs(ps("yockw")) do
        kill(pid)
    end
    cp({ force = true },
        string.format([[%s %s]], pathf("$/../yock/*"), secret:read("default.path")))
    ctx.exit(2)
end)

job("clean", function(ctx)
    rm({
        safe = false
    }, yock_path, pathf("$/../tmp"))
end)

job("remote", function(ctx)
    if ctx.flags["r"] ~= "1" then
        ctx.exit(1)
    end
    local c = conf.create("../auto/secret.ini", [[
[vm]
user =
pwd =
ip =
port = 22
network = "tcp"
redirect = true
enable = true
]])
    go(function()
        local cloud = c:read("cloud")
        if cloud ~= nil and cloud["enable"] then
            ssh(cloud, function(s)
                local remote_lua = random.str(8) .. ".lua"
                local tmp_zip = random.str(8)
                s:Put("../auto/remote.lua", remote_lua)
                local o = strings.ToLower(s:OS())
                if strings.Contains(o, "linux") then
                    s:Put(string.format("../linux-%s.tar.gz", ctx.flags["v"]), tmp_zip .. ".tar")
                elseif strings.Contains(o, "windows") then
                    s:Put(string.format("../windows-%s.zip", ctx.flags["v"]), tmp_zip .. ".zip")
                elseif strings.Contains(o, "darwin") then
                    s:Put(string.format("../darwin-%s.tar.gz", ctx.flags["v"]), tmp_zip .. ".tar")
                end
                s:Sh("../auto/remote.sh", tmp_zip, remote_lua)
            end)
        end
        notify("cloud")
    end)
    local vm = c:read("vm")
    if vm ~= nil and vm["enable"] then
        ssh(vm, function(s)
            local remote_lua = random.str(8) .. ".lua"
            local tmp_zip = random.str(8)
            s:Put("../auto/remote.lua", remote_lua)
            local o = strings.ToLower(s:OS())
            if strings.Contains(o, "linux") then
                s:Put(string.format("../linux-%s.tar.gz", ctx.flags["v"]), tmp_zip .. ".tar")
            elseif strings.Contains(o, "windows") then
                s:Put(string.format("../windows-%s.zip", ctx.flags["v"]), tmp_zip .. ".zip")
            elseif strings.Contains(o, "darwin") then
                s:Put(string.format("../darwin-%s.tar.gz", ctx.flags["v"]), tmp_zip .. ".tar")
            end
            s:Sh("../auto/remote.sh", tmp_zip, remote_lua)
        end)
    end

    wait("cloud")
    -- TODO
    ---@diagnostic disable: undefined-global
    -- sandbox(s, function()
    --     mkdir("/")
    --     tarc("yock.tar", ".")
    -- end)
end)

jobs("all", "build", "remote", "clean")
jobs("alldev", "build", "depoly-dev", "remote", "clean")
jobs("dist", "build")
