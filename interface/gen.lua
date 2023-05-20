local proto = import("./proto")
local protoc = proto.protoc

local root = "."
local target = "yock.proto"

job("dart", function(cenv)
    local worksapce = "/dart"
    protoc({
        plugin = "dart",
        out = root .. worksapce,
        proto_path = root,
        proto = target
    })
    return true
end)

job("py", function(cenv)
    local worksapce = "/python"
    protoc({
        plugin = "python",
        out = root .. worksapce,
        grpc_out = root .. worksapce,
        proto_path = root,
        proto = target,
        spec = {
            pyi = root .. worksapce
        }
    })
    return true
end)

job("java", function(cenv)
    local worksapce = "/java/src/main/java/"
    local plugin_path = ""
    protoc({
        plugin = "java",
        plugin_path = plugin_path,
        out = root .. worksapce,
        grpc_out = root .. worksapce,
        proto_path = root,
        proto = target
    })
    return true
end)

job("go", function(cenv)
    local worksapce = "/go"
    protoc({
        plugin = "golang",
        out = root .. worksapce,
        grpc_out = root .. worksapce,
        spec = {
            "--go_opt=paths=source_relative",
            "--go-grpc_opt=paths=source_relative"
        },
        proto = target
    })
    return true
end)

job("c", function(cenv)
    return true
end)

job("csharp", function(cenv)
    return true
end)
