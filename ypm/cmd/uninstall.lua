--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

argsparse(env, {
    m = flag_type.str,
    g = flag_type.bool,
    wd = flag_type.str
})

ypm:open()
local module = env.flags["m"]
local g = env.flags["g"]
local wd = env.flags["wd"]

if g then
    ypm:rm_module(module)
else
    local fp = jsonfile:open(path.join(wd, "modules.json"))
    fp.buf["dependency"][module] = nil
    fp:write()
    fp:close()
    rm({
        safe = false
    }, path.join(wd, "yock_modules", module))
end

ypm:close()
