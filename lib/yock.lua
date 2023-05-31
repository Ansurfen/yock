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

---@param e err
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
