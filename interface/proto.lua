local version = "v1"

local protoc = function(tbl)
    local is_plugin = function(name)
        if tbl ~= nil and tbl.plugin == name then
            return true
        end
        return false
    end
    local build_protoc = function(...)
        return cmdf("protoc", ...)
    end
    optional({
        case(is_plugin("dart"), function()
            mkdir(tbl["out"])
            sh({
                redirect = true,
                debug = true
            }, build_protoc("--dart_out=grpc:" .. tbl["out"], "-I" .. tbl["proto_path"], tbl["proto"]))
            cd(tbl["out"])
            safe_write("pubspec.yaml", [[
name: yocki
environment:
    sdk: '^2.12.0'
dependencies:
    grpc: ^3.1.0
    protobuf: ^2.1.0]])
        end),
        case(is_plugin("go"), function()

        end),
        case(is_plugin("java"), function()
            sh({
                redirect = true,
                debug = true
            }, build_protoc("--plugin=protoc-gen-grpc-java=" .. tbl["plugin_path"],
                "--grpc-java-out" .. tbl["grpc_out"], "--java_out=" .. tbl["out"], "--proto_path=" .. tbl["proto_path"],
                tbl["proto"]))
        end),
        case(is_plugin("python"), function()
            mkdir(tbl["out"])
            sh({
                redirect = true,
                debug = true,

            }, cmdf("python -m grpc_tools.protoc", "-I" .. tbl["proto_path"],
                "--python_out=" .. tbl["out"], "--pyi_out=" .. tbl["spec"]["pyi"], "--grpc_python_out=" .. tbl
                ["grpc_out"], tbl["proto"]))
        end),
        case(is_plugin("golang"), function()
            mkdir(tbl["out"])
            local spec = ""
            for _, value in ipairs(tbl["spec"]) do
                spec = spec .. value .. " "
            end
            sh({
                redirect = true,
                debug = true,
            }, build_protoc("--go_out=" .. tbl["out"], "--go-grpc_out=" .. tbl["grpc_out"], spec , tbl["proto"]))
        end)
    }, function()
        print("no support the plugin")
    end)
end

return {
    protoc = protoc
}
