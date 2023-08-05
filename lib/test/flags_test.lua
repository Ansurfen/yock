-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

env.args = { "-a", "hello", "-c", "-b", "a", "-e", "1", "-b", "b" }

argsparse(env, {
    a = flag_type.str,
    b = flag_type.arr,
    c = flag_type.bool,
    d = flag_type.bool,
    e = flag_type.num
})

table.dump(env)
