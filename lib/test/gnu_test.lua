-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: param-type-mismatch

job("default", function(cenv)
    clear()
    -- chmod("main.go", 0777)
    print(pwd())
    cd("..")
    print(pwd())
    print(touch("tmp.txt"))
    print("tmp.txt: ", cat("tmp.txt"))
    rm({
        safe = false
    }, "tmp.txt")
    return true
end)

job("whoami", function(cenv)
    print(whoami())
    return true
end)

job("ls", function(cenv)
    print(ls({
        dir = ".",
        str = true
    }))
    table.dump(ls("."))
    return true
end)

job("awk", function(cenv)
    awk({
        prog = {
            "../bin/test.awk",
            "../bin/test2.awk"
        },
        file = {
            "awk_test.txt"
        },
        var = {
            name = "ansurfen",
            age = 20
        }
    })
    return true
end)

job("sed", function(cenv)
    local out, err = sed({
        old = "(.*)",
        new = "//$1",
        file = { "t.txt" },
    })
    print(out, err)
    return true
end)

job("grep", function(cenv)
    grep({
        case = true,
        color = "never",
        pattern = "get",
        file = { "awk_test.txt" }
    })
    return true
end)

job("alias", function(cenv)
    alias("CC", "go")
    sh("$CC -v")
    unalias("CC")
    sh("$CC -v")
    return true
end)

job("sudo", function(cenv)
    sudo("go -v")
    return true
end)

job("find", function(cenv)
    print(find("gnu_test.lua"))
    local tbl, err = find({
        pattern = "rg$",
        dir = false
    }, "../../bin")
    yassert(err)
    table.dump(tbl)
    return true
end)

job("echo", function(cenv)
    echo("$GOPATH")
    print(echo("$GOPATH not auto print"))
    write("file.txt", "hello world")
    echo({ fd = { "file.txt" }, mode = "a" }, "Hello World")
    return true
end)

job("ps", function(cenv)
    -- local process = ps({ user = true, mem = true, cpu = true, time = true })
    -- local mapValue = reflect.ValueOf(process)
    -- local iter = mapValue:MapRange()
    -- while iter:Next() do
    --     local v = iter:Value():Interface()
    --     print(v.Mem, v.CPU, v.Start, v.Cmd)
    -- end
    nohup("test.exe -p 9090")
    local procs = pgrep("test")
    for i = 1, #procs, 1 do
        print(procs[i].Pid, procs[i].Name)
    end
    kill("test")
    procs = pgrep("test")
    print(#procs)
    for i = 1, #procs, 1 do
        print(procs[i].Pid, procs[i].Name)
    end
    return true
end)

job("whereis", function(cenv)
    print(whereis("go"))
    return true
end)

job("export", function(cenv)
    export("a", "b")
    export("a:c")
    unset("a")
    return true
end)

job("net", function(cenv)
    ifconfig()
    return true
end)

job("sys-test", function(cenv)
    local service = "TestService"
    local err = systemctl.create(service, {
        service = {
            execStart = "test.exe -p 9090"
        }
    })

    yassert(err)

    local s, err = systemctl.status(service)
    yassert(err)
    print(s:PID(), s:Name(), s:Status())
    -- systemctl.start(service)
    -- systemctl.status(service)
    -- systemctl.stop(service)
    err = systemctl.delete(service)
    yassert(err)
    s, err = systemctl.status(service)
    if s == nil then
        print("删了")
    end
    return true
end)

job("sys-ls", function(cenv)
    local services = systemctl.list("service", "all")
    for _, srv in ipairs(services) do
        print(srv:PID(), srv:Name(), srv:Status())
    end
    return true
end)

job("curl", function(cenv)
    local data, err = curl({}, "")
    yassert(err)
    print(data)
    return true
end)

job("iptables-ls", function(cenv)
    local data, err = iptables.list({
        legacy = true,
        name = ""
    })
    yassert(err)
    for _, v in ipairs(data) do
        print(v:Name(), v:Proto(), v:Action())
    end
    return true
end)

job("iptables-test", function(cenv)
    local data, err = iptables.list({
        legacy = true,
        name = "MyRule"
    })
    if err ~= nil then
        print("not found")
    end
    err = iptables.add({
        name = "MyRule",
        chain = "input",
        protocol = "tcp",
        destination = "8080",
        action = "drop"
    })
    data, err = iptables.list({
        legacy = true,
        name = "MyRule"
    })
    yassert(err)
    for _, v in ipairs(data) do
        print(v:Name(), v:Proto(), v:Action())
    end
    err = iptables.del({
        name = "MyRule",
        chain = "input",
        protocol = "tcp",
        destination = "8080",
        action = "drop"
    })
    yassert(err)
    data, err = iptables.list({
        legacy = true,
        name = "MyRule"
    })
    if err == nil then
        for _, v in ipairs(data) do
            print(v:Name(), v:Proto(), v:Action())
        end
    end
    return true
end)
