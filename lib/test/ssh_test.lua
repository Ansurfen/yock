local s = ssh({
    user = "ubuntu",
    pwd = "root",
    ip = "192.168.127.128",
    network = "tcp",
    redirect = true,
}, function(client)
    -- client:Shell()
    client:Exec({ "echo 'root' | sudo -S ls" })
end)
