option({
    yockw = {
        self_boot = false,
        port = 8080,
    }
})
yockw.metrics.counter.inc("http_handle")
