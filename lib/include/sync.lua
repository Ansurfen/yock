-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---
---@class waitGroup
---
---{{.wait_group}}
---
local waitGroup = {}

---
---{{.wait_group_add}}
---
---@param delta number
---
function waitGroup:Add(delta)
end

---
---{{.wait_group_done}}
---
function waitGroup:Done()
end

---
---{{.wait_group_wait}}
---
function waitGroup:Wait()
end

---
---{{.sync}}
---
sync = {}

---
---{{.sync_new}}
---
---@return waitGroup
---
function sync.new()
    return {}
end
