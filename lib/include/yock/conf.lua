-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---conf can open and parse configuration with easy and fast way,
---supporting yaml, yml, json, toml, hcl, tfvars, ini, properties,
---props, prop, dotenv, env file format.
---@class conf
---@field buf table # saves objected data, and can visit by table
---@field viper Viper # Viper is a prioritized configuration registry. It maintains a set of configuration sources, fetches values to populate those, and provides them according to the source's priority.
---@field open fun(file: string): conf # must open specified file, otherwise it would panic.
---@field create fun(file: string, tmpl: string): conf # create returns parsed conf object, and creates a new file and panics when file isn't exist.
---@field read fun(self: conf, k: string): table|nil # read returns value by json path
---@field write fun(self: conf, k: string, v: any) # writes v to specified k, and required to call the save function for persisting on configuration file.
---@field save fun(self: conf) # save persists data based-on memory into configuration.
conf = {}

---create returns parsed conf object, and creates a new file and panics when file isn't exist.
---@param file string
---@param tmpl string
---@return conf
function conf.create(file, tmpl) end

---open must open specified file, otherwise it would panic.
---@param file string
---@return conf
function conf.open(file) end

---read returns value by json path
---@param k string
---@return table|nil
function conf:read(k) end

---writes v to specified k, and required to call the save function for persisting on configuration file.
---@param k string
---@param v any
function conf:write(k, v) end

---save persists data based-on memory into configuration.
function conf:save() end