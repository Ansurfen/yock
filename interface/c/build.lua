--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

sh({
    redirect = true,
    debug = true
}, "gcc ./yock_test.c ./libyock/cJSON.c -L ./libyock -lyock -o yock")
