--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

prompt({
    desc = {
        use = "ypm",
        short = "ypm is package manager for yock"
    },
    sub = {
        import("cmd/init"),
        import("cmd/add"),
        import("cmd/rm"),
        import("cmd/proxy"),
        import("cmd/install"),
        import("cmd/uninstall"),
        import("cmd/cache"),
        import("cmd/version")
    },
})
