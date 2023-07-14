-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

signal = {}

--- Ignored reports whether sig is currently ignored.
---@param sig osSignal
---@return boolean
function signal.Ignored(sig) end

--- Notify causes package signal to relay incoming signals to c.
--- If no signals are provided, all incoming signals will be relayed to c.
--- Otherwise, just the provided signals will.
---
--- Package signal will not block sending to c: the caller must ensure
--- that c has sufficient buffer space to keep up with the expected
--- signal rate. For a channel used for notification of just one signal value,
--- a buffer of size 1 is sufficient.
---
--- It is allowed to call Notify multiple times with the same channel:
--- each call expands the set of signals sent to that channel.
--- The only way to remove signals from the set is to call Stop.
---
--- It is allowed to call Notify multiple times with different channels
--- and the same signals: each channel receives copies of incoming
--- signals independently.
---@param c any
---@vararg any
function signal.Notify(c, ...) end

--- Ignore causes the provided signals to be ignored. If they are received by
--- the program, nothing will happen. Ignore undoes the effect of any prior
--- calls to Notify for the provided signals.
--- If no signals are provided, all incoming signals will be ignored.
---@vararg any
function signal.Ignore(...) end

--- Reset undoes the effect of any prior calls to Notify for the provided
--- signals.
--- If no signals are provided, all signal handlers will be reset.
---@vararg any
function signal.Reset(...) end

--- Stop causes package signal to stop relaying incoming signals to c.
--- It undoes the effect of all prior calls to Notify using c.
--- When Stop returns, it is guaranteed that c will receive no more signals.
---@param c any
function signal.Stop(c) end

--- NotifyContext returns a copy of the parent context that is marked done
--- (its Done channel is closed) when one of the listed signals arrives,
--- when the returned stop function is called, or when the parent context's
--- Done channel is closed, whichever happens first.
---
--- The stop function unregisters the signal behavior, which, like signal.Reset,
--- may restore the default behavior for a given signal. For example, the default
--- behavior of a Go program receiving os.Interrupt is to exit. Calling
--- NotifyContext(parent, os.Interrupt) will change the behavior to cancel
--- the returned context. Future interrupts received will not trigger the default
--- (exit) behavior until the returned stop function is called.
---
--- The stop function releases resources associated with it, so code should
--- call stop as soon as the operations running in this Context complete and
--- signals no longer need to be diverted to the context.
---@param parent contextContext
---@vararg any
---@return contextContext, any
function signal.NotifyContext(parent, ...) end
