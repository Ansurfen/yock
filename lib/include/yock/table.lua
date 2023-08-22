-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---dump presents data between table on terminal
---
---### Example:
---```lua
--- table.dump({1, 2, 3})
---```
---@param tbl table
function table.dump(tbl) end

---clone returns a table where values is the
---same with received tbl, which is deep copy.
---@return table
function table.clone(tbl) end
