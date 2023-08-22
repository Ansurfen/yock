-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class wait_group
local wait_group = {}

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
function wait_group:Add(delta) end

--- Done decrements the WaitGroup counter by one.
function wait_group:Done() end

--- Wait blocks until the WaitGroup counter is zero.
function wait_group:Wait() end

sync = {}

---A WaitGroup waits for a collection of goroutines to finish.
---The main goroutine calls Add to set the number of goroutines to wait for.
---Then each of the goroutines runs and calls Done when finished.
---At the same time, Wait can be used to block until all goroutines have finished.
---@return wait_group
function sync.wait_group() end

---A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
---@return syncMutex
function sync.mutex() end

---@class chan
---@field send fun(self: chan, v: any)
---@field receive fun(self: chan): boolean, any
---@field close fun(self: chan)

---@class channel
---@field make fun(self: channel): chan
---@field select fun(self: channel, ...)
channel = {}

---make returns a new chan
---@return chan
function channel.make() end

---select just like golang's select grammatical sugar,
---which is designed for handling chan operation. Therefore,
---every case must be a available chan operation, send or receive.
---select will listen all message from given case and execute
---corresponding function. If there are multiple chan is ready at the moment,
---select will choose to execute with random. On the contrary, if none are ready,
---select will execute default case when it exists.
---### Syntax:
---* {"|<-", chan, fun(ok: bool, v: any)}, callback when message was received
---* {"<-|", chan, v: any}, send message from chan
---* {"default", fun()}, if none are ready, execute default case
---
---### Example:
---```lua
---local ch = channel.make()
---go(function()
---    while true do
---        channel.select({ "|<-", ch, function(ok, v) -- wait message from chan
---            print(ok, v)
---        end }, { "default", function() -- if none, execute default case
---            print("do nothing")
---            time.Sleep(500 * time.Millisecond)
---        end })
---    end
---end)
---go(function()
---    local i = 1
---    while true do
---        channel.select({ "<-|", ch, i }) -- continuous send i as message per second
---        i = i + 1
---        time.Sleep(1 * time.Second)
---    end
---end)
---time.Sleep(20 * time.Second)
---```
---@vararg any
function channel.select(...) end
