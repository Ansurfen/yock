--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

prompt({
    desc = {
        use = "ypm",
        short = "ypm is package manager for yock"
    },
    sub = {
        import("./new"),
        import("./rm"),
        import("./install"),
        import("./uninstall")
    },
    flags = {}
})
