package cmd

import "errors"

const (
	// net

	ErrInvalidURL     = "invalid url"
	ErrInvalidMethod  = "invalid method"
	BadCreateRequest  = "error createing request"
	ErrBadSendRequest = "error sending request"

	// os

	ErrNoSupportPlatform = "not support platform"

	// io

	ErrBadCreateDir  = "fail to create dir"
	ErrBadCreateFile = "fail to create file"
)

var (
	ErrGeneral = errors.New("error happen")
)
