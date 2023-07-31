// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package process

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/ansurfen/yock/ycho"
	"github.com/fsnotify/fsnotify"
)

type OSNotify struct {
	fsn        *fsnotify.Watcher
	fsCallback map[string][]int
	fsCbLock   *sync.Mutex

	handles    []*OSHandle
	handleLock *sync.Mutex
}

func NewOSNotify() *OSNotify {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	return &OSNotify{
		fsn:        watch,
		fsCallback: make(map[string][]int),
		fsCbLock:   &sync.Mutex{},
		handleLock: &sync.Mutex{},
	}
}

func (n *OSNotify) Remove(id int) {
	n.handleLock.Lock()
	defer n.handleLock.Unlock()
	if id >= 0 && id < len(n.handles) {
		n.handles[id].enable = false
	}
}

type OSHandle struct {
	handle func()
	enable bool
}

func (n *OSNotify) AddFunc(paths []string, cmd func()) int {
	n.handleLock.Lock()
	n.handles = append(n.handles, &OSHandle{
		enable: true,
		handle: cmd,
	})
	id := len(n.handles) - 1
	n.handleLock.Unlock()

	n.fsCbLock.Lock()
	defer n.fsCbLock.Unlock()
	for _, path := range paths {
		n.fsn.Add(path)
		n.fsCallback[path] = append(n.fsCallback[path], id)
	}
	return id
}

func (n *OSNotify) Listen() {
	for {
		select {
		// file system
		case event, ok := <-n.fsn.Events:
			if !ok {
				continue
			}
			cur := event.Name
			prev := filepath.Join(cur, "..")
			for cur != prev {
				if cb, ok := n.fsCallback[cur]; ok {
					for _, c := range cb {
						handle := n.handles[c]
						if handle.enable {
							n.handles[c].handle()
						}
					}
				}
				cur = filepath.Join(cur, "..")
				prev = filepath.Join(prev, "..")
			}
		case err, ok := <-n.fsn.Errors:
			if !ok {
				// ycho.Error()
				continue
			}
			ycho.Error(err)

		default:
			// port listen

			// hardware listen

			// script
			time.Sleep(500 * time.Microsecond)
		}
	}
}
