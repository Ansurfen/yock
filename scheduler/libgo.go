// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yocks

import (
	"github.com/ansurfen/yock/lib/go/archive"
	"github.com/ansurfen/yock/lib/go/bufio"
	compresslib "github.com/ansurfen/yock/lib/go/compress"
	"github.com/ansurfen/yock/lib/go/fmt"
	iolib "github.com/ansurfen/yock/lib/go/io"
	netlib "github.com/ansurfen/yock/lib/go/net"
	"github.com/ansurfen/yock/lib/go/os"
	"github.com/ansurfen/yock/lib/go/path"
	reflectlib "github.com/ansurfen/yock/lib/go/reflect"
	"github.com/ansurfen/yock/lib/go/regexp"
	"github.com/ansurfen/yock/lib/go/strconv"
	libstrings "github.com/ansurfen/yock/lib/go/strings"
	libsync "github.com/ansurfen/yock/lib/go/sync"
	libtime "github.com/ansurfen/yock/lib/go/time"
	"github.com/ansurfen/yock/lib/go/unicode"
)

var libgo = []loader{
	reflectlib.LoadReflect,
	fmt.LoadFmt,
	netlib.LoadNet,
	path.LoadPath,
	regexp.LoadRegexp,
	libstrings.LoadStrings,
	libtime.LoadTime,
	libsync.LoadSync,
	iolib.LoadIo,
	bufio.LoadBufio,
	unicode.LoadUnicode,
	oslib.LoadOs,
	strconv.LoadStrconv,
	compresslib.LoadCompress,
	archive.LoadArchive,
}
