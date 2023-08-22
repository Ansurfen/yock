-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-doc-field

---@meta _

---@class json
json = {}

---encode marshals table into json string
---### Example:
---```lua
---print(json.encode({ 1, 2, 3 })) -- input: [1,2,3]
---
---print(json.encode({ a = 10, b = { c = "d" } })) -- input: {"a":10,"b":{"c":"d"}}
---```
---@param v any
---@vararg string
---@return string
function json.encode(v, ...) end

---decode unmarshals json string to table
---@param str string
---@return table
function json.decode(str) end

---create opens json file to be specified and returns json_object
---and create json file when it don't exist. The second parameter
---indicates content to write file when create.
---### Example:
---```lua
---local jf = json.create("./test.json", "{}")
---```
---@param file string
---@param str? string
---@return json_object
function json.create(file, str) end

---open must opens an existed file. if not, it'll panic.
---@param file string
---@return json_object
function json.open(file) end

---unmarshals json string to object and returns
---@param str string
---@return json_object
function json.from_str(str) end

---@class json_object
---@field buf table
---@field file string
json_object = {}

---@param k string
---@return boolean
function json_object:getbool(k) end

---@param k string
---@return number
function json_object:getnumber(k) end

---@param k string
---@return string
function json_object:getstr(k) end

---@param k string
---@return table
function json_object:gettable(k) end

---get could visit value by json path.
---### Example:
---```lua
---local jf = json.from_str([[{ "a" = { "b" = 10 } }]])
---print(jf:get("a.b")) -- input: 10
---```
---@param k string
---@return any
function json_object:get(k) end

---rawget returns value according to key.
---It's different with `get`, and not any handling.
---@param k string
---@return any
function json_object:rawget(k) end

---set could set value by json path.
---### Example:
---```lua
---local jf = json.from_str([[{ "a": { "b": 10 } }]])
---jf:set("a.b", 11)
---```
---@param k string
---@param v any
function json_object:set(k, v) end

---rawget sets value according to key.
---It's different with `set`, and not any handling.
---@param k string
---@param v any
function json_object:rawset(k, v) end

---string returns json string by self.
---@return string
function json_object:string() end

---save persists json into file to be specified.
---In general, it's the same with calling `json.open` or
---`json.create` to indicate. You also could reset
---file field to change it. Pretty is optional, and it's
---false in default and will formats json string when be
---set true.
---
---### Example:
---```lua
---local jf = json.open("./test.json")
---jf:save(true)
---# to begin
---{"a": {"b": 10}}
---# nowadays
---{
---    "a": {
---        "b": 10
---    }
---}
---```
---@param pretty? boolean
function json_object:save(pretty) end
