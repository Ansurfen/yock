// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import "time"

type FileSystem struct {
	volumes map[string]*Volume
	fcache  *FileSystemCache
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		volumes: make(map[string]*Volume),
		fcache:  newFCache(),
	}
}

func (fs *FileSystem) Mount(volume string) {
	fs.volumes[volume] = newVolume(volume, VolumeMeta{
		createAt: time.Now(),
	})
}

func (fs *FileSystem) Unmount(volume string) {
}

func (fs *FileSystem) FindFile(Volume, filename string) (FileInfo, bool) {
	return FileInfo{}, false
}

func (fs *FileSystem) CreateFile(volume string, file FileInfo) {
}

func (fs *FileSystem) UpdateFile(volume string, file File) {
}

func (fs *FileSystem) RemoveFile(volume string, file File) {
}

func (fs *FileSystem) OpenFile(volume, file string) {
}

func (fs *FileSystem) Swap() {

}
