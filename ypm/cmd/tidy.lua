--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

return {
    desc = { use = "tidy" },
    run = function(cmd, args)
        local proxies, err = find({
            dir = false,
            pattern = "\\.lua"
        }, pathf("#1", "../../proxy"))
        if err ~= nil or #proxies == 0 then
            cp(cat(pathf("#1", "../../template/defaultSource.tpl")), pathf("#1", "../../proxy"))
        end
        cp(pathf("~/lib/include"), pathf("$"))
    end
}
