package util

import "errors"

var (
	ErrFileNotExist = errors.New("file not exist")
	ErrInvalidFile  = errors.New("invalid file")

	ErrInvalidPort = errors.New("invalid port")

	ErrArgsTooLittle = errors.New("arguments is too little")

	ErrInvalidModuleName = errors.New("invalid module name")

	ErrGeneral = errors.New("err happen")
)
