--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

argsparse(env, {
    k = flag_type.str,
    v = flag_type.str
})

ypm:open()
ypm:set_cache(env.flags["k"], env.flags["v"])
ypm:close()
