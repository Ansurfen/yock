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
end)

job("csharp", function(cenv)
end)
