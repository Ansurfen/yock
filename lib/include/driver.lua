-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

-- dns, plugin, and driver are all derivatives of the dependency analysis pattern.
-- They are now abandoned, see pack/dependency.go for details.

---@meta _

---{{.pull}}
---@param tbl table
function pull(tbl)
end

---@param plugin string
---@return string,string
function parse_plugin(plugin)
    return "", ""
end

---@param file string
---@return string
function export_builder(file)
    return ""
end

---@class pluginlist
plugin_list = {}

---@param path string
---@return boolean
function plugin_list:IsExist(path)
    return false
end

---@param pid string
---@param path string
function plugin_list:AddPlugin(pid, path)
end

---@class exportOpt
---@field update fun()
---@field install fun()
---@field uninstall fun()
---@field init fun(env: any)
local exportOpt = {}

---
---{{.export}}
---
---@param opt exportOpt
---
function export(opt)
end

---@param opt table
---@vararg string
function installs(opt, ...)
end

---@param plugin string
---@param opt table
function install(plugin, opt)
end

---@param file string
---@return string
function load_plugin(file)
    return ""
end

plugins = {}

---@param opt table
function plugin(opt)
end

---
---{{.ldns}}
---
---@class ldns
---
ldns = {}

---
---{{.lsdn_get_driver}}
---
---@param domain string
---
function ldns:GetDriver(domain)
end

---
---{{.ldns_get_plugin}}
---
---@param domain string
---
function ldns:GetPlugin(domain)
end

---
---{{.lsdn_put_plugin}}
---
---@param domain string
---@param url string
---@param path string
---
function ldns:PutPlugin(domain, url, path)
end

---
---{{.lsdn_put_driver}}
---
---@param domain string
---@param url string
---@param path string
---
function ldns:PutDriver(domain, url, path)
end

---
---{{.lsdn_alias_driver}}
---
---@param domain string
---@param alias string
---
function ldns:AliasDriver(domain, alias)
end

---
---{{.lsdn_alias_plugin}}
---
---@param domain string
---@param alias string
---
function ldns:AliasPlugin(domain, alias)
end

---
---{{.gdns}}
---
---@class gdns
---
gdns = {}

---
---{{.gdns_get_driver}}
---
---@param domain string
---
function gdns:GetDriver(domain)
end

---
---{{.gdns_get_plugin}}
---
---@param domain string
---@return table
---
function gdns:GetPlugin(domain)
    return {}
end

---@param domain string
---@param url string
---@param path string
function gdns:UpdatePlugin(domain, url, path)
end

---
---{{.driver}}
---
---@param callback fun(...): ...:any
---
function driver(callback)
end

---
---{{.uninit_driver}}
---
---@param fn string
---@return fun(opt: table, ...:any)
---
function uninit_driver(fn)
    ---@param opt table
    ---@vararg any
    return function(opt, ...)

    end
end

---
---{{.set_driver}}
---
---@param driver string
---@param name string
---@return string
---
function set_driver(driver, name)
    return ""
end

---
---{{.exec_driver}}
---
---@param did string
---@param opt table
---@vararg any
---
function exec_driver(did, opt, ...)
end
