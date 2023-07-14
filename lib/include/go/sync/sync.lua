-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

sync = {}

--- NewCond returns a new Cond with Locker l.
---@param l syncLocker
---@return syncCond
function sync.NewCond(l) end

--- Once is an object that will perform exactly one action.
---
--- A Once must not be copied after first use.
---
--- In the terminology of the Go memory model,
--- the return from f “synchronizes before”
--- the return from any call of once.Do(f).
---@class syncOnce
local syncOnce = {}

--- Do calls the function f if and only if Do is being called for the
--- first time for this instance of Once. In other words, given
---
---	var once Once
---
--- if once.Do(f) is called multiple times, only the first call will invoke f,
--- even if f has a different value in each invocation. A new instance of
--- Once is required for each function to execute.
---
--- Do is intended for initialization that must be run exactly once. Since f
--- is niladic, it may be necessary to use a function literal to capture the
--- arguments to a function to be invoked by Do:
---
---	config.once.Do(func() { config.init(filename) })
---
--- Because no call to Do returns until the one call to f returns, if f causes
--- Do to be called, it will deadlock.
---
--- If f panics, Do considers it to have returned; future calls of Do return
--- without calling f.
---@param f any
function syncOnce:Do(f) end

--- A Pool is a set of temporary objects that may be individually saved and
--- retrieved.
---
--- Any item stored in the Pool may be removed automatically at any time without
--- notification. If the Pool holds the only reference when this happens, the
--- item might be deallocated.
---
--- A Pool is safe for use by multiple goroutines simultaneously.
---
--- Pool's purpose is to cache allocated but unused items for later reuse,
--- relieving pressure on the garbage collector. That is, it makes it easy to
--- build efficient, thread-safe free lists. However, it is not suitable for all
--- free lists.
---
--- An appropriate use of a Pool is to manage a group of temporary items
--- silently shared among and potentially reused by concurrent independent
--- clients of a package. Pool provides a way to amortize allocation overhead
--- across many clients.
---
--- An example of good use of a Pool is in the fmt package, which maintains a
--- dynamically-sized store of temporary output buffers. The store scales under
--- load (when many goroutines are actively printing) and shrinks when
--- quiescent.
---
--- On the other hand, a free list maintained as part of a short-lived object is
--- not a suitable use for a Pool, since the overhead does not amortize well in
--- that scenario. It is more efficient to have such objects implement their own
--- free list.
---
--- A Pool must not be copied after first use.
---
--- In the terminology of the Go memory model, a call to Put(x) “synchronizes before”
--- a call to Get returning that same value x.
--- Similarly, a call to New returning x “synchronizes before”
--- a call to Get returning that same value x.
---@class syncPool
---@field New any
local syncPool = {}

--- Put adds x to the pool.
---@param x any
function syncPool:Put(x) end

--- Get selects an arbitrary item from the Pool, removes it from the
--- Pool, and returns it to the caller.
--- Get may choose to ignore the pool and treat it as empty.
--- Callers should not assume any relation between values passed to Put and
--- the values returned by Get.
---
--- If Get would otherwise return nil and p.New is non-nil, Get returns
--- the result of calling p.New.
---@return any
function syncPool:Get() end

--- A RWMutex is a reader/writer mutual exclusion lock.
--- The lock can be held by an arbitrary number of readers or a single writer.
--- The zero value for a RWMutex is an unlocked mutex.
---
--- A RWMutex must not be copied after first use.
---
--- If a goroutine holds a RWMutex for reading and another goroutine might
--- call Lock, no goroutine should expect to be able to acquire a read lock
--- until the initial read lock is released. In particular, this prohibits
--- recursive read locking. This is to ensure that the lock eventually becomes
--- available; a blocked Lock call excludes new readers from acquiring the
--- lock.
---
--- In the terminology of the Go memory model,
--- the n'th call to Unlock “synchronizes before” the m'th call to Lock
--- for any n < m, just as for Mutex.
--- For any call to RLock, there exists an n such that
--- the n'th call to Unlock “synchronizes before” that call to RLock,
--- and the corresponding call to RUnlock “synchronizes before”
--- the n+1'th call to Lock.
---@class syncRWMutex
local syncRWMutex = {}

--- RUnlock undoes a single RLock call;
--- it does not affect other simultaneous readers.
--- It is a run-time error if rw is not locked for reading
--- on entry to RUnlock.
function syncRWMutex:RUnlock() end

--- Lock locks rw for writing.
--- If the lock is already locked for reading or writing,
--- Lock blocks until the lock is available.
function syncRWMutex:Lock() end

--- TryLock tries to lock rw for writing and reports whether it succeeded.
---
--- Note that while correct uses of TryLock do exist, they are rare,
--- and use of TryLock is often a sign of a deeper problem
--- in a particular use of mutexes.
---@return boolean
function syncRWMutex:TryLock() end

--- Unlock unlocks rw for writing. It is a run-time error if rw is
--- not locked for writing on entry to Unlock.
---
--- As with Mutexes, a locked RWMutex is not associated with a particular
--- goroutine. One goroutine may RLock (Lock) a RWMutex and then
--- arrange for another goroutine to RUnlock (Unlock) it.
function syncRWMutex:Unlock() end

--- RLocker returns a Locker interface that implements
--- the Lock and Unlock methods by calling rw.RLock and rw.RUnlock.
---@return syncLocker
function syncRWMutex:RLocker() end

--- RLock locks rw for reading.
---
--- It should not be used for recursive read locking; a blocked Lock
--- call excludes new readers from acquiring the lock. See the
--- documentation on the RWMutex type.
function syncRWMutex:RLock() end

--- TryRLock tries to lock rw for reading and reports whether it succeeded.
---
--- Note that while correct uses of TryRLock do exist, they are rare,
--- and use of TryRLock is often a sign of a deeper problem
--- in a particular use of mutexes.
---@return boolean
function syncRWMutex:TryRLock() end

--- A WaitGroup waits for a collection of goroutines to finish.
--- The main goroutine calls Add to set the number of
--- goroutines to wait for. Then each of the goroutines
--- runs and calls Done when finished. At the same time,
--- Wait can be used to block until all goroutines have finished.
---
--- A WaitGroup must not be copied after first use.
---
--- In the terminology of the Go memory model, a call to Done
--- “synchronizes before” the return of any Wait call that it unblocks.
---@class syncWaitGroup
local syncWaitGroup = {}

--- Add adds delta, which may be negative, to the WaitGroup counter.
--- If the counter becomes zero, all goroutines blocked on Wait are released.
--- If the counter goes negative, Add panics.
---
--- Note that calls with a positive delta that occur when the counter is zero
--- must happen before a Wait. Calls with a negative delta, or calls with a
--- positive delta that start when the counter is greater than zero, may happen
--- at any time.
--- Typically this means the calls to Add should execute before the statement
--- creating the goroutine or other event to be waited for.
--- If a WaitGroup is reused to wait for several independent sets of events,
--- new Add calls must happen after all previous Wait calls have returned.
--- See the WaitGroup example.
---@param delta number
function syncWaitGroup:Add(delta) end

--- Done decrements the WaitGroup counter by one.
function syncWaitGroup:Done() end

--- Wait blocks until the WaitGroup counter is zero.
function syncWaitGroup:Wait() end

--- Cond implements a condition variable, a rendezvous point
--- for goroutines waiting for or announcing the occurrence
--- of an event.
---
--- Each Cond has an associated Locker L (often a *Mutex or *RWMutex),
--- which must be held when changing the condition and
--- when calling the Wait method.
---
--- A Cond must not be copied after first use.
---
--- In the terminology of the Go memory model, Cond arranges that
--- a call to Broadcast or Signal “synchronizes before” any Wait call
--- that it unblocks.
---
--- For many simple use cases, users will be better off using channels than a
--- Cond (Broadcast corresponds to closing a channel, and Signal corresponds to
--- sending on a channel).
---
--- For more on replacements for sync.Cond, see [Roberto Clapis's series on
--- advanced concurrency patterns], as well as [Bryan Mills's talk on concurrency
--- patterns].
---
--- [Roberto Clapis's series on advanced concurrency patterns]: https://blogtitle.github.io/categories/concurrency/
--- [Bryan Mills's talk on concurrency patterns]: https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view
---@class syncCond
---@field L syncLocker
local syncCond = {}

--- Wait atomically unlocks c.L and suspends execution
--- of the calling goroutine. After later resuming execution,
--- Wait locks c.L before returning. Unlike in other systems,
--- Wait cannot return unless awoken by Broadcast or Signal.
---
--- Because c.L is not locked while Wait is waiting, the caller
--- typically cannot assume that the condition is true when
--- Wait returns. Instead, the caller should Wait in a loop:
---
---	c.L.Lock()
---	for !condition() {
---	    c.Wait()
---	}
---	... make use of condition ...
---	c.L.Unlock()
function syncCond:Wait() end

--- Signal wakes one goroutine waiting on c, if there is any.
---
--- It is allowed but not required for the caller to hold c.L
--- during the call.
---
--- Signal() does not affect goroutine scheduling priority; if other goroutines
--- are attempting to lock c.L, they may be awoken before a "waiting" goroutine.
function syncCond:Signal() end

--- Broadcast wakes all goroutines waiting on c.
---
--- It is allowed but not required for the caller to hold c.L
--- during the call.
function syncCond:Broadcast() end

--- Map is like a Go map[interface{}]interface{} but is safe for concurrent use
--- by multiple goroutines without additional locking or coordination.
--- Loads, stores, and deletes run in amortized constant time.
---
--- The Map type is specialized. Most code should use a plain Go map instead,
--- with separate locking or coordination, for better type safety and to make it
--- easier to maintain other invariants along with the map content.
---
--- The Map type is optimized for two common use cases: (1) when the entry for a given
--- key is only ever written once but read many times, as in caches that only grow,
--- or (2) when multiple goroutines read, write, and overwrite entries for disjoint
--- sets of keys. In these two cases, use of a Map may significantly reduce lock
--- contention compared to a Go map paired with a separate Mutex or RWMutex.
---
--- The zero Map is empty and ready for use. A Map must not be copied after first use.
---
--- In the terminology of the Go memory model, Map arranges that a write operation
--- “synchronizes before” any read operation that observes the effect of the write, where
--- read and write operations are defined as follows.
--- Load, LoadAndDelete, LoadOrStore, Swap, CompareAndSwap, and CompareAndDelete
--- are read operations; Delete, LoadAndDelete, Store, and Swap are write operations;
--- LoadOrStore is a write operation when it returns loaded set to false;
--- CompareAndSwap is a write operation when it returns swapped set to true;
--- and CompareAndDelete is a write operation when it returns deleted set to true.
---@class syncMap
local syncMap = {}

--- CompareAndSwap swaps the old and new values for key
--- if the value stored in the map is equal to old.
--- The old value must be of a comparable type.
---@param key any
---@param old any
---@param new any
---@return boolean
function syncMap:CompareAndSwap(key, old, new) end

--- Store sets the value for a key.
---@param key any
---@param value any
function syncMap:Store(key, value) end

--- LoadOrStore returns the existing value for the key if present.
--- Otherwise, it stores and returns the given value.
--- The loaded result is true if the value was loaded, false if stored.
---@param key any
---@param value any
---@return any, boolean
function syncMap:LoadOrStore(key, value) end

--- LoadAndDelete deletes the value for a key, returning the previous value if any.
--- The loaded result reports whether the key was present.
---@param key any
---@return any, boolean
function syncMap:LoadAndDelete(key) end

--- Delete deletes the value for a key.
---@param key any
function syncMap:Delete(key) end

--- Swap swaps the value for a key and returns the previous value if any.
--- The loaded result reports whether the key was present.
---@param key any
---@param value any
---@return any, boolean
function syncMap:Swap(key, value) end

--- CompareAndDelete deletes the entry for key if its value is equal to old.
--- The old value must be of a comparable type.
---
--- If there is no current value for key in the map, CompareAndDelete
--- returns false (even if the old value is the nil interface value).
---@param key any
---@param old any
---@return boolean
function syncMap:CompareAndDelete(key, old) end

--- Range calls f sequentially for each key and value present in the map.
--- If f returns false, range stops the iteration.
---
--- Range does not necessarily correspond to any consistent snapshot of the Map's
--- contents: no key will be visited more than once, but if the value for any key
--- is stored or deleted concurrently (including by f), Range may reflect any
--- mapping for that key from any point during the Range call. Range does not
--- block other methods on the receiver; even f itself may call any method on m.
---
--- Range may be O(N) with the number of elements in the map even if f returns
--- false after a constant number of calls.
---@param f any
function syncMap:Range(f) end

--- Load returns the value stored in the map for a key, or nil if no
--- value is present.
--- The ok result indicates whether value was found in the map.
---@param key any
---@return any, boolean
function syncMap:Load(key) end

--- A Mutex is a mutual exclusion lock.
--- The zero value for a Mutex is an unlocked mutex.
---
--- A Mutex must not be copied after first use.
---
--- In the terminology of the Go memory model,
--- the n'th call to Unlock “synchronizes before” the m'th call to Lock
--- for any n < m.
--- A successful call to TryLock is equivalent to a call to Lock.
--- A failed call to TryLock does not establish any “synchronizes before”
--- relation at all.
---@class syncMutex
local syncMutex = {}

--- Unlock unlocks m.
--- It is a run-time error if m is not locked on entry to Unlock.
---
--- A locked Mutex is not associated with a particular goroutine.
--- It is allowed for one goroutine to lock a Mutex and then
--- arrange for another goroutine to unlock it.
function syncMutex:Unlock() end

--- Lock locks m.
--- If the lock is already in use, the calling goroutine
--- blocks until the mutex is available.
function syncMutex:Lock() end

--- TryLock tries to lock m and reports whether it succeeded.
---
--- Note that while correct uses of TryLock do exist, they are rare,
--- and use of TryLock is often a sign of a deeper problem
--- in a particular use of mutexes.
---@return boolean
function syncMutex:TryLock() end

--- A Locker represents an object that can be locked and unlocked.
---@class syncLocker
local syncLocker = {}
