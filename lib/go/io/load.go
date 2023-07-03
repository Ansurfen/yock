// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package iolib

import (
	"io"

	yocki "github.com/ansurfen/yock/interface"
	fslib "github.com/ansurfen/yock/lib/go/io/fs"
	ioutillib "github.com/ansurfen/yock/lib/go/io/ioutil"
)

func LoadIo(yocks yocki.YockScheduler) {
	fslib.LoadFs(yocks)
	ioutillib.LoadIoutil(yocks)
	lib := yocks.OpenLib("io")
	lib.SetField(map[string]any{
		// functions
		"CopyN":            io.CopyN,
		"Copy":             io.Copy,
		"CopyBuffer":       io.CopyBuffer,
		"Pipe":             io.Pipe,
		"ReadAtLeast":      io.ReadAtLeast,
		"WriteString":      io.WriteString,
		"NewSectionReader": io.NewSectionReader,
		"MultiReader":      io.MultiReader,
		"LimitReader":      io.LimitReader,
		"ReadAll":          io.ReadAll,
		"ReadFull":         io.ReadFull,
		"NewOffsetWriter":  io.NewOffsetWriter,
		"NopCloser":        io.NopCloser,
		"TeeReader":        io.TeeReader,
		"MultiWriter":      io.MultiWriter,
		// constants
		"SeekStart":   io.SeekStart,
		"SeekCurrent": io.SeekCurrent,
		"SeekEnd":     io.SeekEnd,
		// variable
		"ErrShortWrite":    io.ErrShortWrite,
		"ErrShortBuffer":   io.ErrShortBuffer,
		"EOF":              io.EOF,
		"ErrUnexpectedEOF": io.ErrUnexpectedEOF,
		"ErrNoProgress":    io.ErrNoProgress,
		"Discard":          io.Discard,
		"ErrClosedPipe":    io.ErrClosedPipe,
	})
}
