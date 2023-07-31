// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import "github.com/ansurfen/yock/util/cache"

var _ cache.Entry[string, *Volume] = (*VolumeEntry)(nil)

type VolumeEntry struct {
	name   string
	volume *Volume
}

func (entry *VolumeEntry) SetKey(k string) {
	entry.name = k
}

func (entry *VolumeEntry) Key() string { return entry.name }

func (entry *VolumeEntry) SetValue(v *Volume) {
	entry.volume = v
}

func (entry *VolumeEntry) Value() *Volume { return entry.volume }

func (entry *VolumeEntry) Free() {
	entry.volume.SwapIn()
	// TODO: delete in filesystem
}

type FileSystemCache struct {
	volumes cache.Cache[string, *Volume]
}

func newVolumeEntry() cache.Entry[string, *Volume] {
	return &VolumeEntry{}
}

func newFCache() *FileSystemCache {
	return &FileSystemCache{
		volumes: cache.NewLRU(10, newVolumeEntry),
	}
}

func (fcache *FileSystemCache) Clean() {

}
