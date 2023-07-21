--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

export(strf("PATH:%s", pathf("~")))
export(strf("PATH:%s", pathf("@/mnt")))
sh("./yock mount ypm ypm/ctl.lua")
chmod(pathf("@/mnt/ypm"), 0775)
