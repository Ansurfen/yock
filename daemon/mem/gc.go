// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mem

import "github.com/ansurfen/yock/daemon/fs"

type MemWatch struct {
	fcache *fs.FileSystemCache
}

func (gc *MemWatch) GabargeCollect() error {
	gc.fcache.Clean()
	return nil
}
