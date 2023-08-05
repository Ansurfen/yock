--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "add", short = "Add doc to include directory" },
    run = function(cmd, args)
        if #args == 0 then
            yassert("arguments too little")
        end
        local name = args[1]
        cp(pathf("~/yock_modules", name, "doc"), pathf("$/include", name))
    end
}
