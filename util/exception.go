// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "errors"

var (
	ErrArgsTooLittle = errors.New("arguments is too little")

	ErrInvalidModuleName = errors.New("invalid module name")

	ErrDumplicateJobName = errors.New("dumplicate job name")

	ErrCreateSession  = errors.New("fail to create session")
	ErrExecuteCommand = errors.New("fail to execute command")
	ErrAllocTerm      = errors.New("fail to allocate term")
	ErrAllocShell     = errors.New("fail to allocate shell")

	// net
	ErrInvalidPort      = errors.New("invalid port")
	ErrInvalidURL       = errors.New("invalid url")
	ErrInvalidMethod    = errors.New("invalid method")
	ErrBadCreateRequest = errors.New("error createing request")
	ErrBadSendRequest   = errors.New("error sending request")

	// os

	ErrNoSupportPlatform = errors.New("not support the platform")
	ErrNoSupportHardward = errors.New("not support the hardward")

	ErrInvalidPath = errors.New("invalid path")

	// io

	ErrBadCreateDir  = errors.New("fail to create dir")
	ErrBadCreateFile = errors.New("fail to create file")

	ErrInvalidFile  = errors.New("invalid file")
	ErrFileNotExist = errors.New("file not exist")
	ErrFileExist    = errors.New("file exist")

	// yock plugin
	ErrPluginExist = errors.New("plugin exist already")
	ErrDomainExist = errors.New("domain exist already")
	ErrAliasExist  = errors.New("alias exist already")

	ErrGeneral = errors.New("err happen")
)
