-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

sync = {}

---{{.syncNewCond}}
---@param l syncLocker
---@return syncCond
function sync.NewCond(l)
end

---@class syncOnce
local syncOnce = {}

---{{.syncOnceDo}}
---@param f any
function syncOnce:Do(f)
end

---@class syncPool
---@field New any
local syncPool = {}

---{{.syncPoolPut}}
---@param x any
function syncPool:Put(x)
end

---{{.syncPoolGet}}
---@return any
function syncPool:Get()
end

---@class syncRWMutex
local syncRWMutex = {}

---{{.syncRWMutexRLock}}
function syncRWMutex:RLock()
end

---{{.syncRWMutexTryRLock}}
---@return boolean
function syncRWMutex:TryRLock()
end

---{{.syncRWMutexRUnlock}}
function syncRWMutex:RUnlock()
end

---{{.syncRWMutexLock}}
function syncRWMutex:Lock()
end

---{{.syncRWMutexTryLock}}
---@return boolean
function syncRWMutex:TryLock()
end

---{{.syncRWMutexUnlock}}
function syncRWMutex:Unlock()
end

---{{.syncRWMutexRLocker}}
---@return syncLocker
function syncRWMutex:RLocker()
end

---@class syncWaitGroup
local syncWaitGroup = {}

---{{.syncWaitGroupAdd}}
---@param delta number
function syncWaitGroup:Add(delta)
end

---{{.syncWaitGroupDone}}
function syncWaitGroup:Done()
end

---{{.syncWaitGroupWait}}
function syncWaitGroup:Wait()
end

---@class syncCond
---@field L syncLocker
local syncCond = {}

---{{.syncCondWait}}
function syncCond:Wait()
end

---{{.syncCondSignal}}
function syncCond:Signal()
end

---{{.syncCondBroadcast}}
function syncCond:Broadcast()
end

---@class syncMap
local syncMap = {}

---{{.syncMapLoadOrStore}}
---@param key any
---@param value any
---@return any, boolean
function syncMap:LoadOrStore(key, value)
end

---{{.syncMapDelete}}
---@param key any
function syncMap:Delete(key)
end

---{{.syncMapSwap}}
---@param key any
---@param value any
---@return any, boolean
function syncMap:Swap(key, value)
end

---{{.syncMapCompareAndSwap}}
---@param key any
---@param old any
---@param new any
---@return boolean
function syncMap:CompareAndSwap(key, old, new)
end

---{{.syncMapRange}}
---@param f any
function syncMap:Range(f)
end

---{{.syncMapLoad}}
---@param key any
---@return any, boolean
function syncMap:Load(key)
end

---{{.syncMapStore}}
---@param key any
---@param value any
function syncMap:Store(key, value)
end

---{{.syncMapLoadAndDelete}}
---@param key any
---@return any, boolean
function syncMap:LoadAndDelete(key)
end

---{{.syncMapCompareAndDelete}}
---@param key any
---@param old any
---@return boolean
function syncMap:CompareAndDelete(key, old)
end

---@class syncLocker
local syncLocker = {}

---@class syncMutex
local syncMutex = {}

---{{.syncMutexLock}}
function syncMutex:Lock()
end

---{{.syncMutexTryLock}}
---@return boolean
function syncMutex:TryLock()
end

---{{.syncMutexUnlock}}
function syncMutex:Unlock()
end
