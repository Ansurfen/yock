package cmd

import "errors"

var (
	// net

	ErrInvalidURL       = errors.New("invalid url")
	ErrInvalidMethod    = errors.New("invalid method")
	ErrBadCreateRequest = errors.New("error createing request")
	ErrBadSendRequest   = errors.New("error sending request")

	// os

	ErrNoSupportPlatform = errors.New("not support platform")

	// io

	ErrBadCreateDir  = errors.New("fail to create dir")
	ErrBadCreateFile = errors.New("fail to create file")

	ErrGeneral = errors.New("error happen")
)
