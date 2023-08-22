// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/ansurfen/yock/auto/generator/archive"
)

func main() {
	archive.SetInfo("The Yock Authors", "MIT")
	archive.EnableYockComment()
	archive.ExportGoFile()
	archive.LoadDir("bytes", `D:\D\langs\go\src\bytes`)
	// archive.LoadDir("path", `D:\D\langs\go\src\path`)
	// archive.LoadDir("gin", `D:\D\langs\go\pkg\mod\github.com\gin-gonic\gin@v1.9.1`)
	// archive.LoadDir("tar", `D:\D\langs\go\src\archive\tar`)
	// archive.LoadDir("zip", `D:\D\langs\go\src\archive\zip`)
	// archive.LoadDir("bufio", `D:\D\langs\go\src\bufio`)
	// archive.LoadDir("bzip2", `D:\D\langs\go\src\compress\bzip2`)
	// archive.LoadDir("flate", `D:\D\langs\go\src\compress\flate`)
	// archive.LoadDir("gzip", `D:\D\langs\go\src\compress\gzip`)
	// archive.LoadDir("lzw", `D:\D\langs\go\src\compress\lzw`)
	// archive.LoadDir("zlib", `D:\D\langs\go\src\compress\zlib`)
	// archive.LoadDir("filepath", `D:\D\langs\go\src\path\filepath`)
	// archive.LoadDir("os", `D:\D\langs\go\src\os`)
	// archive.LoadDir("exec", `D:\D\langs\go\src\os\exec`)
	// archive.LoadDir("signal", `D:\D\langs\go\src\os\signal`)
	// archive.LoadDir("user", `D:\D\langs\go\src\os\user`)
	// archive.LoadDir("net", `D:\D\langs\go\src\net`)
	// archive.LoadDir("http", `D:\D\langs\go\src\net\http`)
	// archive.LoadDir("mail", `D:\D\langs\go\src\net\mail`)
	// archive.LoadDir("netip", `D:\D\langs\go\src\net\netip`)
	// archive.LoadDir("rpc", `D:\D\langs\go\src\net\rpc`)
	// archive.LoadDir("smtp", `D:\D\langs\go\src\net\smtp`)
	// archive.LoadDir("url", `D:\D\langs\go\src\net\url`)
	// archive.LoadDir("textproto", `D:\D\langs\go\src\net\textproto`)
	// archive.LoadDir("reflect", `D:\D\langs\go\src\reflect`)
	// archive.LoadDir("time", `D:\D\langs\go\src\time`)
	// archive.LoadDir("sync", `D:\D\langs\go\src\sync`)
	// archive.LoadDir("atomic", `D:\D\langs\go\src\sync\atomic`)
	// archive.LoadDir("io", `D:\D\langs\go\src\io`)
	// archive.LoadDir("fs", `D:\D\langs\go\src\io\fs`)
	// archive.LoadDir("ioutil", `D:\D\langs\go\src\io\ioutil`)
	archive.Export("tmp")
}
