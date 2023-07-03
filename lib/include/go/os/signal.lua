-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

signal = {}

---{{.signalNotifyContext}}
---@param parent contextContext
---@vararg any
---@return contextContext, any
function signal.NotifyContext(parent, ...)
end

---{{.signalNotify}}
---@param c any
---@vararg any
function signal.Notify(c, ...)
end

---{{.signalReset}}
---@vararg any
function signal.Reset(...)
end

---{{.signalStop}}
---@param c any
function signal.Stop(c)
end

---{{.signalIgnore}}
---@vararg any
function signal.Ignore(...)
end

---{{.signalIgnored}}
---@param sig osSignal
---@return boolean
function signal.Ignored(sig)
end
