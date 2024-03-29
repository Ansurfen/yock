option({
    sync = false,
    jobs = { "c", "go" },
})

job("_proto_go", function()
    optional({
        case(Windows(), function()

        end)
    })
    print("编译protoc文件")
    notify("job._proto_go")
end)

job("c", function()
    go(function()
        http({
        }, "cJSON.h", "cJSON.c")
        notify("curl finish")
    end)
    wait("job._proto_go")
    print("protoc生成后开始挪动文件")
    local i = 0
    while i ~= 2 do
        print("c, " .. i)
        time.Sleep(5 * time.Second)
        i = i + 1
    end
    waits("job._proto_go", "curl finish")
    notify("c fine")
end)

job("go", function()
    wait("job._proto_go")
    local i = 0
    while i ~= 5 do
        print("go, " .. i)
        time.Sleep(3 * time.Second)
        i = i + 1
    end
    notify("go fine")
end)

job("clean", function()
    waits("go fine", "c fine")
    rm({
        safe = true
    }, "cJSON.h", "cJSON.c")
end)

jobs("cgo", "c", "go")
