--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "rm" },
    run = function(cmd, args)
        if #args == 0 then
            yassert("arguments too little")
        end
        local name = args[1]
        rm({ safe = false }, pathf("$/include", name))
    end,
}
