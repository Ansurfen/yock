// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/ansurfen/yock/daemon/util"
)

const (
	user1 = "ff:ff:ff:ff:ff:ff"
	user2 = "00:11:22:33:44:55"
	user3 = "01:23:45:67:89:10"
)

func TestSharedFile(t *testing.T) {
	file := SharedFile{
		Name: "%1.txt",
		Files: map[string]*atomicFileSet{
			user1: {
				atomicFile: newSingleFile(FileInfo{
					Owner:    util.ID,
					Size:     1024,
					CreateAt: time.Now().Unix(),
					Path:     "./testdata/a/1.txt",
				}),
			},
			user2: {
				atomicFile: newSingleFile(FileInfo{
					Owner:    util.ID,
					Size:     1024,
					CreateAt: time.Now().Unix(),
					Path:     "./testdata/a/1.txt",
				}),
			},
			user3: {
				atomicFile: newSingleFile(FileInfo{
					Owner:    util.ID,
					Size:     1024,
					CreateAt: time.Now().Unix(),
					Path:     "./testdata/a/1.txt",
				}),
			},
		},
	}
	file.Files[user1].Clone(user2)
	fmt.Println(file.Files[user1])
	// test to open a file multiple times by the same user
	fds, _ := file.Open(user1, util.ID)
	fd := fds[0]
	fmt.Println(file.Open(user1, util.ID))
	// test if opening the same file by different users will be rejected
	fmt.Println(file.Open(user1, user1))
	// test file lock count
	file.Close(fd)
	file.Close(fd)
	file.Close(fd)
	// fmt.Println(file.Open(user1, user1))
	fmt.Println(file.Open("*", user2))
	fmt.Println(file.Open(user1, user1))
	fmt.Println(file.Open(user1, util.ID))
	// fmt.Println(file.Open("*", util.ID))
	// fmt.Println(file.Open(user1, user2))
	raw, err := json.Marshal(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))
}
