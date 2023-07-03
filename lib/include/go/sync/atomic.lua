-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

atomic = {}

---{{.atomicLoadUint64}}
---@param addr number
---@return number
function atomic.LoadUint64(addr)
end

---{{.atomicAddInt32}}
---@param addr any
---@param delta any
---@return any
function atomic.AddInt32(addr, delta)
end

---{{.atomicStoreInt64}}
---@param addr number
---@param val number
function atomic.StoreInt64(addr, val)
end

---{{.atomicSwapUintptr}}
---@param addr any
---@param new any
---@return any
function atomic.SwapUintptr(addr, new)
end

---{{.atomicLoadInt32}}
---@param addr any
---@return any
function atomic.LoadInt32(addr)
end

---{{.atomicStoreInt32}}
---@param addr any
---@param val any
function atomic.StoreInt32(addr, val)
end

---{{.atomicCompareAndSwapUint32}}
---@param addr any
---@param old any
---@param new any
---@return boolean
function atomic.CompareAndSwapUint32(addr, old, new)
end

---{{.atomicSwapUint32}}
---@param addr any
---@param new any
---@return any
function atomic.SwapUint32(addr, new)
end

---{{.atomicAddUint32}}
---@param addr any
---@param delta any
---@return any
function atomic.AddUint32(addr, delta)
end

---{{.atomicAddInt64}}
---@param addr number
---@param delta number
---@return number
function atomic.AddInt64(addr, delta)
end

---{{.atomicCompareAndSwapUint64}}
---@param addr number
---@param old number
---@param new number
---@return boolean
function atomic.CompareAndSwapUint64(addr, old, new)
end

---{{.atomicCompareAndSwapInt64}}
---@param addr number
---@param old number
---@param new number
---@return boolean
function atomic.CompareAndSwapInt64(addr, old, new)
end

---{{.atomicLoadUint32}}
---@param addr any
---@return any
function atomic.LoadUint32(addr)
end

---{{.atomicCompareAndSwapInt32}}
---@param addr any
---@param old any
---@param new any
---@return boolean
function atomic.CompareAndSwapInt32(addr, old, new)
end

---{{.atomicCompareAndSwapPointer}}
---@param addr unsafePointer
---@param old unsafePointer
---@param new unsafePointer
---@return boolean
function atomic.CompareAndSwapPointer(addr, old, new)
end

---{{.atomicCompareAndSwapUintptr}}
---@param addr any
---@param old any
---@param new any
---@return boolean
function atomic.CompareAndSwapUintptr(addr, old, new)
end

---{{.atomicStoreUintptr}}
---@param addr any
---@param val any
function atomic.StoreUintptr(addr, val)
end

---{{.atomicSwapUint64}}
---@param addr number
---@param new number
---@return number
function atomic.SwapUint64(addr, new)
end

---{{.atomicLoadInt64}}
---@param addr number
---@return number
function atomic.LoadInt64(addr)
end

---{{.atomicStorePointer}}
---@param addr unsafePointer
---@param val unsafePointer
function atomic.StorePointer(addr, val)
end

---{{.atomicSwapPointer}}
---@param addr unsafePointer
---@param new unsafePointer
---@return unsafePointer
function atomic.SwapPointer(addr, new)
end

---{{.atomicSwapInt32}}
---@param addr any
---@param new any
---@return any
function atomic.SwapInt32(addr, new)
end

---{{.atomicStoreUint32}}
---@param addr any
---@param val any
function atomic.StoreUint32(addr, val)
end

---{{.atomicLoadPointer}}
---@param addr unsafePointer
---@return unsafePointer
function atomic.LoadPointer(addr)
end

---{{.atomicSwapInt64}}
---@param addr number
---@param new number
---@return number
function atomic.SwapInt64(addr, new)
end

---{{.atomicLoadUintptr}}
---@param addr any
---@return any
function atomic.LoadUintptr(addr)
end

---{{.atomicStoreUint64}}
---@param addr number
---@param val number
function atomic.StoreUint64(addr, val)
end

---{{.atomicAddUint64}}
---@param addr number
---@param delta number
---@return number
function atomic.AddUint64(addr, delta)
end

---{{.atomicAddUintptr}}
---@param addr any
---@param delta any
---@return any
function atomic.AddUintptr(addr, delta)
end

---@class atomicValue
local atomicValue = {}

---{{.atomicValueLoad}}
---@return any
function atomicValue:Load()
end

---{{.atomicValueStore}}
---@param val any
function atomicValue:Store(val)
end

---{{.atomicValueSwap}}
---@param new any
---@return any
function atomicValue:Swap(new)
end

---{{.atomicValueCompareAndSwap}}
---@param old any
---@param new any
---@return boolean
function atomicValue:CompareAndSwap(old, new)
end

---@class atomicInt64
local atomicInt64 = {}

---{{.atomicInt64Store}}
---@param val number
function atomicInt64:Store(val)
end

---{{.atomicInt64Swap}}
---@param new number
---@return number
function atomicInt64:Swap(new)
end

---{{.atomicInt64CompareAndSwap}}
---@param old number
---@param new number
---@return boolean
function atomicInt64:CompareAndSwap(old, new)
end

---{{.atomicInt64Add}}
---@param delta number
---@return number
function atomicInt64:Add(delta)
end

---{{.atomicInt64Load}}
---@return number
function atomicInt64:Load()
end

---@class atomicUint64
local atomicUint64 = {}

---{{.atomicUint64Load}}
---@return number
function atomicUint64:Load()
end

---{{.atomicUint64Store}}
---@param val number
function atomicUint64:Store(val)
end

---{{.atomicUint64Swap}}
---@param new number
---@return number
function atomicUint64:Swap(new)
end

---{{.atomicUint64CompareAndSwap}}
---@param old number
---@param new number
---@return boolean
function atomicUint64:CompareAndSwap(old, new)
end

---{{.atomicUint64Add}}
---@param delta number
---@return number
function atomicUint64:Add(delta)
end

---@class atomicUintptr
local atomicUintptr = {}

---{{.atomicUintptrLoad}}
---@return any
function atomicUintptr:Load()
end

---{{.atomicUintptrStore}}
---@param val any
function atomicUintptr:Store(val)
end

---{{.atomicUintptrSwap}}
---@param new any
---@return any
function atomicUintptr:Swap(new)
end

---{{.atomicUintptrCompareAndSwap}}
---@param old any
---@param new any
---@return boolean
function atomicUintptr:CompareAndSwap(old, new)
end

---{{.atomicUintptrAdd}}
---@param delta any
---@return any
function atomicUintptr:Add(delta)
end

---@class atomicBool
local atomicBool = {}

---{{.atomicBoolLoad}}
---@return boolean
function atomicBool:Load()
end

---{{.atomicBoolStore}}
---@param val boolean
function atomicBool:Store(val)
end

---{{.atomicBoolSwap}}
---@param new boolean
---@return boolean
function atomicBool:Swap(new)
end

---{{.atomicBoolCompareAndSwap}}
---@param old boolean
---@param new boolean
---@return boolean
function atomicBool:CompareAndSwap(old, new)
end

---@class atomicPointer
local atomicPointer = {}

---@class atomicInt32
local atomicInt32 = {}

---{{.atomicInt32Store}}
---@param val any
function atomicInt32:Store(val)
end

---{{.atomicInt32Swap}}
---@param new any
---@return any
function atomicInt32:Swap(new)
end

---{{.atomicInt32CompareAndSwap}}
---@param old any
---@param new any
---@return boolean
function atomicInt32:CompareAndSwap(old, new)
end

---{{.atomicInt32Add}}
---@param delta any
---@return any
function atomicInt32:Add(delta)
end

---{{.atomicInt32Load}}
---@return any
function atomicInt32:Load()
end

---@class atomicUint32
local atomicUint32 = {}

---{{.atomicUint32Swap}}
---@param new any
---@return any
function atomicUint32:Swap(new)
end

---{{.atomicUint32CompareAndSwap}}
---@param old any
---@param new any
---@return boolean
function atomicUint32:CompareAndSwap(old, new)
end

---{{.atomicUint32Add}}
---@param delta any
---@return any
function atomicUint32:Add(delta)
end

---{{.atomicUint32Load}}
---@return any
function atomicUint32:Load()
end

---{{.atomicUint32Store}}
---@param val any
function atomicUint32:Store(val)
end
