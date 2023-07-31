// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import (
	"sync"
	"time"

	du "github.com/ansurfen/yock/daemon/util"
	"github.com/ansurfen/yock/util/container"
)

type Volume struct {
	index    container.Trie[File]
	createAt time.Time
	name     string
	mut      *sync.Mutex
}

func NewVolume(name string) *Volume {
	return &Volume{
		name:     name,
		index:    container.MapTrieOf[File](),
		mut:      &sync.Mutex{},
		createAt: time.Now(),
	}
}

func (v *Volume) Name() string {
	return v.name
}

func (v *Volume) Put(entry DirectoryEntry) {
	v.mut.Lock()
	defer v.mut.Unlock()
	if node, ok := v.index.FindNode(entry.Dir); !ok {
		err := v.index.Insert(entry.Dir, &SharedFile{
			Files: map[string]*atomicFileSet{
				du.ID: {
					atomicFile: atomicFile{
						Meta: entry.Info,
					},
				},
			},
			Name: entry.Dir,
		})
		if err != nil {
			panic(err)
		}
	} else {
		node.Value().Append(du.ID, entry.Info)
	}
}

func (v *Volume) Get(dir string) File {
	f, ok := v.index.FindNode(dir)
	if ok {
		return f.Value()
	}
	return nil
}

func (v *Volume) List(dir string) []string {
	return v.index.FindChildren(dir)
}

func (v *Volume) SwapIn() {}

func (v *Volume) SwapOut() {}
