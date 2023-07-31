--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

for pid, _ in pairs(ps("yockd")) do
    kill(pid)
end
