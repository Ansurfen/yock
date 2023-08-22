local c = conf.open("secret.ini")
local vm = c:read("vm")
if vm ~= nil then
    ssh(vm, function(client)
        client:Get("/home/ubuntu/yock/yock", "yock")
    end)
end
