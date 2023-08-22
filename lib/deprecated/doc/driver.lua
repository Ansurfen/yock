-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

-- dns, plugin, and driver are all derivatives of the dependency analysis pattern.
-- They are now abandoned, see pack/dependency.go for details.

---@meta _

---@param tbl table
function pull(tbl) end

---@param plugin string
---@return string,string
function parse_plugin(plugin) end

---@param file string
---@return string
function export_builder(file) end

---@class pluginlist
plugin_list = {}

---@param path string
---@return boolean
function plugin_list:IsExist(path) end

---@param pid string
---@param path string
function plugin_list:AddPlugin(pid, path) end

---@class exportOpt
---@field update fun()
---@field install fun()
---@field uninstall fun()
---@field init fun(env: any)
local exportOpt = {}

---@param opt exportOpt
function export(opt) end

---@param opt table
---@vararg string
function installs(opt, ...) end

---@param plugin string
---@param opt table
function install(plugin, opt) end

---@param file string
---@return string
function load_plugin(file) end

plugins = {}

---@param opt table
function plugin(opt) end

---@class ldns
ldns = {}

---@param domain string
function ldns:GetDriver(domain) end

---@param domain string
function ldns:GetPlugin(domain) end

---@param domain string
---@param url string
---@param path string
function ldns:PutPlugin(domain, url, path) end

---@param domain string
---@param url string
---@param path string
function ldns:PutDriver(domain, url, path) end

---@param domain string
---@param alias string
function ldns:AliasDriver(domain, alias) end

---@param domain string
---@param alias string
function ldns:AliasPlugin(domain, alias) end

---@class gdns
gdns = {}

---@param domain string
function gdns:GetDriver(domain) end

---@param domain string
---@return table
function gdns:GetPlugin(domain) end

---@param domain string
---@param url string
---@param path string
function gdns:UpdatePlugin(domain, url, path) end

---@param callback fun(...): ...:any
function driver(callback) end

---@param fn string
---@return fun(opt: table, ...:any)
function uninit_driver(fn)
    ---@param opt table
    ---@vararg any
    return function(opt, ...) end
end

---@param driver string
---@param name string
---@return string
function set_driver(driver, name) end

---@param did string
---@param opt table
---@vararg any
function exec_driver(did, opt, ...) end
