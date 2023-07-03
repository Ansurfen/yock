local proto = import("./proto")
local protoc = proto.protoc

local root = "."
local target = "yocki.proto"

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
    local workspace = "/c/libyock"
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
