--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

local version = "1.0.0"

return {
    desc = { use = "version", short = "Present ypm's version" },
    run = function(cmd, args)
        print(string.format("version: %s", version))
    end,
}
