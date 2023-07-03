// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package archive

import (
	yocki "github.com/ansurfen/yock/interface"
	tarlib "github.com/ansurfen/yock/lib/go/archive/tar"
	ziplib "github.com/ansurfen/yock/lib/go/archive/zip"
)

func LoadArchive(yocks yocki.YockScheduler) {
	ziplib.LoadZip(yocks)
	tarlib.LoadTar(yocks)
}
