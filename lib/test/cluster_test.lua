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

job("host1", function(ctx)
    print("host1\n")
    argsparse(ctx, {
        ip = flag_type.str,
    })
    table.dump(ctx)
    optional({
        case(Windows(), function()
            optional({
                case(is_localhost(ctx.flags["ip"]), function()
                    print("localhost")
                    -- exec({
                    --     redirect = true
                    -- }, "go run . run .\\test\\goroutine_test.lua multi ")
                end),
                case(not is_localhost(ctx.flags["ip"]), function()
                    print("ssh")
                end)
            })
        end),
        case(Linux(), function()
            print(777)
        end)
    })
end)

job("host2", function(ctx)
    print("host2\n")
    argsparse(ctx, {
        ip = flag_type.arr,
    })
    -- table.dump(cenv)
end)

job("host3", function(ctx)
    print("host3")
end)
