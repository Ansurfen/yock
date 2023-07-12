--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: lowercase-global

-- require layer
--
-- yock overloads package.path to ensure the validity
-- of relative and absolute path search modules.
package.path = "?.lua"

env.yock_tmp = path.join(env.yock_path, "tmp")
env.yock_bin = path.join(env.yock_path, "bin")

---@param e err|string|nil
---@param msg? any
function yassert(e, msg)
    if e ~= nil then
        if msg ~= nil then
            ycho:Fatal(msg)
        else
            ycho:Fatal(e)
        end
    end
end

---@param file string
---@param data string
write = function(file, data)
    local _, err = write_file(file, data)
    -- local _, err = echo({
    --     fd = { file },
    --     mode = "c|t|rw"
    -- }, data)
    yassert(err)
end
