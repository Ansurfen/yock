-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

atomic = {}

--- StoreUint64 atomically stores val into *addr.
--- Consider using the more ergonomic and less error-prone [Uint64.Store] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param val number
function atomic.StoreUint64(addr, val) end

--- SwapPointer atomically stores new into *addr and returns the previous *addr value.
--- Consider using the more ergonomic and less error-prone [Pointer.Swap] instead.
---@param addr unsafePointer
---@param new unsafePointer
---@return unsafePointer
function atomic.SwapPointer(addr, new) end

--- CompareAndSwapUint64 executes the compare-and-swap operation for a uint64 value.
--- Consider using the more ergonomic and less error-prone [Uint64.CompareAndSwap] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param old number
---@param new number
---@return boolean
function atomic.CompareAndSwapUint64(addr, old, new) end

--- StoreInt32 atomically stores val into *addr.
--- Consider using the more ergonomic and less error-prone [Int32.Store] instead.
---@param addr any
---@param val any
function atomic.StoreInt32(addr, val) end

--- CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
--- Consider using the more ergonomic and less error-prone [Int32.CompareAndSwap] instead.
---@param addr any
---@param old any
---@param new any
---@return boolean
function atomic.CompareAndSwapInt32(addr, old, new) end

--- AddUint64 atomically adds delta to *addr and returns the new value.
--- To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
--- In particular, to decrement x, do AddUint64(&x, ^uint64(0)).
--- Consider using the more ergonomic and less error-prone [Uint64.Add] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param delta number
---@return number
function atomic.AddUint64(addr, delta) end

--- LoadUint64 atomically loads *addr.
--- Consider using the more ergonomic and less error-prone [Uint64.Load] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@return number
function atomic.LoadUint64(addr) end

--- LoadUint32 atomically loads *addr.
--- Consider using the more ergonomic and less error-prone [Uint32.Load] instead.
---@param addr any
---@return any
function atomic.LoadUint32(addr) end

--- StoreInt64 atomically stores val into *addr.
--- Consider using the more ergonomic and less error-prone [Int64.Store] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param val number
function atomic.StoreInt64(addr, val) end

--- SwapInt64 atomically stores new into *addr and returns the previous *addr value.
--- Consider using the more ergonomic and less error-prone [Int64.Swap] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param new number
---@return number
function atomic.SwapInt64(addr, new) end

--- CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
--- Consider using the more ergonomic and less error-prone [Int64.CompareAndSwap] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param old number
---@param new number
---@return boolean
function atomic.CompareAndSwapInt64(addr, old, new) end

--- LoadPointer atomically loads *addr.
--- Consider using the more ergonomic and less error-prone [Pointer.Load] instead.
---@param addr unsafePointer
---@return unsafePointer
function atomic.LoadPointer(addr) end

--- CompareAndSwapPointer executes the compare-and-swap operation for a unsafe.Pointer value.
--- Consider using the more ergonomic and less error-prone [Pointer.CompareAndSwap] instead.
---@param addr unsafePointer
---@param old unsafePointer
---@param new unsafePointer
---@return boolean
function atomic.CompareAndSwapPointer(addr, old, new) end

--- SwapUint32 atomically stores new into *addr and returns the previous *addr value.
--- Consider using the more ergonomic and less error-prone [Uint32.Swap] instead.
---@param addr any
---@param new any
---@return any
function atomic.SwapUint32(addr, new) end

--- LoadInt32 atomically loads *addr.
--- Consider using the more ergonomic and less error-prone [Int32.Load] instead.
---@param addr any
---@return any
function atomic.LoadInt32(addr) end

--- SwapInt32 atomically stores new into *addr and returns the previous *addr value.
--- Consider using the more ergonomic and less error-prone [Int32.Swap] instead.
---@param addr any
---@param new any
---@return any
function atomic.SwapInt32(addr, new) end

--- LoadUintptr atomically loads *addr.
--- Consider using the more ergonomic and less error-prone [Uintptr.Load] instead.
---@param addr any
---@return any
function atomic.LoadUintptr(addr) end

--- SwapUint64 atomically stores new into *addr and returns the previous *addr value.
--- Consider using the more ergonomic and less error-prone [Uint64.Swap] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param new number
---@return number
function atomic.SwapUint64(addr, new) end

--- AddUintptr atomically adds delta to *addr and returns the new value.
--- Consider using the more ergonomic and less error-prone [Uintptr.Add] instead.
---@param addr any
---@param delta any
---@return any
function atomic.AddUintptr(addr, delta) end

--- LoadInt64 atomically loads *addr.
--- Consider using the more ergonomic and less error-prone [Int64.Load] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@return number
function atomic.LoadInt64(addr) end

--- CompareAndSwapUintptr executes the compare-and-swap operation for a uintptr value.
--- Consider using the more ergonomic and less error-prone [Uintptr.CompareAndSwap] instead.
---@param addr any
---@param old any
---@param new any
---@return boolean
function atomic.CompareAndSwapUintptr(addr, old, new) end

--- AddUint32 atomically adds delta to *addr and returns the new value.
--- To subtract a signed positive constant value c from x, do AddUint32(&x, ^uint32(c-1)).
--- In particular, to decrement x, do AddUint32(&x, ^uint32(0)).
--- Consider using the more ergonomic and less error-prone [Uint32.Add] instead.
---@param addr any
---@param delta any
---@return any
function atomic.AddUint32(addr, delta) end

--- SwapUintptr atomically stores new into *addr and returns the previous *addr value.
--- Consider using the more ergonomic and less error-prone [Uintptr.Swap] instead.
---@param addr any
---@param new any
---@return any
function atomic.SwapUintptr(addr, new) end

--- AddInt64 atomically adds delta to *addr and returns the new value.
--- Consider using the more ergonomic and less error-prone [Int64.Add] instead
--- (particularly if you target 32-bit platforms; see the bugs section).
---@param addr number
---@param delta number
---@return number
function atomic.AddInt64(addr, delta) end

--- StoreUint32 atomically stores val into *addr.
--- Consider using the more ergonomic and less error-prone [Uint32.Store] instead.
---@param addr any
---@param val any
function atomic.StoreUint32(addr, val) end

--- StorePointer atomically stores val into *addr.
--- Consider using the more ergonomic and less error-prone [Pointer.Store] instead.
---@param addr unsafePointer
---@param val unsafePointer
function atomic.StorePointer(addr, val) end

--- CompareAndSwapUint32 executes the compare-and-swap operation for a uint32 value.
--- Consider using the more ergonomic and less error-prone [Uint32.CompareAndSwap] instead.
---@param addr any
---@param old any
---@param new any
---@return boolean
function atomic.CompareAndSwapUint32(addr, old, new) end

--- AddInt32 atomically adds delta to *addr and returns the new value.
--- Consider using the more ergonomic and less error-prone [Int32.Add] instead.
---@param addr any
---@param delta any
---@return any
function atomic.AddInt32(addr, delta) end

--- StoreUintptr atomically stores val into *addr.
--- Consider using the more ergonomic and less error-prone [Uintptr.Store] instead.
---@param addr any
---@param val any
function atomic.StoreUintptr(addr, val) end

--- An Int64 is an atomic int64. The zero value is zero.
---@class atomicInt64
local atomicInt64 = {}

--- Add atomically adds delta to x and returns the new value.
---@param delta number
---@return number
function atomicInt64:Add(delta) end

--- Load atomically loads and returns the value stored in x.
---@return number
function atomicInt64:Load() end

--- Store atomically stores val into x.
---@param val number
function atomicInt64:Store(val) end

--- Swap atomically stores new into x and returns the previous value.
---@param new number
---@return number
function atomicInt64:Swap(new) end

--- CompareAndSwap executes the compare-and-swap operation for x.
---@param old number
---@param new number
---@return boolean
function atomicInt64:CompareAndSwap(old, new) end

--- An Uint32 is an atomic uint32. The zero value is zero.
---@class atomicUint32
local atomicUint32 = {}

--- Swap atomically stores new into x and returns the previous value.
---@param new any
---@return any
function atomicUint32:Swap(new) end

--- CompareAndSwap executes the compare-and-swap operation for x.
---@param old any
---@param new any
---@return boolean
function atomicUint32:CompareAndSwap(old, new) end

--- Add atomically adds delta to x and returns the new value.
---@param delta any
---@return any
function atomicUint32:Add(delta) end

--- Load atomically loads and returns the value stored in x.
---@return any
function atomicUint32:Load() end

--- Store atomically stores val into x.
---@param val any
function atomicUint32:Store(val) end

--- A Value provides an atomic load and store of a consistently typed value.
--- The zero value for a Value returns nil from Load.
--- Once Store has been called, a Value must not be copied.
---
--- A Value must not be copied after first use.
---@class reflectValue
local reflectValue = {}

--- Load returns the value set by the most recent Store.
--- It returns nil if there has been no call to Store for this Value.
---@return any
function reflectValue:Load() end

--- Store sets the value of the Value v to val.
--- All calls to Store for a given Value must use values of the same concrete type.
--- Store of an inconsistent type panics, as does Store(nil).
---@param val any
function reflectValue:Store(val) end

--- Swap stores new into Value and returns the previous value. It returns nil if
--- the Value is empty.
---
--- All calls to Swap for a given Value must use values of the same concrete
--- type. Swap of an inconsistent type panics, as does Swap(nil).
---@param new any
---@return any
function reflectValue:Swap(new) end

--- CompareAndSwap executes the compare-and-swap operation for the Value.
---
--- All calls to CompareAndSwap for a given Value must use values of the same
--- concrete type. CompareAndSwap of an inconsistent type panics, as does
--- CompareAndSwap(old, nil).
---@param old any
---@param new any
---@return boolean
function reflectValue:CompareAndSwap(old, new) end

--- A Pointer is an atomic pointer of type *T. The zero value is a nil *T.
---@class atomicPointer
local atomicPointer = {}

--- An Uint64 is an atomic uint64. The zero value is zero.
---@class atomicUint64
local atomicUint64 = {}

--- Load atomically loads and returns the value stored in x.
---@return number
function atomicUint64:Load() end

--- Store atomically stores val into x.
---@param val number
function atomicUint64:Store(val) end

--- Swap atomically stores new into x and returns the previous value.
---@param new number
---@return number
function atomicUint64:Swap(new) end

--- CompareAndSwap executes the compare-and-swap operation for x.
---@param old number
---@param new number
---@return boolean
function atomicUint64:CompareAndSwap(old, new) end

--- Add atomically adds delta to x and returns the new value.
---@param delta number
---@return number
function atomicUint64:Add(delta) end

--- An Uintptr is an atomic uintptr. The zero value is zero.
---@class atomicUintptr
local atomicUintptr = {}

--- Swap atomically stores new into x and returns the previous value.
---@param new any
---@return any
function atomicUintptr:Swap(new) end

--- CompareAndSwap executes the compare-and-swap operation for x.
---@param old any
---@param new any
---@return boolean
function atomicUintptr:CompareAndSwap(old, new) end

--- Add atomically adds delta to x and returns the new value.
---@param delta any
---@return any
function atomicUintptr:Add(delta) end

--- Load atomically loads and returns the value stored in x.
---@return any
function atomicUintptr:Load() end

--- Store atomically stores val into x.
---@param val any
function atomicUintptr:Store(val) end

--- A Bool is an atomic boolean value.
--- The zero value is false.
---@class atomicBool
local atomicBool = {}

--- CompareAndSwap executes the compare-and-swap operation for the boolean value x.
---@param old boolean
---@param new boolean
---@return boolean
function atomicBool:CompareAndSwap(old, new) end

--- Load atomically loads and returns the value stored in x.
---@return boolean
function atomicBool:Load() end

--- Store atomically stores val into x.
---@param val boolean
function atomicBool:Store(val) end

--- Swap atomically stores new into x and returns the previous value.
---@param new boolean
---@return boolean
function atomicBool:Swap(new) end

--- An Int32 is an atomic int32. The zero value is zero.
---@class atomicInt32
local atomicInt32 = {}

--- Swap atomically stores new into x and returns the previous value.
---@param new any
---@return any
function atomicInt32:Swap(new) end

--- CompareAndSwap executes the compare-and-swap operation for x.
---@param old any
---@param new any
---@return boolean
function atomicInt32:CompareAndSwap(old, new) end

--- Add atomically adds delta to x and returns the new value.
---@param delta any
---@return any
function atomicInt32:Add(delta) end

--- Load atomically loads and returns the value stored in x.
---@return any
function atomicInt32:Load() end

--- Store atomically stores val into x.
---@param val any
function atomicInt32:Store(val) end
