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
    ctx.exit(2)
end)

local SD_LICENSE = [[MIT License
Copyright (c) 2018 Gregory

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.]]

job("tidy", function(ctx)
    local wd = pwd()
    local o  = ctx.flags["os"]

    rm({ safe = false }, "../auto/bin/dist")

    local install_deps = function(os, target)
        cd(pathf("../auto/bin"))
        ctx.set_os(os)
        mkdir("./dist/awk", "./dist/sed", "./dist/grep")
        mkdir(pathf(target, "awk"), pathf(target, "sed"), pathf(target, "grep"))
        for _, v in ipairs(find({ dir = false, pattern = "awk" }, ".")) do
            if strings.Contains(v, env.platform.Arch) and strings.Contains(v, os) then
                local dir    = string.format("./dist/awk/%s", os)
                local _, err = uncompress(v, dir)
                yassert(err)
                local res, err = find({
                    pattern = "goawk" .. ctx.platform:Exf(),
                    dir = false
                }, dir)
                yassert(err)
                for _, value in ipairs(res) do
                    if path.base(value) == "goawk" .. ctx.platform:Exf() then
                        mv(value, pathf(target, "awk"))
                    end
                end
                res, err = find({
                    pattern = "LICENSE.txt",
                    dir = false
                }, dir)
                yassert(err)
                if #res > 0 then
                    mv(res[1], pathf(target, "awk"))
                end
                break
            end
        end
        for _, v in ipairs(find({ dir = false, pattern = "sd" }, ".")) do
            if strings.Contains(v, os) then
                local dir    = string.format("./dist/sed/%s", os)
                local _, err = uncompress(v, dir)
                yassert(err)
                local res, err = find({
                    pattern = "sd" .. ctx.platform:Exf(),
                    dir = false
                }, dir)
                yassert(err)
                for _, value in ipairs(res) do
                    if path.base(value) == "sd" .. ctx.platform:Exf() then
                        mv(value, pathf(target, "sed"))
                    end
                end
                err = write(pathf(target, "sed", "LICENSE"), SD_LICENSE)
                yassert(err)
                break
            end
        end
        for _, v in ipairs(find({ dir = false, pattern = "grep" }, ".")) do
            if strings.Contains(v, os) then
                local dir    = string.format("./dist/grep/%s", os)
                local _, err = uncompress(v, dir)
                yassert(err)
                local res, err = find({
                    pattern = "rg" .. ctx.platform:Exf() .. "$",
                    dir = false
                }, dir)
                yassert(err)
                for _, value in ipairs(res) do
                    if path.base(value) == "rg" .. ctx.platform:Exf() then
                        mv(value, pathf(target, "grep"))
                    end
                end
                res, err = find({
                    pattern = "COPYING",
                    dir = false
                }, dir)
                yassert(err)
                if #res > 0 then
                    mv(res[1], pathf(target, "grep"))
                end
                res, err = find({
                    pattern = "LICENSE-MIT",
                    dir = false
                }, dir)
                yassert(err)
                if #res > 0 then
                    mv(res[1], pathf(target, "grep"))
                end
                res, err = find({
                    pattern = "UNLICENSE",
                    dir = false
                }, dir)
                yassert(err)
                if #res > 0 then
                    mv(res[1], pathf(target, "grep"))
                end
                break
            end
        end
        cd(wd)
    end

    if o == "all" then
        for _, os in ipairs({ "linux", "darwin", "windows" }) do
            install_deps(os, pathf(wd, "../tmp/", os, "yock", "bin"))
        end
    else
        install_deps(o, pathf(yock_path, "bin"))
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
    }, yock_path, pathf("$/../tmp"), pathf("$/../auto/bin/dist"))
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

job("pack", function(ctx)
    local os = ctx.flags["os"]
    if os == "all" then
        for _, o in ipairs({ "linux", "windows", "darwin" }) do
            ctx.set_os(o)
            local target = string.format(pathf("$/../tmp/%s/yock"), o)
            cp({ recurse = true, force = true }, { [yock_path .. "/*"] = target })
            compress(target, pathf("..", string.format("%s-%s%s", o, ctx.flags["v"], ctx.platform:Zip())))
        end
    else
        compress(yock_path, pathf("..", string.format("%s-%s", os, ctx.flags["v"]) .. ctx.platform:Zip()))
    end
    ctx.exit(2)
end)

jobs("all", "build", "tidy", "pack", "remote", "clean")
jobs("alldev", "build", "tidy", "pack", "depoly-dev", "remote", "clean")
jobs("dist", "build")
