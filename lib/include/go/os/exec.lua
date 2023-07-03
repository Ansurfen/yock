-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class exec
---@field ErrWaitDelay any
---@field ErrDot any
---@field ErrNotFound any
exec = {}

---{{.execCommand}}
---@param name string
---@vararg string
---@return execCmd
function exec.Command(name, ...)
end

---{{.execLookPath}}
---@param file string
---@return string, err
function exec.LookPath(file)
end

---{{.execCommandContext}}
---@param ctx contextContext
---@param name string
---@vararg string
---@return execCmd
function exec.CommandContext(ctx, name, ...)
end

---@class execError
---@field Name string
---@field Err err
local execError = {}

---{{.execErrorError}}
---@return string
function execError:Error()
end

---{{.execErrorUnwrap}}
---@return err
function execError:Unwrap()
end

---@class execCmd
---@field Path string
---@field Args any
---@field Env any
---@field Dir string
---@field Stdin any
---@field Stdout any
---@field Stderr any
---@field ExtraFiles any
---@field SysProcAttr any
---@field Process any
---@field ProcessState any
---@field Err err
---@field Cancel any
---@field WaitDelay any
local execCmd = {}

---{{.execCmdStdinPipe}}
---@return any, err
function execCmd:StdinPipe()
end

---{{.execCmdStdoutPipe}}
---@return any, err
function execCmd:StdoutPipe()
end

---{{.execCmdStderrPipe}}
---@return any, err
function execCmd:StderrPipe()
end

---{{.execCmdWait}}
---@return err
function execCmd:Wait()
end

---{{.execCmdRun}}
---@return err
function execCmd:Run()
end

---{{.execCmdOutput}}
---@return byte[], err
function execCmd:Output()
end

---{{.execCmdString}}
---@return string
function execCmd:String()
end

---{{.execCmdEnviron}}
---@return string[]
function execCmd:Environ()
end

---{{.execCmdStart}}
---@return err
function execCmd:Start()
end

---{{.execCmdCombinedOutput}}
---@return byte[], err
function execCmd:CombinedOutput()
end

---@class execExitError
---@field Stderr any
local execExitError = {}

---{{.execExitErrorError}}
---@return string
function execExitError:Error()
end
