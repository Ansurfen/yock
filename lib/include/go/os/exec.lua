-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class exec
---@field ErrWaitDelay any
---@field ErrDot any
---@field ErrNotFound any
---@field ErrNotFound any
---@field ErrNotFound any
---@field ErrNotFound any
exec = {}

--- Command returns the Cmd struct to execute the named program with
--- the given arguments.
---
--- It sets only the Path and Args in the returned structure.
---
--- If name contains no path separators, Command uses LookPath to
--- resolve name to a complete path if possible. Otherwise it uses name
--- directly as Path.
---
--- The returned Cmd's Args field is constructed from the command name
--- followed by the elements of arg, so arg should not include the
--- command name itself. For example, Command("echo", "hello").
--- Args[0] is always name, not the possibly resolved Path.
---
--- On Windows, processes receive the whole command line as a single string
--- and do their own parsing. Command combines and quotes Args into a command
--- line string with an algorithm compatible with applications using
--- CommandLineToArgvW (which is the most common way). Notable exceptions are
--- msiexec.exe and cmd.exe (and thus, all batch files), which have a different
--- unquoting algorithm. In these or other similar cases, you can do the
--- quoting yourself and provide the full command line in SysProcAttr.CmdLine,
--- leaving Args empty.
---@param name string
---@vararg string
---@return execCmd
function exec.Command(name, ...) end

--- CommandContext is like Command but includes a context.
---
--- The provided context is used to interrupt the process
--- (by calling cmd.Cancel or os.Process.Kill)
--- if the context becomes done before the command completes on its own.
---
--- CommandContext sets the command's Cancel function to invoke the Kill method
--- on its Process, and leaves its WaitDelay unset. The caller may change the
--- cancellation behavior by modifying those fields before starting the command.
---@param ctx contextContext
---@param name string
---@vararg string
---@return execCmd
function exec.CommandContext(ctx, name, ...) end

--- LookPath searches for an executable named file in the
--- directories named by the PATH environment variable.
--- LookPath also uses PATHEXT environment variable to match
--- a suitable candidate.
--- If file contains a slash, it is tried directly and the PATH is not consulted.
--- Otherwise, on success, the result is an absolute path.
---
--- In older versions of Go, LookPath could return a path relative to the current directory.
--- As of Go 1.19, LookPath will instead return that path along with an error satisfying
--- errors.Is(err, ErrDot). See the package documentation for more details.
---@param file string
---@return string, err
function exec.LookPath(file) end

--- Cmd represents an external command being prepared or run.
---
--- A Cmd cannot be reused after calling its Run, Output or CombinedOutput
--- methods.
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

--- StderrPipe returns a pipe that will be connected to the command's
--- standard error when the command starts.
---
--- Wait will close the pipe after seeing the command exit, so most callers
--- need not close the pipe themselves. It is thus incorrect to call Wait
--- before all reads from the pipe have completed.
--- For the same reason, it is incorrect to use Run when using StderrPipe.
--- See the StdoutPipe example for idiomatic usage.
---@return any, err
function execCmd:StderrPipe() end

--- Output runs the command and returns its standard output.
--- Any returned error will usually be of type *ExitError.
--- If c.Stderr was nil, Output populates ExitError.Stderr.
---@return byte[], err
function execCmd:Output() end

--- CombinedOutput runs the command and returns its combined standard
--- output and standard error.
---@return byte[], err
function execCmd:CombinedOutput() end

--- String returns a human-readable description of c.
--- It is intended only for debugging.
--- In particular, it is not suitable for use as input to a shell.
--- The output of String may vary across Go releases.
---@return string
function execCmd:String() end

--- StdoutPipe returns a pipe that will be connected to the command's
--- standard output when the command starts.
---
--- Wait will close the pipe after seeing the command exit, so most callers
--- need not close the pipe themselves. It is thus incorrect to call Wait
--- before all reads from the pipe have completed.
--- For the same reason, it is incorrect to call Run when using StdoutPipe.
--- See the example for idiomatic usage.
---@return any, err
function execCmd:StdoutPipe() end

--- Run starts the specified command and waits for it to complete.
---
--- The returned error is nil if the command runs, has no problems
--- copying stdin, stdout, and stderr, and exits with a zero exit
--- status.
---
--- If the command starts but does not complete successfully, the error is of
--- type *ExitError. Other error types may be returned for other situations.
---
--- If the calling goroutine has locked the operating system thread
--- with runtime.LockOSThread and modified any inheritable OS-level
--- thread state (for example, Linux or Plan 9 name spaces), the new
--- process will inherit the caller's thread state.
---@return err
function execCmd:Run() end

--- Start starts the specified command but does not wait for it to complete.
---
--- If Start returns successfully, the c.Process field will be set.
---
--- After a successful call to Start the Wait method must be called in
--- order to release associated system resources.
---@return err
function execCmd:Start() end

--- Wait waits for the command to exit and waits for any copying to
--- stdin or copying from stdout or stderr to complete.
---
--- The command must have been started by Start.
---
--- The returned error is nil if the command runs, has no problems
--- copying stdin, stdout, and stderr, and exits with a zero exit
--- status.
---
--- If the command fails to run or doesn't complete successfully, the
--- error is of type *ExitError. Other error types may be
--- returned for I/O problems.
---
--- If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits
--- for the respective I/O loop copying to or from the process to complete.
---
--- Wait releases any resources associated with the Cmd.
---@return err
function execCmd:Wait() end

--- StdinPipe returns a pipe that will be connected to the command's
--- standard input when the command starts.
--- The pipe will be closed automatically after Wait sees the command exit.
--- A caller need only call Close to force the pipe to close sooner.
--- For example, if the command being run will not exit until standard input
--- is closed, the caller must close the pipe.
---@return any, err
function execCmd:StdinPipe() end

--- Environ returns a copy of the environment in which the command would be run
--- as it is currently configured.
---@return string[]
function execCmd:Environ() end

--- An ExitError reports an unsuccessful exit by a command.
---@class execExitError
---@field Stderr any
local execExitError = {}


---@return string
function execExitError:Error() end

--- Error is returned by LookPath when it fails to classify a file as an
--- executable.
---@class execError
---@field Name string
---@field Err err
local execError = {}


---@return string
function execError:Error() end


---@return err
function execError:Unwrap() end
