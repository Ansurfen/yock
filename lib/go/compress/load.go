// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package compresslib

import (
	yocki "github.com/ansurfen/yock/interface"
	bzip2lib "github.com/ansurfen/yock/lib/go/compress/bzip2"
	gziplib "github.com/ansurfen/yock/lib/go/compress/gzip"
	lzwlib "github.com/ansurfen/yock/lib/go/compress/lzw"
	zliblib "github.com/ansurfen/yock/lib/go/compress/zlib"
)

func LoadCompress(yocks yocki.YockScheduler) {
	lzwlib.LoadLzw(yocks)
	gziplib.LoadGzip(yocks)
	bzip2lib.LoadBzip2(yocks)
	zliblib.LoadZlib(yocks)
}
