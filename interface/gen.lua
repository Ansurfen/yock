local proto = import("./proto")
local protoc = proto.protoc

local root = "."
local target = "yock.proto"

job("dart", function(cenv)
    local workspace = "/dart"
    protoc({
        plugin = "dart",
        out = root .. workspace,
        proto_path = root,
        proto = target
    })
    return true
end)

job("py", function(cenv)
    local workspace = "/python/yocki"
    protoc({
        plugin = "python",
        out = root .. workspace,
        grpc_out = root .. workspace,
        proto_path = root,
        proto = target,
        spec = {
            pyi = root .. workspace
        }
    })
    return true
end)

job("java", function(cenv)
    local workspace = "/java/src/main/java/"
    local plugin_path = ""
    protoc({
        plugin = "java",
        plugin_path = plugin_path,
        out = root .. workspace,
        grpc_out = root .. workspace,
        proto_path = root,
        proto = target
    })
    return true
end)

job("go", function(cenv)
    local workspace = "/go"
    protoc({
        plugin = "golang",
        out = root .. workspace,
        grpc_out = root .. workspace,
        spec = {
            "--go_opt=paths=source_relative",
            "--go-grpc_opt=paths=source_relative"
        },
        proto = target
    })
    return true
end)

job("c", function(cenv)
    local repo = "https://raw.githubusercontent.com/DaveGamble/cJSON/master/"
    local libs = { "cJSON.c", "cJSON.h" }
    local wd, err = pwd()
    yassert(err)
    local workspace = "/c"
    for _, lib in ipairs(libs) do
        http({
            save = true,
            filename = function(s)
                return path.join(wd, workspace, "libyock", lib)
            end,
            debug = true
        }, repo .. lib)
    end
    workspace = "/c/libyock"
    protoc({
        plugin = "golang",
        out = root .. workspace,
        grpc_out = root .. workspace,
        spec = {
            "--go_opt=paths=source_relative",
            "--go-grpc_opt=paths=source_relative"
        },
        proto = target
    })
    cd("./c/libyock")
    sh({
            debug = true,
            redirect = true
        }, [[./sd.exe 'package \w+' 'package main' .\yock_grpc.pb.go]],
        [[./sd.exe 'package \w+' 'package main' .\yock.pb.go]],
        "yock run gen.lua all")

    return true
end)

job("csharp", function(cenv)
    return true
end)
