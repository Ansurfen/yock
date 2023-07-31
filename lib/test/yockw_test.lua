option({
    ycho = {
        stdout = true
    },
    yockw = {
        self_boot = false,
        port = 3376,
        metrics = {
            resolved = {},
            counter = {
                {
                    name = "CPU_Usage",
                    help = "To monitor CPU usage"
                },
                {
                    name = "CPU_Usage",
                    help = "To monitor CPU usage",
                    label = { "type", "rate" }
                },
            },
            gauge = {
                {
                    name = "Memory_Usage",
                    help = "To monitor Memory Usage"
                }
            },
            histogram = {
                {
                    name = "Memory_Usage",
                    help = "To monitor Memory Usage",
                    buckets = { 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10 }
                }
            },
            summary = {
                {
                    name = "Memory_Usage",
                    help = "To monitor Memory Usage",
                    objectives = {
                        ["0.5"] = 0.05,
                        ["0.9"] = 0.01,
                        ["0.99"] = 0.001
                    }
                }
            }
        }
    }
})

yockw.metrics.counter.new({
    name = "CPU Usage",
    help = "To monitor CPU usage"
})
yockw.metrics.counter.add("CPU_Usage", 10)
yockw.metrics.counter.inc("CPU_Usage")
yockw.metrics.counter.rm("CPU_Usage")
table.dump(yockw.metrics.counter.ls())
yockw.metrics.counter_vec.new({
    name = "CPU_Usage2",
    help = "To monitor CPU usage",
    label = { "type", "rate" }
})
yockw.metrics.counter_vec.add("CPU_Usage2", 10, {
    "amd64", "10"
})
yockw.metrics.counter_vec.add("CPU_Usage2", 10, {
    type = "amd64",
    rate = "10",
})
table.dump(yockw.metrics.counter_vec.ls())
