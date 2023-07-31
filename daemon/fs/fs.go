package fs

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	du "github.com/ansurfen/yock/daemon/util"
	"github.com/ansurfen/yock/util"
)

type FileSystem struct {
	volumes map[string]*Volume
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		volumes: make(map[string]*Volume),
	}
}

func (fs *FileSystem) Put(src, dst string) error {
	vol, path := SplitPath(dst)
	if _, ok := fs.volumes[vol]; !ok {
		fs.volumes[vol] = NewVolume(vol)
	}
	for _, entry := range ParseDir(src, path) {
		fs.volumes[vol].Put(entry)
	}
	return nil
}

func copyFile(src string, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	util.Mkdirs(filepath.Dir(dst))
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) Get(src, dst string) error {
	vol, path := SplitPath(src)
	if v, ok := fs.volumes[vol]; ok {
		for _, p := range v.List(FormatPath(path)) {
			file := v.Get(FormatPath(path) + p)
			err := copyFile(file.Info(du.ID).Path, filepath.Join(dst, ResolvePath(FormatPath(path)+p)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (fs *FileSystem) Find(path string) File {
	vol, path := SplitPath(path)
	if v, ok := fs.volumes[vol]; ok {
		return v.Get(path)
	}
	return nil
}

func (fs *FileSystem) List(dir string) (ret []string) {
	vol, path := SplitPath(dir)
	if v, ok := fs.volumes[vol]; ok {
		for _, p := range v.List(FormatPath(path)) {
			ret = append(ret, ResolvePath(path)+ResolvePath(p))
		}
	}
	return
}

type DirectoryEntry struct {
	Dir  string
	Info FileInfo
}

func ParseDir(real, virtual string) (ret []DirectoryEntry) {
	real, err := filepath.Abs(real)
	if err != nil {
		panic(err)
	}
	pathInfo, err := os.Stat(real)
	if err != nil {
		panic(err)
	}
	if pathInfo.IsDir() {
		err = filepath.Walk(real, func(fullpath string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			relPath, err := filepath.Rel(real, fullpath)
			if err != nil {
				return err
			}
			pathInfo, err := os.Stat(fullpath)
			if err != nil {
				return err
			}
			raw, err := util.ReadStraemFromFile(fullpath)
			if err != nil {
				panic(err)
			}
			if !pathInfo.IsDir() {
				dir := FormatPath(filepath.Join(virtual, relPath))
				if len(dir) > 0 {
					dir = dir[:len(dir)-1]
				}
				ret = append(ret, DirectoryEntry{
					Dir: dir,
					Info: FileInfo{
						Owner:    du.ID,
						Path:     fullpath,
						CreateAt: time.Now().Unix(),
						Size:     pathInfo.Size(),
						Hash:     util.SHA256(string(raw)),
					},
				})
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	} else {
		raw, err := util.ReadStraemFromFile(real)
		if err != nil {
			panic(err)
		}
		dir := FormatPath(filepath.Join(virtual, filepath.Base(real)))
		if len(dir) > 0 {
			dir = dir[:len(dir)-1]
		}
		ret = append(ret, DirectoryEntry{
			Dir: dir,
			Info: FileInfo{
				Owner:    du.ID,
				Path:     real,
				CreateAt: time.Now().Unix(),
				Size:     pathInfo.Size(),
				Hash:     util.SHA256(string(raw)),
			},
		})
	}
	return
}
