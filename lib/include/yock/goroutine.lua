-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---
---{{.go}}
---
---@param callback fun()
---@async
---
function go(callback)
end

---
---{{.wait}}
---
---@param sig string
---
function wait(sig)
end

---
---{{.waits}}
---
---@vararg string
---
function waits(...)
end

---
---{{.notify}}
---
---@param sig string
---
function notify(sig)
end
