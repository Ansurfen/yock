-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class cpu
---@field physics_core number
---@field logical_core number
---@field info table
cpu = {}

---@param percpu boolean
---@return table,err
function cpu.times(percpu)
    return {}
end

---@param interval number
---@param percpu boolean
---@return table,err
function cpu.percent(interval, percpu)
    return {}
end

disk = {}

---@vararg string
---@return table, err
function disk.info(...)
    return {}
end

---@param all boolean
---@return table, err
function disk.partitions(all)
    return {}
end

---@param path string
---@return table, err
function disk.usage(path)
    return {}
end

mem = {}

---@return table
function mem.info()
    return {}
end

---@return table
function mem.swap()
    return {}
end

host = {}

---@return string
function host.boot_time()
    return ""
end

---@return string, string, string, err
function host.info()
    return "", "", ""
end

net = {}

---@param pernic boolean
---@return table, err
function net.io(pernic)
    return {}
end

---@return table, err
function net.interfaces()
    return {}
end

---@param kind string
---@return table, err
function net.connections(kind)
    return {}
end
