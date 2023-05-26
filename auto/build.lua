print("start to build")

local zip_name = "release"
local wd, err = pwd()
yassert(err)
local yock_path = path.join(wd, "../yock")
mkdir(yock_path)

job("build", function(cenv)
    parse_flags(cenv, {
        o = flag_type.string_type,
        os = flag_type.string_type
    })
    local os = env.platform.OS
    os = assign.string(os, cenv.flags["os"])
    optional({
        case(os == "windows", function()
            err = exec({
                    debug = true,
                    redirect = true
                }, "go env -w GOOS=windows",
                [[go build -o ../yock/yock.exe -ldflags "-X 'github.com/ansurfen/yock/util.YockBuild=release'" .]])
        end),
    }, function() -- ? PosixOS: linux, darwin, etc.
        err = exec({
                debug = true
            }, "go env -w GOOS=linux",
            [[go build -o ../yock/yock -ldflags "-X 'github.com/ansurfen/yock/util.YockBuild=release'" .]])
    end)
    yassert(err)
    local yock_lib_path = path.join(yock_path, "lib")
    cp(path.join(wd, "../lib"), yock_lib_path)
    mkdir(path.join(yock_path, "ypm"))
    mkdir(path.join(yock_lib_path, "boot"))
    cp(path.join(wd, "../ypm/ypm.lua"), path.join(yock_lib_path, "boot"))
    cp(path.join(wd, "../ypm/include/ypm.lua"), path.join(yock_lib_path, "include"))
    cp(path.join(wd, "../yock-todo/ypm/source"), path.join(yock_path, "ypm"))
    rm({
        safe = false
    }, path.join(yock_lib_path, "test"),  path.join(yock_lib_path, "bash"))
    zip_name = assign.string(zip_name, cenv.flags.o)
    err = zip(path.join(wd, "../" .. zip_name .. ".zip"), yock_path)
    yassert(err)
    return true
end)

job("zip", function(cenv)
    return true
end)

job("clean", function(cenv)
    rm({
        safe = false
    }, yock_path)
    return true
end)

jobs("all", "build", "zip")
jobs("dist", "build")
