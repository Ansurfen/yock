// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import "os"

type File interface {
	Open() error
	Close() error
	Info() FileInfo
}

type TruncateFile struct {
	ptr  *os.File
	info FileInfo
}

type ExpansiveFile struct {
	infos []FileInfo
}

type FileInfo struct {
	Owner    string `json:"owner"`
	Size     int64  `json:"size"`
	Hash     string `json:"hash"`
	CreateAt string `json:"createAt"`
}

func (info FileInfo) String() string {
	return ""
}
