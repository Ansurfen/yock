--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

option({
    strict = true,
    ycho = {
        stdout = true
    }
})

rm({ safe = false }, pathf("@"))
