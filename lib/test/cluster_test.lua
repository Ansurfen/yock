job_option({
    flags = {
        host1 = {
            ip = "localhost"
        },
        host2 = {
            ip = "localhost"
        },
        host3 = {
            ip = "198.162.0.2"
        }
    }
})

job("host1", function(cenv)
    print("host1\n")
    argsparse(cenv, {
        ip = flag_type.str,
    })
    table.dump(cenv)
    optional({
        case(Windows(), function()
            optional({
                case(is_localhost(cenv.flags["ip"]), function()
                    print("localhost")
                    -- exec({
                    --     redirect = true
                    -- }, "go run . run .\\test\\goroutine_test.lua multi ")
                end),
                case(not is_localhost(cenv.flags["ip"]), function()
                    print("ssh")
                end)
            })
        end),
        case(Linux(), function()
            print(777)
        end)
    })
    return true
end)

job("host2", function(cenv)
    print("host2\n")
    argsparse(cenv, {
        ip = flag_type.arr,
    })
    -- table.dump(cenv)
    return true
end)

job("host3", function(cenv)
    print("host3")
    return true
end)
