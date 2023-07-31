// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	du "github.com/ansurfen/yock/daemon/util"
	"github.com/ansurfen/yock/util"
)

type File interface {
	Open(owner, wanter string) ([]string, error)
	Close(fd string)
	Info(owner string) FileInfo
	Append(owner string, meta FileInfo)
}

const gen_fd = "abcdefg"

type SharedFile struct {
	Files map[string]*atomicFileSet
	Name  string
}

func encodeFd(file, owner, consumer string) string {
	return util.EncodeAESWithKey(gen_fd, file+"#"+owner+"#"+consumer)
}

func decodeFd(fd string) (file, owner, consumer string) {
	res := strings.Split(util.DecodeAESWithKey(gen_fd, fd), "#")
	if len(res) != 3 {
		return
	}
	return res[0], res[1], res[2]
}

func (fp *SharedFile) Append(producer string, meta FileInfo) {
	if cf, ok := fp.Files[producer]; ok {
		cf.atomicFile.Meta = meta
	} else {
		fp.Files[producer] = &atomicFileSet{
			atomicFile: atomicFile{
				Meta: meta,
			},
		}
	}
}

func (fp *SharedFile) Open(producer, consumer string) (ret []string, err error) {
	if producer == "*" {
		for name, f := range fp.Files {
			err = f.Open(consumer)
			if err != nil {
				for _, r := range ret {
					fp.Close(r)
				}
				ret = nil
				return
			}
			ret = append(ret, encodeFd(fp.Name, name, consumer))
		}
	}
	if file, ok := fp.Files[producer]; ok {
		err = file.Open(consumer)
		if err != nil {
			return
		}
		ret = append(ret, util.EncodeAESWithKey(gen_fd, fp.Name+"#"+producer+"#"+consumer))
	}
	return
}

func (file *SharedFile) Close(fd string) {
	_, owner, consumer := decodeFd(fd)
	if f, ok := file.Files[owner]; ok {
		f.mut.Unlock(consumer)
	}
}

func (file *SharedFile) Info(producer string) FileInfo {
	return file.Files[producer].Meta
}

type atomicFileSet struct {
	atomicFile
	Replication []atomicFile
}

type atomicFile struct {
	mut  *fileLocker
	Meta FileInfo
	ptr  *os.File
}

func newSingleFile(meta FileInfo) atomicFile {
	return atomicFile{
		Meta: meta,
		mut:  &fileLocker{},
	}
}

func (file *atomicFileSet) Clone(newOwner string) {
	found := false
	for _, replication := range file.Replication {
		if replication.Meta.Owner == newOwner {
			found = true
			break
		}
	}
	if !found {
		meta := file.Meta
		meta.Owner = newOwner
		file.Replication = append(file.Replication, atomicFile{
			Meta: meta,
			mut:  &fileLocker{},
		})
	}
}

type fileLocker struct {
	holder string
	count  int
}

func (locker *fileLocker) Unlock(owner string) {
	if locker.holder == owner {
		locker.count--
		if locker.count == 0 {
			locker.holder = ""
			locker.count = 0
		}
	}
}

func (locker *fileLocker) Lock(consumer string) bool {
	if len(locker.holder) == 0 {
		locker.holder = consumer
		locker.count = 1
		return true
	}
	if locker.holder == consumer {
		locker.count++
		return true
	}
	return false
}

func (locker *fileLocker) IsLock(owner string) bool {
	return len(locker.holder) != 0 && locker.holder != owner
}

func (fp *atomicFile) Open(consumer string) error {
	if fp.mut.IsLock(consumer) {
		return fmt.Errorf("file is opened")
	}
	if fp.Meta.Owner == du.ID {
		if !fp.mut.Lock(consumer) {
			return fmt.Errorf("file fails to lock")
		}
		path := decode(fp.Meta.Path)
		ptr, err := os.Open(path)
		if err != nil {
			return err
		}
		fp.ptr = ptr
	} else {
		return fmt.Errorf("file not found")
	}
	return nil
}

func (f *SharedFile) Save() string {
	raw, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

type FileInfo struct {
	Owner    string `json:"owner"`
	Size     int64  `json:"size"`
	Hash     string `json:"hash"`
	CreateAt int64  `json:"createAt"`
	Path     string `json:"path"`
}

func decode(path string) string {
	return path
}
