-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: param-type-mismatch

job("default", function(ctx)
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
end)

job("whoami", function(ctx)
    print(whoami())
end)

job("ls", function(ctx)
    print(ls({
        dir = ".",
        str = true
    }))
    table.dump(ls("."))
end)

job("awk", function(ctx)
    local res, err = awk({
        prog = {
            "./static/test.awk",
            "./static/test2.awk"
        },
        file = {
            "./static/awk_test.txt"
        },
        var = {
            name = "ansurfen",
            age = 20
        }
    })
    yassert(err)
    print(res)
end)

job("sed", function(ctx)
    local out, err = sed({
        old = "(.*)",
        new = "//$1",
        file = { "./static/sed_test.txt" },
    })
    print(out, err)
end)

job("grep", function(ctx)
    local res, err = grep({
        case = true,
        color = "never",
        pattern = "get",
        file = { "./static/awk_test.txt" }
    })
    yassert(err)
    print(res)
end)

job("alias", function(ctx)
    alias("CC", "go")
    sh("$CC -v")
    unalias("CC")
    sh("$CC -v")
end)

job("sudo", function(ctx)
    sudo("go -v")
end)

job("find", function(ctx)
    print(find("gnu_test.lua"))
    local tbl, err = find({
        pattern = "rg$",
        dir = false
    }, "../../bin")
    yassert(err)
    table.dump(tbl)
end)

job("echo", function(ctx)
    echo("$GOPATH")
    print(echo("$GOPATH not auto print"))
    write("file.txt", "hello world")
    echo({ fd = { "file.txt" }, mode = "a" }, "Hello World")
end)

job("nohup", function(ctx)
    yassert(nohup("test.exe -p 9090"))
    for _, info in ipairs(lsof(9090)) do
        kill(info.pid)
    end
    table.dump(lsof(9090))
end)

job("ps", function(ctx)
    for pid, info in pairs(ps()) do
        print(pid, info.cmd, info.name)
    end
    print("num of cmd with vscode: ", #ps("vscode"))
    print("num of process of which the pid is 1: ", #ps(1))
end)

job("pgrep", function(ctx)
    for _, proc in ipairs(pgrep("vscode")) do
        print(proc.name, proc.pid)
    end
end)

job("whereis", function(ctx)
    print(whereis("go"))
end)

job("export", function(ctx)
    export("a", "b")
    export("a:c")
    unset("a")
end)

job("net", function(ctx)
    table.dump(ifconfig())
end)

job("testService", function(ctx)
    local service = "TestService"
    local err = systemctl.create(service, {
        service = {
            execStart = "test.exe -p 9090"
        }
    })
    yassert(err)
    local s, err = systemctl.status(service)
    yassert(err)
    print("PID", "Name", "Status")
    print(s.pid, s.name, s.status)
    -- systemctl.start(service)
    -- systemctl.status(service)
    -- systemctl.stop(service)
    err = systemctl.delete(service)
    yassert(err)
    s, err = systemctl.status(service)
    yassert(s == nil, function()
        print(strf("%s was deleted", service))
    end)
end)

job("services", function(ctx)
    local services = systemctl.list("service", "all")
    for _, srv in ipairs(services) do
        print(srv.pid, srv.name, srv.status)
    end
end)

job("curl", function(ctx)
    local data, err = curl({}, "")
    yassert(err)
    print(data)
end)

job("iptlist", function(ctx)
    local ruels, err = iptables.list({
        legacy = true,
        name = ""
    })
    yassert(err, function()
        for _, rule in ipairs(ruels) do
            print(rule.name, rule.proto, rule.action)
        end
    end)
end)

job("ipttest", function(ctx)
    local data, err = iptables.list({
        legacy = true,
        name = "MyRule"
    })
    yassert(#data == 0, function()
        print("not found rule")
        err = iptables.add({
            name = "MyRule",
            chain = "input",
            protocol = "tcp",
            destination = "8080",
            action = "drop"
        })
        yassert(err, function()
            print("add rule...")
        end)
    end)
    data, err = iptables.list({
        legacy = true,
        name = "MyRule"
    })
    yassert(err == nil, function()
        print("rule:", data.name, data.proto, data.action, data.src, data.dst)
    end)
    err = iptables.del({
        name = "MyRule",
        chain = "input",
        protocol = "tcp",
        destination = "8080",
        action = "drop"
    })
    yassert(err, function()
        print("delete rule...")
    end)
    data, err = iptables.list({
        legacy = true,
        name = "MyRule"
    })
    yassert(err ~= nil or #data == 0, function()
        print("rule was deleted")
    end)
end)

job("lsof", function(ctx)
    print("PID", "Proto", "State", "Local")
    for _, info in ipairs(lsof(58838)) do
        print(info.pid, info.proto, info.state, info.Local)
    end
    table.dump(ps(15764))
    table.dump(ps("python"))
    kill(15764)
end)

job("read", function(ctx)
    read("name")
    sh([[echo "Hello $name"]])
end)
