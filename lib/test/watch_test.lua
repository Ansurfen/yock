-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

job("cpu", function(cenv)
    table.dump(cpu.percent(3 * time.Second, true))
    table.dump(cpu.times(true))
    table.dump(cpu.times(false))
    table.dump(cpu)
end)

job("disk", function(cenv)
    table.dump(disk.info())
    table.dump(disk.partitions(true))
    table.dump(disk.usage("D:"))
end)

job("mem", function(cenv)
    table.dump(mem.swap())
    table.dump(mem.info())
end)

job("host", function(cenv)
    print(host.boot_time())
    print(host.info())
end)

job("net", function(cenv)
    table.dump(net.io(true))
    table.dump(net.interfaces())
end)

jobs("all", "cpu", "disk", "mem", "host", "net")
