option({
    yockd = {
        self_boot = false,
        port = 1314,
        name = "master"
    },
    ycho = {
        stdout = true
    },
    yockw = {
        self_boot = false,
        port = 8080,
    }
})

yockw.metrics.counter.new({
    name = "http_handle",
    help = "listen http"
})

yockw.metrics.counter.new({
    name = "file_handle",
    help = "listen file modify"
})

local wd = pwd()
local pid, err = yockd.process.spawn("cron", "*/1 * * * *", strf("yock run %s\\http.lua", wd))
if err ~= nil then
    yassert(err:Error())
end
print(pid)
pid, err = yockd.process.spawn("fs", pathf(wd, "tmp"), strf("yock run %s\\fs.lua", wd))
if err ~= nil then
    yassert(err:Error())
end
print(pid)

table.dump(yockd.process.list())
