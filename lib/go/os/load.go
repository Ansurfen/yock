// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oslib

import (
	"os"

	yocki "github.com/ansurfen/yock/interface"
	execlib "github.com/ansurfen/yock/lib/go/os/exec"
	userlib "github.com/ansurfen/yock/lib/go/os/user"
)

func LoadOs(yocks yocki.YockScheduler) {
	userlib.LoadUser(yocks)
	execlib.LoadExec(yocks)
	lib := yocks.OpenLib("os")
	lib.SetField(map[string]any{
		// functions
		"Getpagesize":       os.Getpagesize,
		"Getppid":           os.Getppid,
		"Chtimes":           os.Chtimes,
		"SameFile":          os.SameFile,
		"Getpid":            os.Getpid,
		"IsPermission":      os.IsPermission,
		"WriteFile":         os.WriteFile,
		"Lstat":             os.Lstat,
		"Lchown":            os.Lchown,
		"Getgid":            os.Getgid,
		"IsPathSeparator":   os.IsPathSeparator,
		"Expand":            os.Expand,
		"UserCacheDir":      os.UserCacheDir,
		"Symlink":           os.Symlink,
		"Getuid":            os.Getuid,
		"Pipe":              os.Pipe,
		"Exit":              os.Exit,
		"IsTimeout":         os.IsTimeout,
		"FindProcess":       os.FindProcess,
		"StartProcess":      os.StartProcess,
		"Link":              os.Link,
		"Getegid":           os.Getegid,
		"NewSyscallError":   os.NewSyscallError,
		"DirFS":             os.DirFS,
		"Stat":              os.Stat,
		"Readlink":          os.Readlink,
		"Getgroups":         os.Getgroups,
		"Geteuid":           os.Geteuid,
		"Executable":        os.Executable,
		"UserConfigDir":     os.UserConfigDir,
		"Truncate":          os.Truncate,
		"PathSeparator":     os.PathSeparator,
		"PathListSeparator": os.PathListSeparator,
		"ModeDir":           os.ModeDir,
		"ModeAppend":        os.ModeAppend,
		"ModeExclusive":     os.ModeExclusive,
		"ModeTemporary":     os.ModeTemporary,
		"ModeSymlink":       os.ModeSymlink,
		"ModeDevice":        os.ModeDevice,
		"ModeNamedPipe":     os.ModeNamedPipe,
		"ModeSocket":        os.ModeSocket,
		"ModeSetuid":        os.ModeSetuid,
		"ModeSetgid":        os.ModeSetgid,
		"ModeCharDevice":    os.ModeCharDevice,
		"ModeSticky":        os.ModeSticky,
		"ModeIrregular":     os.ModeIrregular,
		"ModeType":          os.ModeType,
		"ModePerm":          os.ModePerm,
		"Create":            os.Create,
		"Open":              os.Open,
		"OpenFile":          os.OpenFile,
		"Stdin":             os.Stdin,
		"Stdout":            os.Stdout,
		"Stderr":            os.Stderr,
		"Args":              os.Args,
	})
}
