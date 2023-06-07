local protoc = function(...)
    return sh({
        redirect = true
    }, cmdf("protoc", ...))
end

local root = "../dameon"
local worksapce = "/interface"
local target = "yock.proto"

protoc("--go_out=" .. root .. worksapce,
    "--go_opt=paths=source_relative",
    "--go-grpc_out=" .. root .. worksapce,
    "--go-grpc_opt=paths=source_relative",
    target)
