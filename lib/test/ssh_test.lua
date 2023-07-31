ssh({
    user = "ubuntu",
    pwd = "root",
    ip = "192.168.127.128",
    port = 22,
    network = "tcp",
    redirect = true,
}, function(client)
    print(client:OS())
    -- client:Shell()
    -- client:Exec({ "echo 'root' | sudo -S ls" })
end)

-- s:Get("release.tar", "a.tar")
-- s:Put("../../yock.tar", "yock.tar")
