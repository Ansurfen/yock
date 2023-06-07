// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

type FileInfo struct {
	owner    string
	size     int64
	hash     string
	createAt string
}
